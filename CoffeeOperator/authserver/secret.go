package authserver

import (
	"context"
	"fmt"
	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	SecretName = "coffee-auth-server-secret"
)

var secretAttributes = []string{
	"COFFEE_AUTH_HOST", "COFFEE_AUTH_PORT", "COFFEE_AUTH_PROTOCOL", "COFFEE_AUTH_PATH",
}

func (r *Reconciler) CheckIfSecretExists() error {
	_, err := r.getSecret()
	return err
}

func (r *Reconciler) IsSecretValid() error {
	secret, err := r.getSecret()
	if err != nil {
		return err
	}

	return checkAttributes(secret)
}

func checkAttributes(secret *v1.Secret) error {
	for _, attribute := range secretAttributes {
		if _, hasAttribute := secret.Data[attribute]; !hasAttribute {
			return fmt.Errorf("'%s' is missing from '%s'", attribute, SecretName)
		}
	}
	return nil
}

func (r *Reconciler) getSecret() (*v1.Secret, error) {
	var authServerSecret v1.Secret
	err := r.Get(context.TODO(), client.ObjectKey{Namespace: coffeev1.Namespace, Name: SecretName}, &authServerSecret)
	return &authServerSecret, err
}
