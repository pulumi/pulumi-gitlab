// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ProjectTagCommitGetArgs : Pulumi.ResourceArgs
    {
        [Input("authorEmail")]
        public Input<string>? AuthorEmail { get; set; }

        [Input("authorName")]
        public Input<string>? AuthorName { get; set; }

        [Input("authoredDate")]
        public Input<string>? AuthoredDate { get; set; }

        [Input("committedDate")]
        public Input<string>? CommittedDate { get; set; }

        [Input("committerEmail")]
        public Input<string>? CommitterEmail { get; set; }

        [Input("committerName")]
        public Input<string>? CommitterName { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("parentIds")]
        private InputList<string>? _parentIds;
        public InputList<string> ParentIds
        {
            get => _parentIds ?? (_parentIds = new InputList<string>());
            set => _parentIds = value;
        }

        [Input("shortId")]
        public Input<string>? ShortId { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        public ProjectTagCommitGetArgs()
        {
        }
    }
}