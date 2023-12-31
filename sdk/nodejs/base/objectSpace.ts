// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ObjectSpace extends pulumi.CustomResource {
    /**
     * Get an existing ObjectSpace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ObjectSpace {
        return new ObjectSpace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:base:ObjectSpace';

    /**
     * Returns true if the given object is an instance of ObjectSpace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectSpace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectSpace.__pulumiType;
    }

    public /*out*/ readonly access_key!: pulumi.Output<string | undefined>;
    public readonly cloudspace_id!: pulumi.Output<string | undefined>;
    public /*out*/ readonly creation_time!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly domain!: pulumi.Output<string | undefined>;
    public readonly external_network!: pulumi.Output<number | undefined>;
    public readonly letsencrypt!: pulumi.Output<boolean | undefined>;
    public readonly letsencrypt_email!: pulumi.Output<string | undefined>;
    public readonly location!: pulumi.Output<string>;
    public /*out*/ readonly objectspace_id!: pulumi.Output<string>;
    public /*out*/ readonly objectspace_name!: pulumi.Output<string>;
    public /*out*/ readonly secret!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly subnet!: pulumi.Output<string | undefined>;
    public readonly token!: pulumi.Output<string>;
    public /*out*/ readonly update_time!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ObjectSpace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectSpaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
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
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["external_network"] = args ? args.external_network : undefined;
            resourceInputs["letsencrypt"] = args ? args.letsencrypt : undefined;
            resourceInputs["letsencrypt_email"] = args ? args.letsencrypt_email : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["access_key"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["objectspace_id"] = undefined /*out*/;
            resourceInputs["objectspace_name"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
        } else {
            resourceInputs["access_key"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["external_network"] = undefined /*out*/;
            resourceInputs["letsencrypt"] = undefined /*out*/;
            resourceInputs["letsencrypt_email"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["objectspace_id"] = undefined /*out*/;
            resourceInputs["objectspace_name"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subnet"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectSpace.__pulumiType, name, resourceInputs, opts);
    }
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
