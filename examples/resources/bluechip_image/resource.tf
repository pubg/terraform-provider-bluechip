resource "bluechip_image" "current" {
  metadata {
    name      = "my-test"
    namespace = "default"
  }
  spec {
    app         = "my-test"
    timestamp   = 1398329823
    commit_hash = "1234567890"
    repository  = "test"
    tag         = "test"
    branch      = "test"
  }
}
