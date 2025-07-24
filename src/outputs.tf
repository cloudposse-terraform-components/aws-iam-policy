output "json" {
  description = "JSON body of the IAM policy document"
  value       = module.iam_policy.json
}

output "policy_arn" {
  description = "ARN of created IAM policy"
  value       = module.iam_policy.policy_arn
}
