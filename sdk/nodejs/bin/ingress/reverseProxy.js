"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReverseProxy = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
class ReverseProxy extends pulumi.CustomResource {
    /**
     * Get an existing ReverseProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, opts) {
        return new ReverseProxy(name, undefined, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of ReverseProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReverseProxy.__pulumiType;
    }
    /**
     * Create a ReverseProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name, args, opts) {
        let resourceInputs = {};
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
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
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
            resourceInputs["customerID"] = (args === null || args === void 0 ? void 0 : args.customerID) ? pulumi.secret(args.customerID) : undefined;
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
            resourceInputs["token"] = (args === null || args === void 0 ? void 0 : args.token) ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = (args === null || args === void 0 ? void 0 : args.url) ? pulumi.secret(args.url) : undefined;
            resourceInputs["back_end"] = undefined /*out*/;
            resourceInputs["front_end"] = undefined /*out*/;
            resourceInputs["reverseproxy_id"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        else {
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
exports.ReverseProxy = ReverseProxy;
/** @internal */
ReverseProxy.__pulumiType = 'vco:ingress:ReverseProxy';
//# sourceMappingURL=reverseProxy.js.map