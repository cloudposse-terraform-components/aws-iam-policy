package test

import (
	"encoding/json"
	"testing"

	"github.com/cloudposse/test-helpers/pkg/atmos"
	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

type ComponentSuite struct {
	helper.TestSuite
}

func (s *ComponentSuite) TestBasic() {
	const component = "example/basic"
	const stack = "default-test"
	const awsRegion = "us-east-2"

	defer s.DestroyAtmosComponent(s.T(), component, stack, nil)
	options, _ := s.DeployAtmosComponent(s.T(), component, stack, nil)
	assert.NotNil(s.T(), options)
	policyArn := atmos.Output(s.T(), options, "policy_arn")

	policy := aws.GetIamPolicyDocument(s.T(), awsRegion, policyArn)
	policyMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(policy), &policyMap)
	assert.NoError(s.T(), err)

	expectedPolicy := map[string]interface{}{
		"Id": "EC2DescribeInstances",
		"Statement": []map[string]interface{}{
			{
				"Action":   "ec2:DescribeInstances",
				"Effect":   "Allow",
				"Resource": "*",
				"Sid":      "EC2DescribeInstances",
			},
			{
				"Action": []string{
					"s3:PutObject",
					"s3:ListMultipartUploadParts",
					"s3:ListBucketVersions",
					"s3:ListBucketMultipartUploads",
					"s3:ListBucket",
					"s3:HeadObject",
					"s3:GetObject",
				},
				"Effect":   "Allow",
				"Resource": "*",
				"Sid":      "S3ReadWrite",
			},
			{
				"Action":   "kms:*",
				"Effect":   "Deny",
				"Resource": "*",
				"Sid":      "DenyKmsDecrypt",
			},
		},
		"Version": "2012-10-17",
	}
	assert.Equal(s.T(), expectedPolicy, policyMap)

	s.DriftTest(component, stack, nil)
}

func (s *ComponentSuite) TestEnabledFlag() {
	const component = "example/disabled"
	const stack = "default-test"
	s.VerifyEnabledFlag(component, stack, nil)
}

func TestRunSuite(t *testing.T) {
	suite := new(ComponentSuite)
	helper.Run(t, suite)
}
