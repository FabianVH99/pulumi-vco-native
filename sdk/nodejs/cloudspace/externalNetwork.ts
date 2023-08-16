// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ExternalNetwork extends pulumi.CustomResource {
    /**
     * Get an existing ExternalNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExternalNetwork {
        return new ExternalNetwork(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:cloudspace:ExternalNetwork';

    /**
     * Returns true if the given object is an instance of ExternalNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalNetwork.__pulumiType;
    }

    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly external_network_id!: pulumi.Output<string>;
    public readonly external_network_ip!: pulumi.Output<string>;
    public readonly external_network_type!: pulumi.Output<string>;
    public readonly metric!: pulumi.Output<number>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ExternalNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExternalNetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.external_network_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_id'");
            }
            if ((!args || args.external_network_ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_ip'");
            }
            if ((!args || args.external_network_type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_type'");
            }
            if ((!args || args.metric === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metric'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["external_network_id"] = args ? args.external_network_id : undefined;
            resourceInputs["external_network_ip"] = args ? args.external_network_ip : undefined;
            resourceInputs["external_network_type"] = args ? args.external_network_type : undefined;
            resourceInputs["metric"] = args ? args.metric : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
        } else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["external_network_id"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["external_network_type"] = undefined /*out*/;
            resourceInputs["metric"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExternalNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExternalNetwork resource.
 */
export interface ExternalNetworkArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    external_network_id: pulumi.Input<string>;
    external_network_ip: pulumi.Input<string>;
    external_network_type: pulumi.Input<string>;
    metric: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
