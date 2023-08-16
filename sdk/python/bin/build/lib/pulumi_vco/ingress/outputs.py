# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'BackEnd',
    'FrontEnd',
    'HealthCheck',
    'LetsEncrypt',
    'Options',
    'ReverseProxyBackend',
    'ReverseProxyFrontEnd',
    'ServerPoolHost',
    'StickySessionCookie',
    'TLS',
]

@pulumi.output_type
class BackEnd(dict):
    def __init__(__self__, *,
                 serverpool_id: str,
                 target_port: int,
                 serverpool_name: Optional[str] = None):
        pulumi.set(__self__, "serverpool_id", serverpool_id)
        pulumi.set(__self__, "target_port", target_port)
        if serverpool_name is not None:
            pulumi.set(__self__, "serverpool_name", serverpool_name)

    @property
    @pulumi.getter
    def serverpool_id(self) -> str:
        return pulumi.get(self, "serverpool_id")

    @property
    @pulumi.getter
    def target_port(self) -> int:
        return pulumi.get(self, "target_port")

    @property
    @pulumi.getter
    def serverpool_name(self) -> Optional[str]:
        return pulumi.get(self, "serverpool_name")


@pulumi.output_type
class FrontEnd(dict):
    def __init__(__self__, *,
                 port: int,
                 ip_address: Optional[str] = None,
                 tls: Optional['outputs.TLS'] = None):
        pulumi.set(__self__, "port", port)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def ip_address(self) -> Optional[str]:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def tls(self) -> Optional['outputs.TLS']:
        return pulumi.get(self, "tls")


@pulumi.output_type
class HealthCheck(dict):
    def __init__(__self__, *,
                 interval: Optional[int] = None,
                 path: Optional[str] = None,
                 port: Optional[int] = None,
                 scheme: Optional[str] = None,
                 timeout: Optional[int] = None):
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if scheme is not None:
            pulumi.set(__self__, "scheme", scheme)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def scheme(self) -> Optional[str]:
        return pulumi.get(self, "scheme")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")


@pulumi.output_type
class LetsEncrypt(dict):
    def __init__(__self__, *,
                 email: str,
                 enabled: bool):
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")


@pulumi.output_type
class Options(dict):
    def __init__(__self__, *,
                 health_check: Optional['outputs.HealthCheck'] = None,
                 sticky_session_cookie: Optional['outputs.StickySessionCookie'] = None):
        if health_check is not None:
            pulumi.set(__self__, "health_check", health_check)
        if sticky_session_cookie is not None:
            pulumi.set(__self__, "sticky_session_cookie", sticky_session_cookie)

    @property
    @pulumi.getter
    def health_check(self) -> Optional['outputs.HealthCheck']:
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter
    def sticky_session_cookie(self) -> Optional['outputs.StickySessionCookie']:
        return pulumi.get(self, "sticky_session_cookie")


@pulumi.output_type
class ReverseProxyBackend(dict):
    def __init__(__self__, *,
                 scheme: str,
                 serverpool_id: str,
                 target_port: int,
                 options: Optional['outputs.Options'] = None):
        pulumi.set(__self__, "scheme", scheme)
        pulumi.set(__self__, "serverpool_id", serverpool_id)
        pulumi.set(__self__, "target_port", target_port)
        if options is not None:
            pulumi.set(__self__, "options", options)

    @property
    @pulumi.getter
    def scheme(self) -> str:
        return pulumi.get(self, "scheme")

    @property
    @pulumi.getter
    def serverpool_id(self) -> str:
        return pulumi.get(self, "serverpool_id")

    @property
    @pulumi.getter
    def target_port(self) -> int:
        return pulumi.get(self, "target_port")

    @property
    @pulumi.getter
    def options(self) -> Optional['outputs.Options']:
        return pulumi.get(self, "options")


@pulumi.output_type
class ReverseProxyFrontEnd(dict):
    def __init__(__self__, *,
                 domain: str,
                 letsencrypt: 'outputs.LetsEncrypt',
                 scheme: str,
                 http_port: Optional[int] = None,
                 https_port: Optional[int] = None,
                 ip_address: Optional[str] = None):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "letsencrypt", letsencrypt)
        pulumi.set(__self__, "scheme", scheme)
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def letsencrypt(self) -> 'outputs.LetsEncrypt':
        return pulumi.get(self, "letsencrypt")

    @property
    @pulumi.getter
    def scheme(self) -> str:
        return pulumi.get(self, "scheme")

    @property
    @pulumi.getter
    def http_port(self) -> Optional[int]:
        return pulumi.get(self, "http_port")

    @property
    @pulumi.getter
    def https_port(self) -> Optional[int]:
        return pulumi.get(self, "https_port")

    @property
    @pulumi.getter
    def ip_address(self) -> Optional[str]:
        return pulumi.get(self, "ip_address")


@pulumi.output_type
class ServerPoolHost(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 host_id: Optional[str] = None):
        if address is not None:
            pulumi.set(__self__, "address", address)
        if host_id is not None:
            pulumi.set(__self__, "host_id", host_id)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def host_id(self) -> Optional[str]:
        return pulumi.get(self, "host_id")


@pulumi.output_type
class StickySessionCookie(dict):
    def __init__(__self__, *,
                 http_only: Optional[bool] = None,
                 name: Optional[str] = None,
                 same_site: Optional[str] = None,
                 secure: Optional[bool] = None):
        if http_only is not None:
            pulumi.set(__self__, "http_only", http_only)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if same_site is not None:
            pulumi.set(__self__, "same_site", same_site)
        if secure is not None:
            pulumi.set(__self__, "secure", secure)

    @property
    @pulumi.getter
    def http_only(self) -> Optional[bool]:
        return pulumi.get(self, "http_only")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def same_site(self) -> Optional[str]:
        return pulumi.get(self, "same_site")

    @property
    @pulumi.getter
    def secure(self) -> Optional[bool]:
        return pulumi.get(self, "secure")


@pulumi.output_type
class TLS(dict):
    def __init__(__self__, *,
                 domain: Optional[str] = None,
                 is_enabled: Optional[bool] = None,
                 tls_termination: Optional[bool] = None):
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if tls_termination is not None:
            pulumi.set(__self__, "tls_termination", tls_termination)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def is_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def tls_termination(self) -> Optional[bool]:
        return pulumi.get(self, "tls_termination")


