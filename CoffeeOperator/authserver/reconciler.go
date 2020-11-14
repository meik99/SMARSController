package authserver

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconciler struct {
	client.Client
}

func NewReconciler(clt client.Client) *Reconciler {
	return &Reconciler{
		clt,
	}
}
