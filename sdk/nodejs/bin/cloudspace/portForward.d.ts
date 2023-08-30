import * as pulumi from "@pulumi/pulumi";
export declare class PortForward extends pulumi.CustomResource {
    /**
     * Get an existing PortForward resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PortForward;
    /**
     * Returns true if the given object is an instance of PortForward.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is PortForward;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly local_port: pulumi.Output<number>;
    readonly nested_cs_id: pulumi.Output<string | undefined>;
    readonly portforward_id: pulumi.Output<string>;
    readonly protocol: pulumi.Output<string>;
    readonly public_ip: pulumi.Output<string>;
    readonly public_port: pulumi.Output<number>;
    readonly till_public_port: pulumi.Output<number | undefined>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    readonly vm_id: pulumi.Output<number>;
    /**
     * Create a PortForward resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortForwardArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a PortForward resource.
 */
export interface PortForwardArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    local_port: pulumi.Input<number>;
    nested_cs_id?: pulumi.Input<string>;
    protocol: pulumi.Input<string>;
    public_ip: pulumi.Input<string>;
    public_port: pulumi.Input<number>;
    till_public_port?: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vm_id: pulumi.Input<number>;
}
