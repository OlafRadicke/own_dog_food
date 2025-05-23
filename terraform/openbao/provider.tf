
provider "vault" {
  address = "http://openbao.openbao:8200"
  token   = var.vault_token
  #   token   = "your-vault-token"
}
