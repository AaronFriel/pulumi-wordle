# Wordle Pulumi Provider

Play Wordle while writing infrastructure as code. If you're working from an office, your boss won't
even know that you aren't deploying state of the art cloud managed NAT gateways!

Every day the game resets, and the word list should be the same as the original Wordle's. But unlike
the real Wordle this one gives you unlimited retries, just like deploying real infrastructure.

An example of using the single resource defined in this example is in `examples/simple`.

```
$ pulumi up
Please choose a stack, or create a new one: dev
Previewing update (dev)

View Live: https://app.pulumi.com/friel/simple/dev/previews/4cc0fa07-4d90-4c55-980c-e930400d11f1

     Type                    Name        Plan       Info
     pulumi:pulumi:Stack     simple-dev
 ~   └─ wordle:index:Wordle  my-random   update     [diff: ~word]

Outputs:
  ~ output: "🟫🟫🟨🟫🟫\n🟫🟫🟫🟫🟫\n" => output<string>

Resources:
    ~ 1 to update
    1 unchanged

Do you want to perform this update? yes
Updating (dev)

View Live: https://app.pulumi.com/friel/simple/dev/updates/12

     Type                    Name        Status      Info
     pulumi:pulumi:Stack     simple-dev
 ~   └─ wordle:index:Wordle  my-random   updated     [diff: ~word]

Outputs:
  ~ output: "🟫🟫🟨🟫🟫\n🟫🟫🟫🟫🟫\n" => "🟫🟫🟨🟫🟫\n🟫🟫🟫🟫🟫\n🟫🟫🟫🟫🟫\n"

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
$ cd examples/simple
$ yarn link @frielforreal/pulumi-wordle
$ yarn install
$ pulumi stack init test
$ pulumi up
```

## References

Other resources for learning about the Pulumi resource model:
* [Pulumi Kubernetes provider](https://github.com/pulumi/pulumi-kubernetes/blob/master/provider/pkg/provider/provider.go)
* [Pulumi Terraform Remote State provider](https://github.com/pulumi/pulumi-terraform/blob/master/provider/cmd/pulumi-resource-terraform/provider.go)
* [Dynamic Providers](https://www.pulumi.com/docs/intro/concepts/programming-model/#dynamicproviders)
