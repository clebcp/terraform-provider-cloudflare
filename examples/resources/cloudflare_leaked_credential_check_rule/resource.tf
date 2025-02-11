resource "cloudflare_leaked_credential_check_rule" "example_leaked_credential_check_rule" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "18a14bafaa8eb1df04ce683ec18c765e"
  password = "lookup_json_string(http.request.body.raw, \"secret\")"
  username = "lookup_json_string(http.request.body.raw, \"user\")"
}
