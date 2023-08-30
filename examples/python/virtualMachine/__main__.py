"""A Python Pulumi program"""

import pulumi
from pulumi import ResourceOptions
from pulumi_vco import base, cloudspace, virtual_machine

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
    name="pulumi_python_cloudspace_vm",
    private_network="192.168.10.0/24",
    external_network_id=13,
    private=False,
    local_domain="pulumipydomain",
))

pulumi_virtual_machine = cloudspace.VirtualMachine("Python-vm", cloudspace.VirtualMachineArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    name = "Pulumi_python_vm",
    description = "vm_test",
    vcpus = 4,
    memory = 2048,
    image_id = 258,
    disk_size = 250,
    boot_type = "bios",
))

pulumi_virtual_machine_portforward = cloudspace.PortForward("Python-pf", cloudspace.PortForwardArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    local_port = 22,
    public_port = 22,
    protocol = "tcp",
    public_ip = pulumi_cloudspace.external_network_ip,
    vm_id = pulumi_virtual_machine.vm_id,
))

pulumi_virtual_machine_nic = virtual_machine.VirtualMachineNIC("Python-vm-nic", virtual_machine.VirtualMachineNICArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    external_network_id=pulumi_cloudspace.external_network_id,
    vm_id = pulumi_virtual_machine.vm_id,
), opts=ResourceOptions(depends_on=[pulumi_virtual_machine_portforward]))

pulumi_virtual_machine_cd = virtual_machine.VirtualMachineCD("Python-vm-cd", virtual_machine.VirtualMachineCDArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    cdrom_id=9000,
    vm_id = pulumi_virtual_machine.vm_id,
), opts=ResourceOptions(depends_on=[pulumi_virtual_machine_nic]))

pulumi_disk = base.Disk("python-disk", base.DiskArgs(
    
    url=pulumi_url,
    customer_id=pulumi_customerId,
    location=pulumi_location,
    token=pulumi_token,
    disk_name="pulumi_py_disk",
    description="test",
    disk_size=100,
))

pulumi_virtual_machine_disk = virtual_machine.VirtualMachineDisk("Python-vm-disk", virtual_machine.VirtualMachineDiskArgs(
    url= pulumi_url,
    customer_id=pulumi_customerId,
    token=pulumi_token,
    cloudspace_id=pulumi_cloudspace.cloudspace_id,
    disk_id=pulumi_disk.disk_id,
    vm_id = pulumi_virtual_machine.vm_id,
), opts=ResourceOptions(depends_on=[pulumi_virtual_machine_cd]))

pulumi.export('cloudspace_id', pulumi_cloudspace.cloudspace_id)
pulumi.export('vm_id', pulumi_virtual_machine.vm_id)
pulumi.export('disk_id', pulumi_disk.disk_id)
pulumi.export('portforward_id', pulumi_virtual_machine_portforward.portforward_id)