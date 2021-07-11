terraform {
  required_providers {
    upstash = {
      version = "0.1"
      source  = "github.com/threpio/terraform-providers-upstash"
    }
  }
}

provider "upstash" {}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}
