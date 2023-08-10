// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Resources.Outputs
{

    [OutputType]
    public sealed class ResourceLimits
    {
        public readonly double External_network_quota;
        public readonly double Memory_quota;
        public readonly double Public_ip_quota;
        public readonly double Vcpu_quota;
        public readonly double Vdisk_space_quota;

        [OutputConstructor]
        private ResourceLimits(
            double external_network_quota,

            double memory_quota,

            double public_ip_quota,

            double vcpu_quota,

            double vdisk_space_quota)
        {
            External_network_quota = external_network_quota;
            Memory_quota = memory_quota;
            Public_ip_quota = public_ip_quota;
            Vcpu_quota = vcpu_quota;
            Vdisk_space_quota = vdisk_space_quota;
        }
    }
}
