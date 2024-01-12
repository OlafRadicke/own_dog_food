package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func PulumiNamespace(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/pulumi/namespace.yaml",
	})
	return err
}
