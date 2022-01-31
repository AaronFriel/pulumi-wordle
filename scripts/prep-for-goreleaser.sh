#!/bin/bash

PACKAGE_VERSION=$(pulumictl get version --language javascript)

sed -e 's/\0.0.2/'$PACKAGE_VERSION'/g' < provider/cmd/pulumi-resource-wordle/package.json > package.json
