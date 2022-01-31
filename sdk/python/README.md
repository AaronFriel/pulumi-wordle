<p align="center">
  <img width="460" height="300" src="pulumi-wordle-logo-transparent.svg">
</p>

# Wordle Pulumi Provider

Play Wordle while writing infrastructure as code. If you're working from an office, your boss won't
even know that you aren't deploying state of the art cloud managed NAT gateways!

Every day the game resets, and the word list should be the same as the original Wordle's. But unlike
the real Wordle this one gives you unlimited retries, just like deploying real infrastructure.

## How to use

This provider is not yet published, so you'll need to build this locally and put the bin dir on your PATH for examples to work!

## Example

An example of using the single resource defined in this example is in `examples/ts`.

```ts
import * as wordle from "@frielforreal/pulumi-wordle";

const random = new wordle.Wordle("wordle", { word: "raise" });
```

```
$ PATH=../../bin/:$PATH pulumi up --skip-preview
Please choose a stack, or create a new one: dev
Updating (dev)

View Live: https://app.pulumi.com/friel/ts/dev/updates/3

     Type                    Name        Status      Info
     pulumi:pulumi:Stack     simple-dev              2 warnings
 ~   └─ wordle:index:Wordle  wordle      updated     [diff: ~word]

Outputs:
  ~ output: [
        [0]: "🟫🟩🟩🟨🟫"
        [1]: "🟫🟩🟩🟩🟩"
      + [2]: "🟩🟩🟩🟩🟩"
    ]

Resources:
    ~ 1 updated
    1 unchanged

Duration: 2s
```

## Build and Test

```bash
# build and install the resource provider plugin
$ make build install

# test
$ cd examples/ts
$ yarn link @frielforreal/pulumi-wordle
$ yarn install
$ pulumi stack init test
$ PATH=../../bin:$PATH pulumi up
```

## References

Other resources for learning about the Pulumi resource model:
* [Pulumi Kubernetes provider](https://github.com/pulumi/pulumi-kubernetes/blob/master/provider/pkg/provider/provider.go)
* [Pulumi Terraform Remote State provider](https://github.com/pulumi/pulumi-terraform/blob/master/provider/cmd/pulumi-resource-terraform/provider.go)
* [Dynamic Providers](https://www.pulumi.com/docs/intro/concepts/programming-model/#dynamicproviders)
