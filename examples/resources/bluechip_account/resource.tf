resource "bluechip_account" "current" {
  metadata {
    name = "test2"
    namespace = "default"
  }
  spec {
    account_id = "12398213"
    display_name = "test"
    description = "test"
    alias = "test"
    vendor = "AWS"
    regions = ["test"]
  }
}
