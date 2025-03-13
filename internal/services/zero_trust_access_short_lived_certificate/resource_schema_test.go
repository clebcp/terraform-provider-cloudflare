// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_short_lived_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustAccessShortLivedCertificateModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_short_lived_certificate.ZeroTrustAccessShortLivedCertificateModel)(nil)
	schema := zero_trust_access_short_lived_certificate.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
