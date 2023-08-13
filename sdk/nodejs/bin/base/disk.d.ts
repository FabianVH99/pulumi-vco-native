import * as pulumi from "@pulumi/pulumi";
export declare class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Disk;
    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Disk;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly description: pulumi.Output<string>;
    readonly detach: pulumi.Output<boolean | undefined>;
    readonly disk_id: pulumi.Output<number>;
    readonly disk_name: pulumi.Output<string>;
    readonly disk_size: pulumi.Output<number>;
    readonly disk_type: pulumi.Output<string>;
    readonly exposed: pulumi.Output<boolean>;
    readonly iops: pulumi.Output<number | undefined>;
    readonly location: pulumi.Output<string>;
    readonly model: pulumi.Output<string>;
    readonly order: pulumi.Output<string>;
    readonly permanently: pulumi.Output<boolean | undefined>;
    readonly port: pulumi.Output<number>;
    readonly status: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    readonly vm_id: pulumi.Output<number>;
    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    customerID: pulumi.Input<string>;
    description: pulumi.Input<string>;
    detach?: pulumi.Input<boolean>;
    disk_name: pulumi.Input<string>;
    disk_size: pulumi.Input<number>;
    disk_type?: pulumi.Input<string>;
    iops?: pulumi.Input<number>;
    location: pulumi.Input<string>;
    permanently?: pulumi.Input<boolean>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vm_id?: pulumi.Input<string>;
}
