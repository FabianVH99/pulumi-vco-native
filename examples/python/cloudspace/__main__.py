"""A Python Pulumi program"""

import pulumi
from pulumi_vco import base

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

pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)