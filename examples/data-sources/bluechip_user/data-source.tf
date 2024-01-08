data "bluechip_user" "current" {
  metadata {
    name = "my-test"
  }
  depends_on = [bluechip_user.current]
}
