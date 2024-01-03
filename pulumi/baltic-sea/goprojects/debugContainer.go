package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateDebugContainer(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "debug-container", &yaml.ConfigFileArgs{
		File: "yaml/debugging/fedora-deployment.yaml",
	})
	if err != nil {
		return err
	}
	return nil

}
