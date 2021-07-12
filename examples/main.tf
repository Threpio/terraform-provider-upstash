terraform {
  required_providers {
    upstash = {
      version = "0.1"
      source  = "hashicorp.com/edu/upstash"
    }
  }
}

provider "upstash" {
  email = "theoandresier@gmail.com"
  api_key = "e16b3125-167d-4e28-b8d9-b128f22a0b41"
}

module "upstash_databases" {
  source = "./database"
  database_name = "test-tha2"
}

output "database" {
  value = module.upstash_databases.all_databases
}

