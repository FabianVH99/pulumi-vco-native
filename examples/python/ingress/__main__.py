"""A Python Pulumi program"""

import pulumi
from pulumi_vco import base, ingress

config = pulumi.Config()

pulumi_token = config.require("token")
pulumi_url = config.require("url")
pulumi_customerId = config.require("customerId")
pulumi_location = config.require("location")

pulumi_cloudspace = base.Cloudspace("Python-cloudspace", base.CloudspaceArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    location=pulumi_location,
    token=pulumi_token,
    name="pulumi_python_ingress",
    private_network="192.168.10.0/24",
    external_network_id=13,
    private=False,
    local_domain="pulumipydomain",
))

pulumi_serverpool = ingress.ServerPool("Python-sv", ingress.ServerPoolArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id = pulumi_cloudspace.cloudspace_id,
    name = "Pulumi_python_sv",
    description = "sv_pool_test",
))

pulumi_host = ingress.Host("Python-host", ingress.HostArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id = pulumi_cloudspace.cloudspace_id,
    serverpool_id= pulumi_serverpool.serverpool_id,
    address="192.168.10.10",
))

pulumi_loadbalancer = ingress.LoadBalancer("Python-lb", ingress.LoadBalancerArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id = pulumi_cloudspace.cloudspace_id,
    name = "Pulumi_python_lb",
    type = "udp",
    port = 23,
    serverpool_id = pulumi_serverpool.serverpool_id,
    target_port = 23,
    ip_address = pulumi_cloudspace.external_network_ip,
    is_enabled = True,
    domain = "whiteskycloud-2.try-dns.whitesky.cloud",
    tls_termination = True,
))

pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)
pulumi.export('serverpool_id', pulumi_serverpool.serverpool_id)
pulumi.export('loadbalancer_id', pulumi_loadbalancer.loadbalancer_id)
