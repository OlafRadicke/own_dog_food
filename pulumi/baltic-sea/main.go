package main

import (
	// "fmt"
	"baltic-sea/goprojects"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createNamespace(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/namespace/pulumi-apps.yaml",
	})
	return err
}

var balticSea = func(ctx *pulumi.Context) error {
	nameSpaceName := "pulumi-apps"

	err := createNamespace(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createDebugContainer(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createPrometheus(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createTifTest(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createTifTestIngress(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateTheIndependentFriendDe(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateTheIndependentFriendDeIngress(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createTestOlafRadickeDE(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	// err = createPrometheusIngress(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = createKeycloakPostres(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = installKeycloakOperator(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = createKeycloak(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	pulumi.Run(balticSea)
	// pulumi.Run(mySetup)
}
