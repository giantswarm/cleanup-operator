package controller

import (
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"github.com/giantswarm/operatorkit/resource"
	"github.com/giantswarm/operatorkit/resource/wrapper/metricsresource"
	"github.com/giantswarm/operatorkit/resource/wrapper/retryresource"

	"github.com/giantswarm/cleanup-operator/service/controller/resource/app"
)

type cleanerResourceSetConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

func newCleanerResourceSet(config cleanerResourceSetConfig) (*controller.ResourceSet, error) {
	var err error

	var appResource resource.Interface
	{
		c := app.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		appResource, err = app.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resources := []resource.Interface{
		appResource,
	}

	{
		c := retryresource.WrapConfig{
			Logger: config.Logger,
		}

		resources, err = retryresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	{
		c := metricsresource.WrapConfig{}

		resources, err = metricsresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	// handlesFunc defines which objects you want to get into your controller, e.g. which objects you want to watch.
	handlesFunc := func(obj interface{}) bool {
		// TODO: By default this will handle all objects of the type your controller is watching.
		// Your controller is watching a certain kubernetes type, so why do we need to check again?
		// Because there might be a change in the object structure - e.g. the type `AWSConfig` object might have the field
		// availabilityZones recently, but older ones don't, and you don't want to handle those.
		//
		// Normally we use this to filter objects containing the expected `versionbundle` version.
		// So two versions of your operator don't accidentally reconcile the same CR.
		return true
	}

	var resourceSet *controller.ResourceSet
	{
		c := controller.ResourceSetConfig{
			Handles:   handlesFunc,
			Logger:    config.Logger,
			Resources: resources,
		}

		resourceSet, err = controller.NewResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return resourceSet, nil
}
