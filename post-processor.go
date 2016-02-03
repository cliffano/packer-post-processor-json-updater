package main

import (
	"regexp"

	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

type Config struct {
	AmiId map[string][]string `mapstructure:"ami_id"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

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

	if p.config.AmiId == nil {
		p.config.AmiId = make(map[string][]string, 0)
	}

	return nil
}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {

	if p.config.AmiId != nil {
		r, _ := regexp.Compile("ami-[a-z0-9]+")
		amiId := r.FindString(artifact.Id())

		for file, properties := range p.config.AmiId {
			err := UpdateJsonFile(file, properties, amiId, ui)
			if err != nil {
				return artifact, false, err
			}
		}
	}

	return artifact, true, nil
}
