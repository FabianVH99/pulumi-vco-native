// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:ingress:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    public /*out*/ readonly back_end!: pulumi.Output<outputs.ingress.BackEnd>;
    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly domain!: pulumi.Output<string | undefined>;
    public /*out*/ readonly front_end!: pulumi.Output<outputs.ingress.FrontEnd>;
    public readonly ip_address!: pulumi.Output<string | undefined>;
    public readonly is_enabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly loadbalancer_id!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number>;
    public readonly serverpool_id!: pulumi.Output<string>;
    public readonly target_port!: pulumi.Output<number>;
    public readonly tls_termination!: pulumi.Output<boolean | undefined>;
    public readonly token!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
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
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["ip_address"] = args ? args.ip_address : undefined;
            resourceInputs["is_enabled"] = args ? args.is_enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["serverpool_id"] = args ? args.serverpool_id : undefined;
            resourceInputs["target_port"] = args ? args.target_port : undefined;
            resourceInputs["tls_termination"] = args ? args.tls_termination : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["back_end"] = undefined /*out*/;
            resourceInputs["front_end"] = undefined /*out*/;
            resourceInputs["loadbalancer_id"] = undefined /*out*/;
        } else {
            resourceInputs["back_end"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["front_end"] = undefined /*out*/;
            resourceInputs["ip_address"] = undefined /*out*/;
            resourceInputs["is_enabled"] = undefined /*out*/;
            resourceInputs["loadbalancer_id"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["serverpool_id"] = undefined /*out*/;
            resourceInputs["target_port"] = undefined /*out*/;
            resourceInputs["tls_termination"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
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
