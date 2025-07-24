package test

import (
	"strings"
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

	assert.True(s.T(), strings.Contains(policy, "ec2:DescribeInstances"))
	assert.True(s.T(), strings.Contains(policy, "kms:*"))
	assert.True(s.T(), strings.Contains(policy, "s3:GetObject"))
	assert.True(s.T(), strings.Contains(policy, "s3:ListBucket"))
	assert.True(s.T(), strings.Contains(policy, "s3:ListBucketMultipartUploads"))
	assert.True(s.T(), strings.Contains(policy, "s3:ListBucketVersions"))
	assert.True(s.T(), strings.Contains(policy, "s3:ListMultipartUploadParts"))
	assert.True(s.T(), strings.Contains(policy, "s3:PutObject"))
	assert.True(s.T(), strings.Contains(policy, "s3:HeadObject"))

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
