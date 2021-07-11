terraform {
  required_providers {
    upstash = {
      version = "0.1"
      source  = "github.com/threpio/terraform-providers-upstash"
    }
  }
}

variable "database_name" {
  type    = string
  default = "testdb"
}

data "upstash_databases" "all" {}

# Returns all databases
output "all_databases" {
  value = data.upstash_databases.all.databases
}

# Only returns one
output "database" {
  value = {
    for database in data.upstash_databases.all.databases :
    database.database_id => database
    if database.database_name == var.database_name
  }
}
