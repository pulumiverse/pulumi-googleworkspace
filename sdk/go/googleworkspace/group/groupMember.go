// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package group

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Group Member resource manages Google Workspace Groups Members. Group Member resides under the `https://www.googleapis.com/auth/admin.directory.group` client scope.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-googleWorkspace/sdk/go/googleWorkspace"
// 	"github.com/pulumi/pulumi-googleworkspace/sdk/go/googleworkspace/group"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sales, err := group.NewGroup(ctx, "sales", &group.GroupArgs{
// 			Email: pulumi.String("sales@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		michael, err := googleWorkspace.NewUser(ctx, "michael", &googleWorkspace.UserArgs{
// 			PrimaryEmail: pulumi.String("michael.scott@example.com"),
// 			Password:     pulumi.String("34819d7beeabb9260a5c854bc85b3e44"),
// 			HashFunction: pulumi.String("MD5"),
// 			Name: &UserNameArgs{
// 				FamilyName: pulumi.String("Scott"),
// 				GivenName:  pulumi.String("Michael"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = group.NewGroupMember(ctx, "manager", &group.GroupMemberArgs{
// 			GroupId: sales.ID(),
// 			Email:   michael.PrimaryEmail,
// 			Role:    pulumi.String("MANAGER"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import googleworkspace:group/groupMember:GroupMember manager groups/01abcde23fg4h5i/members/123456789012345678901
// ```
type GroupMember struct {
	pulumi.CustomResourceState

	// Defines mail delivery preferences of member. Acceptable values are:`ALL_MAIL`: All messages, delivered as soon as they arrive. `DAILY`: No more than one message a day. `DIGEST`: Up to 25 messages bundled into a single message. `DISABLED`: Remove subscription. `NONE`: No messages. Defaults to `ALL_MAIL`.
	DeliverySettings pulumi.StringPtrOutput `pulumi:"deliverySettings"`
	// The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Email pulumi.StringOutput `pulumi:"email"`
	// ETag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The ID of this resource.
	Id pulumi.StringOutput `pulumi:"id"`
	// The unique ID of the group member. A member id can be used as a member request URI's memberKey.
	MemberId pulumi.StringOutput `pulumi:"memberId"`
	// The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are: `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members. `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list. `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER. Defaults to `MEMBER`.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// Status of member.
	Status pulumi.StringOutput `pulumi:"status"`
	// The type of group member. Acceptable values are: `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID. `GROUP`: The member is another group. `USER`: The member is a user. Defaults to `USER`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewGroupMember registers a new resource with the given unique name, arguments, and options.
func NewGroupMember(ctx *pulumi.Context,
	name string, args *GroupMemberArgs, opts ...pulumi.ResourceOption) (*GroupMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	var resource GroupMember
	err := ctx.RegisterResource("googleworkspace:group/groupMember:GroupMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMember gets an existing GroupMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMemberState, opts ...pulumi.ResourceOption) (*GroupMember, error) {
	var resource GroupMember
	err := ctx.ReadResource("googleworkspace:group/groupMember:GroupMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMember resources.
type groupMemberState struct {
	// Defines mail delivery preferences of member. Acceptable values are:`ALL_MAIL`: All messages, delivered as soon as they arrive. `DAILY`: No more than one message a day. `DIGEST`: Up to 25 messages bundled into a single message. `DISABLED`: Remove subscription. `NONE`: No messages. Defaults to `ALL_MAIL`.
	DeliverySettings *string `pulumi:"deliverySettings"`
	// The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Email *string `pulumi:"email"`
	// ETag of the resource.
	Etag *string `pulumi:"etag"`
	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	GroupId *string `pulumi:"groupId"`
	// The ID of this resource.
	Id *string `pulumi:"id"`
	// The unique ID of the group member. A member id can be used as a member request URI's memberKey.
	MemberId *string `pulumi:"memberId"`
	// The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are: `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members. `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list. `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER. Defaults to `MEMBER`.
	Role *string `pulumi:"role"`
	// Status of member.
	Status *string `pulumi:"status"`
	// The type of group member. Acceptable values are: `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID. `GROUP`: The member is another group. `USER`: The member is a user. Defaults to `USER`.
	Type *string `pulumi:"type"`
}

type GroupMemberState struct {
	// Defines mail delivery preferences of member. Acceptable values are:`ALL_MAIL`: All messages, delivered as soon as they arrive. `DAILY`: No more than one message a day. `DIGEST`: Up to 25 messages bundled into a single message. `DISABLED`: Remove subscription. `NONE`: No messages. Defaults to `ALL_MAIL`.
	DeliverySettings pulumi.StringPtrInput
	// The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Email pulumi.StringPtrInput
	// ETag of the resource.
	Etag pulumi.StringPtrInput
	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	GroupId pulumi.StringPtrInput
	// The ID of this resource.
	Id pulumi.StringPtrInput
	// The unique ID of the group member. A member id can be used as a member request URI's memberKey.
	MemberId pulumi.StringPtrInput
	// The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are: `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members. `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list. `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER. Defaults to `MEMBER`.
	Role pulumi.StringPtrInput
	// Status of member.
	Status pulumi.StringPtrInput
	// The type of group member. Acceptable values are: `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID. `GROUP`: The member is another group. `USER`: The member is a user. Defaults to `USER`.
	Type pulumi.StringPtrInput
}

func (GroupMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberState)(nil)).Elem()
}

type groupMemberArgs struct {
	// Defines mail delivery preferences of member. Acceptable values are:`ALL_MAIL`: All messages, delivered as soon as they arrive. `DAILY`: No more than one message a day. `DIGEST`: Up to 25 messages bundled into a single message. `DISABLED`: Remove subscription. `NONE`: No messages. Defaults to `ALL_MAIL`.
	DeliverySettings *string `pulumi:"deliverySettings"`
	// The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Email string `pulumi:"email"`
	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	GroupId string `pulumi:"groupId"`
	// The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are: `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members. `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list. `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER. Defaults to `MEMBER`.
	Role *string `pulumi:"role"`
	// The type of group member. Acceptable values are: `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID. `GROUP`: The member is another group. `USER`: The member is a user. Defaults to `USER`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a GroupMember resource.
type GroupMemberArgs struct {
	// Defines mail delivery preferences of member. Acceptable values are:`ALL_MAIL`: All messages, delivered as soon as they arrive. `DAILY`: No more than one message a day. `DIGEST`: Up to 25 messages bundled into a single message. `DISABLED`: Remove subscription. `NONE`: No messages. Defaults to `ALL_MAIL`.
	DeliverySettings pulumi.StringPtrInput
	// The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The email must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Email pulumi.StringInput
	// Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.
	GroupId pulumi.StringInput
	// The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. Acceptable values are: `MANAGER`: This role is only available if the Google Groups for Business is enabled using the Admin Console. A `MANAGER` role can do everything done by an `OWNER` role except make a member an `OWNER` or delete the group. A group can have multiple `MANAGER` members. `MEMBER`: This role can subscribe to a group, view discussion archives, and view the group's membership list. `OWNER`: This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER. Defaults to `MEMBER`.
	Role pulumi.StringPtrInput
	// The type of group member. Acceptable values are: `CUSTOMER`: The member represents all users in a domain. An email address is not returned and the ID returned is the customer ID. `GROUP`: The member is another group. `USER`: The member is a user. Defaults to `USER`.
	Type pulumi.StringPtrInput
}

func (GroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberArgs)(nil)).Elem()
}

type GroupMemberInput interface {
	pulumi.Input

	ToGroupMemberOutput() GroupMemberOutput
	ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput
}

func (*GroupMember) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMember)(nil))
}

