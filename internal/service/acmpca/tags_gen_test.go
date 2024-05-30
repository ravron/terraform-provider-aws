// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package acmpca_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	tfacmpca "github.com/hashicorp/terraform-provider-aws/internal/service/acmpca"
)

func expectFullTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullTags(tfacmpca.ServicePackage(context.Background()), resourceAddress, knownValue)
}
