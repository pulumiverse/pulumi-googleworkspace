// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getGroupSettings(args: GetGroupSettingsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupSettingsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("googleworkspace:index/getGroupSettings:getGroupSettings", {
        "email": args.email,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupSettings.
 */
export interface GetGroupSettingsArgs {
    /**
     * The group's email address.
     */
    email: string;
}

/**
 * A collection of values returned by getGroupSettings.
 */
export interface GetGroupSettingsResult {
    /**
     * Identifies whether members external to your organization can join the group. If true, Google Workspace users external to your organization can become members of this group. If false, users not belonging to the organization are not allowed to become members of this group.
     */
    readonly allowExternalMembers: boolean;
    /**
     * Allows posting from web. If true, allows any member to post to the group forum. If false, Members only use Gmail to communicate with the group.
     */
    readonly allowWebPosting: boolean;
    /**
     * Allows the group to be archived only. If true, Group is archived and the group is inactive. New messages to this group are rejected. The older archived messages are browsable and searchable. If true, the `whoCanPostMessage` property is set to `NONE_CAN_POST`. If reverted from true to false, `whoCanPostMessage` is set to `ALL_MANAGERS_CAN_POST`. If false, The group is active and can receive messages. When false, updating `whoCanPostMessage` to `NONE_CAN_POST`, results in an error.
     */
    readonly archiveOnly: boolean;
    /**
     * Set the content of custom footer text. The maximum number of characters is 1,000.
     */
    readonly customFooterText: string;
    /**
     * An email address used when replying to a message if the `replyTo` property is set to `REPLY_TO_CUSTOM`. This address is defined by an account administrator. When the group's `replyTo` property is set to `REPLY_TO_CUSTOM`, the `customReplyTo` property holds a custom email address used when replying to a message, the `customReplyTo` property must have a text value or an error is returned.
     */
    readonly customReplyTo: string;
    /**
     * Specifies whether the group has a custom role that's included in one of the settings being merged.
     */
    readonly customRolesEnabledForSettingsToBeMerged: boolean;
    /**
     * When a message is rejected, this is text for the rejection notification sent to the message's author. By default, this property is empty and has no value in the API's response body. The maximum notification text size is 10,000 characters. Requires `sendMessageDenyNotification` property to be true.
     */
    readonly defaultMessageDenyNotificationText: string;
    /**
     * Description of the group. The maximum group description is no more than 300 characters.
     */
    readonly description: string;
    /**
     * The group's email address.
     */
    readonly email: string;
    /**
     * Specifies whether a collaborative inbox will remain turned on for the group.
     */
    readonly enableCollaborativeInbox: boolean;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * Whether to include custom footer.
     */
    readonly includeCustomFooter: boolean;
    /**
     * Enables the group to be included in the Global Address List. If true, the group is included in the Global Address List. If false, it is not included in the Global Address List.
     */
    readonly includeInGlobalAddressList: boolean;
    /**
     * Allows the Group contents to be archived. If true, archive messages sent to the group. If false, Do not keep an archive of messages sent to this group. If false, previously archived messages remain in the archive.
     */
    readonly isArchived: boolean;
    /**
     * Enables members to post messages as the group. If true, group member can post messages using the group's email address instead of their own email address. Message appear to originate from the group itself. Any message moderation settings on individual users or new members do not apply to posts made on behalf of the group. If false, members can not post in behalf of the group's email address.
     */
    readonly membersCanPostAsTheGroup: boolean;
    /**
     * Moderation level of incoming messages. Possible values are: `MODERATE_ALL_MESSAGES`: All messages are sent to the group owner's email address for approval. If approved, the message is sent to the group. `MODERATE_NON_MEMBERS`: All messages from non group members are sent to the group owner's email address for approval. If approved, the message is sent to the group. `MODERATE_NEW_MEMBERS`: All messages from new members are sent to the group owner's email address for approval. If approved, the message is sent to the group. `MODERATE_NONE`: No moderator approval is required. Messages are delivered directly to the group.Note: When the `whoCanPostMessage` is set to `ANYONE_CAN_POST`, we recommend the `messageModerationLevel` be set to `MODERATE_NON_MEMBERS` to protect the group from possible spam.When `memberCanPostAsTheGroup` is true, any message moderation settings on individual users or new members will not apply to posts made on behalf of the group.
     */
    readonly messageModerationLevel: string;
    /**
     * Name of the group, which has a maximum size of 75 characters.
     */
    readonly name: string;
    /**
     * The primary language for group. For a group's primary language use the language tags from the Google Workspace languages found at Google Workspace Email Settings API Email Language Tags.
     */
    readonly primaryLanguage: string;
    /**
     * Specifies who receives the default reply. Possible values are: `REPLY_TO_CUSTOM`: For replies to messages, use the group's custom email address. When set to `REPLY_TO_CUSTOM`, the `customReplyTo` property holds the custom email address used when replying to a message, the customReplyTo property must have a value. Otherwise an error is returned. `REPLY_TO_SENDER`: The reply sent to author of message. `REPLY_TO_LIST`: This reply message is sent to the group. `REPLY_TO_OWNER`: The reply is sent to the owner(s) of the group. This does not include the group's managers. `REPLY_TO_IGNORE`: Group users individually decide where the message reply is sent. `REPLY_TO_MANAGERS`: This reply message is sent to the group's managers, which includes all managers and the group owner.
     */
    readonly replyTo: string;
    /**
     * Allows a member to be notified if the member's message to the group is denied by the group owner. If true, when a message is rejected, send the deny message notification to the message author. The `defaultMessageDenyNotificationText` property is dependent on the `sendMessageDenyNotification` property being true. If false, when a message is rejected, no notification is sent.
     */
    readonly sendMessageDenyNotification: boolean;
    /**
     * Specifies moderation levels for messages detected as spam. Possible values are: `ALLOW`: Post the message to the group. `MODERATE`: Send the message to the moderation queue. This is the default. `SILENTLY_MODERATE`: Send the message to the moderation queue, but do not send notification to moderators. `REJECT`: Immediately reject the message.
     */
    readonly spamModerationLevel: string;
    /**
     * Specifies who can moderate metadata. Possible values are: `ALL_MEMBERS`, `OWNERS_AND_MANAGERS`, `MANAGERS_ONLY`, `OWNERS_ONLY`, `NONE`
     */
    readonly whoCanAssistContent: string;
    /**
     * Permission to contact owner of the group via web UI. Possible values are: `ALL_IN_DOMAIN_CAN_CONTACT`, `ALL_MANAGERS_CAN_CONTACT`, `ALL_MEMBERS_CAN_CONTACT`, `ANYONE_CAN_CONTACT`
     */
    readonly whoCanContactOwner: string;
    /**
     * Specifies the set of users for whom this group is discoverable. Possible values are: `ANYONE_CAN_DISCOVER`, `ALL_IN_DOMAIN_CAN_DISCOVER`, `ALL_MEMBERS_CAN_DISCOVER`
     */
    readonly whoCanDiscoverGroup: string;
    /**
     * Permission to join group. Possible values are: `ANYONE_CAN_JOIN`: Any Internet user, both inside and outside your domain, can join the group. `ALL_IN_DOMAIN_CAN_JOIN`: Anyone in the account domain can join. This includes accounts with multiple domains. `INVITED_CAN_JOIN`: Candidates for membership can be invited to join. `CAN_REQUEST_TO_JOIN`: Non members can request an invitation to join.
     */
    readonly whoCanJoin: string;
    /**
     * Permission to leave the group. Possible values are: `ALL_MANAGERS_CAN_LEAVE`, `ALL_MEMBERS_CAN_LEAVE`, `NONE_CAN_LEAVE`
     */
    readonly whoCanLeaveGroup: string;
    /**
     * Specifies who can moderate content. Possible values are: `ALL_MEMBERS`, `OWNERS_AND_MANAGERS`, `OWNERS_ONLY`, `NONE`
     */
    readonly whoCanModerateContent: string;
    /**
     * Specifies who can manage members. Possible values are: `ALL_MEMBERS`, `OWNERS_AND_MANAGERS`, `OWNERS_ONLY`, `NONE`
     */
    readonly whoCanModerateMembers: string;
    /**
     * Permissions to post messages. Possible values are: `NONE_CAN_POST`: The group is disabled and archived. No one can post a message to this group. * When archiveOnly is false, updating whoCanPostMessage to NONE*CAN*POST, results in an error. * If archiveOnly is reverted from true to false, whoCanPostMessages is set to ALL*MANAGERS*CAN_POST. `ALL_MANAGERS_CAN_POST`: Managers, including group owners, can post messages. `ALL_MEMBERS_CAN_POST`: Any group member can post a message. `ALL_OWNERS_CAN_POST`: Only group owners can post a message. `ALL_IN_DOMAIN_CAN_POST`: Anyone in the account can post a message. `ANYONE_CAN_POST`: Any Internet user who outside your account can access your Google Groups service and post a message. *Note: When `whoCanPostMessage` is set to `ANYONE_CAN_POST`, we recommend the`messageModerationLevel` be set to `MODERATE_NON_MEMBERS` to protect the group from possible spam. Users not belonging to the organization are not allowed to become members of this group.
     */
    readonly whoCanPostMessage: string;
    /**
     * Permissions to view group messages. Possible values are: `ANYONE_CAN_VIEW`: Any Internet user can view the group's messages. `ALL_IN_DOMAIN_CAN_VIEW`: Anyone in your account can view this group's messages. `ALL_MEMBERS_CAN_VIEW`: All group members can view the group's messages. `ALL_MANAGERS_CAN_VIEW`: Any group manager can view this group's messages.
     */
    readonly whoCanViewGroup: string;
    /**
     * Permissions to view membership. Possible values are: `ALL_IN_DOMAIN_CAN_VIEW`: Anyone in the account can view the group members list. If a group already has external members, those members can still send email to this group. `ALL_MEMBERS_CAN_VIEW`: The group members can view the group members list. `ALL_MANAGERS_CAN_VIEW`: The group managers can view group members list.
     */
    readonly whoCanViewMembership: string;
}

export function getGroupSettingsOutput(args: GetGroupSettingsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupSettingsResult> {
    return pulumi.output(args).apply(a => getGroupSettings(a, opts))
}

/**
 * A collection of arguments for invoking getGroupSettings.
 */
export interface GetGroupSettingsOutputArgs {
    /**
     * The group's email address.
     */
    email: pulumi.Input<string>;
}
