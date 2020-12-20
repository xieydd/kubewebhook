package model

import (
	admissionv1 "k8s.io/api/admission/v1"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// AdmissionReviewVersion reprensents the version of the admission review.
type AdmissionReviewVersion string

const (
	// AdmissionReviewVersionV1beta1 is the version of the v1beta1 webhooks admission review.
	AdmissionReviewVersionV1beta1 AdmissionReviewVersion = "v1beta1"

	// AdmissionReviewVersionV1 is the version of the v1 webhooks admission review.
	AdmissionReviewVersionV1 AdmissionReviewVersion = "v1"
)

// AdmissionReviewOp represents an admission review operation.
type AdmissionReviewOp string

const (
	// OperationUnknown is an unknown operation.
	OperationUnknown AdmissionReviewOp = "unknown"
	// OperationCreate is a create operation.
	OperationCreate AdmissionReviewOp = "create"
	// OperationUpdate is a update operation.
	OperationUpdate AdmissionReviewOp = "update"
	// OperationDelete is a delete operation.
	OperationDelete AdmissionReviewOp = "delete"
	// OperationConnect is a connect operation.
	OperationConnect AdmissionReviewOp = "connect"
)

// AdmissionReview represents a request admission review.
type AdmissionReview struct {
	OriginalAdmissionReview runtime.Object

	ID           string
	Name         string
	Namespace    string
	Operation    AdmissionReviewOp
	Version      AdmissionReviewVersion
	RequestGVR   *metav1.GroupVersionResource
	RequestGVK   *metav1.GroupVersionKind
	OldObjectRaw []byte
	NewObjectRaw []byte
	DryRun       bool
}

// NewAdmissionReviewV1Beta1 returns a new AdmissionReview from a admission/v1beta/admissionReview.
func NewAdmissionReviewV1Beta1(ar *admissionv1beta1.AdmissionReview) AdmissionReview {
	// Default false.
	dryRun := false
	if ar.Request.DryRun != nil {
		dryRun = *ar.Request.DryRun
	}

	return AdmissionReview{
		OriginalAdmissionReview: ar,
		ID:                      string(ar.Request.UID),
		Name:                    ar.Request.Name,
		Version:                 AdmissionReviewVersionV1beta1,
		Namespace:               ar.Request.Namespace,
		Operation:               v1Beta1OperationToModel(ar.Request.Operation),
		OldObjectRaw:            ar.Request.OldObject.Raw,
		NewObjectRaw:            ar.Request.Object.Raw,
		RequestGVR:              ar.Request.RequestResource,
		RequestGVK:              ar.Request.RequestKind,
		DryRun:                  dryRun,
	}
}

func v1Beta1OperationToModel(op admissionv1beta1.Operation) AdmissionReviewOp {
	switch op {
	case admissionv1beta1.Create:
		return OperationCreate
	case admissionv1beta1.Update:
		return OperationUpdate
	case admissionv1beta1.Delete:
		return OperationDelete
	case admissionv1beta1.Connect:
		return OperationConnect
	}

	return OperationUnknown
}

// NewAdmissionReviewV1 returns a new AdmissionReview from a admission/v1/admissionReview.
func NewAdmissionReviewV1(ar *admissionv1.AdmissionReview) AdmissionReview {
	// Default false.
	dryRun := false
	if ar.Request.DryRun != nil {
		dryRun = *ar.Request.DryRun
	}

	return AdmissionReview{
		OriginalAdmissionReview: ar,
		ID:                      string(ar.Request.UID),
		Name:                    ar.Request.Name,
		Namespace:               ar.Request.Namespace,
		Version:                 AdmissionReviewVersionV1,
		Operation:               v1OperationToModel(ar.Request.Operation),
		OldObjectRaw:            ar.Request.OldObject.Raw,
		NewObjectRaw:            ar.Request.Object.Raw,
		DryRun:                  dryRun,
	}
}

func v1OperationToModel(op admissionv1.Operation) AdmissionReviewOp {
	switch op {
	case admissionv1.Create:
		return OperationCreate
	case admissionv1.Update:
		return OperationUpdate
	case admissionv1.Delete:
		return OperationDelete
	case admissionv1.Connect:
		return OperationConnect
	}

	return OperationUnknown
}
