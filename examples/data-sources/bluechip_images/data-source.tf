data "bluechip_images" "current" {
  filter {
    operator = "equal"
    field      = "spec.commitHash"
    value   = "6874ece755439b5b3473b5b910fb4938751d6689"
  }
  namespace = "pubg"
}
