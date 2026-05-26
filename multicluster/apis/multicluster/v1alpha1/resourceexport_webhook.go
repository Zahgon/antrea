/*
Copyright 2021 Antrea Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func (r *ResourceExport) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:webhook:path=/mutate-multicluster-crd-antrea-io-v1alpha1-resourceexport,mutating=true,failurePolicy=fail,sideEffects=None,groups=multicluster.crd.antrea.io,resources=resourceexports,verbs=create;update,versions=v1alpha1,name=mresourceexport.kb.io,admissionReviewVersions={v1,v1beta1}

type ResourceExportCustomDefaulter struct{}

var _ admission.Defaulter[*ResourceExport] = &ResourceExportCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind ResourceExport.
func (d *ResourceExportCustomDefaulter) Default(_ context.Context, r *ResourceExport) error {
	_ = "STUB: not implemented"
	return nil
}

// Set default values

func (d *ResourceExportCustomDefaulter) applyDefaults(r *ResourceExport) {
	_ = "STUB: not implemented"
	return
}

// Only mutate ResourceExport created for ClusterNetworkPolicy resources

// Add domain qualified finalizer for ResourceExports to avoid Kubernetes from reporting errors:
//  "prefer a domain-qualified finalizer name to avoid accidental conflicts with other finalizer writers"
// https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#finalizers
// Note for ResourceExports created before Antrea v2.2, LegacyResourceExportFinalizer may still be present
// and needs to be removed before a ResourceExport is deleted.
