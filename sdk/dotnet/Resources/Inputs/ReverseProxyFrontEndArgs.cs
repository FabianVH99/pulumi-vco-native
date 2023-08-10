// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Resources.Inputs
{

    public sealed class ReverseProxyFrontEndArgs : global::Pulumi.ResourceArgs
    {
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("http_port")]
        public Input<int>? Http_port { get; set; }

        [Input("https_port")]
        public Input<int>? Https_port { get; set; }

        [Input("ip_address")]
        public Input<string>? Ip_address { get; set; }

        [Input("letsencrypt", required: true)]
        public Input<Inputs.LetsEncryptArgs> Letsencrypt { get; set; } = null!;

        [Input("scheme", required: true)]
        public Input<string> Scheme { get; set; } = null!;

        public ReverseProxyFrontEndArgs()
        {
        }
        public static new ReverseProxyFrontEndArgs Empty => new ReverseProxyFrontEndArgs();
    }
}
