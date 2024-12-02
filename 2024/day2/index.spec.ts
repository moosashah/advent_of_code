import { readTextFile } from "../utils";
import { expect, test } from "bun:test";

const splitToLines = (content: string): string[] =>
  content.split("\n").filter(Boolean);

const isSafe = (levels: string[]): boolean => {
  const diffs = levels.slice(1).map((level, idx) => +level - +levels[idx]);
  if (!diffs.length) return false;

  const asc = diffs.every((diff) => diff >= 1 && diff <= 3);
  const desc = diffs.every((diff) => diff <= -1 && diff >= -3);
  return asc || desc;
};

const isSafeWithDampener = (levels: string[]): boolean => {
  if (isSafe(levels)) return true;
  return levels.some((_, i) => {
    const modifiedLine = [...levels];
    modifiedLine.splice(i, 1);
    return isSafe(modifiedLine);
  });
};

test("part 1 example", () => {
  const content = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`;

  let safe = 0;

  const lines = splitToLines(content);

  lines.forEach((line) => {
    const levels = line.split(" ");
    isSafe(levels) && safe++;
  });

  expect(safe).toBe(2);
});

test("part 1", async () => {
  const path = "./input.txt";

  const content = await readTextFile(path);

  let safe = 0;

  const lines = splitToLines(content);

  lines.forEach((line) => {
    const levels = line.split(" ");
    isSafe(levels) && safe++;
  });

  expect(safe).toBe(534);
});

test("part 2", async () => {
  const path = "./input.txt";

  const content = await readTextFile(path);

  let safe = 0;

  const lines = splitToLines(content);

  lines.forEach((line) => {
    const levels = line.split(" ");
    isSafeWithDampener(levels) && safe++;
  });

  expect(safe).toBe(577);
});
