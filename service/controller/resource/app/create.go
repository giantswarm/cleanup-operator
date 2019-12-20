package app

import (
	"context"
	"fmt"
	"time"

	"github.com/giantswarm/cleanup-operator/service/controller/key"
	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	app, err := key.ToApp(obj)
	if err != nil {
		return microerror.Mask(err)
	}
	logger := r.logger.With("namespace", app.Namespace, "app", app.Name)

	logger.LogCtx(ctx, "level", "debug", "message", "checking app")
	ttl := key.TTL(app.GetObjectMeta())
	expires := app.GetCreationTimestamp().Add(ttl)
	if time.Now().After(expires) {
		r.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf(
			"app has TTL=%s and has expired on %s, removing it",
			ttl, expires,
		))
		err := r.k8sclient.G8sClient().ApplicationV1alpha1().Apps(app.Namespace).
			Delete(app.Name, &metav1.DeleteOptions{})
		if apierrors.IsNotFound(err) {
			r.logger.LogCtx(ctx, "level", "debug", "message", "app was already gone")
		} else if err != nil {
			return microerror.Mask(err)
		} else {
			r.logger.LogCtx(ctx, "level", "debug", "message", "deleted app")
		}
	}

	return nil
}
