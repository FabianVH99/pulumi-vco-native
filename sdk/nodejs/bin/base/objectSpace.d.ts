import * as pulumi from "@pulumi/pulumi";
export declare class ObjectSpace extends pulumi.CustomResource {
    /**
     * Get an existing ObjectSpace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ObjectSpace;
    /**
     * Returns true if the given object is an instance of ObjectSpace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ObjectSpace;
    readonly access_key: pulumi.Output<string | undefined>;
    readonly cloudspace_id: pulumi.Output<string | undefined>;
    readonly creation_time: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly domain: pulumi.Output<string | undefined>;
    readonly external_network: pulumi.Output<number | undefined>;
    readonly letsencrypt: pulumi.Output<boolean | undefined>;
    readonly letsencrypt_email: pulumi.Output<string | undefined>;
    readonly location: pulumi.Output<string>;
    readonly objectspace_id: pulumi.Output<string>;
    readonly objectspace_name: pulumi.Output<string>;
    readonly secret: pulumi.Output<string | undefined>;
    readonly status: pulumi.Output<string>;
    readonly subnet: pulumi.Output<string | undefined>;
    readonly token: pulumi.Output<string>;
    readonly update_time: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a ObjectSpace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectSpaceArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a ObjectSpace resource.
 */
export interface ObjectSpaceArgs {
    cloudspace_id?: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    external_network?: pulumi.Input<number>;
    letsencrypt?: pulumi.Input<boolean>;
    letsencrypt_email?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    name: pulumi.Input<string>;
    subnet?: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
