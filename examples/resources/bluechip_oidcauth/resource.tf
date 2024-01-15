resource "bluechip_oidcauth" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    username_claim  = "sub"
    username_prefix = "string"
    issuer          = "https://accounts.google.com/"
    client_id       = "string"
    required_claims = ["string"]
    groups_claim    = "string"
    groups_prefix   = "string"
    attribute_mapping {
      from               = "namespace_path"
      from_path_resolver = "bare"
      to                 = "namespace_path"
    }
    attribute_mapping {
      from               = "project_path"
      from_path_resolver = "bare"
      to                 = "project_path"
    }
    attribute_mapping {
      from               = "pipeline_source"
      from_path_resolver = "bare"
      to                 = "pipeline_source"
    }
    attribute_mapping {
      from               = "ref_path"
      from_path_resolver = "bare"
      to                 = "ref_path"
    }
  }
}
