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

__all__ = ['ReverseProxyArgs', 'ReverseProxy']

@pulumi.input_type
class ReverseProxyArgs:
    def __init__(__self__, *,
                 back_end: pulumi.Input['ReverseProxyBackendArgs'],
                 cloudspace_id: pulumi.Input[str],
                 customer_id: pulumi.Input[str],
                 front_end: pulumi.Input['ReverseProxyFrontEndArgs'],
                 name: pulumi.Input[str],
                 token: pulumi.Input[str],
                 url: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReverseProxy resource.
        """
        pulumi.set(__self__, "back_end", back_end)
        pulumi.set(__self__, "cloudspace_id", cloudspace_id)
        pulumi.set(__self__, "customer_id", customer_id)
        pulumi.set(__self__, "front_end", front_end)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "url", url)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def back_end(self) -> pulumi.Input['ReverseProxyBackendArgs']:
        return pulumi.get(self, "back_end")

    @back_end.setter
    def back_end(self, value: pulumi.Input['ReverseProxyBackendArgs']):
        pulumi.set(self, "back_end", value)

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
    def front_end(self) -> pulumi.Input['ReverseProxyFrontEndArgs']:
        return pulumi.get(self, "front_end")

    @front_end.setter
    def front_end(self, value: pulumi.Input['ReverseProxyFrontEndArgs']):
        pulumi.set(self, "front_end", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

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
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


class ReverseProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 back_end: Optional[pulumi.Input[pulumi.InputType['ReverseProxyBackendArgs']]] = None,
                 cloudspace_id: Optional[pulumi.Input[str]] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 front_end: Optional[pulumi.Input[pulumi.InputType['ReverseProxyFrontEndArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ReverseProxy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReverseProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ReverseProxy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ReverseProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReverseProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 back_end: Optional[pulumi.Input[pulumi.InputType['ReverseProxyBackendArgs']]] = None,
                 cloudspace_id: Optional[pulumi.Input[str]] = None,
                 customer_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 front_end: Optional[pulumi.Input[pulumi.InputType['ReverseProxyFrontEndArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReverseProxyArgs.__new__(ReverseProxyArgs)

            if back_end is None and not opts.urn:
                raise TypeError("Missing required property 'back_end'")
            __props__.__dict__["back_end"] = back_end
            if cloudspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'cloudspace_id'")
            __props__.__dict__["cloudspace_id"] = cloudspace_id
            if customer_id is None and not opts.urn:
                raise TypeError("Missing required property 'customer_id'")
            __props__.__dict__["customer_id"] = None if customer_id is None else pulumi.Output.secret(customer_id)
            __props__.__dict__["description"] = description
            if front_end is None and not opts.urn:
                raise TypeError("Missing required property 'front_end'")
            __props__.__dict__["front_end"] = front_end
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = None if url is None else pulumi.Output.secret(url)
            __props__.__dict__["reverseproxy_id"] = None
            __props__.__dict__["type"] = None
        super(ReverseProxy, __self__).__init__(
            'vco:ingress:ReverseProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReverseProxy':
        """
        Get an existing ReverseProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReverseProxyArgs.__new__(ReverseProxyArgs)

        __props__.__dict__["back_end"] = None
        __props__.__dict__["cloudspace_id"] = None
        __props__.__dict__["customer_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["front_end"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["reverseproxy_id"] = None
        __props__.__dict__["token"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["url"] = None
        return ReverseProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def back_end(self) -> pulumi.Output['outputs.ReverseProxyBackend']:
        return pulumi.get(self, "back_end")

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
    def description(self) -> pulumi.Output[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def front_end(self) -> pulumi.Output['outputs.ReverseProxyFrontEnd']:
        return pulumi.get(self, "front_end")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def reverseproxy_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "reverseproxy_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

