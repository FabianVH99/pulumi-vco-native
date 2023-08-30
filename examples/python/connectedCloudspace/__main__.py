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

pulumi_cloudspace_2 = base.Cloudspace("Python-cloudspace-2", base.CloudspaceArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    location=pulumi_location,
    token=pulumi_token,
    name="pulumi_python_cloudspace_2",
    private_network="192.168.11.0/24",
    external_network_id=13,
    private=False,
    local_domain="pulumipydomain",
))

pulumi_connected_cloudspace = cloudspace.ConnectedCloudspace("Python-connected-cloudspace", cloudspace.ConnectedCloudspaceArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    local_cloudspace_ip_=pulumi_cloudspace.external_network_ip,
    connected_cloudspace_id=pulumi_cloudspace_2.cloudspace_id,
    remote_cloudspace_ip_=pulumi_cloudspace_2.external_network_ip,
))

pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)
pulumi.export('cloudspace_id_2', pulumi_cloudspace_2.cloudspace_id)