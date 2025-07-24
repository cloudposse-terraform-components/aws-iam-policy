variable "region" {
  type        = string
  description = "AWS Region"
}

variable "iam_policy" {
  type = list(object({
    policy_id = optional(string, null)
    version   = optional(string, null)
    statements = list(object({
      sid           = optional(string, null)
      effect        = optional(string, null)
      actions       = optional(list(string), null)
      not_actions   = optional(list(string), null)
      resources     = optional(list(string), null)
      not_resources = optional(list(string), null)
      conditions = optional(list(object({
        test     = string
        variable = string
        values   = list(string)
      })), [])
      principals = optional(list(object({
        type        = string
        identifiers = list(string)
      })), [])
      not_principals = optional(list(object({
        type        = string
        identifiers = list(string)
      })), [])
    }))
  }))
  description = <<-EOT
    IAM policy as list of Terraform objects, compatible with Terraform `aws_iam_policy_document` data source
    except that `source_policy_documents` and `override_policy_documents` are not included.
    Use inputs `iam_source_policy_documents` and `iam_override_policy_documents` for that.
    EOT
  default     = []
  nullable    = false
}

variable "description" {
  type        = string
  description = "Description of created IAM policy"
  default     = null
}


variable "iam_source_policy_documents" {
  type        = list(string)
  description = <<-EOT
    List of IAM policy documents (as JSON strings) that are merged together into the exported document.
    Statements defined in `iam_source_policy_documents` must have unique SIDs and be distinct from SIDs
    in `iam_policy`.
    Statements in these documents will be overridden by statements with the same SID in `iam_override_policy_documents`.
    EOT
  default     = null
}

variable "iam_override_policy_documents" {
  type        = list(string)
  description = <<-EOT
    List of IAM policy documents (as JSON strings) that are merged together into the exported document with higher precedence.
    In merging, statements with non-blank SIDs will override statements with the same SID
    from earlier documents in the list and from other "source" documents.
    EOT
  default     = null
}
