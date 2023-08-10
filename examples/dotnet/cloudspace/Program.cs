 // This is an example program showing how to create a cloudspace resource in C#
 // It is recommended to read the resource docs in order to understand
 // how they work and what inputs they require
 using System.Collections.Generic;
 using Pulumi;
 using Pulumi.Vco.Resources;

   return await Deployment.RunAsync(() =>
   {
       //Load sensitive information from pulumi configuration variables
       var config = new Config();
       var url = config.Require("url");
       var token = config.Require("token");
       var customerId = config.Require("customerId");
       var location = config.Require("location");

       //Create a new cloudspace resource using all required values
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