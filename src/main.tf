locals {
  enabled = module.this.enabled
}

module "iam_policy" {
  source  = "cloudposse/iam-policy/aws"
  version = "2.0.2"

  iam_policy                    = var.iam_policy
  description                   = var.description
  iam_source_policy_documents   = var.iam_source_policy_documents
  iam_override_policy_documents = var.iam_override_policy_documents

  iam_policy_enabled = true
  context            = module.this.context
}
