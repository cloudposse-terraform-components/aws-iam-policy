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

  context = module.this.context
}

module "policy_label" {
  source  = "cloudposse/label/null"
  version = "0.25.0"

  context = module.this.context
}

resource "aws_iam_policy" "ecs_exec_edi" {
  count  = local.enabled ? 1 : 0
  name   = module.policy_label_edi.id
  policy = one(data.aws_iam_policy_document.ecs_exec_edi[*].json)
}
