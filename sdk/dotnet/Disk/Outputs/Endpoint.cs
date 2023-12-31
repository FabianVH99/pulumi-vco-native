// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Disk.Outputs
{

    [OutputType]
    public sealed class Endpoint
    {
        public readonly string Address;
        public readonly string Name;
        public readonly int Port;
        public readonly string Private_address;
        public readonly int Private_port;
        public readonly string Psk;
        public readonly string User;

        [OutputConstructor]
        private Endpoint(
            string address,

            string name,

            int port,

            string private_address,

            int private_port,

            string psk,

            string user)
        {
            Address = address;
            Name = name;
            Port = port;
            Private_address = private_address;
            Private_port = private_port;
            Psk = psk;
            User = user;
        }
    }
}
