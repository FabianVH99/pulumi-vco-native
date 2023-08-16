// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class VirtualMachineCD extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineCD resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachineCD {
        return new VirtualMachineCD(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:virtual_machine:VirtualMachineCD';

    /**
     * Returns true if the given object is an instance of VirtualMachineCD.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachineCD {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineCD.__pulumiType;
    }

    public readonly cdrom_id!: pulumi.Output<number>;
    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public /*out*/ readonly description!: pulumi.Output<string>;
    public /*out*/ readonly disk_size!: pulumi.Output<string>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;
    public readonly vm_id!: pulumi.Output<number>;

    /**
     * Create a VirtualMachineCD resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineCDArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cdrom_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cdrom_id'");
            }
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.vm_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vm_id'");
            }
            resourceInputs["cdrom_id"] = args ? args.cdrom_id : undefined;
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["vm_id"] = args ? args.vm_id : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["cdrom_id"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vm_id"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualMachineCD.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachineCD resource.
 */
export interface VirtualMachineCDArgs {
    cdrom_id: pulumi.Input<number>;
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vm_id: pulumi.Input<number>;
}
