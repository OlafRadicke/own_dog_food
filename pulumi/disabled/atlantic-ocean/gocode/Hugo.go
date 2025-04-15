package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Hugo(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "hugo", &yaml.ConfigFileArgs{
		File: "yaml/hugo/*.yaml",
	})
	if err != nil {
		return err
	}
	return nil

}
