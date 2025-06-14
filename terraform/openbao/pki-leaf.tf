
resource "vault_mount" "service_01" {
  depends_on = [
    vault_pki_secret_backend_root_cert.root_ca,
  ]
  path                      = "service_01"
  type                      = "pki"
  description               = "service 01 leaf certs"
  default_lease_ttl_seconds = "2592000" # 30 days
  max_lease_ttl_seconds     = "2592000" # 30 days
}


# # Role for server certs
# # This creates certs of machinename.mydomain.com
# resource "vault_pki_secret_backend_role" "role-server-cer-01" {
#   #   backend            = vault_mount.service_01.path
#   backend            = vault_mount.policy_ca_01.path
#   name               = "Service 01"
#   allowed_domains    = ["service-01.irish.sea"]
#   allow_subdomains   = true
#   allow_glob_domains = false
#   allow_any_name     = false
#   enforce_hostnames  = true
#   allow_ip_sans      = true
#   server_flag        = true
#   client_flag        = false
#   ou                 = ["irish sea"]
#   organization       = ["own dog food"]
#   country            = ["DE"]
#   locality           = ["Krefeld"]
#   province           = ["NRW"]
#   ttl                = vault_mount.service_01.default_lease_ttl_seconds
#   no_store           = true
# }
#   max_ttl            = vault_mount.service_01.default_lease_ttl_seconds

# Alternativ: vault_pki_secret_backend_sign

resource "vault_pki_secret_backend_intermediate_cert_request" "service-01" {
  backend            = vault_mount.service_01.path
  common_name        = "service-01.irish.sea"
  format             = "pem"
  private_key_format = "pkcs8"
  type               = "internal"
}

# resource "vault_pki_secret_backend_sign" "service-01" {
#   backend     = vault_mount.policy_ca_01.path
#   name        = "service-01"
#   common_name = "service-01.irish.sea"
#   csr         = vault_pki_secret_backend_intermediate_cert_request.service-01.csr
#   issuer_ref  = vault_pki_secret_backend_root_sign_intermediate.policy_ca_01.id
#   format      = "pem_bundle"
#   ttl         = "8760h"
# }

# resource "vault_pki_secret_backend_cert" "service-02" {
#   backend     = vault_mount.policy_ca_01.path
#   name        = "service-02"
#   common_name = "service-02.irish.sea"
#   alt_names   = ["www.service-01.irish.sea"]
#   ttl         = "8760h"
# }

# resource "vault_pki_secret_backend_role" "server-01" {
#   #   backend            = vault_mount.service_01.path
#   #   backend            = vault_mount.policy_ca_01.path
#   backend            = "policy_ca_01"
#   name               = "Service 01"
#   allowed_domains    = ["service-01.irish.sea"]
#   allow_subdomains   = true
#   allow_glob_domains = false
#   allow_any_name     = false
#   enforce_hostnames  = true
#   allow_ip_sans      = true
#   server_flag        = true
#   client_flag        = false
#   ou                 = ["irish sea"]
#   organization       = ["own dog food"]
#   country            = ["DE"]
#   locality           = ["Krefeld"]
#   province           = ["NRW"]
#   ttl                = vault_mount.service_01.default_lease_ttl_seconds
#   no_store           = true
# }

# resource "vault_pki_secret_backend_cert" "app" {
#   backend     = vault_mount.policy_ca_01.path
#   name        = vault_pki_secret_backend_role.server-01.name
#   common_name = "service-01.irish.sea"
# }
