"""A Python Pulumi program"""

import pulumi
from pulumi_vco import base, disk

config = pulumi.Config()

pulumi_token = config.require("token")
pulumi_url = config.require("url")
pulumi_customerId = config.require("customerId")
pulumi_location = config.require("location")

pulumi_cloudspace = base.Cloudspace("python-cloudspace", base.CloudspaceArgs(
    url=pulumi_url,
    customer_id=pulumi_customerId,
    location=pulumi_location,
    token=pulumi_token,
    name="pulumi_python_disk",
    private_network="192.168.10.0/24",
    external_network_id=13,
    private=False,
    local_domain="pulumipydomain",
))

pulumi_disk = base.Disk("python-disk", base.DiskArgs(
    
    url=pulumi_url,
    customer_id=pulumi_customerId,
    location=pulumi_location,
    token=pulumi_token,
    disk_name="pulumi_disk",
    description="test",
    disk_size=100,
))

pulumi_exposed_disk = disk.ExposedDisk("python-exposed-disk", disk.ExposedDiskArgs(
    
    url=pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    disk_id=pulumi_disk.disk_id,
))

pulumi.export('disk_id', pulumi_disk.disk_id)
pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)
