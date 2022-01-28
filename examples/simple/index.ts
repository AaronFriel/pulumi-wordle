import * as wordle from "@frielforreal/pulumi-wordle";

const random = new wordle.Wordle("wordle", { word: "raise" });

export const output = random.result;
