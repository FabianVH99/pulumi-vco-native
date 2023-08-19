using System.Collections.Generic;
using Pulumi;
using Pulumi.Vco.Base;
using Pulumi.Vco.Cloudspace;
using Pulumi.Vco.Virtual_machine;

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
        Name = "Pulumi_dotnet_cloudspace_vm",
        Private_network = "192.168.10.0/24",
        Location = location,
        External_network_id = 13,
        Private = false,
        Local_domain = "pulumi-domain",
    });

    var virtualMachine = new VirtualMachine("pulumi-vm", new VirtualMachineArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Name = "Pulumi_dotnet_vm",
        Description = "vm_test",
        Vcpus = 4,
        Memory = 2048,
        Image_id = 258,
        Disk_size = 250,
        Boot_type = "bios",
    });

    var virtualMachineCD = new VirtualMachineCD("pulumi-vm-cd", new VirtualMachineCDArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Vm_id = virtualMachine.Vm_id,
        Cdrom_id = 9000,
    });

    var disk = new Disk("pulumi-disk", new DiskArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Location = location,
        Disk_name = "pulumi_disk_vm",
        Description = "test_vm_resource",
        Disk_size = 100,
    });

    var virtualMachineNIC = new VirtualMachineNIC("pulumi-vm-nic", new VirtualMachineNICArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Vm_id = virtualMachine.Vm_id,
        External_network_id = cloudspace.External_network_id,
    });

    var virtualMachineDisk = new VirtualMachineDisk("pulumi-vm-disk", new VirtualMachineDiskArgs
    {
        Url = url,
        Token = token,
        CustomerID = customerId,
        Cloudspace_id = cloudspace.Cloudspace_id,
        Vm_id = virtualMachine.Vm_id,
        Disk_id = disk.Disk_id,
    }, new CustomResourceOptions { DependsOn = { virtualMachineNIC } });

    return new Dictionary<string, object?>
    {
        ["cloudspace_id"] = cloudspace.Cloudspace_id,
        ["vm_id"] = virtualMachine.Vm_id,
        ["disk_id"] = disk.Disk_id,
    };
});
