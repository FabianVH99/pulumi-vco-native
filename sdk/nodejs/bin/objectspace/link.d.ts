import * as pulumi from "@pulumi/pulumi";
export declare class Link extends pulumi.CustomResource {
    /**
     * Get an existing Link resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Link;
    /**
     * Returns true if the given object is an instance of Link.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Link;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly objectspace_id: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a Link resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a Link resource.
 */
export interface LinkArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    objectspace_id: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
