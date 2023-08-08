// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Resources
{
    [VcoResourceType("vco:resources:LoadBalancer")]
    public partial class LoadBalancer : global::Pulumi.CustomResource
    {
        [Output("back_end")]
        public Output<Outputs.BackEndState> Back_end { get; private set; } = null!;

        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("front_end")]
        public Output<Outputs.FrontEnd> Front_end { get; private set; } = null!;

        [Output("loadbalancer_id")]
        public Output<string> Loadbalancer_id { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs args, CustomResourceOptions? options = null)
            : base("vco:resources:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:resources:LoadBalancer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, options);
        }
    }

    public sealed class LoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudspace_id", required: true)]
        public Input<string> Cloudspace_id { get; set; } = null!;

        [Input("customerID", required: true)]
        public Input<string> CustomerID { get; set; } = null!;

        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public LoadBalancerArgs()
        {
        }
        public static new LoadBalancerArgs Empty => new LoadBalancerArgs();
    }
}