resource "vault_mount" "root_ca" {
  path        = "root_ca"
  type        = "pki"
  description = "PKI Secrets Engine"
}


# Create an self sign root certificate
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


# Set the issuer of the policy ca
# TODO: is it right?
resource "vault_pki_secret_backend_issuer" "root_ca" {
  depends_on = [
    vault_pki_secret_backend_root_cert.root_ca
  ]
  backend     = vault_mount.root_ca.path
  issuer_ref  = vault_pki_secret_backend_root_cert.root_ca.issuer_id
  issuer_name = "root-ca"
}


# resource "vault_pki_secret_backend_config_issuers" "config" {
#   depends_on                    = [vault_pki_secret_backend_issuer.policy_ca_01]
#   backend                       = vault_mount.policy_ca_01.path
#   default                       = vault_pki_secret_backend_issuer.policy_ca_01.issuer_id
#   default_follows_latest_issuer = true
# }
