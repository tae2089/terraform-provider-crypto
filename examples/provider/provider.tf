terraform {
  required_providers {
    crypto = {
      source = "hashicorp.com/tae2089/crypto"
    }
  }
  required_version = ">= 1.5.0"
}
# Configure the connection details for the Inventory service
provider "crypto" {}