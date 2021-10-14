package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/servicediscovery"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func TestAccAWSServiceDiscoveryDnsNamespaceDataSource_private(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_service_discovery_dns_namespace.test"
	resourceName := "aws_service_discovery_private_dns_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(servicediscovery.EndpointsID, t)
			testAccPreCheckAWSServiceDiscovery(t)
		},
		ErrorCheck: acctest.ErrorCheck(t, servicediscovery.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsServiceDiscoveryPrivateDnsNamespaceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hosted_zone", resourceName, "hosted_zone"),
				),
			},
		},
	})
}

func TestAccAWSServiceDiscoveryDnsNamespaceDataSource_public(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_service_discovery_dns_namespace.test"
	resourceName := "aws_service_discovery_public_dns_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(servicediscovery.EndpointsID, t)
			testAccPreCheckAWSServiceDiscovery(t)
		},
		ErrorCheck: acctest.ErrorCheck(t, servicediscovery.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsServiceDiscoveryPublicDnsNamespaceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hosted_zone", resourceName, "hosted_zone"),
				),
			},
		},
	})
}

func testAccCheckAwsServiceDiscoveryPrivateDnsNamespaceConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_service_discovery_private_dns_namespace" "test" {
  name = "%[1]s.tf"
  vpc  = aws_vpc.test.id
}

data "aws_service_discovery_dns_namespace" "test" {
  name = aws_service_discovery_private_dns_namespace.test.name
  type = "DNS_PRIVATE"
}
`, rName)
}

func testAccCheckAwsServiceDiscoveryPublicDnsNamespaceConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_service_discovery_public_dns_namespace" "test" {
  name = "%[1]s.tf"
}

data "aws_service_discovery_dns_namespace" "test" {
  name = aws_service_discovery_public_dns_namespace.test.name
  type = "DNS_PUBLIC"
}
`, rName)
}
