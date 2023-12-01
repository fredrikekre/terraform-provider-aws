package function_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestARNBuildFunction_known(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testARNBuildFunctionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "arn:aws:iam::444455556666:role/example"),
				),
			},
		},
	})
}

func testARNBuildFunctionConfig() string {
	return `
terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
		}
	}
}

output "test" {
	value = provider::aws::arn_build("aws", "iam", "", "444455556666", "role/example")
}`
}
