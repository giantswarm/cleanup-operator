package controller

import (
	"github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/giantswarm/cleanup-operator/pkg/project"
	"github.com/giantswarm/cleanup-operator/service/controller/key"
)

type CleanerConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type Cleaner struct {
	*controller.Controller
}

func NewCleaner(config CleanerConfig) (*Cleaner, error) {
	var err error

	resourceSets, err := newCleanerResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(v1alpha1.App)
			},
			MatchLabels: map[string]string{
				key.CleanerLabel: "true",
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/cleanup-operator-controller.
			Name: project.Name() + "-controller",
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &Cleaner{
		Controller: operatorkitController,
	}

	return c, nil
}

func newCleanerResourceSets(config CleanerConfig) ([]*controller.ResourceSet, error) {
	var err error

	var resourceSet *controller.ResourceSet
	{
		c := cleanerResourceSetConfig{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		resourceSet, err = newCleanerResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resourceSets := []*controller.ResourceSet{
		resourceSet,
	}

	return resourceSets, nil
}
