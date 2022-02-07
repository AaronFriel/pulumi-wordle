using Pulumi;
using Pulumi.Wordle;
using System.Collections.Immutable;

class WordleServer : Stack
{
    [Output] public Output<ImmutableArray<string>> result { get; set; }

    public WordleServer()
    {
        var wordle = new Wordle("wordle", new WordleArgs{
            Word = "azure"
        });

        this.result = wordle.Result;
    }
}
