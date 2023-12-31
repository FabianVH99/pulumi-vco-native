// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Cloudspace extends pulumi.CustomResource {
    /**
     * Get an existing Cloudspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cloudspace {
        return new Cloudspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:base:Cloudspace';

    /**
     * Returns true if the given object is an instance of Cloudspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cloudspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cloudspace.__pulumiType;
    }

    public readonly cdrom_id!: pulumi.Output<number | undefined>;
    public /*out*/ readonly cloudspace_id!: pulumi.Output<string>;
    public /*out*/ readonly cloudspace_mode!: pulumi.Output<string>;
    public /*out*/ readonly creation_time!: pulumi.Output<number>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly disk_size!: pulumi.Output<number | undefined>;
    public readonly external_network_id!: pulumi.Output<number>;
    public /*out*/ readonly external_network_ip!: pulumi.Output<string>;
    public readonly host!: pulumi.Output<string | undefined>;
    public readonly image_id!: pulumi.Output<number | undefined>;
    public readonly local_domain!: pulumi.Output<string | undefined>;
    public readonly location!: pulumi.Output<string>;
    public readonly memory!: pulumi.Output<number | undefined>;
    public readonly memory_quota!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly parent_cloudspace_id!: pulumi.Output<string | undefined>;
    public readonly private!: pulumi.Output<boolean>;
    public readonly private_network!: pulumi.Output<string>;
    public readonly public_ip_quota!: pulumi.Output<number | undefined>;
    public /*out*/ readonly router_type!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;
    public /*out*/ readonly update_time!: pulumi.Output<number>;
    public readonly url!: pulumi.Output<string>;
    public readonly vcpu_quota!: pulumi.Output<number | undefined>;
    public readonly vcpus!: pulumi.Output<number | undefined>;
    public readonly vdisk_space_quota!: pulumi.Output<number | undefined>;

    /**
     * Create a Cloudspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.external_network_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_id'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.private === undefined) && !opts.urn) {
                throw new Error("Missing required property 'private'");
            }
            if ((!args || args.private_network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'private_network'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cdrom_id"] = args ? args.cdrom_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["disk_size"] = args ? args.disk_size : undefined;
            resourceInputs["external_network_id"] = args ? args.external_network_id : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["image_id"] = args ? args.image_id : undefined;
            resourceInputs["local_domain"] = args ? args.local_domain : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["memory_quota"] = args ? args.memory_quota : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent_cloudspace_id"] = args ? args.parent_cloudspace_id : undefined;
            resourceInputs["private"] = args ? args.private : undefined;
            resourceInputs["private_network"] = args ? args.private_network : undefined;
            resourceInputs["public_ip_quota"] = args ? args.public_ip_quota : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["vcpu_quota"] = args ? args.vcpu_quota : undefined;
            resourceInputs["vcpus"] = args ? args.vcpus : undefined;
            resourceInputs["vdisk_space_quota"] = args ? args.vdisk_space_quota : undefined;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["cloudspace_mode"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["router_type"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
        } else {
            resourceInputs["cdrom_id"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["cloudspace_mode"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["external_network_id"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["image_id"] = undefined /*out*/;
            resourceInputs["local_domain"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["memory"] = undefined /*out*/;
            resourceInputs["memory_quota"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent_cloudspace_id"] = undefined /*out*/;
            resourceInputs["private"] = undefined /*out*/;
            resourceInputs["private_network"] = undefined /*out*/;
            resourceInputs["public_ip_quota"] = undefined /*out*/;
            resourceInputs["router_type"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vcpu_quota"] = undefined /*out*/;
            resourceInputs["vcpus"] = undefined /*out*/;
            resourceInputs["vdisk_space_quota"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cloudspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cloudspace resource.
 */
export interface CloudspaceArgs {
    cdrom_id?: pulumi.Input<number>;
    customerID: pulumi.Input<string>;
    disk_size?: pulumi.Input<number>;
    external_network_id: pulumi.Input<number>;
    host?: pulumi.Input<string>;
    image_id?: pulumi.Input<number>;
    local_domain?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    memory?: pulumi.Input<number>;
    memory_quota?: pulumi.Input<number>;
    name: pulumi.Input<string>;
    parent_cloudspace_id?: pulumi.Input<string>;
    private: pulumi.Input<boolean>;
    private_network: pulumi.Input<string>;
    public_ip_quota?: pulumi.Input<number>;
    token: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vcpu_quota?: pulumi.Input<number>;
    vcpus?: pulumi.Input<number>;
    vdisk_space_quota?: pulumi.Input<number>;
}
