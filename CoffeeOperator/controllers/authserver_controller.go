/*
Copyright 2020.

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
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"

	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
)

const (
	InstanceName = "coffee-authserver"
)

// AuthServerReconciler reconciles a AuthServer object
type AuthServerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=coffeetogo.rynkbit.com,resources=authservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coffeetogo.rynkbit.com,resources=authservers/status,verbs=get;update;patch

func (r *AuthServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("authserver", req.NamespacedName)

	// your logic here
	return r.ReconcileAuthServer()
}

func RequeueAfterFiveMinutes() ctrl.Result {
	return ctrl.Result{
		RequeueAfter: 5 * time.Minute,
	}
}

func (r *AuthServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&coffeev1.AuthServer{}).
		Complete(r)
}

func (r *AuthServerReconciler) logError(err error) (ctrl.Result, error) {
	r.Log.Error(err, err.Error())
	return RequeueAfterFiveMinutes(), err
}
