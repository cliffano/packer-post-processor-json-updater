package json_updater

import (
	"context"
	"fmt"
	"regexp"

  "github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

// PostProcessor holds the Config object.
type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

// Configure sets the Config object with configuration values from the Packer
// template.
func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)

	if err != nil {
		return err
	}

	if p.config.AmiID == nil {
		p.config.AmiID = make(map[string][]string, 0)
	}

	return nil
}

// PostProcess parses the AMI ID from the artifact ID, and then passes the AMI ID
// to UpdateJSONFile to be set as the new value of the JSON paths properties in
// Packer template.
// AWS artifact ID output has the format of <region>:<ami_id>,
// for example: ap-southeast-2:ami-4f8fae2c
func (p *PostProcessor) PostProcess(ctx context.Context, ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, bool, error) {

	ui.Say(fmt.Sprintf("%s", artifact.String()))

	if p.config.AmiID != nil {
		r, _ := regexp.Compile("ami-[a-z0-9]+")
		amiID := r.FindString(artifact.Id())

		for file, properties := range p.config.AmiID {
			err := EnsureJSONFileExists(file)
			if err != nil {
				return artifact, false, false, err
			}
			err = UpdateJSONFile(file, properties, amiID, ui)
			if err != nil {
				return artifact, false, false, err
			}
		}
	}

	return artifact, true, false, nil
}
