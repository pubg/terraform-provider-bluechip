resource "bluechip_vendor" "current" {
  metadata {
    name = "current"
  }
  spec {
    display_name = "asdf"
    code_name    = "AWS"
    short_name   = "aws"
    regions      = ["asdf"]
  }
}
