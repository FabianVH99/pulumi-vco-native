import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
export declare class ReverseProxy extends pulumi.CustomResource {
    /**
     * Get an existing ReverseProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReverseProxy;
    /**
     * Returns true if the given object is an instance of ReverseProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ReverseProxy;
    readonly back_end: pulumi.Output<outputs.resources.ReverseProxyBackend>;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly description: pulumi.Output<string>;
    readonly front_end: pulumi.Output<outputs.resources.ReverseProxyFrontEnd>;
    readonly name: pulumi.Output<string>;
    readonly reverseproxy_id: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly type: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a ReverseProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReverseProxyArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a ReverseProxy resource.
 */
export interface ReverseProxyArgs {
    back_end: pulumi.Input<inputs.resources.ReverseProxyBackendArgs>;
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    front_end: pulumi.Input<inputs.resources.ReverseProxyFrontEndArgs>;
    name: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
