package project

import (
	"github.com/giantswarm/versionbundle"
)

func NewVersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "cleanup-operator",
				Description: "An operator that disposes of any evidence of test stuff in your cluster",
				Kind:        versionbundle.KindChanged,
			},
		},
		Components: []versionbundle.Component{},
		Name:       "cleanup-operator",
		Version:    BundleVersion(),
	}
}
