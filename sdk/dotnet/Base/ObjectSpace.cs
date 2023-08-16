// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Base
{
    [VcoResourceType("vco:base:ObjectSpace")]
    public partial class ObjectSpace : global::Pulumi.CustomResource
    {
        [Output("access_key")]
        public Output<string?> Access_key { get; private set; } = null!;

        [Output("cloudspaceID")]
        public Output<string?> CloudspaceID { get; private set; } = null!;

        [Output("creation_time")]
        public Output<string> Creation_time { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        [Output("externalNetwork")]
        public Output<int?> ExternalNetwork { get; private set; } = null!;

        [Output("letsencrypt")]
        public Output<bool?> Letsencrypt { get; private set; } = null!;

        [Output("letsencryptEmail")]
        public Output<string?> LetsencryptEmail { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("objectspace_id")]
        public Output<string> Objectspace_id { get; private set; } = null!;

        [Output("objectspace_name")]
        public Output<string> Objectspace_name { get; private set; } = null!;

        [Output("secret")]
        public Output<string?> Secret { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("subnet")]
        public Output<string?> Subnet { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("update_time")]
        public Output<string> Update_time { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectSpace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectSpace(string name, ObjectSpaceArgs args, CustomResourceOptions? options = null)
            : base("vco:base:ObjectSpace", name, args ?? new ObjectSpaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectSpace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:base:ObjectSpace", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectSpace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectSpace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ObjectSpace(name, id, options);
        }
    }

    public sealed class ObjectSpaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudspaceID")]
        public Input<string>? CloudspaceID { get; set; }

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

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("externalNetwork")]
        public Input<int>? ExternalNetwork { get; set; }

        [Input("letsencrypt")]
        public Input<bool>? Letsencrypt { get; set; }

        [Input("letsencryptEmail")]
        public Input<string>? LetsencryptEmail { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("objectspace_name", required: true)]
        public Input<string> Objectspace_name { get; set; } = null!;

        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

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

        public ObjectSpaceArgs()
        {
        }
        public static new ObjectSpaceArgs Empty => new ObjectSpaceArgs();
    }
}
