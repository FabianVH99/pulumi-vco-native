// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Cloudspace
{
    [VcoResourceType("vco:cloudspace:ConnectedCloudspace")]
    public partial class ConnectedCloudspace : global::Pulumi.CustomResource
    {
        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("connected_cloudspace_id")]
        public Output<string> Connected_cloudspace_id { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("local_cloudspace_ip")]
        public Output<string> Local_cloudspace_ip { get; private set; } = null!;

        [Output("remote_cloudspace_ip")]
        public Output<string> Remote_cloudspace_ip { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectedCloudspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectedCloudspace(string name, ConnectedCloudspaceArgs args, CustomResourceOptions? options = null)
            : base("vco:cloudspace:ConnectedCloudspace", name, args ?? new ConnectedCloudspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectedCloudspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:cloudspace:ConnectedCloudspace", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectedCloudspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectedCloudspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConnectedCloudspace(name, id, options);
        }
    }

    public sealed class ConnectedCloudspaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudspace_id", required: true)]
        public Input<string> Cloudspace_id { get; set; } = null!;

        [Input("connected_cloudspace_id", required: true)]
        public Input<string> Connected_cloudspace_id { get; set; } = null!;

        [Input("customerID", required: true)]
        private Input<string>? _customerID;
        public Input<string>? CustomerID
        {
            get => _customerID;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _customerID = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("local_cloudspace_ip ", required: true)]
        public Input<string> Local_cloudspace_ip  { get; set; } = null!;

        [Input("remote_cloudspace_ip ", required: true)]
        public Input<string> Remote_cloudspace_ip  { get; set; } = null!;

        [Input("token", required: true)]
        private Input<string>? _token;
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url", required: true)]
        private Input<string>? _url;
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ConnectedCloudspaceArgs()
        {
        }
        public static new ConnectedCloudspaceArgs Empty => new ConnectedCloudspaceArgs();
    }
}
