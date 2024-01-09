data "bluechip_rolebindings" "current" {
  filter {
    operator = "equals"
    key      = "metadata.name"
    value   = ""
  }
  namespace = "pubg"
}
