// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Resources
{
    [VcoResourceType("vco:resources:AntiAffinityGroup")]
    public partial class AntiAffinityGroup : global::Pulumi.CustomResource
    {
        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("group_id")]
        public Output<string> Group_id { get; private set; } = null!;

        [Output("spread")]
        public Output<int> Spread { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a AntiAffinityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AntiAffinityGroup(string name, AntiAffinityGroupArgs args, CustomResourceOptions? options = null)
            : base("vco:resources:AntiAffinityGroup", name, args ?? new AntiAffinityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AntiAffinityGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:resources:AntiAffinityGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AntiAffinityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AntiAffinityGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AntiAffinityGroup(name, id, options);
        }
    }

    public sealed class AntiAffinityGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudspace_id", required: true)]
        public Input<string> Cloudspace_id { get; set; } = null!;

        [Input("customerID", required: true)]
        public Input<string> CustomerID { get; set; } = null!;

        [Input("group_id", required: true)]
        public Input<string> Group_id { get; set; } = null!;

        [Input("spread", required: true)]
        public Input<int> Spread { get; set; } = null!;

        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public AntiAffinityGroupArgs()
        {
        }
        public static new AntiAffinityGroupArgs Empty => new AntiAffinityGroupArgs();
    }
}