# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PortForwardArgs', 'PortForward']

@pulumi.input_type
class PortForwardArgs:
    def __init__(__self__, *,
                 cloudspace_id: pulumi.Input[str],
                 customer_id: pulumi.Input[str],
                 local_port: pulumi.Input[int],
                 protocol: pulumi.Input[str],
                 public_ip: pulumi.Input[str],
                 public_port: pulumi.Input[int],
                 token: pulumi.Input[str],
                 url: pulumi.Input[str],
                 vm_id: pulumi.Input[int],
                 nested_cs_id: Optional[pulumi.Input[str]] = None,
                 till_public_port: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PortForward resource.
        """
        pulumi.set(__self__, "cloudspace_id", cloudspace_id)
        pulumi.set(__self__, "customer_id", customer_id)
        pulumi.set(__self__, "local_port", local_port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "public_ip", public_ip)
        pulumi.set(__self__, "public_port", public_port)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "url", url)
        pulumi.set(__self__, "vm_id", vm_id)
        if nested_cs_id is not None:
            pulumi.set(__self__, "nested_cs_id", nested_cs_id)
        if till_public_port is not None:
            pulumi.set(__self__, "till_public_port", till_public_port)

    @property
    @pulumi.getter
    def cloudspace_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cloudspace_id")

    @cloudspace_id.setter
    def cloudspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloudspace_id", value)

    @property
    @pulumi.getter(name="customerID")
    def customer_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "customer_id")

    @customer_id.setter
    def customer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "customer_id", value)

    @property
    @pulumi.getter
    def local_port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "local_port")

    @local_port.setter
    def local_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "local_port", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def public_ip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "public_ip")

    @public_ip.setter
    def public_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_ip", value)

    @property
    @pulumi.getter
    def public_port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "public_port")

    @public_port.setter
    def public_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "public_port", value)

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
    def vm_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "vm_id")

    @vm_id.setter
    def vm_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "vm_id", value)

    @property
    @pulumi.getter
    def nested_cs_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nested_cs_id")

    @nested_cs_id.setter
    def nested_cs_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nested_cs_id", value)

    @property
    @pulumi.getter
    def till_public_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "till_public_port")

    @till_public_port.setter
    def till_public_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "till_public_port", value)


class PortForward(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudspace_id: Optional[pulumi.Input[str]] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 local_port: Optional[pulumi.Input[int]] = None,
                 nested_cs_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 public_ip: Optional[pulumi.Input[str]] = None,
                 public_port: Optional[pulumi.Input[int]] = None,
                 till_public_port: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vm_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a PortForward resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PortForwardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PortForward resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PortForwardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PortForwardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudspace_id: Optional[pulumi.Input[str]] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 local_port: Optional[pulumi.Input[int]] = None,
                 nested_cs_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 public_ip: Optional[pulumi.Input[str]] = None,
                 public_port: Optional[pulumi.Input[int]] = None,
                 till_public_port: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vm_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PortForwardArgs.__new__(PortForwardArgs)

            if cloudspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'cloudspace_id'")
            __props__.__dict__["cloudspace_id"] = cloudspace_id
            if customer_id is None and not opts.urn:
                raise TypeError("Missing required property 'customer_id'")
            __props__.__dict__["customer_id"] = customer_id
            if local_port is None and not opts.urn:
                raise TypeError("Missing required property 'local_port'")
            __props__.__dict__["local_port"] = local_port
            __props__.__dict__["nested_cs_id"] = nested_cs_id
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            if public_ip is None and not opts.urn:
                raise TypeError("Missing required property 'public_ip'")
            __props__.__dict__["public_ip"] = public_ip
            if public_port is None and not opts.urn:
                raise TypeError("Missing required property 'public_port'")
            __props__.__dict__["public_port"] = public_port
            __props__.__dict__["till_public_port"] = till_public_port
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = token
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            if vm_id is None and not opts.urn:
                raise TypeError("Missing required property 'vm_id'")
            __props__.__dict__["vm_id"] = vm_id
            __props__.__dict__["portforward_id"] = None
        super(PortForward, __self__).__init__(
            'vco:resources:PortForward',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PortForward':
        """
        Get an existing PortForward resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PortForwardArgs.__new__(PortForwardArgs)

        __props__.__dict__["cloudspace_id"] = None
        __props__.__dict__["customer_id"] = None
        __props__.__dict__["local_port"] = None
        __props__.__dict__["nested_cs_id"] = None
        __props__.__dict__["portforward_id"] = None
        __props__.__dict__["protocol"] = None
        __props__.__dict__["public_ip"] = None
        __props__.__dict__["public_port"] = None
        __props__.__dict__["till_public_port"] = None
        __props__.__dict__["token"] = None
        __props__.__dict__["url"] = None
        __props__.__dict__["vm_id"] = None
        return PortForward(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cloudspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloudspace_id")

    @property
    @pulumi.getter(name="customerID")
    def customer_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "customer_id")

    @property
    @pulumi.getter
    def local_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "local_port")

    @property
    @pulumi.getter
    def nested_cs_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "nested_cs_id")

    @property
    @pulumi.getter
    def portforward_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portforward_id")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def public_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter
    def public_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter
    def till_public_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "till_public_port")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vm_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "vm_id")

