package app

import (
	"context"
	"fmt"
	"time"

	"github.com/giantswarm/cleanup-operator/service/controller/key"
	"github.com/giantswarm/microerror"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	app, err := key.ToApp(obj)
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("checking app %#q in namespace %#q", app.Name, app.Namespace))
	ttl := key.TTL(app.GetObjectMeta())
	expires := app.GetCreationTimestamp().Add(ttl)
	if time.Now().After(expires) {
		r.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf(
			"app %#q in namespace %#q has TTL=%s and has expired on %s, removing it",
			app.Name, app.Namespace, ttl, expires,
		))
	}

	return nil
}
