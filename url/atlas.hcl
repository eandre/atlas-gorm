env "local" {
  url = "postgresql://atlas-gorm-6kd2:local@localhost:9500/url?sslmode=disable"
  dev = "docker://postgres/15"

  migration {
    dir = "file://migrations"
    format = golang-migrate
  }
}
