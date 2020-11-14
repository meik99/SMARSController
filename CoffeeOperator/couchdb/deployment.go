package couchdb

import (
	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Name             = "coffee-couchdb"
	Replicas         = 1
	Image            = "couchdb:latest"
	SecretVolumeName = "coffee-couchdb-secret-volume"
)

func BuildDeployment() *v1.Deployment {
	return &v1.Deployment{
		ObjectMeta: buildDeploymentObjectMeta(),
		Spec:       buildDeploymentSpec(),
	}
}

func buildDeploymentObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      Name,
		Namespace: coffeev1.Namespace,
		Labels:    buildDeploymentLabels(),
	}
}

func buildDeploymentSpec() v1.DeploymentSpec {
	var replicas int32 = Replicas

	return v1.DeploymentSpec{
		Replicas: &replicas,
		Selector: buildDeploymentSelector(),
		Template: buildDeploymentTemplate(),
	}
}

func buildDeploymentLabels() map[string]string {
	return map[string]string{
		"app":  coffeev1.AppName,
		"type": Name,
	}
}

func buildDeploymentSelector() *metav1.LabelSelector {
	return &metav1.LabelSelector{
		MatchLabels: buildDeploymentLabels(),
	}
}

func buildDeploymentTemplate() v12.PodTemplateSpec {
	return v12.PodTemplateSpec{
		ObjectMeta: buildDeploymentTemplateObjectMeta(),
		Spec:       buildDeploymentPodSpec(),
	}
}

func buildDeploymentTemplateObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Labels: buildDeploymentLabels(),
	}
}

func buildDeploymentPodSpec() v12.PodSpec {
	return v12.PodSpec{
		Containers: []v12.Container{
			buildContainer(),
		},
	}
}

func buildContainer() v12.Container {
	return v12.Container{
		Name:  Name,
		Image: Image,
		EnvFrom: []v12.EnvFromSource{
			buildEnvsFromSecret(),
		},
	}
}

func buildEnvsFromSecret() v12.EnvFromSource {
	return v12.EnvFromSource{
		SecretRef: &v12.SecretEnvSource{
			LocalObjectReference: v12.LocalObjectReference{
				Name: SecretName,
			},
		},
	}
}