func (i *GroupMember) ToGroupMemberOutput() GroupMemberOutput {
	return i.ToGroupMemberOutputWithContext(context.Background())
}

func (i *GroupMember) ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberOutput)
}

func (i *GroupMember) ToGroupMemberPtrOutput() GroupMemberPtrOutput {
	return i.ToGroupMemberPtrOutputWithContext(context.Background())
}

func (i *GroupMember) ToGroupMemberPtrOutputWithContext(ctx context.Context) GroupMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberPtrOutput)
}

type GroupMemberPtrInput interface {
	pulumi.Input

	ToGroupMemberPtrOutput() GroupMemberPtrOutput
	ToGroupMemberPtrOutputWithContext(ctx context.Context) GroupMemberPtrOutput
}

type groupMemberPtrType GroupMemberArgs

func (*groupMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMember)(nil))
}

func (i *groupMemberPtrType) ToGroupMemberPtrOutput() GroupMemberPtrOutput {
	return i.ToGroupMemberPtrOutputWithContext(context.Background())
}

func (i *groupMemberPtrType) ToGroupMemberPtrOutputWithContext(ctx context.Context) GroupMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberPtrOutput)
}

// GroupMemberArrayInput is an input type that accepts GroupMemberArray and GroupMemberArrayOutput values.
// You can construct a concrete instance of `GroupMemberArrayInput` via:
//
//          GroupMemberArray{ GroupMemberArgs{...} }
type GroupMemberArrayInput interface {
	pulumi.Input

	ToGroupMemberArrayOutput() GroupMemberArrayOutput
	ToGroupMemberArrayOutputWithContext(context.Context) GroupMemberArrayOutput
}

