using System.Collections.Generic;
using Pulumi;
using Pulumi.Vco.Base;
using Pulumi.Vco.Objectspace;

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
        Name = "Pulumi_dotnet_cloudspace",
        Private_network = "192.168.10.0/24",
        Location = location,
        External_network_id = 13,
        Private = false,
        Local_domain = "pulumi-domain",
    });

    var objectspace = new ObjectSpace("pulumi-objectspace", new ObjectSpaceArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Name = "Pulumi_objectspace_dotnet",
        Location = location,
    });

    var objectspaceLink = new Link("pulumi-objectspace-link", new LinkArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Objectspace_id = objectspace.Objectspace_id,
    });

    return new Dictionary<string, object?>
    {
        ["pulumi-cloudspace.cloudspace_id"] = cloudspace.Cloudspace_id,
        ["pulumi-objectspace.objectspace_id"] = objectspace.Objectspace_id
    };
});
