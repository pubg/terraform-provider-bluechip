resource "bluechip_cidr" "current" {
  metadata {
    name      = "my-test"
    namespace = "default"
  }
  spec {
    ipv4_cidrs = ["1.1.1.1/32"]
    ipv6_cidrs = ["::1/32"]
  }
}
