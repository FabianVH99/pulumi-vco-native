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
from ._inputs import *

__all__ = ['CloudspaceArgs', 'Cloudspace']

@pulumi.input_type
class CloudspaceArgs:
    def __init__(__self__, *,
                 customer_id: pulumi.Input[str],
                 external_network_id: pulumi.Input[int],
                 firewall_custom: pulumi.Input['FirewallCustomArgs'],
                 host: pulumi.Input[str],
                 local_domain: pulumi.Input[str],
                 location: pulumi.Input[str],
                 memory_quota: pulumi.Input[int],
                 name: pulumi.Input[str],
                 parent_cloudspace_id: pulumi.Input[str],
                 private: pulumi.Input[bool],
                 private_network: pulumi.Input[str],
                 public_ip_quota: pulumi.Input[int],
                 token: pulumi.Input[str],
                 url: pulumi.Input[str],
                 vcpu_quota: pulumi.Input[int],
                 vdisk_space_quota: pulumi.Input[int]):
        """
        The set of arguments for constructing a Cloudspace resource.
        """
        pulumi.set(__self__, "customer_id", customer_id)
        pulumi.set(__self__, "external_network_id", external_network_id)
        pulumi.set(__self__, "firewall_custom", firewall_custom)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "local_domain", local_domain)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "memory_quota", memory_quota)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "parent_cloudspace_id", parent_cloudspace_id)
        pulumi.set(__self__, "private", private)
        pulumi.set(__self__, "private_network", private_network)
        pulumi.set(__self__, "public_ip_quota", public_ip_quota)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "url", url)
        pulumi.set(__self__, "vcpu_quota", vcpu_quota)
        pulumi.set(__self__, "vdisk_space_quota", vdisk_space_quota)

    @property
    @pulumi.getter(name="customerID")
    def customer_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "customer_id")

    @customer_id.setter
    def customer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "customer_id", value)

    @property
    @pulumi.getter(name="externalNetworkID")
    def external_network_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "external_network_id")

    @external_network_id.setter
    def external_network_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "external_network_id", value)

    @property
    @pulumi.getter
    def firewall_custom(self) -> pulumi.Input['FirewallCustomArgs']:
        return pulumi.get(self, "firewall_custom")

    @firewall_custom.setter
    def firewall_custom(self, value: pulumi.Input['FirewallCustomArgs']):
        pulumi.set(self, "firewall_custom", value)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def local_domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "local_domain")

    @local_domain.setter
    def local_domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_domain", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def memory_quota(self) -> pulumi.Input[int]:
        return pulumi.get(self, "memory_quota")

    @memory_quota.setter
    def memory_quota(self, value: pulumi.Input[int]):
        pulumi.set(self, "memory_quota", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent_cloudspace_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "parent_cloudspace_id")

    @parent_cloudspace_id.setter
    def parent_cloudspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_cloudspace_id", value)

    @property
    @pulumi.getter
    def private(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "private")

    @private.setter
    def private(self, value: pulumi.Input[bool]):
        pulumi.set(self, "private", value)

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> pulumi.Input[str]:
        return pulumi.get(self, "private_network")

    @private_network.setter
    def private_network(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_network", value)

    @property
    @pulumi.getter
    def public_ip_quota(self) -> pulumi.Input[int]:
        return pulumi.get(self, "public_ip_quota")

    @public_ip_quota.setter
    def public_ip_quota(self, value: pulumi.Input[int]):
        pulumi.set(self, "public_ip_quota", value)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Input[str]:
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[str]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def vcpu_quota(self) -> pulumi.Input[int]:
        return pulumi.get(self, "vcpu_quota")

    @vcpu_quota.setter
    def vcpu_quota(self, value: pulumi.Input[int]):
        pulumi.set(self, "vcpu_quota", value)

    @property
    @pulumi.getter
    def vdisk_space_quota(self) -> pulumi.Input[int]:
        return pulumi.get(self, "vdisk_space_quota")

    @vdisk_space_quota.setter
    def vdisk_space_quota(self, value: pulumi.Input[int]):
        pulumi.set(self, "vdisk_space_quota", value)


class Cloudspace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 external_network_id: Optional[pulumi.Input[int]] = None,
                 firewall_custom: Optional[pulumi.Input[pulumi.InputType['FirewallCustomArgs']]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 local_domain: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 memory_quota: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_cloudspace_id: Optional[pulumi.Input[str]] = None,
                 private: Optional[pulumi.Input[bool]] = None,
                 private_network: Optional[pulumi.Input[str]] = None,
                 public_ip_quota: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vcpu_quota: Optional[pulumi.Input[int]] = None,
                 vdisk_space_quota: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Cloudspace resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudspaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Cloudspace resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CloudspaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudspaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 external_network_id: Optional[pulumi.Input[int]] = None,
                 firewall_custom: Optional[pulumi.Input[pulumi.InputType['FirewallCustomArgs']]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 local_domain: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 memory_quota: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_cloudspace_id: Optional[pulumi.Input[str]] = None,
                 private: Optional[pulumi.Input[bool]] = None,
                 private_network: Optional[pulumi.Input[str]] = None,
                 public_ip_quota: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vcpu_quota: Optional[pulumi.Input[int]] = None,
                 vdisk_space_quota: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudspaceArgs.__new__(CloudspaceArgs)

            if customer_id is None and not opts.urn:
                raise TypeError("Missing required property 'customer_id'")
            __props__.__dict__["customer_id"] = customer_id
            if external_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'external_network_id'")
            __props__.__dict__["external_network_id"] = external_network_id
            if firewall_custom is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_custom'")
            __props__.__dict__["firewall_custom"] = firewall_custom
            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            if local_domain is None and not opts.urn:
                raise TypeError("Missing required property 'local_domain'")
            __props__.__dict__["local_domain"] = local_domain
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if memory_quota is None and not opts.urn:
                raise TypeError("Missing required property 'memory_quota'")
            __props__.__dict__["memory_quota"] = memory_quota
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if parent_cloudspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_cloudspace_id'")
            __props__.__dict__["parent_cloudspace_id"] = parent_cloudspace_id
            if private is None and not opts.urn:
                raise TypeError("Missing required property 'private'")
            __props__.__dict__["private"] = private
            if private_network is None and not opts.urn:
                raise TypeError("Missing required property 'private_network'")
            __props__.__dict__["private_network"] = private_network
            if public_ip_quota is None and not opts.urn:
                raise TypeError("Missing required property 'public_ip_quota'")
            __props__.__dict__["public_ip_quota"] = public_ip_quota
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = token
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            if vcpu_quota is None and not opts.urn:
                raise TypeError("Missing required property 'vcpu_quota'")
            __props__.__dict__["vcpu_quota"] = vcpu_quota
            if vdisk_space_quota is None and not opts.urn:
                raise TypeError("Missing required property 'vdisk_space_quota'")
            __props__.__dict__["vdisk_space_quota"] = vdisk_space_quota
            __props__.__dict__["cloudspace_id"] = None
            __props__.__dict__["cloudspace_mode"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["external_network_id"] = None
            __props__.__dict__["external_network_ip"] = None
            __props__.__dict__["private_network"] = None
            __props__.__dict__["resource_limits"] = None
            __props__.__dict__["router_type"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_time"] = None
        super(Cloudspace, __self__).__init__(
            'vco:resources:Cloudspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cloudspace':
        """
        Get an existing Cloudspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CloudspaceArgs.__new__(CloudspaceArgs)

        __props__.__dict__["cloudspace_id"] = None
        __props__.__dict__["cloudspace_mode"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["customer_id"] = None
        __props__.__dict__["external_network_id"] = None
        __props__.__dict__["external_network_ip"] = None
        __props__.__dict__["firewall_custom"] = None
        __props__.__dict__["host"] = None
        __props__.__dict__["local_domain"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["memory_quota"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parent_cloudspace_id"] = None
        __props__.__dict__["private"] = None
        __props__.__dict__["private_network"] = None
        __props__.__dict__["public_ip_quota"] = None
        __props__.__dict__["resource_limits"] = None
        __props__.__dict__["router_type"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["token"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["url"] = None
        __props__.__dict__["vcpu_quota"] = None
        __props__.__dict__["vdisk_space_quota"] = None
        return Cloudspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cloudspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloudspace_id")

    @property
    @pulumi.getter
    def cloudspace_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloudspace_mode")

    @property
    @pulumi.getter
    def creation_time(self) -> pulumi.Output[int]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="customerID")
    def customer_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "customer_id")

    @property
    @pulumi.getter
    def external_network_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "external_network_id")

    @property
    @pulumi.getter
    def external_network_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "external_network_ip")

    @property
    @pulumi.getter
    def firewall_custom(self) -> pulumi.Output['outputs.FirewallCustom']:
        return pulumi.get(self, "firewall_custom")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def local_domain(self) -> pulumi.Output[str]:
        return pulumi.get(self, "local_domain")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def memory_quota(self) -> pulumi.Output[int]:
        return pulumi.get(self, "memory_quota")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent_cloudspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "parent_cloudspace_id")

    @property
    @pulumi.getter
    def private(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "private")

    @property
    @pulumi.getter
    def private_network(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_network")

    @property
    @pulumi.getter
    def public_ip_quota(self) -> pulumi.Output[int]:
        return pulumi.get(self, "public_ip_quota")

    @property
    @pulumi.getter
    def resource_limits(self) -> pulumi.Output['outputs.ResourceLimits']:
        return pulumi.get(self, "resource_limits")

    @property
    @pulumi.getter
    def router_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "router_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def update_time(self) -> pulumi.Output[int]:
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vcpu_quota(self) -> pulumi.Output[int]:
        return pulumi.get(self, "vcpu_quota")

    @property
    @pulumi.getter
    def vdisk_space_quota(self) -> pulumi.Output[int]:
        return pulumi.get(self, "vdisk_space_quota")

