package couchdb

import (
	"context"
	"fmt"
	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	SecretName = "coffee-couchdb-secret"
	DbUser     = "DB_USER"
	DbPassword = "DB_PASSWORD"
)

func (r *Reconciler) CheckIfSecretExists() error {
	_, err := r.getSecret()
	return err
}

func (r *Reconciler) IsSecretValid() error {
	secret, err := r.getSecret()
	if err != nil {
		return err
	}

	if _, hasKey := secret.Data[DbUser]; !hasKey {
		return fmt.Errorf("'%s' is missing from '%s'", DbUser, SecretName)
	}
	if _, hasKey := secret.Data[DbPassword]; !hasKey {
		return fmt.Errorf("'%s' is missing from '%s'", DbPassword, SecretName)
	}

	return nil
}

func (r *Reconciler) getSecret() (*v1.Secret, error) {
	var dbSecret v1.Secret
	err := r.Get(context.TODO(), client.ObjectKey{Namespace: coffeev1.Namespace, Name: SecretName}, &dbSecret)
	return &dbSecret, err
}
