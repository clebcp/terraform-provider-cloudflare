// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessTagResultDataSourceEnvelope struct {
	Result AccessTagDataSourceModel `json:"result,computed"`
}

type AccessTagResultListDataSourceEnvelope struct {
	Result *[]*AccessTagDataSourceModel `json:"result,computed"`
}

type AccessTagDataSourceModel struct {
	AccountID types.String                       `tfsdk:"account_id" path:"account_id"`
	TagName   types.String                       `tfsdk:"tag_name" path:"tag_name"`
	Name      types.String                       `tfsdk:"name" json:"name,computed"`
	AppCount  types.Int64                        `tfsdk:"app_count" json:"app_count"`
	CreatedAt timetypes.RFC3339                  `tfsdk:"created_at" json:"created_at"`
	UpdatedAt timetypes.RFC3339                  `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy *AccessTagFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type AccessTagFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
