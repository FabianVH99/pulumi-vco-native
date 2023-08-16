import * as pulumi from "@pulumi/pulumi";
export declare class Cloudspace extends pulumi.CustomResource {
    /**
     * Get an existing Cloudspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cloudspace;
    /**
     * Returns true if the given object is an instance of Cloudspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Cloudspace;
    readonly cdrom_id: pulumi.Output<number | undefined>;
    readonly cloudspace_id: pulumi.Output<string>;
    readonly cloudspace_mode: pulumi.Output<string>;
    readonly creation_time: pulumi.Output<number>;
    readonly customerID: pulumi.Output<string>;
    readonly disk_size: pulumi.Output<number | undefined>;
    readonly external_network_id: pulumi.Output<number>;
    readonly external_network_ip: pulumi.Output<string>;
    readonly host: pulumi.Output<string | undefined>;
    readonly image_id: pulumi.Output<number | undefined>;
    readonly local_domain: pulumi.Output<string>;
    readonly location: pulumi.Output<string>;
    readonly memory: pulumi.Output<number | undefined>;
    readonly memory_quota: pulumi.Output<number | undefined>;
    readonly name: pulumi.Output<string>;
    readonly parent_cloudspace_id: pulumi.Output<string | undefined>;
    readonly private: pulumi.Output<boolean>;
    readonly private_network: pulumi.Output<string>;
    readonly public_ip_quota: pulumi.Output<number | undefined>;
    readonly router_type: pulumi.Output<string>;
    readonly status: pulumi.Output<string>;
    readonly token: pulumi.Output<string>;
    readonly type: pulumi.Output<string | undefined>;
    readonly update_time: pulumi.Output<number>;
    readonly url: pulumi.Output<string>;
    readonly vcpu_quota: pulumi.Output<number | undefined>;
    readonly vcpus: pulumi.Output<number | undefined>;
    readonly vdisk_space_quota: pulumi.Output<number | undefined>;
    /**
     * Create a Cloudspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudspaceArgs, opts?: pulumi.CustomResourceOptions);
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
