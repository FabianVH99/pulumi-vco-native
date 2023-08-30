"""A Python Pulumi program"""

import pulumi
from pulumi_vco import base, objectspace

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
    name="pulumi_python_objspaceLink",
    private_network="192.168.10.0/24",
    external_network_id=13,
    private=False,
    local_domain="pulumipydomain",
))


pulumi_objectspace = base.ObjectSpace("Python-objectspace", base.ObjectSpaceArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    location=pulumi_location,
    token=pulumi_token,
    objectspace_name="pulumi_python_objectspace",
))

pulumi_objectspaceLink = objectspace.Link("Python-link", objectspace.LinkArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    objectspace_id=pulumi_objectspace.objectspace_id
))

pulumi.export('objectspace_id', pulumi_objectspace.objectspace_id)
pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)
