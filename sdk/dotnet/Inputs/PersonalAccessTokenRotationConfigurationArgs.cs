// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class PersonalAccessTokenRotationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The duration (in days) the new token should be valid for.
        /// </summary>
        [Input("expirationDays", required: true)]
        public Input<int> ExpirationDays { get; set; } = null!;

        /// <summary>
        /// The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
        /// </summary>
        [Input("rotateBeforeDays", required: true)]
        public Input<int> RotateBeforeDays { get; set; } = null!;

        public PersonalAccessTokenRotationConfigurationArgs()
        {
        }
        public static new PersonalAccessTokenRotationConfigurationArgs Empty => new PersonalAccessTokenRotationConfigurationArgs();
    }
}
