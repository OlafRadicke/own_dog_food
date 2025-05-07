terraform {
  backend "local" {
    path = "/terraform/state/openbao.tfstate"
  }
}