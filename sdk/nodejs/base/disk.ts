// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:base:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    public /*out*/ readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public readonly detach!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly disk_id!: pulumi.Output<number>;
    public readonly disk_name!: pulumi.Output<string>;
    public readonly disk_size!: pulumi.Output<number>;
    public readonly disk_type!: pulumi.Output<string>;
    public /*out*/ readonly exposed!: pulumi.Output<boolean>;
    public readonly iops!: pulumi.Output<number | undefined>;
    public /*out*/ readonly iotune!: pulumi.Output<outputs.base.IOTune>;
    public readonly location!: pulumi.Output<string>;
    public /*out*/ readonly model!: pulumi.Output<string>;
    public /*out*/ readonly order!: pulumi.Output<string>;
    public readonly permanently!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly port!: pulumi.Output<number>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;
    public readonly vm_id!: pulumi.Output<number>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.disk_name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disk_name'");
            }
            if ((!args || args.disk_size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disk_size'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["customerID"] = args ? args.customerID : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["detach"] = args ? args.detach : undefined;
            resourceInputs["disk_name"] = args ? args.disk_name : undefined;
            resourceInputs["disk_size"] = args ? args.disk_size : undefined;
            resourceInputs["disk_type"] = args ? args.disk_type : undefined;
            resourceInputs["iops"] = args ? args.iops : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["permanently"] = args ? args.permanently : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vm_id"] = args ? args.vm_id : undefined;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["disk_id"] = undefined /*out*/;
            resourceInputs["exposed"] = undefined /*out*/;
            resourceInputs["iotune"] = undefined /*out*/;
            resourceInputs["model"] = undefined /*out*/;
            resourceInputs["order"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["detach"] = undefined /*out*/;
            resourceInputs["disk_id"] = undefined /*out*/;
            resourceInputs["disk_name"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["disk_type"] = undefined /*out*/;
            resourceInputs["exposed"] = undefined /*out*/;
            resourceInputs["iops"] = undefined /*out*/;
            resourceInputs["iotune"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["model"] = undefined /*out*/;
            resourceInputs["order"] = undefined /*out*/;
            resourceInputs["permanently"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vm_id"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Disk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    customerID: pulumi.Input<string>;
    description: pulumi.Input<string>;
    detach?: pulumi.Input<boolean>;
    disk_name: pulumi.Input<string>;
    disk_size: pulumi.Input<number>;
    disk_type?: pulumi.Input<string>;
    iops?: pulumi.Input<number>;
    location: pulumi.Input<string>;
    permanently?: pulumi.Input<boolean>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vm_id?: pulumi.Input<string>;
}