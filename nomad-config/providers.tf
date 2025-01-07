terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}


# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = var.do_token
}

data "digitalocean_droplet" "nomad" {
  name = "nomad-dev-server"
}
provider "nomad" {
  address = "http://${data.digitalocean_droplet.nomad.ipv4_address}:4646"
  secret_id   = var.nomad_token

}

