import * as pulumi from "@pulumi/pulumi";
export declare class VirtualMachineNIC extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineNIC resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachineNIC;
    /**
     * Returns true if the given object is an instance of VirtualMachineNIC.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is VirtualMachineNIC;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly device_name: pulumi.Output<string>;
    readonly external_cloudspace_id: pulumi.Output<string>;
    readonly external_network_id: pulumi.Output<number>;
    readonly external_network_ip: pulumi.Output<string | undefined>;
    readonly external_network_type: pulumi.Output<string | undefined>;
    readonly ip_address: pulumi.Output<string>;
    readonly mac_address: pulumi.Output<string>;
    readonly model: pulumi.Output<string>;
    readonly network_id: pulumi.Output<number>;
    readonly nic_type: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    readonly vm_id: pulumi.Output<number>;
    /**
     * Create a VirtualMachineNIC resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineNICArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a VirtualMachineNIC resource.
 */
export interface VirtualMachineNICArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    external_cloudspace_id?: pulumi.Input<string>;
    external_network_id: pulumi.Input<number>;
    external_network_ip?: pulumi.Input<string>;
    external_network_type?: pulumi.Input<string>;
    model?: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vm_id: pulumi.Input<number>;
}
