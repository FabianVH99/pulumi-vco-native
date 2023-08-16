import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
export declare class ServerPool extends pulumi.CustomResource {
    /**
     * Get an existing ServerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerPool;
    /**
     * Returns true if the given object is an instance of ServerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ServerPool;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly description: pulumi.Output<string>;
    readonly hosts: pulumi.Output<outputs.ingress.ServerPoolHost[] | undefined>;
    readonly name: pulumi.Output<string>;
    readonly serverpool_id: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a ServerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerPoolArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a ServerPool resource.
 */
export interface ServerPoolArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    description: pulumi.Input<string>;
    name: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
