data "bluechip_oidcauth" "current" {
  metadata {
    name = "my-test"
  }
  depends_on = [bluechip_oidcauth.current]
}
