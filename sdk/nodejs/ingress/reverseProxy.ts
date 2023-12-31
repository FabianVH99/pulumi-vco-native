// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ReverseProxy extends pulumi.CustomResource {
    /**
     * Get an existing ReverseProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReverseProxy {
        return new ReverseProxy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:ingress:ReverseProxy';

    /**
     * Returns true if the given object is an instance of ReverseProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReverseProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReverseProxy.__pulumiType;
    }

    public /*out*/ readonly back_end!: pulumi.Output<outputs.ingress.ReverseProxyBackend | undefined>;
    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly domain!: pulumi.Output<string>;
    public readonly email!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean>;
    public /*out*/ readonly front_end!: pulumi.Output<outputs.ingress.ReverseProxyFrontEnd | undefined>;
    public readonly health_check_scheme!: pulumi.Output<string | undefined>;
    public readonly http_only!: pulumi.Output<boolean | undefined>;
    public readonly http_port!: pulumi.Output<number | undefined>;
    public readonly https_port!: pulumi.Output<number | undefined>;
    public readonly interval!: pulumi.Output<number | undefined>;
    public readonly ip_address!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly path!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number | undefined>;
    public /*out*/ readonly reverseproxy_id!: pulumi.Output<string>;
    public readonly same_site!: pulumi.Output<string | undefined>;
    public readonly scheme!: pulumi.Output<string>;
    public readonly secure!: pulumi.Output<boolean | undefined>;
    public readonly serverpool_id!: pulumi.Output<string>;
    public readonly stickySession_name!: pulumi.Output<string | undefined>;
    public readonly target_port!: pulumi.Output<number>;
    public readonly timeout!: pulumi.Output<number | undefined>;
    public readonly token!: pulumi.Output<string>;
    public /*out*/ readonly type!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ReverseProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReverseProxyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.scheme === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheme'");
            }
            if ((!args || args.serverpool_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverpool_id'");
            }
            if ((!args || args.target_port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target_port'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["health_check_scheme"] = args ? args.health_check_scheme : undefined;
            resourceInputs["http_only"] = args ? args.http_only : undefined;
            resourceInputs["http_port"] = args ? args.http_port : undefined;
            resourceInputs["https_port"] = args ? args.https_port : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["ip_address"] = args ? args.ip_address : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["same_site"] = args ? args.same_site : undefined;
            resourceInputs["scheme"] = args ? args.scheme : undefined;
            resourceInputs["secure"] = args ? args.secure : undefined;
            resourceInputs["serverpool_id"] = args ? args.serverpool_id : undefined;
            resourceInputs["stickySession_name"] = args ? args.stickySession_name : undefined;
            resourceInputs["target_port"] = args ? args.target_port : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["back_end"] = undefined /*out*/;
            resourceInputs["front_end"] = undefined /*out*/;
            resourceInputs["reverseproxy_id"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["back_end"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["front_end"] = undefined /*out*/;
            resourceInputs["health_check_scheme"] = undefined /*out*/;
            resourceInputs["http_only"] = undefined /*out*/;
            resourceInputs["http_port"] = undefined /*out*/;
            resourceInputs["https_port"] = undefined /*out*/;
            resourceInputs["interval"] = undefined /*out*/;
            resourceInputs["ip_address"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["reverseproxy_id"] = undefined /*out*/;
            resourceInputs["same_site"] = undefined /*out*/;
            resourceInputs["scheme"] = undefined /*out*/;
            resourceInputs["secure"] = undefined /*out*/;
            resourceInputs["serverpool_id"] = undefined /*out*/;
            resourceInputs["stickySession_name"] = undefined /*out*/;
            resourceInputs["target_port"] = undefined /*out*/;
            resourceInputs["timeout"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReverseProxy.__pulumiType, name, resourceInputs, opts);
    }
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
