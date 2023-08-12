import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
export declare class ExposedDisk extends pulumi.CustomResource {
    /**
     * Get an existing ExposedDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExposedDisk;
    /**
     * Returns true if the given object is an instance of ExposedDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ExposedDisk;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly disk_id: pulumi.Output<number>;
    readonly endpoint: pulumi.Output<outputs.disk.Endpoint>;
    readonly iops: pulumi.Output<number | undefined>;
    readonly max_connections: pulumi.Output<number | undefined>;
    readonly protocol: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a ExposedDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExposedDiskArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a ExposedDisk resource.
 */
export interface ExposedDiskArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    disk_id: pulumi.Input<number>;
    iops?: pulumi.Input<number>;
    max_connections?: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
