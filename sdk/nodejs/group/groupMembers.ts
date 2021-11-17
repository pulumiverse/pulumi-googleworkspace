// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Group Members resource manages Google Workspace Groups Members. Group Members resides under the `https://www.googleapis.com/auth/admin.directory.group` client scope.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as googleWorkspace from "@pulumi/googleWorkspace";
 *
 * const salesGroup = new googleWorkspace.group.Group("salesGroup", {email: "sales@example.com"});
 * const michael = new googleWorkspace.User("michael", {
 *     primaryEmail: "michael.scott@example.com",
 *     password: "34819d7beeabb9260a5c854bc85b3e44",
 *     hashFunction: "MD5",
 *     name: {
 *         familyName: "Scott",
 *         givenName: "Michael",
 *     },
 * });
 * const frank = new googleWorkspace.User("frank", {
 *     primaryEmail: "frank.scott@example.com",
 *     password: "2095312189753de6ad47dfe20cbe97ec",
 *     hashFunction: "MD5",
 *     name: {
 *         familyName: "Scott",
 *         givenName: "Frank",
 *     },
 * });
 * const salesGroupMembers = new googleWorkspace.group.GroupMembers("salesGroupMembers", {
 *     groupId: salesGroup.id,
 *     members: [
 *         {
 *             email: michael.primaryEmail,
 *             role: "MANAGER",
 *         },
 *         {
 *             email: frank.primaryEmail,
 *             role: "MEMBER",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import googleworkspace:group/groupMembers:GroupMembers sales groups/01abcde23fg4h5i
 * ```
 */
export class GroupMembers extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembersState, opts?: pulumi.CustomResourceOptions): GroupMembers {
        return new GroupMembers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'googleworkspace:group/groupMembers:GroupMembers';

    /**
     * Returns true if the given object is an instance of GroupMembers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembers.__pulumiType;
    }

    /**
     * ETag of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The ID of this resource.
     */
    public /*out*/ readonly id!: pulumi.Output<string>;
    /**
     * The members of the group
     */
    public readonly members!: pulumi.Output<outputs.group.GroupMembersMember[] | undefined>;

    /**
     * Create a GroupMembers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMembersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMembersArgs | GroupMembersState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMembersState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["id"] = state ? state.id : undefined;
            inputs["members"] = state ? state.members : undefined;
        } else {
            const args = argsOrState as GroupMembersArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupMembers.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembers resources.
 */
export interface GroupMembersState {
    /**
     * ETag of the resource.
     */
    etag?: pulumi.Input<string>;
    /**
     * Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<string>;
    /**
     * The members of the group
     */
    members?: pulumi.Input<pulumi.Input<inputs.group.GroupMembersMember>[]>;
}

/**
 * The set of arguments for constructing a GroupMembers resource.
 */
export interface GroupMembersArgs {
    /**
     * Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
     */
    groupId: pulumi.Input<string>;
    /**
     * The members of the group
     */
    members?: pulumi.Input<pulumi.Input<inputs.group.GroupMembersMember>[]>;
}
