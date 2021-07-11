terraform {
  required_providers {
    upstash = {
      version = "0.1"
      source  = "threpio/upstash"
    }
  }
}

provider "upstash" {}

module "psl" {
  source = "./database"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}
