job "example" {
  namespace   = "example"
  datacenters = ["dc1"]

  group "example" {

    count = 1

    network {
      port "vs" {
        static = 2000 # static for now
      }
    }

    task "example" {
      driver = "docker"

      # Render environment variables from Nomad Variables
      template {
        destination = "${NOMAD_SECRETS_DIR}/env.vars"
        env         = true
        change_mode = "restart"
        data        = <<EOF
          {{- with nomadVar "nomad/jobs/example/example/example" -}}
          DATABASE_NAME = {{ .DATABASE_NAME }}
          DATABASE_USER  = {{ .DATABASE_USER }}
          DATABASE_PASSWORD = {{ .DATABASE_PASSWORD }}
          DATABASE_HOST = {{ .DATABASE_HOST }}
          DATABASE_PORT = {{ .DATABASE_PORT }}
          DATABASE_SSL_MODE = {{ .DATABASE_SSL_MODE }}
          VAULT_ADDR = {{ .VAULT_ADDR }}
          VAULT_TOKEN = {{ .VAULT_TOKEN }}
          VAULT_TRANSIT_KEY_NAME = {{ .VAULT_TRANSIT_KEY_NAME }}
          {{- end -}}
          EOF
      }


      config {
        image          = "assetsatlas/example-service:1.0"
        ports          = ["vs"]
        auth_soft_fail = true
      }

      resources {
        cpu    = 160
        memory = 256
      }

      env {
        LOG_LEVEL = "debug"
      }
    }
  }
}