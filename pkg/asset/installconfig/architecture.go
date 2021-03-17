package installconfig

import (
	"sort"
	"unsafe"

	"github.com/pkg/errors"
	survey "gopkg.in/AlecAivazis/survey.v1"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/baremetal"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/none"
)

type architecture struct {
	Architecture string
}

var _ asset.Asset = (*architecture)(nil)

// Dependencies returns no dependencies.
func (a *architecture) Dependencies() []asset.Asset {
	return []asset.Asset{
		&platform{},
	}
}

// Generate queries for the base domain from the user.
func (a *architecture) Generate(parents asset.Parents) error {
	platform := &platform{}
	parents.Get(platform)

	var archlist []types.Architecture
	switch platform.CurrentName() {
	case aws.Name:
		archlist = types.ArchitectureListAWS
	case libvirt.Name, none.Name, baremetal.Name:
		archlist = types.ArchitectureList
	default:
		archlist = []types.Architecture{types.ArchitectureAMD64}
	}

	if err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Architecture",
				Help:    "The architecture of the control plane and the compute nodes. Heterogenous clusters are not supported",
				Default: types.ArchitectureAMD64,
				Options: *(*[]string)(unsafe.Pointer(&archlist)),
			},
			Validate: survey.ComposeValidators(survey.Required, func(ans interface{}) error {
				choice := ans.(string)
				list := *(*[]string)(unsafe.Pointer(&archlist))
				i := sort.SearchStrings(list, choice)
				if i == len(list) || list[i] != choice {
					return errors.Errorf("invalid architecture %q", choice)
				}
				return nil
			}),
		},
	}, &a.Architecture); err != nil {
		return errors.Wrap(err, "failed UserInput")
	}
	return nil
}

// Name returns the human-friendly name of the asset.
func (a *architecture) Name() string {
	return "Architecture"
}
