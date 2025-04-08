package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"

	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	lifecyclev1alpha1 "github.com/tldr-it-stepankutaj/lifecycle-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type WatchedOperatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *WatchedOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var watched lifecyclev1alpha1.WatchedOperator
	if err := r.Get(ctx, req.NamespacedName, &watched); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var sub operatorsv1alpha1.Subscription
	if err := r.Get(ctx, types.NamespacedName{
		Name:      watched.Spec.Name,
		Namespace: watched.Spec.Namespace,
	}, &sub); err != nil {
		logger.Error(err, "unable to get Subscription")
		return ctrl.Result{}, err
	}

	latestCSV, err := GetLatestCSVFromCatalog(ctx, r.Client, watched.Spec.Name, watched.Spec.Channel, "openshift-marketplace")
	if err != nil {
		logger.Error(err, "unable to get latest CSV from PackageManifest")
		return ctrl.Result{}, err
	}

	currentCSV := sub.Status.InstalledCSV
	upgradeReady := currentCSV != latestCSV

	if upgradeReady && watched.Spec.AutoUpgrade {
		sub.Spec.StartingCSV = latestCSV
		if err := r.Update(ctx, &sub); err != nil {
			logger.Error(err, "failed auto-upgrade")
			return ctrl.Result{}, err
		}
	}

	if upgradeReady && watched.Spec.WebhookURL != "" {
		_ = SendWebhookNotification(ctx, watched.Spec.WebhookURL, watched.Spec.Name, currentCSV, latestCSV, "New version available")
	}

	watched.Status.CurrentVersion = currentCSV
	watched.Status.LatestVersion = latestCSV
	watched.Status.UpgradeReady = upgradeReady

	if err := r.Status().Update(ctx, &watched); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *WatchedOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&lifecyclev1alpha1.WatchedOperator{}).
		Complete(r)
}
