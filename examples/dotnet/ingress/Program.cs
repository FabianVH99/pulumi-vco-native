using System.Collections.Generic;
using System.Net;
using Pulumi;
using Pulumi.Vco.Base;
using Pulumi.Vco.Ingress;

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
        Name = "Pulumi_dotnet_cloudspace_sv",
        Private_network = "192.168.10.0/24",
        Location = location,
        External_network_id = 13,
        Private = false,
        Local_domain = "pulumi-domain",
    });

    var serverPool = new ServerPool("pulumi-sv_pool", new ServerPoolArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Name = "Pulumi_dotnet_sv",
        Description = "sv_pool_test",
    });

    var host = new Host("pulumi-host", new HostArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Serverpool_id = serverPool.Serverpool_id,
        Address = "192.168.10.10",
    });

    var loadBalancer = new LoadBalancer("pulumi-lb", new LoadBalancerArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Name = "Pulumi_dotnet_lb",
        Type = "tcp",
        Port = 23,
        Serverpool_id = serverPool.Serverpool_id,
        Target_port = 23,
        Ip_address = cloudspace.External_network_ip,
        Is_enabled = true,
        Domain = "whiteskycloud-2.try-dns.whitesky.cloud",
        Tls_termination = true,
    });

    var reverseProxy = new ReverseProxy("pulumi-rp", new ReverseProxyArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Name = "Pulumi_dotnet_rp",
        Domain = "Pulumi",
        Scheme = "http",
        Enabled = false,
        Http_port = 25,
        Target_port = 25,
        Serverpool_id = serverPool.Serverpool_id,
    });

    return new Dictionary<string, object?>
   {
      ["cs_id"] =cloudspace.Cloudspace_id,
      ["sv_id"] = serverPool.Serverpool_id,
      ["lb_id"] = loadBalancer.Loadbalancer_id,
      ["rp_id"] = reverseProxy.Reverseproxy_id,
    };
});
