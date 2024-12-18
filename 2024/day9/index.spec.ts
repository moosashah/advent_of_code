import { expect, test } from "bun:test";
import { readTextFile } from "../utils";

test("part 1", async () => {
  const text = await readTextFile("./input.txt");

  const map: string[] = [];

  let fileId = 0;

  for (let i = 0; i < text.length; i++) {
    const n = parseInt(text[i]);

    const definingFile = i % 2 === 0;
    if (definingFile) {
      map.push(...Array.from({ length: n }, () => fileId.toString()));
      fileId++;
    } else {
      map.push(...Array.from({ length: n }, () => "."));
    }
  }

  while (true) {
    const last = map.findLastIndex((element) => element !== ".");
    const space = map.findIndex((element) => element === ".");
    if (space > last) {
      break;
    }

    map[space] = map[last];
    map[last] = ".";
  }

  let answer = 0;

  for (let i = 0; i < map.length; i++) {
    if (map[i] === ".") {
      continue;
    }
    answer += i * +map[i];
  }
  expect(answer).toBe(6399153661894);
});
