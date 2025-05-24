
resource "vault_mount" "policy_ca_01" {
  depends_on = [
    vault_pki_secret_backend_root_cert.root_ca,
  ]
  path                      = "policy_ca_01"
  type                      = vault_mount.root_ca.type
  description               = "intermediate"
  default_lease_ttl_seconds = 2592000
  max_lease_ttl_seconds     = 2592000
}

# Generate a key
resource "vault_pki_secret_backend_key" "policy_ca_01" {
  backend  = vault_mount.policy_ca_01.path
  key_name = "policy_ca_01"
  type     = "internal"
  key_type = "rsa"
  key_bits = "4096"
}

# Generate an CSR
resource "vault_pki_secret_backend_intermediate_cert_request" "policy_ca_01" {
  # depends_on = [
  #   vault_mount.root_ca,
  #   vault_mount.policy_ca_01
  # ]
  common_name = "Irish sea policy CA 01"
  backend     = vault_mount.policy_ca_01.path
  type        = vault_pki_secret_backend_root_cert.root_ca.type
  # type        = "existing"
  # key_ref     = vault_pki_secret_backend_key.policy_ca_01.key_id
  # format             = "pem"
  # private_key_format = "der"
  # key_type           = "rsa"
  # key_bits           = "4096"
  # backend            = vault_mount.root_ca.path
  # backend            = vault_mount.policy_ca_01.path
}

# Root CA Signing of CSR
resource "vault_pki_secret_backend_root_sign_intermediate" "policy_ca_01" {
  depends_on = [
    vault_pki_secret_backend_intermediate_cert_request.policy_ca_01,
    vault_mount.root_ca,
  ]
  backend              = vault_mount.root_ca.path
  csr                  = vault_pki_secret_backend_intermediate_cert_request.policy_ca_01.csr
  format               = "pem"
  common_name          = "Irish sea policy CA 01"
  exclude_cn_from_sans = true
  ou                   = "irish sea"
  organization         = "own dog food"
  country              = "DE"
  locality             = "Krefeld"
  province             = "NRW"
  ttl                  = 252288000 #8 years
}

# Now that CSR is processed and we have a signed cert
# Put the Certificate, and The Root CA into the backend
# mount point.  IF you do not put the CA in here, the
# chained_ca output of a generated cert will only be
# the intermedaite cert and not the whole chain.
resource "vault_pki_secret_backend_intermediate_set_signed" "policy_ca_01" {
  backend     = vault_mount.policy_ca_01.path
  certificate = vault_pki_secret_backend_root_sign_intermediate.policy_ca_01.certificate
}

# Set the issuer of the policy ca
# TODO: is it right?
resource "vault_pki_secret_backend_issuer" "policy_ca_01" {
  backend    = vault_mount.policy_ca_01.path
  issuer_ref = vault_pki_secret_backend_issuer.root_ca.issuer_id
  # issuer_ref  = vault_pki_secret_backend_intermediate_set_signed.policy_ca_01.imported_issuers
  issuer_name = "policy-ca-01"
}
