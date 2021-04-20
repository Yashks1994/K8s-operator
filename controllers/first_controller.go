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
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	crdsv1alpha1 "first-crd/api/v1alpha1"
)

// FirstReconciler reconciles a First object
type FirstReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=crds.my-domain.io,resources=firsts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=crds.my-domain.io,resources=firsts/status,verbs=get;update;patch

func (r *FirstReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("first", req.NamespacedName)

	log.Info("+++++++++ Started Reconciling  ++++++++++")
	defer log.Info("--------   Reconciling Completed ----------")

	instance := &crdsv1alpha1.First{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, could have been deleted after
			// reconcile request
			return ctrl.Result{}, nil
		}

		// error handling the object, requeue the request
		return ctrl.Result{}, err
	}

	if instance.Status.Phase == "" {
		instance.Status.Phase = crdsv1alpha1.PhasePending
	}

	pod := crdsv1alpha1.NewPod(instance)
	fmt.Println(pod)
	err = ctrl.SetControllerReference(instance, pod, r.Scheme)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.Create(context.TODO(), pod)
	if err != nil {
		return ctrl.Result{}, nil
	}

	log.Info("We will create the deployment now")
	deploy := crdsv1alpha1.CreateDeployment(instance)
	fmt.Println(deploy)
	if err != nil {
		return ctrl.Result{}, err
	}
	err = r.Create(context.TODO(), deploy)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *FirstReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crdsv1alpha1.First{}).
		Complete(r)
}
