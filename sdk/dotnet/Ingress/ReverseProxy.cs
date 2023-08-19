// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Ingress
{
    [VcoResourceType("vco:ingress:ReverseProxy")]
    public partial class ReverseProxy : global::Pulumi.CustomResource
    {
        [Output("back_end")]
        public Output<Outputs.ReverseProxyBackend?> Back_end { get; private set; } = null!;

        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("front_end")]
        public Output<Outputs.ReverseProxyFrontEnd?> Front_end { get; private set; } = null!;

        [Output("health_check_scheme")]
        public Output<string?> Health_check_scheme { get; private set; } = null!;

        [Output("http_only")]
        public Output<bool?> Http_only { get; private set; } = null!;

        [Output("http_port")]
        public Output<int?> Http_port { get; private set; } = null!;

        [Output("https_port")]
        public Output<int?> Https_port { get; private set; } = null!;

        [Output("interval")]
        public Output<int?> Interval { get; private set; } = null!;

        [Output("ip_address")]
        public Output<string?> Ip_address { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        [Output("reverseproxy_id")]
        public Output<string> Reverseproxy_id { get; private set; } = null!;

        [Output("same_site")]
        public Output<string?> Same_site { get; private set; } = null!;

        [Output("scheme")]
        public Output<string> Scheme { get; private set; } = null!;

        [Output("secure")]
        public Output<bool?> Secure { get; private set; } = null!;

        [Output("serverpool_id")]
        public Output<string> Serverpool_id { get; private set; } = null!;

        [Output("stickySession_name")]
        public Output<string?> StickySession_name { get; private set; } = null!;

        [Output("target_port")]
        public Output<int> Target_port { get; private set; } = null!;

        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ReverseProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReverseProxy(string name, ReverseProxyArgs args, CustomResourceOptions? options = null)
            : base("vco:ingress:ReverseProxy", name, args ?? new ReverseProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReverseProxy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:ingress:ReverseProxy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ReverseProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReverseProxy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReverseProxy(name, id, options);
        }
    }

    public sealed class ReverseProxyArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudspace_id", required: true)]
        public Input<string> Cloudspace_id { get; set; } = null!;

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

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("health_check_scheme")]
        public Input<string>? Health_check_scheme { get; set; }

        [Input("http_only")]
        public Input<bool>? Http_only { get; set; }

        [Input("http_port")]
        public Input<int>? Http_port { get; set; }

        [Input("https_port")]
        public Input<int>? Https_port { get; set; }

        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("ip_address")]
        public Input<string>? Ip_address { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("same_site")]
        public Input<string>? Same_site { get; set; }

        [Input("scheme", required: true)]
        public Input<string> Scheme { get; set; } = null!;

        [Input("secure")]
        public Input<bool>? Secure { get; set; }

        [Input("serverpool_id", required: true)]
        public Input<string> Serverpool_id { get; set; } = null!;

        [Input("stickySession_name")]
        public Input<string>? StickySession_name { get; set; }

        [Input("target_port", required: true)]
        public Input<int> Target_port { get; set; } = null!;

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

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

        public ReverseProxyArgs()
        {
        }
        public static new ReverseProxyArgs Empty => new ReverseProxyArgs();
    }
}
