import * as outputs from "../types/output";
export declare namespace cloudspace {
    interface CpuTopology {
        cores: number;
        sockets: number;
        threads: number;
    }
    interface NetworkInterface {
        device_name: string;
        external_cloudspace_id?: string;
        ip_address: string;
        mac_address: string;
        model: string;
        network_id: number;
        nic_type: string;
    }
    interface OsAccount {
        login: string;
        password: string;
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
export declare namespace disk {
    interface Endpoint {
        address: string;
        name: string;
        port: number;
        private_address: string;
        private_port: number;
        psk: string;
        user: string;
    }
}
export declare namespace ingress {
    interface BackEnd {
        serverpool_id: string;
        serverpool_name?: string;
        target_port: number;
    }
    interface FrontEnd {
        ip_address?: string;
        port: number;
        tls?: outputs.ingress.TLS;
    }
    interface HealthCheck {
        interval?: number;
        path?: string;
        port?: number;
        scheme?: string;
        timeout?: number;
    }
    interface LetsEncrypt {
        email?: string;
        enabled?: boolean;
    }
    interface Options {
        health_check?: outputs.ingress.HealthCheck;
        sticky_session_cookie?: outputs.ingress.StickySessionCookie;
    }
    interface ReverseProxyBackend {
        options?: outputs.ingress.Options;
        scheme: string;
        serverpool_id: string;
        target_port: number;
    }
    interface ReverseProxyFrontEnd {
        domain?: string;
        http_port?: number;
        https_port?: number;
        ip_address?: string;
        letsencrypt?: outputs.ingress.LetsEncrypt;
        scheme: string;
    }
    interface ServerPoolHost {
        address?: string;
        host_id?: string;
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
}
