// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class BranchProtectionAllowedToUnprotectArgs : Pulumi.ResourceArgs
    {
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        [Input("accessLevelDescription")]
        public Input<string>? AccessLevelDescription { get; set; }

        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public BranchProtectionAllowedToUnprotectArgs()
        {
        }
    }
}