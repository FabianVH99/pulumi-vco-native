[![npm version](https://badge.fury.io/js/@fabianv-cloud%2Fvco.svg)](https://badge.fury.io/js/@fabianv-cloud%2Fvco)
[![PyPI version](https://badge.fury.io/py/pulumi-vco.svg)](https://badge.fury.io/py/pulumi-vco)
[![NuGet version](https://badge.fury.io/nu/Pulumi.Vco.svg)](https://badge.fury.io/nu/Pulumi.Vco)
# Pulumi Vco Native

A Pulumi provider for Whitesky.Cloud Virtual Cloud Operators.
<br><br>
The Whitesky.Cloud Vco Provider for Pulumi enables you to manipulate resources in your vco portal.
Please read the documentation or consult the API docs in your vco portal in order to properly understand how the resources work.

# Install plugin binary

Before using this package you need to install the resource-vco plugin. The latest release of which can be found [here](https://github.com/fabianv-cloud/pulumi-vco-native).
You then have to install the plugin on your system by storing the contents of the zipped file in: <br>
* Windows: %USERPROFILE%\\.pulumi\plugins\resource-vco-{VERSION}
* Linux and MacOS: ~/.pulumi/plugins/resource-vco-{VERSION}
<br>

When you have moved the plugin to the appropriate directory you can install it using: 
```commandline
pulumi plugin install resource vco {VERSION}
```
> **_IMPORTANT:_**  Please make sure the version of the plugin you are using is equivalent to the version of the package you are using. If you use v1.0.0 of the pulumi-vco pip package for example, you should also use the v1.0.0 release of the plugin.

# Install package
You can find links to the various packages at the top of the Readme. Once you have installed the desired package
you can import it into your program. Please consult the [examples](https://github.com/fabianv-cloud/pulumi-vco-native/tree/main/examples)
to figure out how to get started.

## npm

You can install the *@fabianv-cloud/vco* package using either npm or yarn.

npm:
```commandline
npm i @fabianv-cloud/vco
```
yarn:
```commandline
yarn add @fabianv-cloud/vco
```

## PyPI

You can install the *pulumi-vco* package with pip:

```commandline
pip install pulumi-vco
```

## NuGet

You can install the *Pulumi.Vco* package using the NuGet package manager UI or CLI:

.NET CLI:
```commandline
dotnet add package Pulumi.Vco --version 1.0.0
```
