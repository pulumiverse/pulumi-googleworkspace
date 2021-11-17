# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .chrome_policy import *
from .get_domain import *
from .get_domain_alias import *
from .get_group import *
from .get_group_member import *
from .get_group_members import *
from .get_group_settings import *
from .get_org_unit import *
from .get_privileges import *
from .get_role import *
from .get_schema import *
from .get_user import *
from .gmail_send_as_alias import *
from .org_unit import *
from .provider import *
from .schema import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_googleWorkspace.config as __config
    config = __config
    import pulumi_googleWorkspace.domain as __domain
    domain = __domain
    import pulumi_googleWorkspace.group as __group
    group = __group
    import pulumi_googleWorkspace.role as __role
    role = __role
else:
    config = _utilities.lazy_import('pulumi_googleWorkspace.config')
    domain = _utilities.lazy_import('pulumi_googleWorkspace.domain')
    group = _utilities.lazy_import('pulumi_googleWorkspace.group')
    role = _utilities.lazy_import('pulumi_googleWorkspace.role')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "googleWorkspace",
  "mod": "domain/domain",
  "fqn": "pulumi_googleWorkspace.domain",
  "classes": {
   "googleworkspace:domain/domain:Domain": "Domain"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "domain/domainAlias",
  "fqn": "pulumi_googleWorkspace.domain",
  "classes": {
   "googleworkspace:domain/domainAlias:DomainAlias": "DomainAlias"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "group/group",
  "fqn": "pulumi_googleWorkspace.group",
  "classes": {
   "googleworkspace:group/group:Group": "Group"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "group/groupMember",
  "fqn": "pulumi_googleWorkspace.group",
  "classes": {
   "googleworkspace:group/groupMember:GroupMember": "GroupMember"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "group/groupMembers",
  "fqn": "pulumi_googleWorkspace.group",
  "classes": {
   "googleworkspace:group/groupMembers:GroupMembers": "GroupMembers"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "group/groupSettings",
  "fqn": "pulumi_googleWorkspace.group",
  "classes": {
   "googleworkspace:group/groupSettings:GroupSettings": "GroupSettings"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "index/chromePolicy",
  "fqn": "pulumi_googleWorkspace",
  "classes": {
   "googleworkspace:index/chromePolicy:ChromePolicy": "ChromePolicy"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "index/gmailSendAsAlias",
  "fqn": "pulumi_googleWorkspace",
  "classes": {
   "googleworkspace:index/gmailSendAsAlias:GmailSendAsAlias": "GmailSendAsAlias"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "index/orgUnit",
  "fqn": "pulumi_googleWorkspace",
  "classes": {
   "googleworkspace:index/orgUnit:OrgUnit": "OrgUnit"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "index/schema",
  "fqn": "pulumi_googleWorkspace",
  "classes": {
   "googleworkspace:index/schema:Schema": "Schema"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "index/user",
  "fqn": "pulumi_googleWorkspace",
  "classes": {
   "googleworkspace:index/user:User": "User"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "role/role",
  "fqn": "pulumi_googleWorkspace.role",
  "classes": {
   "googleworkspace:role/role:Role": "Role"
  }
 },
 {
  "pkg": "googleWorkspace",
  "mod": "role/roleAssignment",
  "fqn": "pulumi_googleWorkspace.role",
  "classes": {
   "googleworkspace:role/roleAssignment:RoleAssignment": "RoleAssignment"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "googleWorkspace",
  "token": "pulumi:providers:googleWorkspace",
  "fqn": "pulumi_googleWorkspace",
  "class": "Provider"
 }
]
"""
)
