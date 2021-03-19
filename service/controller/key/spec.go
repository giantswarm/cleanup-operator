package key

// LabelsGetter is an object that allows access to its labels via GetLabels
// method.
type LabelsGetter interface {
	GetLabels() map[string]string
}
