"""A Python Pulumi program"""

import pulumi
import pulumi_wordle as wordle

Wordle = wordle.Wordle("example", word="infra")

pulumi.export("result", Wordle.result)
