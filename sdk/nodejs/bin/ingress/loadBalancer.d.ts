import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
export declare class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoadBalancer;
    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is LoadBalancer;
    readonly back_end: pulumi.Output<outputs.ingress.BackEnd>;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly description: pulumi.Output<string | undefined>;
    readonly domain: pulumi.Output<string | undefined>;
    readonly front_end: pulumi.Output<outputs.ingress.FrontEnd>;
    readonly ip_address: pulumi.Output<string | undefined>;
    readonly is_enabled: pulumi.Output<boolean | undefined>;
    readonly loadbalancer_id: pulumi.Output<string>;
    readonly name: pulumi.Output<string>;
    readonly port: pulumi.Output<number>;
    readonly serverpool_id: pulumi.Output<string>;
    readonly target_port: pulumi.Output<number>;
    readonly tls_termination: pulumi.Output<boolean | undefined>;
    readonly token: pulumi.Output<string>;
    readonly type: pulumi.Output<string>;
    readonly url: pulumi.Output<string>;
    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    ip_address?: pulumi.Input<string>;
    is_enabled?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    port: pulumi.Input<number>;
    serverpool_id: pulumi.Input<string>;
    target_port: pulumi.Input<number>;
    tls_termination?: pulumi.Input<boolean>;
    token: pulumi.Input<string>;
    type: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
