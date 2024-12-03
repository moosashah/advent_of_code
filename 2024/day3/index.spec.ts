import { expect, test } from "bun:test";
import { readTextFile } from "../utils";

function extractMulMatches(content: string): string[] {
  return content.match(/mul\(\d{1,3},\d{1,3}\)/g) || [];
}

function removeDontMatches(content: string): string {
  return content.replace(/don't\(\)(.*?)do\(\)/gs, "");
}

function multiplyMatches(matches: string[]): number {
  return matches.reduce((acc, match) => {
    const lp = match.indexOf("(");
    const rp = match.indexOf(")");
    const nums = match
      .slice(lp + 1, rp)
      .split(",")
      .map(Number);
    acc += nums[0] * nums[1];

    return acc;
  }, 0);
}

test("part 1 example", () => {
  const content = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`;

  const matches = extractMulMatches(content);
  expect(matches).toEqual(["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]);

  const total = multiplyMatches(matches);

  expect(total).toBe(161);
});

test("part 1", async () => {
  const content = await readTextFile("./input.txt");

  const matches = extractMulMatches(content);

  const total = multiplyMatches(matches);

  expect(total).toBe(166630675);
});

test("part 2 example", async () => {
  const content =
    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";

  const cleaned = removeDontMatches(content);

  const matches = extractMulMatches(cleaned);

  const total = multiplyMatches(matches);

  expect(total).toBe(48);
});

test("part 2", async () => {
  const content = await readTextFile("./input.txt");

  //catch "don't() until do() and replace with empty string"
  const cleaned = removeDontMatches(content);

  //then I just do another pass and capture mul groups
  const matches = extractMulMatches(cleaned);

  const total = multiplyMatches(matches);

  expect(total).toBe(93465710);
});
