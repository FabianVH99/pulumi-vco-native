 // This is an example program showing how to create a cloudspace resource in C#
 // It is recommended to read the resource docs in order to understand
 // how they work and what inputs they require
 using System.Collections.Generic;
using System.Runtime.Intrinsics.Arm;
using Pulumi;
using Pulumi.Vco.Base;
using Pulumi.Vco.Cloudspace;

return await Deployment.RunAsync(() =>
   {
       //Load sensitive information from pulumi configuration variables
       var config = new Config();
       var url = config.Require("url");
       var token = config.Require("token");
       var customerId = config.Require("customerId");
       var location = config.Require("location");

       //Create a new cloudspace resource using all required values
       
   });