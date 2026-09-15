package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	platformv1alpha1 "github.com/davidvieira51756/cloudnative-go-lab/api/v1alpha1"
)

type AppDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *AppDeploymentReconciler) Reconcile(
	ctx context.Context,
	req ctrl.Request,
) (ctrl.Result, error) {
	var app platformv1alpha1.AppDeployment

	if err := r.Get(ctx, req.NamespacedName, &app); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	var deployment appsv1.Deployment

	err := r.Get(
		ctx,
		types.NamespacedName{
			Name:      app.Name,
			Namespace: app.Namespace,
		},
		&deployment,
	)

	if apierrors.IsNotFound(err) {
		deployment = buildDeployment(&app)

		if err := ctrl.SetControllerReference(
			&app,
			&deployment,
			r.Scheme,
		); err != nil {
			return ctrl.Result{}, err
		}

		if err := r.Create(ctx, &deployment); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	updated := false

	if deployment.Spec.Replicas == nil ||
		*deployment.Spec.Replicas != app.Spec.Replicas {

		replicas := app.Spec.Replicas
		deployment.Spec.Replicas = &replicas
		updated = true
	}

	container := &deployment.Spec.Template.Spec.Containers[0]

	if container.Image != app.Spec.Image {
		container.Image = app.Spec.Image
		updated = true
	}

	if len(container.Ports) == 0 ||
		container.Ports[0].ContainerPort != app.Spec.Port {

		container.Ports = []corev1.ContainerPort{
			{
				ContainerPort: app.Spec.Port,
			},
		}

		updated = true
	}

	if updated {
		if err := r.Update(ctx, &deployment); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func buildDeployment(
	app *platformv1alpha1.AppDeployment,
) appsv1.Deployment {

	labels := map[string]string{
		"app": app.Name,
	}

	replicas := app.Spec.Replicas

	return appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
		},

		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,

			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},

				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "app",
							Image:           app.Spec.Image,
							ImagePullPolicy: corev1.PullIfNotPresent,

							Ports: []corev1.ContainerPort{
								{
									ContainerPort: app.Spec.Port,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *AppDeploymentReconciler) SetupWithManager(
	mgr ctrl.Manager,
) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&platformv1alpha1.AppDeployment{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
