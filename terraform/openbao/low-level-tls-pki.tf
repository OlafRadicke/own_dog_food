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

# resource "vault_pki_secret_backend_config_ca" "root_ca_config" {
#   depends_on = [
#     vault_mount.pki_root_ca,
#     tls_self_signed_cert.root_ca_cert
#   ]
#   backend    = vault_mount.pki_root_ca.path
#   pem_bundle = tls_private_key.root-ca-key.private_key_pem
#   # pem_bundle    = tls_private_key.root-ca-key.private_key_pem_pkcs8

# }
