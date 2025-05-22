
provider "vault" {
  address = "http://openbao.openbao:8200"
  token   = var.vault_token
  #   token   = "your-vault-token"
}


# resource "tls_private_key" "root-ca-key" {
#   algorithm = "RSA"
#   rsa_bits  = 4096
# }


# resource "tls_self_signed_cert" "root_ca_cert" {
#   depends_on = [
#     tls_private_key.root-ca-key,
#   ]
#   private_key_pem = tls_private_key.root-ca-key.private_key_pem
#   # key_algorithm = "RSA"
#   subject {
#     common_name         = "Irish sea root CA 01"
#     organization        = "own dog food"
#     organizational_unit = "irish sea"
#     country             = "DE"
#     locality            = "Krefeld"
#     province            = "NRW"

#   }
#   # 20 Jahre
#   validity_period_hours = 175200
#   allowed_uses = [
#     "cert_signing",
#     "crl_signing",
#     "server_auth",
#   ]
#   is_ca_certificate = true
# }



resource "vault_mount" "pki_root_ca" {
  path        = "pki_root_ca"
  type        = "pki"
  description = "PKI Secrets Engine"
}

resource "vault_pki_secret_backend_config_ca" "root_ca_config" {
  depends_on = [
    vault_mount.pki_root_ca,
    tls_self_signed_cert.root_ca_cert
  ]
  backend    = vault_mount.pki_root_ca.path
  pem_bundle = tls_private_key.root-ca-key.private_key_pem
  # pem_bundle    = tls_private_key.root-ca-key.private_key_pem_pkcs8

}


resource "vault_pki_secret_backend_root_cert" "root_ca" {
  depends_on  = [vault_mount.pki_root_ca]
  backend     = vault_mount.pki_root_ca.path
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
#   backend                       = vault_mount.pki_root_ca.path
#   default                       = vault_pki_secret_backend_issuer.issuer_root_ca.issuer_id
#   default_follows_latest_issuer = true
# }

# https://www.infralovers.com/de/blog/2023-10-16-hashicorp-vault-acme-terraform-configuration/

