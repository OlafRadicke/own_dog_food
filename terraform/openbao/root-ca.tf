# https://developer.hashicorp.com/terraform/tutorials/cloud/dynamic-credentials-no-code


# https://developer.hashicorp.com/vault/tutorials/get-started/learn-terraform
# provider "vault" {
#    auth_login {
#       path = "auth/userpass/login/${var.login_username}"
#       parameters = {
#          password = var.login_password
#       }
#    }
# }



provider "vault" {
  address = "http://openbao.openbao:8200"
  token   = var.vault_token
#   token   = "your-vault-token"
}


# resource "vault_mount" "pki-issuer" {
#   path        = "pki/issuer"
#   type        = "pki"
# }

# resource "vault_generic_secret" "example" {
#   path = "secret/example"
#   data_json = jsonencode({
#     username = "my-user"
#     password = "my-password"
#   })
# }

resource "vault_mount" "pki" {
  path = "pki/issuer"
  type = "pki"
  description = "PKI Secrets Engine"
}

# Workarount https://www.infralovers.com/de/blog/2023-10-16-hashicorp-vault-acme-terraform-configuration/
resource "vault_generic_endpoint" "root_config_cluster" {
  depends_on           = [vault_mount.pki]
  path                 = "${vault_mount.pki.path}/config/cluster"
  ignore_absent_fields = true
  disable_delete       = true

  data_json = <<EOT
{
    "aia_path": "http://openbao.openbao:8200/v1/${vault_mount.pki.path}",
    "path": "http://openbao.openbao:8200/v1/${vault_mount.pki.path}"
}
EOT
}

resource "vault_generic_endpoint" "root_config_urls" {
  depends_on           = [vault_mount.pki, vault_generic_endpoint.root_config_cluster]
  path                 = "${vault_mount.pki.path}/config/urls"
  ignore_absent_fields = true
  disable_delete       = true

  data_json = <<EOT
{
    "enable_templating": true,
    "issuing_certificates": "{{cluster_aia_path}}/issuer/{{issuer_id}}/der",
    "crl_distribution_points": "{{cluster_aia_path}}/issuer/{{issuer_id}}/crl/der",
    "ocsp_servers": "{{cluster_path}}/ocsp"
}
EOT
}

resource "vault_pki_secret_backend_role" "server-pki" {
  backend        = vault_mount.pki.path
  name           = "server-pki"
  no_store       = false
  allow_any_name = true
}

# https://registry.terraform.io/providers/hashicorp/vault/latest/docs/resources/pki_secret_backend_root_cert#backend-1

resource "vault_pki_secret_backend_root_cert" "root_ca" {
  depends_on            = [vault_mount.pki]
  backend               = "http://openbao.openbao:8200"
  type                  = "internal"
  common_name           = "root-ca.irish.sea"
  ttl                   = "315360000"
  format                = "pem"
  private_key_format    = "der"
  key_type              = "rsa"
  key_bits              = 4096
  exclude_cn_from_sans  = true
  ou                    = "irish sea"
  organization          = "own dog food"
}


# resource "vault_pki_secret_backend_role" "example_role" {
#   backend         = vault_mount.pki.path
#   name            = "example-role"
#   allowed_domains = ["example.com"]
#   allow_subdomains = true
#   max_ttl         = "720h"
# }

# resource "vault_pki_secret_backend_cert" "example_cert" {
#   backend     = vault_mount.pki.path
#   name        = vault_pki_secret_backend_role.example_role.name
#   common_name = "app.example.com"
# }
