package json_updater

import (
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/template/interpolate"
)

// Config is the post-processor configuration with interpolation supported.
// Currently only supports AWS AMI ID.
type Config struct {
	AmiID map[string][]string `mapstructure:"ami_id"`
  common.PackerConfig `mapstructure:",squash"`

	ctx interpolate.Context
}
