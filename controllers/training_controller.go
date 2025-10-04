package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	trainingv1 "github.com/tektutor/training-operator/api/v1"
)

// TrainingReconciler reconciles a Training object
type TrainingReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=training.tektutor.org,resources=trainings,verbs=get;list;watch;create;update;patch;delete

func (r *TrainingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var training trainingv1.Training
	if err := r.Get(ctx, req.NamespacedName, &training); err != nil {
		if errors.IsNotFound(err) {
			// Resource deleted
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Extract info from Training spec
	topic := training.Spec.Topic
	fromDate := training.Spec.FromDate
	toDate := training.Spec.ToDate
	duration := training.Spec.Duration
	city := training.Spec.City

	// Print the training calendar entry
	fmt.Printf(
		"Adding to Training Calendar:\nTopic: %s\nFrom: %s\nTo: %s\nDuration: %s\nCity: %s\n\n",
		topic, fromDate, toDate, duration, city,
	)

	return ctrl.Result{}, nil
}

func (r *TrainingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&trainingv1.Training{}).
		Complete(r)
}
