"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.AntiAffinityGroupVM = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
exports.AntiAffinityGroupVM = null;
utilities.lazyLoad(exports, ["AntiAffinityGroupVM"], () => require("./antiAffinityGroupVM"));
const _module = {
    version: utilities.getVersion(),
    construct: (name, type, urn) => {
        switch (type) {
            case "vco:anti_affinity_group:AntiAffinityGroupVM":
                return new exports.AntiAffinityGroupVM(name, undefined, { urn });
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vco", "anti_affinity_group", _module);
//# sourceMappingURL=index.js.map