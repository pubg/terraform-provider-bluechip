resource "bluechip_user" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    password = "tetete"
    groups = ["asdf"]
  }
}
