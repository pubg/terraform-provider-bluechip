resource "bluechip_cluster" "current" {
  metadata {
    name = "test"
    namespace = "default"
  }
  spec {
    project = "pubg"
    environment = "dev"
    organization_unit = "devops"
    platform = "pc"
    pubg {
      infra = "common"
      site = "devops"
    }
    vendor {
      name = "AWS"
      account_id = "12398213"
      engine = "EKS"
      region = "ap-northeast-2"
    }
    kubernetes {
      endpoint = "https://api.devops.dev.pubg.com"
      ca_cert = "-----BEGIN CERTIFI"
      sa_issuer = "https://login.microsoftonline.com/1a27bdbf-e6cc-4e33-85d2-e1c81bad930a/v2.0"
      version = "1.28"
    }
  }
}
