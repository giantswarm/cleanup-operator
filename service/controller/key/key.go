package key

import (
	"time"

	"github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/cleanup-operator/pkg/project"
)

var (
	// CleanerLabel is the label this operator looks for on resources when
	// deciding whether to manage them. In other words, when this label is set
	// on an object it tells the cleanup-operator to watch that object and
	// dispose of it when it outlives its usefulness.
	CleanerLabel = project.Name() + ".giantswarm.io/enabled"

	// TTLLabel is the label that specifies how long a specific object should
	// live as a time.Duration. In other words the object with this label will
	// expire at `Status.CreateTime` + value of `TTLLabel`.
	TTLLabel = project.Name() + ".giantswarm.io/ttl"
)

// ToApp extracts the v1alpha1.App value from the interface wrapper and returns
// it, or returns an error if the interface contains an incorrect type or a nil
// value.
func ToApp(v interface{}) (v1alpha1.App, error) {
	if v == nil {
		return v1alpha1.App{}, microerror.Maskf(invalidArgumentError, "expected '%T', got '%T'", &v1alpha1.App{}, v)
	}

	app, ok := v.(*v1alpha1.App)
	if !ok {
		return v1alpha1.App{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.App{}, v)
	}
	if app == nil {
		return v1alpha1.App{}, microerror.Maskf(invalidArgumentError, "given '%T' cannot be nil", &v1alpha1.App{})
	}

	appCp := app.DeepCopy()

	return *appCp, nil
}

// TTL extracts the value of TTLLabel from an object, parses it and returns a
// time.Duration representing number of seconds an object is set to live. If
// this label is not specified on an object, it will default to an equivalent
// of 8 hours.
func TTL(obj LabelsGetter) time.Duration {
	t, err := time.ParseDuration(obj.GetLabels()[TTLLabel])
	if err != nil {
		return 8 * time.Hour
	}
	return t
}
