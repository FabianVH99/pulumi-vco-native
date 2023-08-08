import * as vco from "@pulumi/vco";

const random = new vco.Random("my-random", { length: 24 });

export const output = random.result;