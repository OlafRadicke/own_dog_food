# https://registry.terraform.io/providers/hashicorp/vault/latest/docs/resources/pki_secret_backend_intermediate_set_signed

resource "vault_mount" "pki_policy_ca_01" {
  depends_on                = [vault_pki_secret_backend_root_cert.root_ca]
  path                      = "pki_policy_ca_01"
  type                      = vault_mount.pki_root_ca.type
  description               = "intermediate"
  default_lease_ttl_seconds = 2592000
  max_lease_ttl_seconds     = 2592000
}


resource "vault_pki_secret_backend_issuer" "policy_ca_01" {
  backend     = vault_pki_secret_backend_root_cert.root_ca.backend
  issuer_ref  = vault_pki_secret_backend_root_cert.root_ca.issuer_id
  issuer_name = "Irish sea root CA 01"
}

resource "vault_pki_secret_backend_config_issuers" "config" {
  backend                       = vault_mount.pki_policy_ca_01.path
  default                       = vault_pki_secret_backend_issuer.policy_ca_01.issuer_id
  default_follows_latest_issuer = true
}


resource "vault_pki_secret_backend_intermediate_cert_request" "csr_policy_ca_01" {
  backend     = vault_mount.pki_policy_ca_01.path
  type        = vault_pki_secret_backend_root_cert.root_ca.type
  common_name = "Irish sea policy CA 01"
}


resource "vault_pki_secret_backend_root_sign_intermediate" "sign_policy_ca_01" {
  depends_on           = [vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01]
  backend              = vault_mount.pki_root_ca.path
  csr                  = vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01.csr
  common_name          = "Irish sea policy CA 01"
  exclude_cn_from_sans = true
  ou                   = "irish sea"
  organization         = "own dog food"
  country              = "DE"
  locality             = "Krefeld"
  province             = "NRW"
  revoke               = true
}


resource "vault_pki_secret_backend_intermediate_set_signed" "policy_ca_01" {
  backend     = vault_mount.pki_policy_ca_01.path
  certificate = vault_pki_secret_backend_root_sign_intermediate.sign_policy_ca_01.certificate
}