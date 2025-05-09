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
#   token   = "your-vault-token"
}

resource "vault_generic_secret" "example" {
  path = "secret/example"
  data_json = jsonencode({
    username = "my-user"
    password = "my-password"
  })
}

# https://registry.terraform.io/providers/hashicorp/vault/latest/docs/resources/pki_secret_backend_root_cert#backend-1

# resource "vault_pki_secret_backend_root_cert" "root_ca" {
#   depends_on            = [vault_mount.pki]
#   backend               = "http://openbao.openbao:8200"
#   type                  = "internal"
#   common_name           = "Root CA"
#   ttl                   = "315360000"
#   format                = "pem"
#   private_key_format    = "der"
#   key_type              = "rsa"
#   key_bits              = 4096
#   exclude_cn_from_sans  = true
#   ou                    = "irish sea"
#   organization          = "own dog food"
# }