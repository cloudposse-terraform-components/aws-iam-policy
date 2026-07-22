locals {
}

module "iam_policy" {
  source  = "cloudposse/iam-policy/aws"
  version = "2.0.2"

  iam_policy                    = var.iam_policy
  description                   = var.description
  iam_source_policy_documents   = var.iam_source_policy_documents
  iam_override_policy_documents = var.iam_override_policy_documents

  # The `cloudposse/iam-policy/aws` module names the policy `module.this.id`.
  # When `use_fullname` is false, restrict the label order to just `name` so the
  # policy is named `var.name`, mirroring the `use_fullname` behavior of the
  # `aws-iam-role` component.
  label_order = var.use_fullname ? null : ["name"]

  iam_policy_enabled = true
  context            = module.this.context
}
