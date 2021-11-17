module github.com/pulumi/pulumi-googleworkspace/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/hashicorp/terraform-provider-googleworkspace/shim => ./shim
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1 // indirect
	github.com/hashicorp/terraform-provider-googleworkspace/shim v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.9.0
	github.com/pulumi/pulumi/sdk/v3 v3.14.1-0.20211007222624-789e39219452
)
