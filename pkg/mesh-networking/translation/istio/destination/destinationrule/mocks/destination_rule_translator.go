// Code generated by MockGen. DO NOT EDIT.
// Source: ./destination_rule_translator.go

// Package mock_destinationrule is a generated GoMock package.
package mock_destinationrule

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockTranslator) Translate(ctx context.Context, in input.LocalSnapshot, destination *v1.Destination, sourceMeshInstallation *v1.MeshSpec_MeshInstallation, reporter reporting.Reporter) *v1alpha3.DestinationRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", ctx, in, destination, sourceMeshInstallation, reporter)
	ret0, _ := ret[0].(*v1alpha3.DestinationRule)
	return ret0
}

// Translate indicates an expected call of Translate
func (mr *MockTranslatorMockRecorder) Translate(ctx, in, destination, sourceMeshInstallation, reporter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), ctx, in, destination, sourceMeshInstallation, reporter)
}

// ShouldTranslate mocks base method
func (m *MockTranslator) ShouldTranslate(destination *v1.Destination, localEventObjs map[schema.GroupVersionKind][]ezkube.ResourceId, remoteEventObjs map[schema.GroupVersionKind][]ezkube.ClusterResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldTranslate", destination, localEventObjs, remoteEventObjs)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldTranslate indicates an expected call of ShouldTranslate
func (mr *MockTranslatorMockRecorder) ShouldTranslate(destination, localEventObjs, remoteEventObjs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldTranslate", reflect.TypeOf((*MockTranslator)(nil).ShouldTranslate), destination, localEventObjs, remoteEventObjs)
}
