resource "bluechip_clusterrolebinding" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    subject_ref {
      kind = "User"
      name = "my-test"
    }
    policy_ref = "admin"
  }
}
