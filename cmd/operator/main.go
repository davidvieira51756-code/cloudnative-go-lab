package main

import (
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	platformv1alpha1 "github.com/davidvieira51756/cloudnative-go-lab/api/v1alpha1"
	"github.com/davidvieira51756/cloudnative-go-lab/internal/controller"
)

func main() {
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	scheme := runtime.NewScheme()

	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		ctrl.Log.Error(err, "failed to add Kubernetes types to scheme")
		os.Exit(1)
	}

	if err := platformv1alpha1.AddToScheme(scheme); err != nil {
		ctrl.Log.Error(err, "failed to add AppDeployment to scheme")
		os.Exit(1)
	}

	manager, err := ctrl.NewManager(
		ctrl.GetConfigOrDie(),
		ctrl.Options{
			Scheme: scheme,
		},
	)

	if err != nil {
		ctrl.Log.Error(err, "failed to create manager")
		os.Exit(1)
	}

	reconciler := &controller.AppDeploymentReconciler{
		Client: manager.GetClient(),
		Scheme: manager.GetScheme(),
	}

	if err := reconciler.SetupWithManager(manager); err != nil {
		ctrl.Log.Error(err, "failed to setup controller")
		os.Exit(1)
	}

	ctrl.Log.Info("starting AppDeployment operator")

	if err := manager.Start(ctrl.SetupSignalHandler()); err != nil {
		ctrl.Log.Error(err, "operator stopped with error")
		os.Exit(1)
	}
}
