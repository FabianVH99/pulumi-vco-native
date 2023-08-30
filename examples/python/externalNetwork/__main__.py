"""A Python Pulumi program"""

import pulumi
from pulumi_vco import base, cloudspace

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
    name="pulumi_python_cloudspace",
    private_network="192.168.10.0/24",
    external_network_id=13,
    private=False,
    local_domain="pulumipydomain",
))

pulumi_nic= cloudspace.ExternalNetwork("Python-nic", cloudspace.ExternalNetworkArgs(
    url = pulumi_url,
    customer_id = pulumi_customerId,
    token = pulumi_token,
    cloudspace_id = pulumi_cloudspace.cloudspace_id,
    external_network_id = pulumi_cloudspace.external_network_id.apply(str),
    external_network_type = "external",
    external_network_ip = "",
    metric = 501,
))

pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)
pulumi.export('nic_id', pulumi_nic.external_network_id)