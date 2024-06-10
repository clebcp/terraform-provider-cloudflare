// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_application"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_ca_certificate"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_custom_page"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_group"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_identity_provider"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_key"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_mutual_tls_certificate"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_mutual_tls_hostname_settings"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_organization"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_policy"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_service_token"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/access_tag"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/account"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/account_member"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/address_map"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/api_shield"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/api_shield_operation"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/api_shield_operation_schema_validation_settings"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/api_shield_schema_validation_settings"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/api_token"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/argo_smart_routing"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/argo_tiered_caching"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/authenticated_origin_pulls"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/authenticated_origin_pulls_certificate"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/bot_management"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/byo_ip_prefix"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/certificate_pack"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/custom_hostname"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/custom_hostname_fallback_origin"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/custom_ssl"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/d1_database"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/device_dex_test"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/device_managed_networks"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/device_posture_integration"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/device_posture_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/device_settings_policy"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/email_routing_address"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/email_routing_catch_all"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/email_routing_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/email_routing_settings"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/fallback_domain"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/filter"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/firewall_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/gre_tunnel"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/healthcheck"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/hostname_tls_setting"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/hyperdrive_config"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/ipsec_tunnel"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/keyless_certificate"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/list"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/list_item"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/load_balancer"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/load_balancer_monitor"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/load_balancer_pool"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/logpull_retention"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/logpush_job"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/logpush_ownership_challenge"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/managed_headers"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/mtls_certificate"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/notification_policy"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/notification_policy_webhooks"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/observatory_scheduled_test"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/origin_ca_certificate"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/page_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/pages_domain"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/pages_project"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/queue"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/r2_bucket"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/rate_limit"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/record"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/regional_hostname"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/regional_tiered_cache"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/ruleset"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/spectrum_application"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/static_route"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/teams_account"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/teams_list"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/teams_location"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/teams_proxy_endpoint"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/teams_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/tiered_cache"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/total_tls"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/tunnel"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/tunnel_config"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/tunnel_route"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/tunnel_virtual_network"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/turnstile_widget"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/url_normalization_settings"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/user_agent_blocking_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/waiting_room"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/waiting_room_event"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/waiting_room_rules"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/waiting_room_setting"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/web3_hostname"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/web_analytics_rule"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/web_analytics_site"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/worker_cron_trigger"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/worker_domain"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/worker_script"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/worker_secret"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/workers_for_platforms_namespace"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/workers_kv"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/workers_kv_namespace"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/zone"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/zone_cache_reserve"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/zone_cache_variants"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/zone_dnssec"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/zone_hold"
	"github.com/stainless-sdks/cloudflare-terraform/internal/resources/zone_lockdown"
)

var _ provider.Provider = &CloudflareProvider{}

