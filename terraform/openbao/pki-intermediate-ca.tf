# https://registry.terraform.io/providers/hashicorp/vault/latest/docs/resources/pki_secret_backend_intermediate_set_signed

resource "vault_mount" "pki_policy_ca_01" {
  depends_on = [
    vault_pki_secret_backend_root_cert.root_ca,
  ]
  path                      = "pki_policy_ca_01"
  type                      = vault_mount.pki_root_ca.type
  description               = "intermediate"
  default_lease_ttl_seconds = 2592000
  max_lease_ttl_seconds     = 2592000
}


resource "vault_pki_secret_backend_issuer" "policy_ca_01" {
  backend     = vault_mount.pki_policy_ca_01.path # Pfad des PKI-Backends
  issuer_ref  = vault_pki_secret_backend_root_cert.pki_root_ca.issuer_id
  issuer_name = "root-ca"
}

# Generate a key
resource "vault_pki_secret_backend_key" "policy_ca_01" {
  backend  = var.issuer.backend
  type     = "internal"
  key_type = "rsa"
  key_bits = "4096"
  key_name = "policy_ca_01"
}


resource "vault_pki_secret_backend_intermediate_cert_request" "policy_ca_01" {
  depends_on  = [vault_mount.pki_policy_ca_01]
  backend     = vault_mount.pki_root_ca.path
  type        = "existing"
  common_name = "Irish sea policy CA 01"
  key_ref     = vault_pki_secret_backend_key.policy_ca_01.key_id
  # format             = "pem"
  # private_key_format = "der"
  # key_type           = "rsa"
  # key_bits           = "4096"
  # backend            = vault_mount.pki_root_ca.path
  # backend            = vault_mount.pki_policy_ca_01.path
}

# Have the Root CA Sign our CSR
resource "vault_pki_secret_backend_root_sign_intermediate" "policy_ca_01" {
  depends_on = [
    vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01,
    vault_mount.pki_root_ca,
  ]
  backend              = vault_mount.pki_policy_ca_01.path
  issuer_ref           = vault_pki_secret_backend_root_cert.root_ca.id
  csr                  = vault_pki_secret_backend_intermediate_cert_request.policy_ca_01.csr
  common_name          = "Irish sea policy CA 01"
  exclude_cn_from_sans = true
  ou                   = "irish sea"
  organization         = "own dog food"
  country              = "DE"
  locality             = "Krefeld"
  province             = "NRW"
  ttl                  = 252288000 #8 years
  # backend              = vault_mount.pki_root_ca.path
  # backend = vault_mount.pki_policy_ca_01.path
  # issuer_ref           = tls_self_signed_cert.root_ca_cert.id
}

# Now that CSR is processed and we have a signed cert
# Put the Certificate, and The Root CA into the backend
# mount point.  IF you do not put the CA in here, the
# chained_ca output of a generated cert will only be
# the intermedaite cert and not the whole chain.
resource "vault_pki_secret_backend_intermediate_set_signed" "policy_ca_01" {
  backend     = vault_mount.pki_policy_ca_01.path
  certificate = vault_pki_secret_backend_root_sign_intermediate.policy_ca_01.certificate
}






# resource "vault_pki_secret_backend_issuer" "policy_ca_01" {
#   backend     = vault_pki_secret_backend_root_cert.root_ca.backend
#   issuer_ref  = vault_pki_secret_backend_root_cert.root_ca.issuer_id
#   issuer_name = "policy_ca_01"
# }

# resource "vault_pki_secret_backend_config_issuers" "config" {
#   depends_on                    = [vault_pki_secret_backend_issuer.policy_ca_01]
#   backend                       = vault_mount.pki_policy_ca_01.path
#   default                       = vault_pki_secret_backend_issuer.policy_ca_01.issuer_id
#   default_follows_latest_issuer = true
# }


# resource "vault_pki_secret_backend_root_sign_intermediate" "sign_policy_ca_01" {
#   depends_on           = [vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01]
#   backend              = vault_mount.pki_root_ca.path
#   csr                  = vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01.csr
#   common_name          = "Irish sea policy CA 01"
#   exclude_cn_from_sans = true
#   ou                   = "irish sea"
#   organization         = "own dog food"
#   country              = "DE"
#   locality             = "Krefeld"
#   province             = "NRW"
#   revoke               = true
# }


# resource "vault_pki_secret_backend_intermediate_set_signed" "policy_ca_01" {
#   backend     = vault_mount.pki_policy_ca_01.path
#   certificate = vault_pki_secret_backend_root_sign_intermediate.sign_policy_ca_01.certificate
# }
