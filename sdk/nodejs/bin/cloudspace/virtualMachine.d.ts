import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
export declare class VirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachine;
    /**
     * Returns true if the given object is an instance of VirtualMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is VirtualMachine;
    readonly acronis: pulumi.Output<boolean | undefined>;
    readonly agent_status: pulumi.Output<string>;
    readonly all_vm_disks: pulumi.Output<boolean | undefined>;
    readonly appliance: pulumi.Output<boolean>;
    readonly boot_disk_id: pulumi.Output<number>;
    readonly boot_type: pulumi.Output<string>;
    readonly cdrom_id: pulumi.Output<number | undefined>;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly cpu_topology: pulumi.Output<outputs.cloudspace.CpuTopology>;
    readonly cpus_pin_status: pulumi.Output<boolean>;
    readonly creation_time: pulumi.Output<number>;
    readonly customerID: pulumi.Output<string>;
    readonly data_disks: pulumi.Output<number[] | undefined>;
    readonly description: pulumi.Output<string>;
    readonly disk_size: pulumi.Output<number | undefined>;
    readonly disks: pulumi.Output<outputs.cloudspace.VmDisk[]>;
    readonly enable_vm_agent: pulumi.Output<boolean | undefined>;
    readonly hostname: pulumi.Output<string>;
    readonly image_id: pulumi.Output<number>;
    readonly impact_alert_hook: pulumi.Output<string | undefined>;
    readonly locked: pulumi.Output<boolean>;
    readonly memory: pulumi.Output<number>;
    readonly name: pulumi.Output<string>;
    readonly network_interfaces: pulumi.Output<outputs.cloudspace.NetworkInterface[]>;
    readonly os_accounts: pulumi.Output<outputs.cloudspace.OsAccount[]>;
    readonly os_image_name: pulumi.Output<string>;
    readonly os_name: pulumi.Output<string | undefined>;
    readonly os_type: pulumi.Output<string | undefined>;
    readonly private_ip: pulumi.Output<string | undefined>;
    readonly snapshot_id: pulumi.Output<string | undefined>;
    readonly stack_id: pulumi.Output<number>;
    readonly status: pulumi.Output<string>;
    readonly storage: pulumi.Output<number>;
    readonly token: pulumi.Output<string>;
    readonly update_time: pulumi.Output<number>;
    readonly url: pulumi.Output<string>;
    readonly user_data: pulumi.Output<string | undefined>;
    readonly vcpus: pulumi.Output<number>;
    readonly vm_id: pulumi.Output<number>;
    /**
     * Create a VirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    acronis?: pulumi.Input<boolean>;
    all_vm_disks?: pulumi.Input<boolean>;
    boot_disk_id?: pulumi.Input<number>;
    boot_type?: pulumi.Input<string>;
    cdrom_id?: pulumi.Input<number>;
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    data_disks?: pulumi.Input<pulumi.Input<number>[]>;
    description: pulumi.Input<string>;
    disk_size?: pulumi.Input<number>;
    enable_vm_agent?: pulumi.Input<boolean>;
    image_id?: pulumi.Input<number>;
    memory: pulumi.Input<number>;
    name: pulumi.Input<string>;
    os_name?: pulumi.Input<string>;
    os_type?: pulumi.Input<string>;
    private_ip?: pulumi.Input<string>;
    snapshot_id?: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    user_data?: pulumi.Input<string>;
    vcpus: pulumi.Input<number>;
}
