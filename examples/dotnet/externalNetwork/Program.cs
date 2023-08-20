using System.Collections.Generic;
using Pulumi;
using Pulumi.Vco.Base;
using Pulumi.Vco.Cloudspace;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var url = config.Require("url");
    var token = config.Require("token");
    var customerId = config.Require("customerId");
    var location = config.Require("location");

    var cloudspace = new Cloudspace("pulumi-cloudspace", new CloudspaceArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Name = "Pulumi_dotnet_cloudspace_vm",
        Private_network = "192.168.10.0/24",
        Location = location,
        External_network_id = 13,
        Private = false,
        Local_domain = "pulumi-domain",
    });

    var externalNetwork = new ExternalNetwork("pulumi-external-network", new ExternalNetworkArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        External_network_id = cloudspace.External_network_id.Apply(id => id.ToString()),
        External_network_type = "external",
        External_network_ip = "",
        Metric = 502,
    });

    return new Dictionary<string, object?>
   {
      ["cloudspace_id"] = cloudspace.Cloudspace_id,
      ["external_network_id"] = externalNetwork.External_network_id,
    };
});
