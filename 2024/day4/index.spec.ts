import { test, expect } from "bun:test";
import { readTextFile } from "../utils";

const XMAS = "XMAS";
const SAMX = "SAMX";

type Direction = {
  rowStep: number;
  colStep: number;
};

const directions: Direction[] = [
  { rowStep: -1, colStep: 0 }, //up
  { rowStep: 0, colStep: 1 }, //right
  { rowStep: -1, colStep: 1 }, //up-right
  { rowStep: -1, colStep: -1 }, //up-left
];

const check = (
  row: number,
  col: number,
  rowStep: number,
  colStep: number,
  input: string[],
): string => {
  let match = "";
  //loop 4 times to get to XMAS or SAMX
  for (let i = 0; i < 4; i++) {
    const currRow = row + rowStep * i;
    const currCol = col + colStep * i;

    if (
      //check for out of bounds cells
      currRow < 0 ||
      currCol < 0 ||
      currRow >= input.length ||
      currCol >= input[currRow].length
    ) {
      //early return if out of bounds
      return match;
    }
    match += input[currRow][currCol];
  }
  return match;
};

test("part 1 example", () => {
  let count = 0;
  const input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
    .split("\n")
    .filter(Boolean);

  for (let row = 0; row < input.length; row++) {
    for (let col = 0; col < input[row].length; col++) {
      for (const { rowStep, colStep } of directions) {
        check(row, col, rowStep, colStep, input) === XMAS && count++;
        check(row, col, rowStep, colStep, input) === SAMX && count++;
      }
    }
  }

  expect(count).toBe(18);
});

test("part 1", async () => {
  const input = (await readTextFile("./input.txt")).split("\n").filter(Boolean);
  let count = 0;
  for (let row = 0; row < input.length; row++) {
    for (let col = 0; col < input[row].length; col++) {
      for (const { rowStep, colStep } of directions) {
        check(row, col, rowStep, colStep, input) === XMAS && count++;
        check(row, col, rowStep, colStep, input) === SAMX && count++;
      }
    }
  }
  expect(count).toBe(2297);
});
