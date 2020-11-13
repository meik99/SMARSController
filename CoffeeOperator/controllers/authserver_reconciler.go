package controllers

import (
	"context"
	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
	"github.com/meik99/CoffeeToGO/CoffeeOperator/authserver"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *AuthServerReconciler) ReconcileAuthServer() (ctrl.Result, error) {
	var instance coffeev1.AuthServer
	err := r.Get(context.TODO(), client.ObjectKey{Name: InstanceName, Namespace: coffeev1.Namespace}, &instance)
	if err != nil {
		return RequeueAfterFiveMinutes(), err
	}

	err = r.createAuthServerDeploymentIfNotExists()
	return RequeueAfterFiveMinutes(), err
}

func (r *AuthServerReconciler) createAuthServerDeploymentIfNotExists() error {
	var authServerDeployment v1.Deployment
	err := r.Get(context.TODO(), client.ObjectKey{Name: authserver.Name, Namespace: coffeev1.Namespace}, &authServerDeployment)
	if err != nil {
		if errors.IsNotFound(err) {
			r.Log.Info("could not find deployment auf AuthServer, creating")
			return r.createAuthServerDeployment()
		} else {
			r.Log.Error(err, err.Error())
			return err
		}
	}
	return nil
}

func (r *AuthServerReconciler) createAuthServerDeployment() error {
	err := r.Create(context.TODO(), authserver.BuildAuthServerDeployment())
	if err != nil {
		r.Log.Error(err, err.Error())
		return err
	}
	return nil
}
