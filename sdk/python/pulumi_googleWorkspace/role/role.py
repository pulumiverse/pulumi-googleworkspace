# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 privileges: pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]] privileges: The set of privileges that are granted to this role.
        :param pulumi.Input[str] description: A short description of the role.
        :param pulumi.Input[str] name: Name of the role.
        """
        pulumi.set(__self__, "privileges", privileges)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]]:
        """
        The set of privileges that are granted to this role.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A short description of the role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_super_admin_role: Optional[pulumi.Input[bool]] = None,
                 is_system_role: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[str] description: A short description of the role.
        :param pulumi.Input[str] etag: ETag of the resource.
        :param pulumi.Input[str] id: ID of the role.
        :param pulumi.Input[bool] is_super_admin_role: Returns true if the role is a super admin role.
        :param pulumi.Input[bool] is_system_role: Returns true if this is a pre-defined system role.
        :param pulumi.Input[str] name: Name of the role.
        :param pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]] privileges: The set of privileges that are granted to this role.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_super_admin_role is not None:
            pulumi.set(__self__, "is_super_admin_role", is_super_admin_role)
        if is_system_role is not None:
            pulumi.set(__self__, "is_system_role", is_system_role)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if privileges is not None:
            pulumi.set(__self__, "privileges", privileges)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A short description of the role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the role.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="isSuperAdminRole")
    def is_super_admin_role(self) -> Optional[pulumi.Input[bool]]:
        """
        Returns true if the role is a super admin role.
        """
        return pulumi.get(self, "is_super_admin_role")

    @is_super_admin_role.setter
    def is_super_admin_role(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_super_admin_role", value)

    @property
    @pulumi.getter(name="isSystemRole")
    def is_system_role(self) -> Optional[pulumi.Input[bool]]:
        """
        Returns true if this is a pre-defined system role.
        """
        return pulumi.get(self, "is_system_role")

    @is_system_role.setter
    def is_system_role(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_system_role", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]]]:
        """
        The set of privileges that are granted to this role.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RolePrivilegeArgs']]]]):
        pulumi.set(self, "privileges", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RolePrivilegeArgs']]]]] = None,
                 __props__=None):
        """
        ## Import

        ```sh
         $ pulumi import googleworkspace:role/role:Role admin 12345678901234567
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A short description of the role.
        :param pulumi.Input[str] name: Name of the role.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RolePrivilegeArgs']]]] privileges: The set of privileges that are granted to this role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        ```sh
         $ pulumi import googleworkspace:role/role:Role admin 12345678901234567
        ```

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RolePrivilegeArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if privileges is None and not opts.urn:
                raise TypeError("Missing required property 'privileges'")
            __props__.__dict__["privileges"] = privileges
            __props__.__dict__["etag"] = None
            __props__.__dict__["id"] = None
            __props__.__dict__["is_super_admin_role"] = None
            __props__.__dict__["is_system_role"] = None
        super(Role, __self__).__init__(
            'googleworkspace:role/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            id: Optional[pulumi.Input[str]] = None,
            is_super_admin_role: Optional[pulumi.Input[bool]] = None,
            is_system_role: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RolePrivilegeArgs']]]]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A short description of the role.
        :param pulumi.Input[str] etag: ETag of the resource.
        :param pulumi.Input[str] id: ID of the role.
        :param pulumi.Input[bool] is_super_admin_role: Returns true if the role is a super admin role.
        :param pulumi.Input[bool] is_system_role: Returns true if this is a pre-defined system role.
        :param pulumi.Input[str] name: Name of the role.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RolePrivilegeArgs']]]] privileges: The set of privileges that are granted to this role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["etag"] = etag
        __props__.__dict__["id"] = id
        __props__.__dict__["is_super_admin_role"] = is_super_admin_role
        __props__.__dict__["is_system_role"] = is_system_role
        __props__.__dict__["name"] = name
        __props__.__dict__["privileges"] = privileges
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A short description of the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        ID of the role.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isSuperAdminRole")
    def is_super_admin_role(self) -> pulumi.Output[bool]:
        """
        Returns true if the role is a super admin role.
        """
        return pulumi.get(self, "is_super_admin_role")

    @property
    @pulumi.getter(name="isSystemRole")
    def is_system_role(self) -> pulumi.Output[bool]:
        """
        Returns true if this is a pre-defined system role.
        """
        return pulumi.get(self, "is_system_role")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Output[Sequence['outputs.RolePrivilege']]:
        """
        The set of privileges that are granted to this role.
        """
        return pulumi.get(self, "privileges")

