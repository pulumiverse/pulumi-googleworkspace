// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleWorkspace.Domain
{
    /// <summary>
    /// Domain Alias resource manages Google Workspace Domain Aliases. Domain Alias resides under the `https://www.googleapis.com/auth/admin.directory.domain` client scope.
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
    ///         var example = new GoogleWorkspace.Domain.DomainAlias("example", new GoogleWorkspace.Domain.DomainAliasArgs
    ///         {
    ///             DomainAliasName = "alias-example.com",
    ///             ParentDomainName = "example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import googleworkspace:domain/domainAlias:DomainAlias example alias-example.com
    /// ```
    /// </summary>
    [GoogleWorkspaceResourceType("googleworkspace:domain/domainAlias:DomainAlias")]
    public partial class DomainAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of the domain alias.
        /// </summary>
        [Output("creationTime")]
        public Output<int> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The domain alias name.
        /// </summary>
        [Output("domainAliasName")]
        public Output<string> DomainAliasName { get; private set; } = null!;

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
        /// The parent domain name that the domain alias is associated with. This can either be a primary or secondary domain name within a customer.
        /// </summary>
        [Output("parentDomainName")]
        public Output<string?> ParentDomainName { get; private set; } = null!;

        /// <summary>
        /// Indicates the verification state of a domain alias.
        /// </summary>
        [Output("verified")]
        public Output<bool> Verified { get; private set; } = null!;


        /// <summary>
        /// Create a DomainAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainAlias(string name, DomainAliasArgs args, CustomResourceOptions? options = null)
            : base("googleworkspace:domain/domainAlias:DomainAlias", name, args ?? new DomainAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainAlias(string name, Input<string> id, DomainAliasState? state = null, CustomResourceOptions? options = null)
            : base("googleworkspace:domain/domainAlias:DomainAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainAlias Get(string name, Input<string> id, DomainAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainAlias(name, id, state, options);
        }
    }

    public sealed class DomainAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain alias name.
        /// </summary>
        [Input("domainAliasName", required: true)]
        public Input<string> DomainAliasName { get; set; } = null!;

        /// <summary>
        /// The parent domain name that the domain alias is associated with. This can either be a primary or secondary domain name within a customer.
        /// </summary>
        [Input("parentDomainName")]
        public Input<string>? ParentDomainName { get; set; }

        public DomainAliasArgs()
        {
        }
    }

    public sealed class DomainAliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time of the domain alias.
        /// </summary>
        [Input("creationTime")]
        public Input<int>? CreationTime { get; set; }

        /// <summary>
        /// The domain alias name.
        /// </summary>
        [Input("domainAliasName")]
        public Input<string>? DomainAliasName { get; set; }

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
        /// The parent domain name that the domain alias is associated with. This can either be a primary or secondary domain name within a customer.
        /// </summary>
        [Input("parentDomainName")]
        public Input<string>? ParentDomainName { get; set; }

        /// <summary>
        /// Indicates the verification state of a domain alias.
        /// </summary>
        [Input("verified")]
        public Input<bool>? Verified { get; set; }

        public DomainAliasState()
        {
        }
    }
}
