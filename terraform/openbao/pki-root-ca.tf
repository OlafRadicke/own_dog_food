resource "vault_mount" "root_ca" {
  path                      = "root_ca"
  type                      = "pki"
  description               = "PKI root CA"
  default_lease_ttl_seconds = "630720000" # 20 Jahre
  max_lease_ttl_seconds     = "630720000" # 20 Jahre
}


# Create an self sign root certificate
resource "vault_pki_secret_backend_root_cert" "root_ca" {
  depends_on         = [vault_mount.root_ca]
  backend            = vault_mount.root_ca.path
  type               = "internal"
  common_name        = "Irish sea root CA 01"
  ttl                = vault_mount.root_ca.default_lease_ttl_seconds
  format             = "pem"
  private_key_format = "der"
  key_type           = "rsa"
  key_bits           = 4096
  # not_after            = "2045-06-08T12:00:00Z"
  exclude_cn_from_sans = true
  ou                   = "irish sea"
  organization         = "own dog food"
  country              = "DE"
  locality             = "Krefeld"
  province             = "NRW"
}

# ttl                   = "315360000"
# ttl                  = "175200h" # 20 Jahre

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


resource "vault_pki_secret_backend_config_issuers" "config" {
  depends_on                    = [vault_pki_secret_backend_root_cert.root_ca]
  backend                       = vault_mount.root_ca.path
  default                       = vault_pki_secret_backend_issuer.root_ca.issuer_id
  default_follows_latest_issuer = true
}
