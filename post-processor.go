package main

import (
	// "regexp"
	// "encoding/json"
	// "fmt"
	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
	"strings"
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

func makeAmiList(a packer.Artifact) (map[string]string, error) {
	type Amis map[string]string
	amis := make(Amis)

	// converting from Artifact interface via Id()
	regionAmiList := strings.Split(a.Id(), ",")

	for _, regionAmi := range regionAmiList {
		// region:ami-id
		amiInfo := strings.Split(regionAmi, ":")
		amis[amiInfo[0]] = amiInfo[1]
	}
	return amis, nil
}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {

	if p.config.AmiId != nil {
		for file, properties := range p.config.AmiId {
			amiList, err := makeAmiList(artifact)
			if err != nil {
				return artifact, false, err
			}

			err = UpdateJsonFile(file, properties, amiList, ui, false)
			if err != nil {
				return artifact, false, err
			}
		}
	}

	return artifact, true, nil
}
