package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func DebugContainer(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "debug-container", &yaml.ConfigFileArgs{
		File: "yaml/debug-container/fedora-deployment.yaml",
	})
	if err != nil {
		return err
	}
	return nil

}
