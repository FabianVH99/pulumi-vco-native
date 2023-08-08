// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Resources.Inputs
{

    public sealed class FirewallCustomArgs : global::Pulumi.ResourceArgs
    {
        [Input("cdrom_id", required: true)]
        public Input<int> Cdrom_id { get; set; } = null!;

        [Input("disk_size", required: true)]
        public Input<int> Disk_size { get; set; } = null!;

        [Input("image_id", required: true)]
        public Input<int> Image_id { get; set; } = null!;

        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("vcpus", required: true)]
        public Input<int> Vcpus { get; set; } = null!;

        public FirewallCustomArgs()
        {
        }
        public static new FirewallCustomArgs Empty => new FirewallCustomArgs();
    }
}