# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "0.0.0"
PLUGIN_VERSION = "0.0.0"


def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "vco Pulumi Package - Development Version"


setup(name='pulumi_vco',
      python_requires='>=3.7',
      version=VERSION,
      description="The Whitesky.Cloud Vco Provider for Pulumi enables you to manipulate resources in the vco portal.",
      long_description=readme(),
      long_description_content_type='text/markdown',
      keywords='whitesky.cloud pulumi vco category/utility kind/native',
      url='https://whitesky.cloud/',
      project_urls={
          'Repository': 'https://github.com/fabianv-cloud/pulumi-vco-native'
      },
      packages=find_packages(),
      package_data={
          'pulumi_vco': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'semver>=2.8.1'
      ],
      zip_safe=False)
