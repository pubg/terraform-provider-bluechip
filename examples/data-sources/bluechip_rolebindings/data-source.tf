data "bluechip_rolebindings" "current" {
  filter {
    operator = "equal"
    key      = "metadata.name"
    value   = ""
  }
  namespace = "pubg"
}
