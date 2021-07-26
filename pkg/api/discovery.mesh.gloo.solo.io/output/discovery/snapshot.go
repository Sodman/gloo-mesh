// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package discovery

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

	discovery_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discovery_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// SnapshotGVKs is a list of the GVKs included in this snapshot
var SnapshotGVKs = []schema.GroupVersionKind{

	schema.GroupVersionKind{
		Group:   "discovery.mesh.gloo.solo.io",
		Version: "v1",
		Kind:    "Destination",
	},
	schema.GroupVersionKind{
		Group:   "discovery.mesh.gloo.solo.io",
		Version: "v1",
		Kind:    "Workload",
	},
	schema.GroupVersionKind{
		Group:   "discovery.mesh.gloo.solo.io",
		Version: "v1",
		Kind:    "Mesh",
	},
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of Destinations with a given set of labels
	Destinations() []LabeledDestinationSet
	// return the set of Workloads with a given set of labels
	Workloads() []LabeledWorkloadSet
	// return the set of Meshes with a given set of labels
	Meshes() []LabeledMeshSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)

	// convert this snapshot to its generic form
	Generic() resource.ClusterSnapshot

	// iterate over the objects contained in the snapshot
	ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject))
}

type snapshot struct {
	name string

	destinations []LabeledDestinationSet
	workloads    []LabeledWorkloadSet
	meshes       []LabeledMeshSet
	clusters     []string
}

