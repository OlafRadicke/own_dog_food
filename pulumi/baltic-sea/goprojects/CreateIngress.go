package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateIngress(ctx *pulumi.Context, nameSpaceName string) error {
	yamlConfig := &yaml.ConfigFileArgs{
		File: "yaml/ingress/*",
	}
	_, err := yaml.NewConfigFile(ctx, "ingress", yamlConfig)
	if err != nil {
		return err
	}

	return nil
}
