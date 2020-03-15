package json_updater

import (
	"github.com/hashicorp/hcl/v2/hcldec"
  "github.com/zclconf/go-cty/cty"
)

// FlatConfig is a flat version of Config.
type FlatConfig struct {
	PackerBuildName     *string             `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string             `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool               `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool               `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string             `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string   `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string            `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	AmiId               map[string][]string `mapstructure:"ami_id" cty:"ami_id"`
}

// FlatMapstructure returns a new FlatConfig.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"ami_id":                     &hcldec.BlockAttrsSpec{TypeName: "ami_id", ElementType: cty.String, Required: true},
	}
	return s
}
