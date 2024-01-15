data "bluechip_rolebinding" "current" {
  metadata {
    name      = "my-test"
    namespace = "default"
  }
}
