import * as pulumi from "@pulumi/pulumi";
import * as vco from "@fabianv-cloud/vco";

// Load sensitive information from pulumi configuration variables
const config = new pulumi.Config();
const url = config.require("url");
const token = config.require("token");
const customerId = config.require("customerId");
const location = config.require("location");

// Create a new cloudspace resource using all required values
const cloudspace = new vco.base.Cloudspace("tsc-cloudspace", {
    url: url,
    token: token,
    customerID: customerId,
    location: location,
    name: "Pulumi_nodejs_cloudspace",
    private_network:"192.168.10.0/24",
    external_network_id: 13,
    private: false,
    local_domain:"pulumi-domain",
});

// Export the cloudspace id as an output. Please note that there is a different between
// cloudspace.Id and cloudspace.Cloudspace_id. The first will return the name of the
// Pulumi resource (in this case, "tsc-cloudspace").
// The second will return the cloudspace id as it is defined in the VCO portal.
export const output = cloudspace.cloudspace_id;
