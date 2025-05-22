# https://registry.terraform.io/providers/hashicorp/vault/latest/docs/resources/pki_secret_backend_intermediate_set_signed

resource "vault_mount" "pki_policy_ca_01" {
  depends_on                = [vault_pki_secret_backend_config_ca.ca_config]
  path                      = "pki_policy_ca_01"
  type                      = vault_mount.pki_root_ca.type
  description               = "intermediate"
  default_lease_ttl_seconds = 2592000
  max_lease_ttl_seconds     = 2592000
}


resource "vault_pki_secret_backend_intermediate_cert_request" "csr_policy_ca_01" {
  depends_on = [vault_mount.pki_policy_ca_01]
  # backend            = vault_mount.pki_policy_ca_01.path
  backend            = vault_mount.pki_root_ca.path
  type               = "internal"
  common_name        = "Irish sea policy CA 01"
  format             = "pem"
  private_key_format = "der"
  key_type           = "rsa"
  key_bits           = "4096"
}


resource "vault_pki_secret_backend_root_sign_intermediate" "csr_policy_ca_01" {
  depends_on = [
    vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01,
    vault_mount.pki_root_ca,
    tls_self_signed_cert.ca_cert,
    vault_pki_secret_backend_config_ca.ca_config,
  ]
  # backend              = vault_mount.pki_policy_ca_01.path
  backend              = vault_mount.pki_root_ca.path
  csr                  = vault_pki_secret_backend_intermediate_cert_request.csr_policy_ca_01.csr
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
resource "vault_pki_secret_backend_intermediate_set_signed" "intermediate" {
  backend = vault_mount.pki_policy_ca_01t.path

  certificate = "${vault_pki_secret_backend_root_sign_intermediate.pki_policy_ca_01.certificate}\n${tls_self_signed_cert.root_ca_cert.cert_pem}"
}



#
# Role for server certs
# This creates certs of machinename.mydomain.com
resource "vault_pki_secret_backend_role" "role-server-cer-01" {
  backend            = vault_mount.pki_policy_ca_01t.path
  name               = "Service 01"
  allowed_domains    = [var.server_cert_domain]
  allow_subdomains   = true
  allow_glob_domains = false
  allow_any_name     = false
  enforce_hostnames  = true
  allow_ip_sans      = true
  server_flag        = true
  client_flag        = false
  ou                 = "irish sea"
  organization       = "own dog food"
  country            = "DE"
  locality           = "Krefeld"
  province           = "NRW"
  # 2 years
  max_ttl = 63113904
  # 30 days
  ttl      = 2592000
  no_store = true

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
