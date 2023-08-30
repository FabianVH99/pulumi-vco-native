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

    var cloudspace_1 = new Cloudspace("pulumi-cloudspace", new CloudspaceArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Name = "Pulumi_dotnet_connectedCloudspace1",
        Private_network = "192.168.10.0/24",
        Location = location,
        External_network_id = 13,
        Private = false,
        Local_domain = "pulumi-domain",
    });

    var cloudspace_2 = new Cloudspace("pulumi-cloudspace-2", new CloudspaceArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Name = "Pulumi_dotnet_connectedCloudspace2",
        Private_network = "192.168.11.0/24",
        Location = location,
        External_network_id = 13,
        Private = false,
        Local_domain = "pulumi-domain-2",
    });

    var connected_cloudspace = new ConnectedCloudspace("pulumi-cloudspace-connection", new ConnectedCloudspaceArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace_1.Cloudspace_id,
        Local_cloudspace_ip = cloudspace_1.External_network_ip,
        Connected_cloudspace_id = cloudspace_2.Cloudspace_id,
        Remote_cloudspace_ip = cloudspace_2.External_network_ip,
    });

    return new Dictionary<string, object?>
    {
        ["cloudspace_1.cloudspace_id"] = cloudspace_1.Cloudspace_id,
        ["cloudspace_2.cloudspace_id"] = cloudspace_2.Cloudspace_id,
    };
});
