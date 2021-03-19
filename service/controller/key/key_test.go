package key_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/cleanup-operator/service/controller/key"
)

func Test_TTL(t *testing.T) {
	testCases := []struct {
		name     string
		input    key.LabelsGetter
		expected time.Duration
	}{
		{
			name: "when a label is specified it converts its value to Duration",
			input: &v1alpha1.App{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						key.TTLLabel: "42s",
					},
				},
			},
			expected: 42000000000,
		},
		{
			name:     "when a label is not specified it returns a default value",
			input:    &v1alpha1.App{},
			expected: 28800000000000,
		},
		{
			name: "when a negative value is specified it returns the default",
			input: &v1alpha1.App{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						key.TTLLabel: "-500",
					},
				},
			},
			expected: 28800000000000,
		},
		{
			name: "when an invalid value is specified it returns the default",
			input: &v1alpha1.App{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						key.TTLLabel: "an hour and two minutes",
					},
				},
			},
			expected: 28800000000000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := key.TTL(tc.input); result != tc.expected {
				t.Fatalf("got: %s, expected: %s", result, tc.expected)
			}
		})
	}
}

func Test_ToApp(t *testing.T) {
	var nilApp *v1alpha1.App

	testCases := []struct {
		name         string
		input        interface{}
		expected     v1alpha1.App
		errorMatcher func(error) bool
	}{
		{
			name: "positive match",
			input: &v1alpha1.App{
				Spec: v1alpha1.AppSpec{
					Name:      "app-name",
					Namespace: "app-namespace",
					Version:   "1.1.1",
				},
			},
			expected: v1alpha1.App{
				Spec: v1alpha1.AppSpec{
					Name:      "app-name",
					Namespace: "app-namespace",
					Version:   "1.1.1",
				},
			},
		},
		{
			name:         "nil interface",
			input:        nil,
			errorMatcher: key.IsInvalidArgument,
		},
		{
			name:         "incorrect type",
			input:        &v1alpha1.AppCatalog{},
			errorMatcher: key.IsWrongTypeError,
		},
		{
			name:         "incorrect value",
			input:        nilApp,
			errorMatcher: key.IsInvalidArgument,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := key.ToApp(tc.input)
			switch {
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("got unexpected error %#v", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("expected an error, got nil")
			case err != nil && !tc.errorMatcher(err):
				t.Fatalf("error %#v does not match one expected by the matcher", err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("got: %#v, expected: %#v", result, tc.expected)
			}
		})
	}
}
