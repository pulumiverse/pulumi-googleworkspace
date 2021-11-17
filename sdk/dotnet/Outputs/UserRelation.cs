// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleWorkspace.Outputs
{

    [OutputType]
    public sealed class UserRelation
    {
        /// <summary>
        /// If the value of type is custom, this property contains the custom type string.
        /// </summary>
        public readonly string? CustomType;
        /// <summary>
        /// The type of relation. Acceptable values: `admin_assistant`, `assistant`, `brother`, `child`, `custom`, `domestic_partner`, `dotted_line_manager`, `exec_assistant`, `father`, `friend`, `manager`, `mother`, `parent`, `partner`, `referred_by`, `relative`, `sister`, `spouse`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The name of the person the user is related to.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private UserRelation(
            string? customType,

            string type,

            string value)
        {
            CustomType = customType;
            Type = type;
            Value = value;
        }
    }
}
