import * as wordle from "@frielforreal/pulumi-wordle";

const Wordle: wordle.Wordle = new wordle.Wordle("wordle", { word: "route" });

export const result = Wordle.result;