func NewSnapshot(
	name string,

	destinations []LabeledDestinationSet,
	workloads []LabeledWorkloadSet,
	meshes []LabeledMeshSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) Snapshot {
	return &snapshot{
		name: name,

		destinations: destinations,
		workloads:    workloads,
		meshes:       meshes,
		clusters:     clusters,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	destinations discovery_mesh_gloo_solo_io_v1_sets.DestinationSet,
	workloads discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet,
	meshes discovery_mesh_gloo_solo_io_v1_sets.MeshSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	partitionedDestinations, err := partitionDestinationsByLabel(labelKey, destinations)
	if err != nil {
		return nil, err
	}
	partitionedWorkloads, err := partitionWorkloadsByLabel(labelKey, workloads)
	if err != nil {
		return nil, err
	}
	partitionedMeshes, err := partitionMeshesByLabel(labelKey, meshes)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedDestinations,
		partitionedWorkloads,
		partitionedMeshes,
		clusters...,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	destinations discovery_mesh_gloo_solo_io_v1_sets.DestinationSet,
	workloads discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet,
	meshes discovery_mesh_gloo_solo_io_v1_sets.MeshSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	labeledDestinations, err := NewLabeledDestinationSet(destinations, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledWorkloads, err := NewLabeledWorkloadSet(workloads, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledMeshes, err := NewLabeledMeshSet(meshes, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledDestinationSet{labeledDestinations},
		[]LabeledWorkloadSet{labeledWorkloads},
		[]LabeledMeshSet{labeledMeshes},
		clusters...,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, cli client.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.destinations {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.workloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncLocalCluster(ctx, cli, errHandler)
}

// apply the desired resources to multiple cluster states; remove stale resources where necessary
func (s *snapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.destinations {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.workloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		Clusters:    s.clusters,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, errHandler)
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

	for _, set := range s.destinations {
		for _, obj := range set.Set().List() {
			cluster := obj.GetClusterName()
			gvk := schema.GroupVersionKind{
				Group:   "discovery.mesh.gloo.solo.io",
				Version: "v1",
				Kind:    "Destination",
			}
			handleObject(cluster, gvk, obj)
		}
	}
	for _, set := range s.workloads {
		for _, obj := range set.Set().List() {
			cluster := obj.GetClusterName()
			gvk := schema.GroupVersionKind{
				Group:   "discovery.mesh.gloo.solo.io",
				Version: "v1",
				Kind:    "Workload",
			}
			handleObject(cluster, gvk, obj)
		}
	}
	for _, set := range s.meshes {
		for _, obj := range set.Set().List() {
			cluster := obj.GetClusterName()
			gvk := schema.GroupVersionKind{
				Group:   "discovery.mesh.gloo.solo.io",
				Version: "v1",
				Kind:    "Mesh",
			}
			handleObject(cluster, gvk, obj)
		}
	}
}

func partitionDestinationsByLabel(labelKey string, set discovery_mesh_gloo_solo_io_v1_sets.DestinationSet) ([]LabeledDestinationSet, error) {
	setsByLabel := map[string]discovery_mesh_gloo_solo_io_v1_sets.DestinationSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Destination", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Destination", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_mesh_gloo_solo_io_v1_sets.NewDestinationSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedDestinations []LabeledDestinationSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledDestinationSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedDestinations = append(partitionedDestinations, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedDestinations, func(i, j int) bool {
		leftLabelValue := partitionedDestinations[i].Labels()[labelKey]
		rightLabelValue := partitionedDestinations[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedDestinations, nil
}

func partitionWorkloadsByLabel(labelKey string, set discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet) ([]LabeledWorkloadSet, error) {
	setsByLabel := map[string]discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Workload", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Workload", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_mesh_gloo_solo_io_v1_sets.NewWorkloadSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedWorkloads []LabeledWorkloadSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledWorkloadSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedWorkloads = append(partitionedWorkloads, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedWorkloads, func(i, j int) bool {
		leftLabelValue := partitionedWorkloads[i].Labels()[labelKey]
		rightLabelValue := partitionedWorkloads[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedWorkloads, nil
}

func partitionMeshesByLabel(labelKey string, set discovery_mesh_gloo_solo_io_v1_sets.MeshSet) ([]LabeledMeshSet, error) {
	setsByLabel := map[string]discovery_mesh_gloo_solo_io_v1_sets.MeshSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_mesh_gloo_solo_io_v1_sets.NewMeshSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshes []LabeledMeshSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshes = append(partitionedMeshes, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshes, func(i, j int) bool {
		leftLabelValue := partitionedMeshes[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshes[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshes, nil
}

func (s snapshot) Destinations() []LabeledDestinationSet {
	return s.destinations
}

func (s snapshot) Workloads() []LabeledWorkloadSet {
	return s.workloads
}

func (s snapshot) Meshes() []LabeledMeshSet {
	return s.meshes
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	destinationSet := discovery_mesh_gloo_solo_io_v1_sets.NewDestinationSet()
	for _, set := range s.destinations {
		destinationSet = destinationSet.Union(set.Set())
	}
	snapshotMap["destinations"] = destinationSet.List()
	workloadSet := discovery_mesh_gloo_solo_io_v1_sets.NewWorkloadSet()
	for _, set := range s.workloads {
		workloadSet = workloadSet.Union(set.Set())
	}
	snapshotMap["workloads"] = workloadSet.List()
	meshSet := discovery_mesh_gloo_solo_io_v1_sets.NewMeshSet()
	for _, set := range s.meshes {
		meshSet = meshSet.Union(set.Set())
	}
	snapshotMap["meshes"] = meshSet.List()

	snapshotMap["clusters"] = s.clusters

	return json.Marshal(snapshotMap)
}

// LabeledDestinationSet represents a set of destinations
// which share a common set of labels.
// These labels are used to find diffs between DestinationSets.
type LabeledDestinationSet interface {
	// returns the set of Labels shared by this DestinationSet
	Labels() map[string]string

	// returns the set of Destinationes with the given labels
	Set() discovery_mesh_gloo_solo_io_v1_sets.DestinationSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledDestinationSet struct {
	set    discovery_mesh_gloo_solo_io_v1_sets.DestinationSet
	labels map[string]string
}

func NewLabeledDestinationSet(set discovery_mesh_gloo_solo_io_v1_sets.DestinationSet, labels map[string]string) (LabeledDestinationSet, error) {
	// validate that each Destination contains the labels, else this is not a valid LabeledDestinationSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Destination %v", k, v, item.Name)
			}
		}
	}

	return &labeledDestinationSet{set: set, labels: labels}, nil
}

func (l *labeledDestinationSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledDestinationSet) Set() discovery_mesh_gloo_solo_io_v1_sets.DestinationSet {
	return l.set
}

func (l labeledDestinationSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_mesh_gloo_solo_io_v1.DestinationList
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
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "Destination",
	}
}

// LabeledWorkloadSet represents a set of workloads
// which share a common set of labels.
// These labels are used to find diffs between WorkloadSets.
type LabeledWorkloadSet interface {
	// returns the set of Labels shared by this WorkloadSet
	Labels() map[string]string

	// returns the set of Workloades with the given labels
	Set() discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledWorkloadSet struct {
	set    discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet
	labels map[string]string
}

func NewLabeledWorkloadSet(set discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet, labels map[string]string) (LabeledWorkloadSet, error) {
	// validate that each Workload contains the labels, else this is not a valid LabeledWorkloadSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Workload %v", k, v, item.Name)
			}
		}
	}

	return &labeledWorkloadSet{set: set, labels: labels}, nil
}

func (l *labeledWorkloadSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledWorkloadSet) Set() discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet {
	return l.set
}

func (l labeledWorkloadSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_mesh_gloo_solo_io_v1.WorkloadList
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
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "Workload",
	}
}

// LabeledMeshSet represents a set of meshes
// which share a common set of labels.
// These labels are used to find diffs between MeshSets.
type LabeledMeshSet interface {
	// returns the set of Labels shared by this MeshSet
	Labels() map[string]string

	// returns the set of Meshes with the given labels
	Set() discovery_mesh_gloo_solo_io_v1_sets.MeshSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshSet struct {
	set    discovery_mesh_gloo_solo_io_v1_sets.MeshSet
	labels map[string]string
}

func NewLabeledMeshSet(set discovery_mesh_gloo_solo_io_v1_sets.MeshSet, labels map[string]string) (LabeledMeshSet, error) {
	// validate that each Mesh contains the labels, else this is not a valid LabeledMeshSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Mesh %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshSet{set: set, labels: labels}, nil
}

func (l *labeledMeshSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshSet) Set() discovery_mesh_gloo_solo_io_v1_sets.MeshSet {
	return l.set
}

func (l labeledMeshSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_mesh_gloo_solo_io_v1.MeshList
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
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "Mesh",
	}
}

type builder struct {
	ctx      context.Context
	name     string
	clusters []string

	destinations discovery_mesh_gloo_solo_io_v1_sets.DestinationSet
	workloads    discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet
	meshes       discovery_mesh_gloo_solo_io_v1_sets.MeshSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		destinations: discovery_mesh_gloo_solo_io_v1_sets.NewDestinationSet(),
		workloads:    discovery_mesh_gloo_solo_io_v1_sets.NewWorkloadSet(),
		meshes:       discovery_mesh_gloo_solo_io_v1_sets.NewMeshSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add Destinations to the collected outputs
	AddDestinations(destinations ...*discovery_mesh_gloo_solo_io_v1.Destination)

	// get the collected Destinations
	GetDestinations() discovery_mesh_gloo_solo_io_v1_sets.DestinationSet

	// add Workloads to the collected outputs
	AddWorkloads(workloads ...*discovery_mesh_gloo_solo_io_v1.Workload)

	// get the collected Workloads
	GetWorkloads() discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet

	// add Meshes to the collected outputs
	AddMeshes(meshes ...*discovery_mesh_gloo_solo_io_v1.Mesh)

	// get the collected Meshes
	GetMeshes() discovery_mesh_gloo_solo_io_v1_sets.MeshSet

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

func (b *builder) AddDestinations(destinations ...*discovery_mesh_gloo_solo_io_v1.Destination) {
	for _, obj := range destinations {
		if obj == nil {
			continue
		}
		b.destinations.Insert(obj)
	}
}
func (b *builder) AddWorkloads(workloads ...*discovery_mesh_gloo_solo_io_v1.Workload) {
	for _, obj := range workloads {
		if obj == nil {
			continue
		}
		b.workloads.Insert(obj)
	}
}
func (b *builder) AddMeshes(meshes ...*discovery_mesh_gloo_solo_io_v1.Mesh) {
	for _, obj := range meshes {
		if obj == nil {
			continue
		}
		b.meshes.Insert(obj)
	}
}

func (b *builder) GetDestinations() discovery_mesh_gloo_solo_io_v1_sets.DestinationSet {
	return b.destinations
}
func (b *builder) GetWorkloads() discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet {
	return b.workloads
}
func (b *builder) GetMeshes() discovery_mesh_gloo_solo_io_v1_sets.MeshSet {
	return b.meshes
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.destinations,
		b.workloads,
		b.meshes,
		b.clusters...,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.destinations,
		b.workloads,
		b.meshes,
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

	b.AddDestinations(other.GetDestinations().List()...)
	b.AddWorkloads(other.GetWorkloads().List()...)
	b.AddMeshes(other.GetMeshes().List()...)
	for _, cluster := range other.Clusters() {
		b.AddCluster(cluster)
	}
}

func (b *builder) Clone() Builder {
	if b == nil {
		return nil
	}
	clone := NewBuilder(b.ctx, b.name)

	for _, destination := range b.GetDestinations().List() {
		clone.AddDestinations(destination.DeepCopy())
	}
	for _, workload := range b.GetWorkloads().List() {
		clone.AddWorkloads(workload.DeepCopy())
	}
	for _, mesh := range b.GetMeshes().List() {
		clone.AddMeshes(mesh.DeepCopy())
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

	for _, obj := range b.GetDestinations().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Destination",
		}
		clusterSnapshots.Insert(cluster, gvk, obj)
	}
	for _, obj := range b.GetWorkloads().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Workload",
		}
		clusterSnapshots.Insert(cluster, gvk, obj)
	}
	for _, obj := range b.GetMeshes().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Mesh",
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

	for _, obj := range b.GetDestinations().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Destination",
		}
		handleObject(cluster, gvk, obj)
	}
	for _, obj := range b.GetWorkloads().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Workload",
		}
		handleObject(cluster, gvk, obj)
	}
	for _, obj := range b.GetMeshes().List() {
		cluster := obj.GetClusterName()
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Mesh",
		}
		handleObject(cluster, gvk, obj)
	}
}
