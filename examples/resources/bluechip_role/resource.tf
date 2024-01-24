resource "bluechip_role" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    statements {
      actions = ["read"]
      paths = ["/**"]
    }
  }
}
