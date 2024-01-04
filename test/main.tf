provider "bluechip" {
  address = ""
  basic_auth {
  }
  aws_auth {
    cluster_name = "bluechip"
    region       = "ap-northeast-2"
    profile       = "pubg-dev"
  }
  jwt_auth {
    token = ""
    auth_method_name = ""
  }
}

#data "bluechip_apiresources" "current" {
#
#}
#
#output "asdf" {
#  value = data.bluechip_apiresources.current
#}

resource "bluechip_cluster" "current" {
  metadata {
  }
}



terraform {
  required_providers {
    bluechip = {
      source = "pubg/bluechip"
    }
  }
}
