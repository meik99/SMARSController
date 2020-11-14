package couchdb

import (
	"context"
	"github.com/go-logr/logr"
	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconciler struct {
	client.Client
	Log logr.Logger
}

func NewCouchDbReconciler(clt client.Client, log logr.Logger) *Reconciler {
	return &Reconciler{
		Client: clt,
		Log:    log,
	}
}

func (r *Reconciler) CreateCouchDbDeploymentIfNotExists() error {
	var deployment v1.Deployment
	err := r.Get(context.TODO(), client.ObjectKey{Name: Name, Namespace: coffeev1.Namespace}, &deployment)
	if err != nil {
		if errors.IsNotFound(err) {
			r.Log.Info("could not find deployment of CouchDB, creating...")
			return r.createCouchDbDeployment()
		} else {
			r.Log.Error(err, err.Error())
			return err
		}
	}
	return nil
}

func (r *Reconciler) createCouchDbDeployment() error {
	err := r.Create(context.TODO(), BuildDeployment())
	if err != nil {
		r.Log.Error(err, err.Error())
		return err
	}
	return nil
}
