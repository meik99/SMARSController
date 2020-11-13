package authserver

import (
	coffeev1 "github.com/meik99/CoffeeToGO/CoffeeOperator/api/v1"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	Name     = "coffee-auth-server"
	Replicas = 1
	AppName  = "coffeetogo"
	Image    = "meik99/coffee-auth-server:latest"

	InitialDelaySeconds = 5
	PeriodSeconds       = 5
	FailureThreshold    = 5

	HealthPath = "/apps/coffeetogo/api/v1/sso/google/health"
	HealthPort = 8080
)

func BuildAuthServerDeployment() *v1.Deployment {
	return buildDeployment()
}

func buildDeployment() *v1.Deployment {
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

func buildDeploymentLabels() map[string]string {
	return map[string]string{
		"app":  AppName,
		"type": Name,
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
		Name:           AppName,
		Image:          Image,
		LivenessProbe:  buildProbe(),
		ReadinessProbe: buildProbe(),
	}
}

func buildProbe() *v12.Probe {
	return &v12.Probe{
		Handler:             buildProbeHandler(),
		InitialDelaySeconds: InitialDelaySeconds,
		PeriodSeconds:       PeriodSeconds,
		FailureThreshold:    FailureThreshold,
	}
}

func buildProbeHandler() v12.Handler {
	return v12.Handler{
		HTTPGet: buildProbeHandlerAction(),
	}
}

func buildProbeHandlerAction() *v12.HTTPGetAction {
	return &v12.HTTPGetAction{
		Path: HealthPath,
		Port: intstr.IntOrString{IntVal: HealthPort},
	}
}
