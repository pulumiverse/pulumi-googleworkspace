// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleWorkspace
{
    /// <summary>
    /// OrgUnit resource manages Google Workspace OrgUnits. Org Unit resides under the `https://www.googleapis.com/auth/admin.directory.orgunit` client scope.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using GoogleWorkspace = Pulumi.GoogleWorkspace;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var org = new GoogleWorkspace.OrgUnit("org", new GoogleWorkspace.OrgUnitArgs
    ///         {
    ///             Description = "paper company",
    ///             ParentOrgUnitPath = "/",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import googleworkspace:index/orgUnit:OrgUnit org "id:01ab2c3d4efg56h"
    /// ```
    /// </summary>
    [GoogleWorkspaceResourceType("googleworkspace:index/orgUnit:OrgUnit")]
    public partial class OrgUnit : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines if a sub-organizational unit can inherit the settings of the parent organization. False means a sub-organizational unit inherits the settings of the nearest parent organizational unit. For more information on inheritance and users in an organization structure, see the [administration help center](https://support.google.com/a/answer/4352075). Defaults to `false`.
        /// </summary>
        [Output("blockInheritance")]
        public Output<bool?> BlockInheritance { get; private set; } = null!;

        /// <summary>
        /// Description of the organizational unit.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// ETag of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        /// <summary>
        /// The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales*support parent path is sales*support.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the organizational unit.
        /// </summary>
        [Output("orgUnitId")]
        public Output<string> OrgUnitId { get; private set; } = null!;

        /// <summary>
        /// The full path to the organizational unit. The orgUnitPath is a derived property. When listed, it is derived from parentOrgunitPath and organizational unit's name. For example, for an organizational unit named 'apps' under parent organization '/engineering', the orgUnitPath is '/engineering/apps'. In order to edit an orgUnitPath, either update the name of the organization or the parentOrgunitPath. A user's organizational unit determines which Google Workspace services the user has access to. If the user is moved to a new organization, the user's access changes. For more information about organization structures, see the [administration help center](https://support.google.com/a/answer/4352075). For more information about moving a user to a different organization, see [chromeosdevices.update a user](https://developers.google.com/admin-sdk/directory/v1/guides/manage-users#update_user).
        /// </summary>
        [Output("orgUnitPath")]
        public Output<string> OrgUnitPath { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the parent organizational unit.
        /// </summary>
        [Output("parentOrgUnitId")]
        public Output<string> ParentOrgUnitId { get; private set; } = null!;

        /// <summary>
        /// The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support organizational unit.
        /// </summary>
        [Output("parentOrgUnitPath")]
        public Output<string> ParentOrgUnitPath { get; private set; } = null!;


        /// <summary>
        /// Create a OrgUnit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgUnit(string name, OrgUnitArgs? args = null, CustomResourceOptions? options = null)
            : base("googleworkspace:index/orgUnit:OrgUnit", name, args ?? new OrgUnitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgUnit(string name, Input<string> id, OrgUnitState? state = null, CustomResourceOptions? options = null)
            : base("googleworkspace:index/orgUnit:OrgUnit", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrgUnit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgUnit Get(string name, Input<string> id, OrgUnitState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgUnit(name, id, state, options);
        }
    }

    public sealed class OrgUnitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if a sub-organizational unit can inherit the settings of the parent organization. False means a sub-organizational unit inherits the settings of the nearest parent organizational unit. For more information on inheritance and users in an organization structure, see the [administration help center](https://support.google.com/a/answer/4352075). Defaults to `false`.
        /// </summary>
        [Input("blockInheritance")]
        public Input<bool>? BlockInheritance { get; set; }

        /// <summary>
        /// Description of the organizational unit.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales*support parent path is sales*support.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique ID of the parent organizational unit.
        /// </summary>
        [Input("parentOrgUnitId")]
        public Input<string>? ParentOrgUnitId { get; set; }

        /// <summary>
        /// The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support organizational unit.
        /// </summary>
        [Input("parentOrgUnitPath")]
        public Input<string>? ParentOrgUnitPath { get; set; }

        public OrgUnitArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class OrgUnitState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if a sub-organizational unit can inherit the settings of the parent organization. False means a sub-organizational unit inherits the settings of the nearest parent organizational unit. For more information on inheritance and users in an organization structure, see the [administration help center](https://support.google.com/a/answer/4352075). Defaults to `false`.
        /// </summary>
        [Input("blockInheritance")]
        public Input<bool>? BlockInheritance { get; set; }

        /// <summary>
        /// Description of the organizational unit.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ETag of the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales*support parent path is sales*support.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique ID of the organizational unit.
        /// </summary>
        [Input("orgUnitId")]
        public Input<string>? OrgUnitId { get; set; }

        /// <summary>
        /// The full path to the organizational unit. The orgUnitPath is a derived property. When listed, it is derived from parentOrgunitPath and organizational unit's name. For example, for an organizational unit named 'apps' under parent organization '/engineering', the orgUnitPath is '/engineering/apps'. In order to edit an orgUnitPath, either update the name of the organization or the parentOrgunitPath. A user's organizational unit determines which Google Workspace services the user has access to. If the user is moved to a new organization, the user's access changes. For more information about organization structures, see the [administration help center](https://support.google.com/a/answer/4352075). For more information about moving a user to a different organization, see [chromeosdevices.update a user](https://developers.google.com/admin-sdk/directory/v1/guides/manage-users#update_user).
        /// </summary>
        [Input("orgUnitPath")]
        public Input<string>? OrgUnitPath { get; set; }

        /// <summary>
        /// The unique ID of the parent organizational unit.
        /// </summary>
        [Input("parentOrgUnitId")]
        public Input<string>? ParentOrgUnitId { get; set; }

        /// <summary>
        /// The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support organizational unit.
        /// </summary>
        [Input("parentOrgUnitPath")]
        public Input<string>? ParentOrgUnitPath { get; set; }

        public OrgUnitState()
        {
            Description = "Managed by Pulumi";
        }
    }
}
