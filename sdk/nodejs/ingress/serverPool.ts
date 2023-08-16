// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ServerPool extends pulumi.CustomResource {
    /**
     * Get an existing ServerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerPool {
        return new ServerPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:ingress:ServerPool';

    /**
     * Returns true if the given object is an instance of ServerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerPool.__pulumiType;
    }

    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public /*out*/ readonly hosts!: pulumi.Output<outputs.ingress.ServerPoolHost[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly serverpool_id!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
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
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["hosts"] = undefined /*out*/;
            resourceInputs["serverpool_id"] = undefined /*out*/;
        } else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["hosts"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serverpool_id"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerPool.__pulumiType, name, resourceInputs, opts);
    }
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
