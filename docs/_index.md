---
title: Wordle
meta_desc: Provides an overview of the Wordle Provider for Pulumi.
layout: overview
---

The Wordle provider for Pulumi can be used to provision a daily Wordle game. Every time you change
the word input, the "result" output is updated!

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as wordle from "@frielforreal/pulumi-wordle";

// important to maximize the # of times Wordle appears in this line
const Wordle: wordle.Wordle = new wordle.Wordle("wordle", {
  word: "trust"
});

export const output = Wordle.result;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_wordle as wordle

Wordle = wordle.Wordle("example", word="infra")

pulumi.export("result", Wordle.result)

```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"

	"github.com/aaronfriel/pulumi/sdk/go/wordle"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		Wordle, err := wordle.NewWordle(ctx, "wordle", &wordle.WordleArgs{
			Word: pulumi.String("cloud"),
		})
		if err != nil {
			return fmt.Errorf("error creating wordle: %v", err)
		}

		ctx.Export("result", Wordle)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Wordle;

class WordleServer : Stack
{
    [Output] public Output<string> result { get; set; }

    public WordleServer()
    {
        var wordle = new Wordle("wordle", new WordleArgs{
            Word = "azure"
        });

        this.result = wordle.Result;
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
