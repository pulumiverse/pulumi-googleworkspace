// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Domain resource manages Google Workspace Domains. Domain resides under the `https://www.googleapis.com/auth/admin.directory.domain` client scope.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as googleworkspace from "@pulumi/googleworkspace";
 *
 * const example = new googleworkspace.domain.Domain("example", {
 *     domainName: "example.com",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import googleworkspace:domain/domain:Domain example example.com
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'googleworkspace:domain/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Creation time of the domain. Expressed in Unix time format.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<number>;
    /**
     * asps.list of domain alias objects.
     */
    public /*out*/ readonly domainAliases!: pulumi.Output<string[]>;
    /**
     * The domain name of the customer.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * ETag of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The ID of this resource.
     */
    public /*out*/ readonly id!: pulumi.Output<string>;
    /**
     * Indicates if the domain is a primary domain.
     */
    public /*out*/ readonly isPrimary!: pulumi.Output<boolean>;
    /**
     * Indicates the verification state of a domain.
     */
    public /*out*/ readonly verified!: pulumi.Output<boolean>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["domainAliases"] = state ? state.domainAliases : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["id"] = state ? state.id : undefined;
            inputs["isPrimary"] = state ? state.isPrimary : undefined;
            inputs["verified"] = state ? state.verified : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["domainAliases"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["isPrimary"] = undefined /*out*/;
            inputs["verified"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Domain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * Creation time of the domain. Expressed in Unix time format.
     */
    creationTime?: pulumi.Input<number>;
    /**
     * asps.list of domain alias objects.
     */
    domainAliases?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The domain name of the customer.
     */
    domainName?: pulumi.Input<string>;
    /**
     * ETag of the resource.
     */
    etag?: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<string>;
    /**
     * Indicates if the domain is a primary domain.
     */
    isPrimary?: pulumi.Input<boolean>;
    /**
     * Indicates the verification state of a domain.
     */
    verified?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The domain name of the customer.
     */
    domainName: pulumi.Input<string>;
}
