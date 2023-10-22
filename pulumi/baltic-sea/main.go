package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createNamespace(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/namespace/pulumi-apps.yaml",
	})
	return err
}

func createPrometheus(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := helm.NewChart(ctx, "prometheus", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("prometheus"),
		Version:   pulumi.String("15.9.1"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://prometheus-community.github.io/helm-charts"),
		},
		Values: pulumi.Map{
			"rbac": pulumi.Map{
				"enabled": pulumi.Bool(false),
			},
			"nodeExporter": pulumi.Map{
				"create": pulumi.Bool(false),
				"service": pulumi.Map{
					"hostPort": pulumi.Int(9110),
				},
			},
		},
	})
	return err
}

func createPrometheusIngress(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := helm.NewChart(ctx, "prometheus-ingress", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("tls_routing"),
		Version:   pulumi.String("0.3.0"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://olafradicke.github.io/own_dog_food/helm/charts/"),
		},
		Values: pulumi.Map{
			"mtls": pulumi.Map{
				"enabled":     pulumi.Bool(false),
				"secretnames": pulumi.String("mtls-ca"),
			},
			"app": pulumi.Map{
				"name":         pulumi.String("prometheus"),
				"fqdn":         pulumi.String("prometheus.olaf-radicke.de"),
				"service":      pulumi.String("prometheus-server"),
				"service_port": pulumi.Int(80),
			},
		},
	})
	return err
}

// Documentation: https://github.com/oauth2-proxy/manifests
func createOauth2Proxy(ctx *pulumi.Context, nameSpaceName string) error {

	cfg := config.New(ctx, "")
	oauth2ProxyclientSecret = cfg.requireSecret("oauth2ProxyclientSecret")

	_, err := helm.NewChart(ctx, "oauth2-proxy", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("oauth2-proxy"),
		Version:   pulumi.String("6.18.0"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://oauth2-proxy.github.io/manifests"),
		},
		Values: pulumi.Map{
			"extraArgs": pulumi.Map{
				"provider":        pulumi.String("keycloak-oidc"),
				"client-id":       pulumi.String("<your client's id>"),
				"client-secret":   pulumi.String(oauth2ProxyclientSecret),
				"redirect-url":    pulumi.String("https://internal.yourcompany.com/oauth2/callback"),
				"oidc-issuer-url": pulumi.String("https://<keycloak host>/auth/realms/<your realm>"),
				"email-domain":    pulumi.String("<yourcompany.com>"),
				// "allowed-role":          pulumi.String("<realm role name>"),
				// "allowed-role":          pulumi.String("<client id>:<client role name>"),
				"allowed-group":         pulumi.String("</group name>"),
				"code-challenge-method": pulumi.String("S256"),
			},
		},
	})
	return err
}

func createKeycloak(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/keycloak/01_keycloakrealmimports.k8s.keycloak.org-v1.yml",
	})
	if err != nil {
		return err
	}
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/keycloak/02_keycloaks.k8s.keycloak.org-v1.yml",
	})
	if err != nil {
		return err
	}
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/keycloak/03_kubernetes.yml",
	})
	if err != nil {
		return err
	}
	return err
}

var balticSea = func(ctx *pulumi.Context) error {
	nameSpaceName := "pulumi-apps"

	err := createNamespace(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createPrometheus(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = createPrometheusIngress(ctx, nameSpaceName)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	pulumi.Run(balticSea)
	// pulumi.Run(mySetup)
}
