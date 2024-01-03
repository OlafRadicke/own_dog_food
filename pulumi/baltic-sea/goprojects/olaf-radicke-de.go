package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateOlafRadickeDE(ctx *pulumi.Context, nameSpaceName string) error {

	yamlConfig := &yaml.ConfigFileArgs{
		File: "yaml/olaf-radicke-de/deployment.yaml",
	}
	_, err := yaml.NewConfigFile(ctx, "olaf-radicke-de-deployment", yamlConfig)
	if err != nil {
		return err
	}

	yamlConfig = &yaml.ConfigFileArgs{
		File: "yaml/olaf-radicke-de/ingress.yaml",
	}
	_, err = yaml.NewConfigFile(ctx, "olaf-radicke-de-ingress", yamlConfig)
	if err != nil {
		return err
	}

	yamlConfig = &yaml.ConfigFileArgs{
		File: "yaml/olaf-radicke-de/service.yaml",
	}
	_, err = yaml.NewConfigFile(ctx, "olaf-radicke-de-service", yamlConfig)
	if err != nil {
		return err
	}

	return nil
}
