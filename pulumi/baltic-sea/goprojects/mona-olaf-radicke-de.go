package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Mona(ctx *pulumi.Context, nameSpaceName string) error {

	yamlConfig := &yaml.ConfigFileArgs{
		File: "yaml/mona-olaf-radicke-de/*.yaml",
	}
	_, err := yaml.NewConfigFile(ctx, "mona-olaf-radicke-de-deployment", yamlConfig)
	if err != nil {
		return err
	}

	return nil
}