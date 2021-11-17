// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package googleworkspace

import (
	"fmt"
	"path/filepath"
	"unicode"

	googleWorkspace "github.com/hashicorp/terraform-provider-googleworkspace/shim"
	"github.com/pulumi/pulumi-googleworkspace/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "googleworkspace"

	// modules:
	mainMod   = "index"
	domainMod = "domain"
	groupMod  = "group"
	roleMod   = "role"
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(googleWorkspace.NewProvider(version.Version))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "googleworkspace",
		Description:          "A Pulumi package for creating and managing Google Workspace cloud resources.",
		Keywords:             []string{"pulumi", "google", "workspace", "gsuite"},
		GitHubOrg:            "hashicorp",
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/pulumi/pulumi-googleworkspace",
		PreConfigureCallback: preConfigureCallback,
		Config: map[string]*tfbridge.SchemaInfo{
			"customerId": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GOOGLE_WORKSPACE_CUSTOMER_ID"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"googleworkspace_chrome_policy":       {Tok: makeResource(mainMod, "ChromePolicy")},
			"googleworkspace_domain":              {Tok: makeResource(domainMod, "Domain")},
			"googleworkspace_domain_alias":        {Tok: makeResource(domainMod, "DomainAlias")},
			"googleworkspace_gmail_send_as_alias": {Tok: makeResource(mainMod, "GmailSendAsAlias")},

			"googleworkspace_group": {
				Tok: makeResource(groupMod, "Group"),

				Fields: map[string]*tfbridge.SchemaInfo{
					"description": {
						Default: managedByPulumi,
					},
				},
			},

			"googleworkspace_group_member":   {Tok: makeResource(groupMod, "GroupMember")},
			"googleworkspace_group_members":  {Tok: makeResource(groupMod, "GroupMembers")},
			"googleworkspace_group_settings": {Tok: makeResource(groupMod, "GroupSettings")},

			"googleworkspace_org_unit": {
				Tok: makeResource(mainMod, "OrgUnit"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"description": {
						Default: managedByPulumi,
					},
				},
			},

			"googleworkspace_role":            {Tok: makeResource(roleMod, "Role")},
			"googleworkspace_role_assignment": {Tok: makeResource(roleMod, "RoleAssignment")},
			"googleworkspace_schema":          {Tok: makeResource(mainMod, "Schema")},
			"googleworkspace_user":            {Tok: makeResource(mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"googleworkspace_domain":         {Tok: makeDataSource(mainMod, "getDomain")},
			"googleworkspace_domain_alias":   {Tok: makeDataSource(mainMod, "getDomainAlias")},
			"googleworkspace_group":          {Tok: makeDataSource(mainMod, "getGroup")},
			"googleworkspace_group_member":   {Tok: makeDataSource(mainMod, "getGroupMember")},
			"googleworkspace_group_members":  {Tok: makeDataSource(mainMod, "getGroupMembers")},
			"googleworkspace_group_settings": {Tok: makeDataSource(mainMod, "getGroupSettings")},
			"googleworkspace_org_unit":       {Tok: makeDataSource(mainMod, "getOrgUnit")},
			"googleworkspace_privileges":     {Tok: makeDataSource(mainMod, "getPrivileges")},
			"googleworkspace_role":           {Tok: makeDataSource(mainMod, "getRole")},
			"googleworkspace_schema":         {Tok: makeDataSource(mainMod, "getSchema")},
			"googleworkspace_user":           {Tok: makeDataSource(mainMod, "getUser")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