type GroupMemberArray []GroupMemberInput

func (GroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMember)(nil)).Elem()
}

func (i GroupMemberArray) ToGroupMemberArrayOutput() GroupMemberArrayOutput {
	return i.ToGroupMemberArrayOutputWithContext(context.Background())
}

func (i GroupMemberArray) ToGroupMemberArrayOutputWithContext(ctx context.Context) GroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberArrayOutput)
}

// GroupMemberMapInput is an input type that accepts GroupMemberMap and GroupMemberMapOutput values.
// You can construct a concrete instance of `GroupMemberMapInput` via:
//
//          GroupMemberMap{ "key": GroupMemberArgs{...} }
type GroupMemberMapInput interface {
	pulumi.Input

	ToGroupMemberMapOutput() GroupMemberMapOutput
	ToGroupMemberMapOutputWithContext(context.Context) GroupMemberMapOutput
}

type GroupMemberMap map[string]GroupMemberInput

func (GroupMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMember)(nil)).Elem()
}

func (i GroupMemberMap) ToGroupMemberMapOutput() GroupMemberMapOutput {
	return i.ToGroupMemberMapOutputWithContext(context.Background())
}

func (i GroupMemberMap) ToGroupMemberMapOutputWithContext(ctx context.Context) GroupMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberMapOutput)
}

type GroupMemberOutput struct{ *pulumi.OutputState }

func (GroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMember)(nil))
}

func (o GroupMemberOutput) ToGroupMemberOutput() GroupMemberOutput {
	return o
}

func (o GroupMemberOutput) ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput {
	return o
}

func (o GroupMemberOutput) ToGroupMemberPtrOutput() GroupMemberPtrOutput {
	return o.ToGroupMemberPtrOutputWithContext(context.Background())
}

func (o GroupMemberOutput) ToGroupMemberPtrOutputWithContext(ctx context.Context) GroupMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupMember) *GroupMember {
		return &v
	}).(GroupMemberPtrOutput)
}

type GroupMemberPtrOutput struct{ *pulumi.OutputState }

func (GroupMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMember)(nil))
}

func (o GroupMemberPtrOutput) ToGroupMemberPtrOutput() GroupMemberPtrOutput {
	return o
}

func (o GroupMemberPtrOutput) ToGroupMemberPtrOutputWithContext(ctx context.Context) GroupMemberPtrOutput {
	return o
}

func (o GroupMemberPtrOutput) Elem() GroupMemberOutput {
	return o.ApplyT(func(v *GroupMember) GroupMember {
		if v != nil {
			return *v
		}
		var ret GroupMember
		return ret
	}).(GroupMemberOutput)
}

type GroupMemberArrayOutput struct{ *pulumi.OutputState }

func (GroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMember)(nil))
}

func (o GroupMemberArrayOutput) ToGroupMemberArrayOutput() GroupMemberArrayOutput {
	return o
}

func (o GroupMemberArrayOutput) ToGroupMemberArrayOutputWithContext(ctx context.Context) GroupMemberArrayOutput {
	return o
}

func (o GroupMemberArrayOutput) Index(i pulumi.IntInput) GroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupMember {
		return vs[0].([]GroupMember)[vs[1].(int)]
	}).(GroupMemberOutput)
}

type GroupMemberMapOutput struct{ *pulumi.OutputState }

func (GroupMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupMember)(nil))
}

func (o GroupMemberMapOutput) ToGroupMemberMapOutput() GroupMemberMapOutput {
	return o
}

func (o GroupMemberMapOutput) ToGroupMemberMapOutputWithContext(ctx context.Context) GroupMemberMapOutput {
	return o
}

func (o GroupMemberMapOutput) MapIndex(k pulumi.StringInput) GroupMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupMember {
		return vs[0].(map[string]GroupMember)[vs[1].(string)]
	}).(GroupMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupMemberOutput{})
	pulumi.RegisterOutputType(GroupMemberPtrOutput{})
	pulumi.RegisterOutputType(GroupMemberArrayOutput{})
	pulumi.RegisterOutputType(GroupMemberMapOutput{})
}
