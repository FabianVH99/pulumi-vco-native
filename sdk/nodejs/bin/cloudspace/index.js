"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.VirtualMachine = exports.PortForward = exports.ExternalNetwork = exports.ConnectedCloudspace = exports.AntiAffinityGroup = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
exports.AntiAffinityGroup = null;
utilities.lazyLoad(exports, ["AntiAffinityGroup"], () => require("./antiAffinityGroup"));
exports.ConnectedCloudspace = null;
utilities.lazyLoad(exports, ["ConnectedCloudspace"], () => require("./connectedCloudspace"));
exports.ExternalNetwork = null;
utilities.lazyLoad(exports, ["ExternalNetwork"], () => require("./externalNetwork"));
exports.PortForward = null;
utilities.lazyLoad(exports, ["PortForward"], () => require("./portForward"));
exports.VirtualMachine = null;
utilities.lazyLoad(exports, ["VirtualMachine"], () => require("./virtualMachine"));
const _module = {
    version: utilities.getVersion(),
    construct: (name, type, urn) => {
        switch (type) {
            case "vco:cloudspace:AntiAffinityGroup":
                return new exports.AntiAffinityGroup(name, undefined, { urn });
            case "vco:cloudspace:ConnectedCloudspace":
                return new exports.ConnectedCloudspace(name, undefined, { urn });
            case "vco:cloudspace:ExternalNetwork":
                return new exports.ExternalNetwork(name, undefined, { urn });
            case "vco:cloudspace:PortForward":
                return new exports.PortForward(name, undefined, { urn });
            case "vco:cloudspace:VirtualMachine":
                return new exports.VirtualMachine(name, undefined, { urn });
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vco", "cloudspace", _module);
//# sourceMappingURL=index.js.map