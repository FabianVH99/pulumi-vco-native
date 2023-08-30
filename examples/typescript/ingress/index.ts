import * as pulumi from "@pulumi/pulumi";
import * as vco from "@fabianv-cloud/vco";

const config = new pulumi.Config();
const url = config.require("url");
const token = config.require("token");
const customerId = config.require("customerId");
const location = config.require("location");

const cloudspace = new vco.base.Cloudspace("tsc-cloudspace", {
    url: url,
    token: token,
    customerID: customerId,
    location: location,
    name: "Pulumi_nodejs_cloudspace_ingress",
    private_network:"192.168.10.0/24",
    external_network_id: 13,
    private: false,
    local_domain:"pulumi-domain",
});

const serverpool = new vco.ingress.ServerPool("tsc-serverpool", {
    url: url,
    token: token,
    customerID: customerId,
    cloudspace_id: cloudspace.cloudspace_id,
    name: "Pulumi_nodejs_serverpool",
    description: "sv_pool_test"
});

const host = new vco.ingress.Host("tsc-host", {
    url: url,
    token: token,
    customerID: customerId,
    cloudspace_id: cloudspace.cloudspace_id,
    serverpool_id: serverpool.serverpool_id,
    address: "192.168.10.10"
});

const loadbalancer = new vco.ingress.LoadBalancer("tsc-loadbalancer", {
    url: url,
    token: token,
    customerID: customerId,
    cloudspace_id: cloudspace.cloudspace_id,
    serverpool_id: serverpool.serverpool_id,
    name: "Pulumi_tsc_lb",
    type: "udp",
    port: 25,
    target_port: 25,
    ip_address: cloudspace.external_network_ip,
    is_enabled: true,
    domain: "whiteskycloud-2.try-dns.whitesky.cloud",
    tls_termination: true
});

const reverseproxy = new vco.ingress.ReverseProxy("tsc-reverseproxy", {
    url: url,
    token: token,
    customerID: customerId,
    cloudspace_id: cloudspace.cloudspace_id,
    serverpool_id: serverpool.serverpool_id,
    name: "Pulumi_tsc_rp",
    scheme: "http",
    http_port: 26,
    target_port: 26,
    ip_address: cloudspace.external_network_ip,
    enabled: true,
    domain: "pulumi",
});

export const cs_id = cloudspace.cloudspace_id;
export const sv_id = serverpool.serverpool_id;
export const lb_id = loadbalancer.loadbalancer_id;
export const rp_id = reverseproxy.reverseproxy_id;