// CloudflareProvider defines the provider implementation.
type CloudflareProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudflareProviderModel describes the provider data model.
type CloudflareProviderModel struct {
	BaseURL        types.String `tfsdk:"base_url" json:"base_url"`
	APIToken       types.String `tfsdk:"api_token" json:"api_token"`
	APIKey         types.String `tfsdk:"api_key" json:"api_key"`
	APIEmail       types.String `tfsdk:"api_email" json:"api_email"`
	UserServiceKey types.String `tfsdk:"user_service_key" json:"user_service_key"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"api_token": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"api_email": schema.StringAttribute{
				Optional: true,
			},
			"user_service_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if !data.APIToken.IsNull() {
		opts = append(opts, option.WithAPIToken(data.APIToken.ValueString()))
	}
	if !data.APIKey.IsNull() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	}
	if !data.APIEmail.IsNull() {
		opts = append(opts, option.WithAPIEmail(data.APIEmail.ValueString()))
	}
	if !data.UserServiceKey.IsNull() {
		opts = append(opts, option.WithUserServiceKey(data.UserServiceKey.ValueString()))
	}

	client := cloudflare.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		account.NewResource,
		account_member.NewResource,
		origin_ca_certificate.NewResource,
		api_token.NewResource,
		zone.NewResource,
		zone_hold.NewResource,
		load_balancer.NewResource,
		load_balancer_monitor.NewResource,
		load_balancer_pool.NewResource,
		zone_cache_reserve.NewResource,
		tiered_cache.NewResource,
		zone_cache_variants.NewResource,
		regional_tiered_cache.NewResource,
		certificate_pack.NewResource,
		total_tls.NewResource,
		argo_smart_routing.NewResource,
		argo_tiered_caching.NewResource,
		custom_ssl.NewResource,
		custom_hostname.NewResource,
		custom_hostname_fallback_origin.NewResource,
		record.NewResource,
		zone_dnssec.NewResource,
		email_routing_settings.NewResource,
		email_routing_rule.NewResource,
		email_routing_catch_all.NewResource,
		email_routing_address.NewResource,
		filter.NewResource,
		zone_lockdown.NewResource,
		firewall_rule.NewResource,
		access_rule.NewResource,
		user_agent_blocking_rule.NewResource,
		healthcheck.NewResource,
		keyless_certificate.NewResource,
		logpush_job.NewResource,
		logpush_ownership_challenge.NewResource,
		logpull_retention.NewResource,
		authenticated_origin_pulls_certificate.NewResource,
		authenticated_origin_pulls.NewResource,
		page_rule.NewResource,
		rate_limit.NewResource,
		waiting_room.NewResource,
		waiting_room_event.NewResource,
		waiting_room_rules.NewResource,
		waiting_room_setting.NewResource,
		web3_hostname.NewResource,
		worker_script.NewResource,
		worker_cron_trigger.NewResource,
		worker_domain.NewResource,
		workers_kv_namespace.NewResource,
		workers_kv.NewResource,
		queue.NewResource,
		api_shield.NewResource,
		api_shield_operation.NewResource,
		api_shield_operation_schema_validation_settings.NewResource,
		api_shield_schema_validation_settings.NewResource,
		managed_headers.NewResource,
		ruleset.NewResource,
		url_normalization_settings.NewResource,
		spectrum_application.NewResource,
		regional_hostname.NewResource,
		address_map.NewResource,
		byo_ip_prefix.NewResource,
		gre_tunnel.NewResource,
		ipsec_tunnel.NewResource,
		static_route.NewResource,
		mtls_certificate.NewResource,
		pages_project.NewResource,
		pages_domain.NewResource,
		list.NewResource,
		list_item.NewResource,
		notification_policy_webhooks.NewResource,
		notification_policy.NewResource,
		d1_database.NewResource,
		r2_bucket.NewResource,
		workers_for_platforms_namespace.NewResource,
		worker_secret.NewResource,
		device_dex_test.NewResource,
		device_managed_networks.NewResource,
		device_settings_policy.NewResource,
		fallback_domain.NewResource,
		device_posture_rule.NewResource,
		device_posture_integration.NewResource,
		access_identity_provider.NewResource,
		access_organization.NewResource,
		access_application.NewResource,
		access_ca_certificate.NewResource,
		access_policy.NewResource,
		access_mutual_tls_certificate.NewResource,
		access_mutual_tls_hostname_settings.NewResource,
		access_group.NewResource,
		access_service_token.NewResource,
		access_key.NewResource,
		access_custom_page.NewResource,
		access_tag.NewResource,
		tunnel.NewResource,
		tunnel_config.NewResource,
		teams_account.NewResource,
		teams_list.NewResource,
		teams_location.NewResource,
		teams_proxy_endpoint.NewResource,
		teams_rule.NewResource,
		tunnel_route.NewResource,
		tunnel_virtual_network.NewResource,
		turnstile_widget.NewResource,
		hyperdrive_config.NewResource,
		web_analytics_site.NewResource,
		web_analytics_rule.NewResource,
		bot_management.NewResource,
		observatory_scheduled_test.NewResource,
		hostname_tls_setting.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
