import * as wordle from "@frielforreal/pulumi-wordle";

const random = new wordle.Wordle("my-random", { word: "truth" });

export const output = random.result;
