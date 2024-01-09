resource "bluechip_cidr" "c1" {
  metadata {
    name      = "cidr1"
    namespace = "default"
    labels    = {
      "foo" = "bar"
    }
    annotations = {
      "office"                        = "true"
      "bluechip.example.com/location" = "seoul"
    }
  }
  spec {
    ipv4_cidrs = ["1.1.1.1/32"]
  }
}

resource "bluechip_cidr" "c2" {
  metadata {
    name      = "cidr2"
    namespace = "default"
    labels    = {
      "foo" = "bar"
    }
    annotations = {
      "office" = "true"
      "bluechip.example.com/location" = "tokyo"
    }
  }
  spec {
    ipv4_cidrs = ["1.1.1.1/32"]
  }
}

data "bluechip_cidrs" "office" {
  namespace = "default"
  filter {
    operator = "equal"
    field    = "metadata.annotations.office"
    value    = "true"
  }
}

data "bluechip_cidrs" "seoul" {
  namespace = "default"

  filter {
    operator = "equal"
    field    = "metadata.annotations.bluechip.example.com/location"
    value    = "true"
  }
}
