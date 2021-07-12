terraform {
  required_providers {
    upstash = {
      version = "0.1"
      source  = "hashicorp.com/edu/upstash"
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

# Only returns database
output "database" {
  value = {
  for database in data.upstash_databases.all.databases :
  database.id => database
  if database.database_name == var.database_name
  }
}