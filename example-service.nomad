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
          EXAMPLE_ENV_VAR = {{ .EXAMPLE_ENV_VAR }}
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