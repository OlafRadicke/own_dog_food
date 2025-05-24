
# resource "vault_mount" "pki_service_01" {
#   depends_on = [
#     vault_pki_secret_backend_root_cert.root_ca,
#   ]
#   path                      = "pki_service_01"
#   type                      = "pki"
#   description               = "leaf certs"
#   default_lease_ttl_seconds = 2592000
#   max_lease_ttl_seconds     = 2592000
# }

# #
# # Role for server certs
# # This creates certs of machinename.mydomain.com
# resource "vault_pki_secret_backend_role" "role-server-cer-01" {
#   depends_on = [
#     vault_mount.pki_service_01,
#     vault_pki_secret_backend_root_sign_intermediate.policy_ca_01
#   ]
#   backend = vault_mount.pki_service_01.path
#   name    = "Service 01"
#   allowed_domains = [
#     "my.fun",
#     "your.fun",
#   ]
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
#   # 2 years
#   max_ttl = 63113904
#   # 30 days
#   ttl      = 2592000
#   no_store = true

# }

# Alternativ: vault_pki_secret_backend_sign
