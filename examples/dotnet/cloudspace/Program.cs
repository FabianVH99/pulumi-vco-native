 // This is an example program showing how to create a resource in C#
 // It is recommended to read the resource docs in order to understand
 // how they work and what inputs they require
using System.Collections.Generic;
using Pulumi;
using Pulumi.Vco.Base;

return await Deployment.RunAsync(() =>
   {
       // Load sensitive information from pulumi configuration variables
       var config = new Config();
       var url = config.Require("url");
       var token = config.Require("token");
       var customerId = config.Require("customerId");
       var location = config.Require("location");

       // Create a new cloudspace resource using all required values
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

       // Export the cloudspace id as an output. Please note that there is a different between
       // cloudspace.Id and cloudspace.Cloudspace_id. The first will return the name of the
       // Pulumi resource (in this case, "pulumi-cloudspace").
       // The second will return the cloudspace id as it is defined in the VCO portal.
       return new Dictionary<string, object?>
       {
           ["pulumi-cloudspace.cloudspace_id"] = cloudspace.Cloudspace_id
       };
   });