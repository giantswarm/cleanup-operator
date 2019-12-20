package key

import "github.com/giantswarm/microerror"

var invalidArgumentError = &microerror.Error{
	Kind: "invalidArgument",
}

// IsInvalidArgument asserts invalidArgumentError which in turn means a
// function was given an invalid argument (like a nil interface).
func IsInvalidArgument(err error) bool {
	return microerror.Cause(err) == invalidArgumentError
}

var wrongTypeError = &microerror.Error{
	Kind: "wrongTypeError",
}

// IsWrongTypeError asserts wrongTypeError which in turn means a function was
// given a wrapped type it did not expect.
func IsWrongTypeError(err error) bool {
	return microerror.Cause(err) == wrongTypeError
}
