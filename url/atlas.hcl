env "local" {
  dev = "docker://postgres/15"

  migration {
    dir = "file://migrations"
    format = golang-migrate
  }
}
