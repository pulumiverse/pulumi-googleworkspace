// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * ```sh
 *  $ pulumi import googleworkspace:role/role:Role admin 12345678901234567
 * ```
 */
export class Role extends pulumi.CustomResource {
    /**
     * Get an existing Role resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleState, opts?: pulumi.CustomResourceOptions): Role {
        return new Role(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'googleworkspace:role/role:Role';

    /**
     * Returns true if the given object is an instance of Role.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Role {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Role.__pulumiType;
    }

    /**
     * A short description of the role.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ETag of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * ID of the role.
     */
    public /*out*/ readonly id!: pulumi.Output<string>;
    /**
     * Returns true if the role is a super admin role.
     */
    public /*out*/ readonly isSuperAdminRole!: pulumi.Output<boolean>;
    /**
     * Returns true if this is a pre-defined system role.
     */
    public /*out*/ readonly isSystemRole!: pulumi.Output<boolean>;
    /**
     * Name of the role.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The set of privileges that are granted to this role.
     */
    public readonly privileges!: pulumi.Output<outputs.role.RolePrivilege[]>;

    /**
     * Create a Role resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleArgs | RoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["id"] = state ? state.id : undefined;
            inputs["isSuperAdminRole"] = state ? state.isSuperAdminRole : undefined;
            inputs["isSystemRole"] = state ? state.isSystemRole : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privileges"] = state ? state.privileges : undefined;
        } else {
            const args = argsOrState as RoleArgs | undefined;
            if ((!args || args.privileges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privileges'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privileges"] = args ? args.privileges : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["isSuperAdminRole"] = undefined /*out*/;
            inputs["isSystemRole"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Role.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Role resources.
 */
export interface RoleState {
    /**
     * A short description of the role.
     */
    description?: pulumi.Input<string>;
    /**
     * ETag of the resource.
     */
    etag?: pulumi.Input<string>;
    /**
     * ID of the role.
     */
    id?: pulumi.Input<string>;
    /**
     * Returns true if the role is a super admin role.
     */
    isSuperAdminRole?: pulumi.Input<boolean>;
    /**
     * Returns true if this is a pre-defined system role.
     */
    isSystemRole?: pulumi.Input<boolean>;
    /**
     * Name of the role.
     */
    name?: pulumi.Input<string>;
    /**
     * The set of privileges that are granted to this role.
     */
    privileges?: pulumi.Input<pulumi.Input<inputs.role.RolePrivilege>[]>;
}

/**
 * The set of arguments for constructing a Role resource.
 */
export interface RoleArgs {
    /**
     * A short description of the role.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    name?: pulumi.Input<string>;
    /**
     * The set of privileges that are granted to this role.
     */
    privileges: pulumi.Input<pulumi.Input<inputs.role.RolePrivilege>[]>;
}
