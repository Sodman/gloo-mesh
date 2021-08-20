// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package local

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/resource"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, gvk schema.GroupVersionKind, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, gvk.String(), sets.Key(obj))
}

// SnapshotGVKs is a list of the GVKs included in this snapshot
var SnapshotGVKs = []schema.GroupVersionKind{

	schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Secret",
	},
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of Secrets with a given set of labels
	Secrets() []LabeledSecretSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, opts output.OutputOpts)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, opts output.OutputOpts)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)

	// convert this snapshot to its generic form
	Generic() resource.ClusterSnapshot

	// iterate over the objects contained in the snapshot
	ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject))
}

type snapshot struct {
	name string

	secrets  []LabeledSecretSet
	clusters []string
}

func NewSnapshot(
	name string,

	secrets []LabeledSecretSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) Snapshot {
	return &snapshot{
		name: name,

		secrets:  secrets,
		clusters: clusters,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	secrets v1_sets.SecretSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	partitionedSecrets, err := partitionSecretsByLabel(labelKey, secrets)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedSecrets,
		clusters...,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	secrets v1_sets.SecretSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	labeledSecrets, err := NewLabeledSecretSet(secrets, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledSecretSet{labeledSecrets},
		clusters...,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, clusterClient client.Client, opts output.OutputOpts) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.secrets {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncLocalCluster(ctx, clusterClient, opts)
}

// apply the desired resources to multiple cluster states; remove stale resources where necessary
func (s *snapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, opts output.OutputOpts) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.secrets {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		Clusters:    s.clusters,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, opts)
}

func (s *snapshot) Generic() resource.ClusterSnapshot {
	clusterSnapshots := resource.ClusterSnapshot{}
	s.ForEachObject(func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject) {
		clusterSnapshots.Insert(cluster, gvk, obj)
	})

	return clusterSnapshots
}

// convert this snapshot to its generic form
func (s *snapshot) ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject)) {

	for _, set := range s.secrets {
		for _, obj := range set.Set().List() {
			cluster := obj.GetClusterName()
			gvk := schema.GroupVersionKind{
				Group:   "",
				Version: "v1",
				Kind:    "Secret",
			}
			handleObject(cluster, gvk, obj)
		}
	}
}

func partitionSecretsByLabel(labelKey string, set v1_sets.SecretSet) ([]LabeledSecretSet, error) {
	setsByLabel := map[string]v1_sets.SecretSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Secret", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Secret", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = v1_sets.NewSecretSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedSecrets []LabeledSecretSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledSecretSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedSecrets = append(partitionedSecrets, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedSecrets, func(i, j int) bool {
		leftLabelValue := partitionedSecrets[i].Labels()[labelKey]
		rightLabelValue := partitionedSecrets[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedSecrets, nil
}

func (s snapshot) Secrets() []LabeledSecretSet {
	return s.secrets
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	secretSet := v1_sets.NewSecretSet()
	for _, set := range s.secrets {
		secretSet = secretSet.Union(set.Set())
	}
	snapshotMap["secrets"] = secretSet.List()

	snapshotMap["clusters"] = s.clusters

	return json.Marshal(snapshotMap)
}

// LabeledSecretSet represents a set of secrets
// which share a common set of labels.
// These labels are used to find diffs between SecretSets.
type LabeledSecretSet interface {
	// returns the set of Labels shared by this SecretSet
	Labels() map[string]string

	// returns the set of Secretes with the given labels
	Set() v1_sets.SecretSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledSecretSet struct {
	set    v1_sets.SecretSet
	labels map[string]string
}

func NewLabeledSecretSet(set v1_sets.SecretSet, labels map[string]string) (LabeledSecretSet, error) {
	// validate that each Secret contains the labels, else this is not a valid LabeledSecretSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Secret %v", k, v, item.Name)
			}
		}
	}

	return &labeledSecretSet{set: set, labels: labels}, nil
}

func (l *labeledSecretSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledSecretSet) Set() v1_sets.SecretSet {
	return l.set
}

func (l labeledSecretSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list v1.SecretList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources: desiredResources,
		ListFunc:  listFunc,
		GVK: schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Secret",
		},
	}
}

type builder struct {
	ctx      context.Context
	name     string
	clusters []string

	secrets v1_sets.SecretSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		secrets: v1_sets.NewSecretSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add Secrets to the collected outputs
	AddSecrets(secrets ...*v1.Secret)

	// get the collected Secrets
	GetSecrets() v1_sets.SecretSet

	// build the collected outputs into a label-partitioned snapshot
	BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error)

	// build the collected outputs into a snapshot with a single partition
	BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error)

	// add a cluster to the collected clusters.
	// this can be used to collect clusters for use with MultiCluster snapshots.
	AddCluster(cluster string)

	// returns the set of clusters currently stored in this builder
	Clusters() []string

	// merge all the resources from another Builder into this one
	Merge(other Builder)

	// create a clone of this builder (deepcopying all resources)
	Clone() Builder

	// convert this snapshot to its generic form
	Generic() resource.ClusterSnapshot

	// iterate over the objects contained in the snapshot
	ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject))
}

func (b *builder) AddSecrets(secrets ...*v1.Secret) {
	for _, obj := range secrets {
		if obj == nil {
			continue
		}
		b.secrets.Insert(obj)
	}
}

func (b *builder) GetSecrets() v1_sets.SecretSet {
	return b.secrets
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.secrets,
		b.clusters...,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.secrets,
		b.clusters...,
	)
}

func (b *builder) AddCluster(cluster string) {
	b.clusters = append(b.clusters, cluster)
}

func (b *builder) Clusters() []string {
	return b.clusters
}

func (b *builder) Merge(other Builder) {
	if other == nil {
		return
	}

	b.AddSecrets(other.GetSecrets().List()...)
	for _, cluster := range other.Clusters() {
		b.AddCluster(cluster)
	}
}

func (b *builder) Clone() Builder {
	if b == nil {
		return nil
	}
	clone := NewBuilder(b.ctx, b.name)

	for _, secret := range b.GetSecrets().List() {
		clone.AddSecrets(secret.DeepCopy())
	}
	for _, cluster := range b.Clusters() {
		clone.AddCluster(cluster)
	}
	return clone
}

// convert this snapshot to its generic form
func (b *builder) Generic() resource.ClusterSnapshot {
	if b == nil {
		return nil
	}
	clusterSnapshots := resource.ClusterSnapshot{}

	for _, obj := range b.GetSecrets().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Secret",
		}
		clusterSnapshots.Insert(cluster, gvk, obj)
	}

	return clusterSnapshots
}

// convert this snapshot to its generic form
func (b *builder) ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject)) {
	if b == nil {
		return
	}

	for _, obj := range b.GetSecrets().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Secret",
		}
		handleObject(cluster, gvk, obj)
	}
}
