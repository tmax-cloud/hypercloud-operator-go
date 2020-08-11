/*


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

package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	claim "tmax.io/apis/claim/v1alpha1"
)

// NamespaceClaimReconciler reconciles a NamespaceClaim object
type NamespaceClaimReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=claim.tmax.io,resources=namespaceclaims,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=claim.tmax.io,resources=namespaceclaims/status,verbs=get;update;patch

func (r *NamespaceClaimReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	reqLogger := r.Log.WithValues("namespaceclaim", req.NamespacedName)
	reqLogger.Info("Reconciling NamespaceClaim")

	namespaceClaim := &claim.NamespaceClaim{}

	if err := r.Get(context.TODO(), req.NamespacedName, namespaceClaim); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("NamespaceClaim resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}

		reqLogger.Error(err, "Failed to get NamespaceClaim")
		return ctrl.Result{}, err
	}

	found := &corev1.Namespace{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: namespaceClaim.Name}, found)

	reqLogger.Info("nsc status:" + string(namespaceClaim.Status.Status))
	if err != nil && !errors.IsNotFound(err) {
		reqLogger.Error(err, "failed to get namespace info")
		return ctrl.Result{}, err
	}

	switch namespaceClaim.Status.Status {
	case "":
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Not found.")
			namespaceClaim.Status.Status = claim.NamespaceClaimStatusTypeAwaiting
		} else {
			namespaceClaim.Status.Status = claim.NamespaceClaimStatusTypeReject
			namespaceClaim.Status.Reason = "Already exist namespace."
		}

	case claim.NamespaceClaimStatusTypeSuccess:
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Create namespace.")

			namespace := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespaceClaim.Name,
				},
			}
			if err := r.Create(context.TODO(), namespace); err != nil {
				reqLogger.Error(err, "Failed to create namespace.")
				namespaceClaim.Status.Status = claim.NamespaceClaimStatueTypeError
				namespaceClaim.Status.Reason = "Failed to create namespace"
				namespaceClaim.Status.Message = err.Error()
			}
		}
	}

	namespaceClaim.Status.LastTransitionTime = metav1.Now()
	if err := r.Update(context.TODO(), namespaceClaim); err != nil {
		reqLogger.Error(err, "Failed to update namespaceclaim status.")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *NamespaceClaimReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&claim.NamespaceClaim{}).
		Complete(r)
}
