import * as pulumi from "@pulumi/pulumi";
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
    readonly back_end: pulumi.Output<outputs.ingress.ReverseProxyBackend | undefined>;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly customerID: pulumi.Output<string>;
    readonly description: pulumi.Output<string | undefined>;
    readonly domain: pulumi.Output<string>;
    readonly email: pulumi.Output<string | undefined>;
    readonly enabled: pulumi.Output<boolean>;
    readonly front_end: pulumi.Output<outputs.ingress.ReverseProxyFrontEnd | undefined>;
    readonly health_check_scheme: pulumi.Output<string | undefined>;
    readonly http_only: pulumi.Output<boolean | undefined>;
    readonly http_port: pulumi.Output<number | undefined>;
    readonly https_port: pulumi.Output<number | undefined>;
    readonly interval: pulumi.Output<number | undefined>;
    readonly ip_address: pulumi.Output<string | undefined>;
    readonly name: pulumi.Output<string>;
    readonly path: pulumi.Output<string | undefined>;
    readonly port: pulumi.Output<number | undefined>;
    readonly reverseproxy_id: pulumi.Output<string>;
    readonly same_site: pulumi.Output<string | undefined>;
    readonly scheme: pulumi.Output<string>;
    readonly secure: pulumi.Output<boolean | undefined>;
    readonly serverpool_id: pulumi.Output<string>;
    readonly stickySession_name: pulumi.Output<string | undefined>;
    readonly target_port: pulumi.Output<number>;
    readonly timeout: pulumi.Output<number | undefined>;
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
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    domain: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    enabled: pulumi.Input<boolean>;
    health_check_scheme?: pulumi.Input<string>;
    http_only?: pulumi.Input<boolean>;
    http_port?: pulumi.Input<number>;
    https_port?: pulumi.Input<number>;
    interval?: pulumi.Input<number>;
    ip_address?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    path?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    same_site?: pulumi.Input<string>;
    scheme: pulumi.Input<string>;
    secure?: pulumi.Input<boolean>;
    serverpool_id: pulumi.Input<string>;
    stickySession_name?: pulumi.Input<string>;
    target_port: pulumi.Input<number>;
    timeout?: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
