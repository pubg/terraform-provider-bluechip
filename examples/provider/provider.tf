provider "bluechip" {
  address = "app.bluechip.com"
  basic_auth {
    username = "myuser"
    password = "mypassword"
  }
}

provider "bluechip" {
  alias   = "token"
  address = "app.bluechip.com"
  token   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}

provider "bluechip" {
  alias   = "aws"
  address = "app.bluechip.com"
  aws_auth {
    cluster_name = "bluechip-prod"
    region       = "us-east-1"
  }
}

provider "bluechip" {
  alias   = "oidc"
  address = "app.bluechip.com"
  oidc_auth {
    validator_name = "kubernetes-centre"
    token          = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}

