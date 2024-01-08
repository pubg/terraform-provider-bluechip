resource "bluechip_oidcauth" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    username_claim= "sub"
    username_prefix= "string"
    issuer = "https://accounts.google.com/"
    client_id = "string"
    required_claims = ["string"]
    groups_claim = "string"
    groups_prefix = "string"
  }
}
