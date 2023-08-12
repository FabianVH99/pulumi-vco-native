import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
export declare namespace base {
}
export declare namespace cloudspace {
}
export declare namespace disk {
}
export declare namespace ingress {
    interface HealthCheckArgs {
        interval?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
        port?: pulumi.Input<number>;
        scheme?: pulumi.Input<string>;
        timeout?: pulumi.Input<number>;
    }
    interface LetsEncryptArgs {
        email: pulumi.Input<string>;
        enabled: pulumi.Input<boolean>;
    }
    interface OptionsArgs {
        health_check?: pulumi.Input<inputs.ingress.HealthCheckArgs>;
        sticky_session_cookie?: pulumi.Input<inputs.ingress.StickySessionCookieArgs>;
    }
    interface ReverseProxyBackendArgs {
        options?: pulumi.Input<inputs.ingress.OptionsArgs>;
        scheme: pulumi.Input<string>;
        serverpool_id: pulumi.Input<string>;
        target_port: pulumi.Input<number>;
    }
    interface ReverseProxyFrontEndArgs {
        domain: pulumi.Input<string>;
        http_port?: pulumi.Input<number>;
        https_port?: pulumi.Input<number>;
        ip_address?: pulumi.Input<string>;
        letsencrypt: pulumi.Input<inputs.ingress.LetsEncryptArgs>;
        scheme: pulumi.Input<string>;
    }
    interface StickySessionCookieArgs {
        http_only?: pulumi.Input<boolean>;
        name?: pulumi.Input<string>;
        same_site?: pulumi.Input<string>;
        secure?: pulumi.Input<boolean>;
    }
}
