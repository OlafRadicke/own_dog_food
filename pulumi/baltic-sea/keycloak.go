package main

import (
	// "fmt"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func installKeycloakOperator(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := yaml.NewConfigFile(ctx, "operator-keycloak-realm-imports", &yaml.ConfigFileArgs{
		File: "yaml/keycloak_operator/01_keycloakrealmimports.k8s.keycloak.org-v1.yml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "operator-keycloaks", &yaml.ConfigFileArgs{
		File: "yaml/keycloak_operator/02_keycloaks.k8s.keycloak.org-v1.yml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "operator-keycloaks-kubernetes", &yaml.ConfigFileArgs{
		File: "yaml/keycloak_operator/03_kubernetes.yml",
	})
	if err != nil {
		return err
	}
	return err
}

func createKeycloakPostres(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "keycloak-01-postgres-config", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/postgres-config.yaml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "keycloak-01-postgres-deployment", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/postgres-deployment.yaml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "keycloak-01-postgres-pvc-pv", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/postgres-pvc-pv.yaml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "keycloak-01-postgres-service", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/postgers-service.yaml",
	})
	if err != nil {
		return err
	}
	return nil

}

func createKeycloak(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "keycloak-01-example-tls-secret", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/example-tls-secret.yaml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "keycloak-01-keycloak-db-secret", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/keycloak-db-secret.yaml",
	})
	if err != nil {
		return err
	}
	_, err = yaml.NewConfigFile(ctx, "keycloak-01-keycloak", &yaml.ConfigFileArgs{
		File: "yaml/keycloak/keycloak.yaml",
	})
	if err != nil {
		return err
	}

	return nil
}
