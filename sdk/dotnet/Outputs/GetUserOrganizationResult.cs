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
    public sealed class GetUserOrganizationResult
    {
        public readonly string CostCenter;
        public readonly string CustomType;
        public readonly string Department;
        public readonly string Description;
        public readonly string Domain;
        public readonly int FullTimeEquivalent;
        public readonly string Location;
        /// <summary>
        /// Holds the given and family names of the user, and the read-only fullName value. The maximum number of characters in the givenName and in the familyName values is 60. In addition, name values support unicode/UTF-8 characters, and can contain spaces, letters (a-z), numbers (0-9), dashes (-), forward slashes (/), and periods (.). Maximum allowed data size for this field is 1Kb.
        /// </summary>
        public readonly string Name;
        public readonly bool Primary;
        public readonly string Symbol;
        public readonly string Title;
        public readonly string Type;

        [OutputConstructor]
        private GetUserOrganizationResult(
            string costCenter,

            string customType,

            string department,

            string description,

            string domain,

            int fullTimeEquivalent,

            string location,

            string name,

            bool primary,

            string symbol,

            string title,

            string type)
        {
            CostCenter = costCenter;
            CustomType = customType;
            Department = department;
            Description = description;
            Domain = domain;
            FullTimeEquivalent = fullTimeEquivalent;
            Location = location;
            Name = name;
            Primary = primary;
            Symbol = symbol;
            Title = title;
            Type = type;
        }
    }
}
