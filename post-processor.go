package main

import (
	"fmt"
	"regexp"

	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

// Config is the post-processor configuration with interpolation supported.
// Currently only supports AWS AMI ID.
type Config struct {
	AmiID map[string][]string `mapstructure:"ami_id"`

	ctx interpolate.Context
}

// PostProcessor holds the Config object.
type PostProcessor struct {
	config Config
}

// Configure sets the Config object with configuration values from the Packer
// template.
func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate: true,
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
func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {

	ui.Say(fmt.Sprintf("[DEBUG] Artifact ID: %s", artifact.Id()))
	ui.Say(fmt.Sprintf("[DEBUG] Artifact String: %s", artifact.String()))

	if p.config.AmiID != nil {
		r, _ := regexp.Compile("ami-[a-z0-9]+")
		amiID := r.FindString(artifact.Id())

		for file, properties := range p.config.AmiID {
			err := UpdateJSONFile(file, properties, amiID, ui)
			if err != nil {
				return artifact, false, err
			}
		}
	}

	return artifact, true, nil
}
