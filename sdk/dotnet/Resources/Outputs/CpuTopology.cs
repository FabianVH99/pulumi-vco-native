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
    public sealed class CpuTopology
    {
        public readonly int Cores;
        public readonly int Sockets;
        public readonly int Threads;

        [OutputConstructor]
        private CpuTopology(
            int cores,

            int sockets,

            int threads)
        {
            Cores = cores;
            Sockets = sockets;
            Threads = threads;
        }
    }
}
