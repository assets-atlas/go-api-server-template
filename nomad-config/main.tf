data "digitalocean_database_cluster" "psql" {
  name = "dev-postgres-cluster"
}

data "digitalocean_droplet" "vault" {
  name = "vault"
}

resource "nomad_namespace" "authentication" {
  name        = "authentication"
  description = "authentication service"
}

resource "nomad_variable" "authentication" {
  path      = "nomad/jobs/authentication/authentication/authentication"
  namespace = nomad_namespace.authentication.name
  items = {
    DATABASE_NAME          = "assetsatlas"
    DATABASE_USER          = data.digitalocean_database_cluster.psql.user
    DATABASE_PASSWORD      = data.digitalocean_database_cluster.psql.password
    DATABASE_HOST          = data.digitalocean_database_cluster.psql.host
    DATABASE_PORT          = data.digitalocean_database_cluster.psql.port
    DATABASE_SSL_MODE      = "require"
    VAULT_ADDR             = "http://${data.digitalocean_droplet.vault.ipv4_address}:8200"
    VAULT_TOKEN            = var.vault_token
    MOUNT_ACCESSOR_ID      = var.mount_accessor_id
    VAULT_TRANSIT_KEY_NAME = var.vault_transit_key_name
    VAULT_OIDC_ROLE_NAME   = var.oidc_role_name
  }
}