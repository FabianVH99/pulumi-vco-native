import * as outputs from "../types/output";
export declare namespace resources {
    interface BackEndState {
        serverpool_id: string;
        serverpool_name: string;
        target_port: number;
    }
    interface CpuTopology {
        cores: number;
        sockets: number;
        threads: number;
    }
    interface Endpoint {
        address: string;
        name: string;
        port: number;
        private_address: string;
        private_port: number;
        psk: string;
        user: string;
    }
    interface FirewallCustom {
        cdrom_id?: number;
        disk_size?: number;
        image_id?: number;
        memory?: number;
        type?: string;
        vcpus?: number;
    }
    interface FrontEnd {
        ip_address?: string;
        port: number;
        tls?: outputs.resources.TLS;
    }
    interface HealthCheck {
        interval?: number;
        path?: string;
        port?: number;
        scheme?: string;
        timeout?: number;
    }
    interface IOTune {
        iops: number;
    }
    interface LetsEncrypt {
        email: string;
        enabled: boolean;
    }
    interface NetworkInterface {
        device_name: string;
        external_cloudspace_id: string;
        ip_address: string;
        mac_address: string;
        model: string;
        network_id: number;
        nic_type: string;
    }
    interface Options {
        health_check?: outputs.resources.HealthCheck;
        sticky_session_cookie?: outputs.resources.StickySessionCookie;
    }
    interface OsAccount {
        login: string;
        password: string;
    }
    interface ResourceLimits {
        external_network_quota: number;
        memory_quota: number;
        public_ip_quota: number;
        vcpu_quota: number;
        vdisk_space_quota: number;
    }
    interface ReverseProxyBackend {
        options?: outputs.resources.Options;
        scheme: string;
        serverpool_id: string;
        target_port: number;
    }
    interface ReverseProxyFrontEnd {
        domain: string;
        http_port?: number;
        https_port?: number;
        ip_address?: string;
        letsencrypt: outputs.resources.LetsEncrypt;
        scheme: string;
    }
    interface ServerPoolHost {
        address: string;
        host_id: string;
    }
    interface StickySessionCookie {
        http_only?: boolean;
        name?: string;
        same_site?: string;
        secure?: boolean;
    }
    interface TLS {
        domain?: string;
        is_enabled?: boolean;
        tls_termination?: boolean;
    }
    interface VmDisk {
        description: string;
        disk_id: number;
        disk_name: string;
        disk_size: number;
        disk_type: string;
        exposed: boolean;
        order: string;
        pci_bus: number;
        pci_slot: number;
        status: string;
    }
}
