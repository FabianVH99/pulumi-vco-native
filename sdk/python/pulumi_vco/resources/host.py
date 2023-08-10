# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HostArgs', 'Host']

@pulumi.input_type
class HostArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 cloudspace_id: pulumi.Input[str],
                 customer_id: pulumi.Input[str],
                 serverpool_id: pulumi.Input[str],
                 token: pulumi.Input[str],
                 url: pulumi.Input[str]):
        """
        The set of arguments for constructing a Host resource.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "cloudspace_id", cloudspace_id)
        pulumi.set(__self__, "customer_id", customer_id)
        pulumi.set(__self__, "serverpool_id", serverpool_id)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

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
    def serverpool_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "serverpool_id")

    @serverpool_id.setter
    def serverpool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "serverpool_id", value)

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


class Host(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 cloudspace_id: Optional[pulumi.Input[str]] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 serverpool_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Host resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Host resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param HostArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 cloudspace_id: Optional[pulumi.Input[str]] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 serverpool_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostArgs.__new__(HostArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            if cloudspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'cloudspace_id'")
            __props__.__dict__["cloudspace_id"] = cloudspace_id
            if customer_id is None and not opts.urn:
                raise TypeError("Missing required property 'customer_id'")
            __props__.__dict__["customer_id"] = customer_id
            if serverpool_id is None and not opts.urn:
                raise TypeError("Missing required property 'serverpool_id'")
            __props__.__dict__["serverpool_id"] = serverpool_id
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = token
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["host_id"] = None
        super(Host, __self__).__init__(
            'vco:resources:Host',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Host':
        """
        Get an existing Host resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HostArgs.__new__(HostArgs)

        __props__.__dict__["address"] = None
        __props__.__dict__["cloudspace_id"] = None
        __props__.__dict__["customer_id"] = None
        __props__.__dict__["host_id"] = None
        __props__.__dict__["serverpool_id"] = None
        __props__.__dict__["token"] = None
        __props__.__dict__["url"] = None
        return Host(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "address")

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
    def host_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter
    def serverpool_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "serverpool_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

