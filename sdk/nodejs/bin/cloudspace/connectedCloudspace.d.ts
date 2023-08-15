import * as pulumi from "@pulumi/pulumi";
export declare class ConnectedCloudspace extends pulumi.CustomResource {
    /**
     * Get an existing ConnectedCloudspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectedCloudspace;
    /**
     * Returns true if the given object is an instance of ConnectedCloudspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ConnectedCloudspace;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly connected_cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly local_cloudspace_ip: pulumi.Output<string>;
    readonly remote_cloudspace_ip: pulumi.Output<string>;
    readonly status: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a ConnectedCloudspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectedCloudspaceArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a ConnectedCloudspace resource.
 */
export interface ConnectedCloudspaceArgs {
    cloudspace_id: pulumi.Input<string>;
    connected_cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    local_cloudspace_ip: pulumi.Input<string>;
    remote_cloudspace_ip: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
