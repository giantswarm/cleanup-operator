package key

import (
	"github.com/giantswarm/cleanup-operator/pkg/project"
)

var (
	// CleanerLabel is the label this operator looks for on resources when
	// deciding whether to manage them. In other words, when this label is set
	// on an object it tells the cleanup-operator to watch that object and
	// dispose of it when it outlives its usefulness.
	CleanerLabel = project.Name() + ".giantswarm.io/enabled"
)
