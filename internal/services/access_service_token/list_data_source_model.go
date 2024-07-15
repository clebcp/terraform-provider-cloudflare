// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_service_token

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessServiceTokensResultListDataSourceEnvelope struct {
	Result *[]*AccessServiceTokensItemsDataSourceModel `json:"result,computed"`
}

type AccessServiceTokensDataSourceModel struct {
	AccountID types.String                                `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                 `tfsdk:"max_items"`
	Items     *[]*AccessServiceTokensItemsDataSourceModel `tfsdk:"items"`
}

type AccessServiceTokensItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id"`
	ClientID  types.String      `tfsdk:"client_id" json:"client_id"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at"`
	Duration  types.String      `tfsdk:"duration" json:"duration"`
	Name      types.String      `tfsdk:"name" json:"name"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at"`
}
