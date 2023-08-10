using System.Collections.Generic;
using Pulumi;
using Pulumi.Vco.Resources;

return await Deployment.RunAsync(() =>
{
   using System.Collections.Generic;
   using Pulumi;
   using Pulumi.Vco.Resources;

   return await Deployment.RunAsync(() =>
   {
       var config = new Config();
       var url = config.Require("url");
       var token = config.Require("token");
       var customerId = config.Require("customerId");
       var location = config.Require("location");

       var cloudspace = new Cloudspace("my-cloudspace", new CloudspaceArgs
       {
           Url = url,
           Token = token,
           CustomerID = customerId,
           Name = "Pulumi_dotnet_cloudspace",
           Location = location,
           PrivateNetwork = "192.168.10.0/24",
           ExternalNetworkID = 13,
           Private = false,
       });
   });

});
