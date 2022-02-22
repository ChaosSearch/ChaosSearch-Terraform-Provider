terraform {
  required_providers {
    chaossearch = {
      version = "~> 0.1.1"
      source = "chaossearch/chaossearch"
    }
  }
}
provider "chaossearch" {
  url               = "https://ap-south-1-aeternum.chaossearch.io"
  access_key_id     = "LCE8T6HRFGJI3ZKBGMGD"
  secret_access_key = "r5MEYkYntYvXqRSBMK6SFLQfPw7hHRQ0v5cqlkIk"
  region            = "ap-south-1"
  login  {
    user_name = "service_user@chaossearch.com"
    password = "thisIsAnEx@mple1!"
    parent_user_id = "be4aeb53-21d5-4902-862c-9c9a17ad6675"
  }

}

#create view
resource "chaossearch_view" "chaossearch-update-view-test" {
  bucket = "Chathura-update-view-test-1"
  case_insensitive = false
  index_pattern   = ".*11"
  index_retention = -1
  overwrite       = true
  sources         = []
  time_field_name = "@timestamp"
  transforms      = []
  filter {
    predicate {
      _type = "chaossumo.query.NIRFrontend.Request.Predicate.Negate"
      pred {
        _type = "chaossumo.query.NIRFrontend.Request.Predicate.TextMatch"
        field = "cs_partition_key_0"
        query = "*bluebike*"
        state {
          _type = "chaossumo.query.QEP.Predicate.TextMatchState.Exact"
        }
      }
    }
  }
}