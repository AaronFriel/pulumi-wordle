# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .wordle import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "wordle",
  "mod": "index",
  "fqn": "pulumi_wordle",
  "classes": {
   "wordle:index:Wordle": "Wordle"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "wordle",
  "token": "pulumi:providers:wordle",
  "fqn": "pulumi_wordle",
  "class": "Provider"
 }
]
"""
)