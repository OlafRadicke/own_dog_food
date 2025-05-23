resource "vault_mount" "root_ca" {
  path        = "root_ca"
  type        = "pki"
  description = "PKI Secrets Engine"
}



resource "vault_pki_secret_backend_root_cert" "root_ca" {
  depends_on  = [vault_mount.root_ca]
  backend     = vault_mount.root_ca.path
  type        = "internal"
  common_name = "Irish sea root CA 01"
  # ttl                   = "315360000"
  ttl                  = "175200h" # 20 Jahre
  format               = "pem"
  private_key_format   = "der"
  key_type             = "rsa"
  key_bits             = 4096
  exclude_cn_from_sans = true
  ou                   = "irish sea"
  organization         = "own dog food"
  country              = "DE"
  locality             = "Krefeld"
  province             = "NRW"
}


# resource "vault_pki_secret_backend_issuer" "issuer_root_ca" {
#   backend     = vault_pki_secret_backend_root_cert.root_ca.backend
#   issuer_ref  = vault_pki_secret_backend_root_cert.root_ca.issuer_id
#   issuer_name = "Irish sea root CA 01 issuer"
# }

# resource "vault_pki_secret_backend_config_issuers" "config" {
#   backend                       = vault_mount.root_ca.path
#   default                       = vault_pki_secret_backend_issuer.issuer_root_ca.issuer_id
#   default_follows_latest_issuer = true
# }

# https://www.infralovers.com/de/blog/2023-10-16-hashicorp-vault-acme-terraform-configuration/

