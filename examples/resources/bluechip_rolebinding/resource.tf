resource "bluechip_rolebinding" "current" {
  metadata {
    name = "my-test"
    namespace = "default"
  }
  spec {
    subject_ref {
      kind = "User"
      name = "my-test"
    }
    policy_ref = "admin"
  }
}
