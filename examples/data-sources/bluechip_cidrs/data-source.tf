data "bluechip_cidrs" "current" {
  namespace = "office"
}

data "bluechip_cidrs" "current2" {
  namespace = "office"

  filter {
    operator = "fuzzy"
    field = "metadata.name"
    value = "console"
  }
}
