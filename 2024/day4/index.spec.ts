import { test, expect } from "bun:test";
import { readTextFile } from "../utils";

const XMAS = "XMAS";
const SAMX = "SAMX";

const right = (row: number, col: number, rows: string[]): string => {
  if (col + 3 < rows[row].length) {
    return (
      rows[row][col] +
      rows[row][col + 1] +
      rows[row][col + 2] +
      rows[row][col + 3]
    );
  }
  return "";
};

const up = (row: number, col: number, rows: string[]): string => {
  if (row - 3 >= 0) {
    return (
      rows[row][col] +
      rows[row - 1][col] +
      rows[row - 2][col] +
      rows[row - 3][col]
    );
  }
  return "";
};

const upRight = (row: number, col: number, rows: string[]): string => {
  if (row - 3 >= 0 && col + 3 < rows[row].length) {
    return (
      rows[row][col] +
      rows[row - 1][col + 1] +
      rows[row - 2][col + 2] +
      rows[row - 3][col + 3]
    );
  }
  return "";
};

const upLeft = (row: number, col: number, rows: string[]): string => {
  if (row - 3 >= 0 && col - 3 >= 0) {
    return (
      rows[row][col] +
      rows[row - 1][col - 1] +
      rows[row - 2][col - 2] +
      rows[row - 3][col - 3]
    );
  }
  return "";
};

test("part 1 example", () => {
  let count = 0;
  const rows = `MMMSXXMASM
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
  for (let row = 0; row < rows.length; row++) {
    for (let col = 0; col < rows[row].length; col++) {
      up(row, col, rows) === XMAS && count++;
      right(row, col, rows) === XMAS && count++;
      upRight(row, col, rows) === XMAS && count++;
      upLeft(row, col, rows) === XMAS && count++;

      up(row, col, rows) === SAMX && count++;
      right(row, col, rows) === SAMX && count++;
      upRight(row, col, rows) === SAMX && count++;
      upLeft(row, col, rows) === SAMX && count++;
    }
  }
  expect(count).toBe(18);
});

test("part 1", async () => {
  const content = (await readTextFile("./input.txt"))
    .split("\n")
    .filter(Boolean);
  let count = 0;
  for (let row = 0; row < content.length; row++) {
    for (let col = 0; col < content[row].length; col++) {
      up(row, col, content) === XMAS && count++;
      right(row, col, content) === XMAS && count++;
      upRight(row, col, content) === XMAS && count++;
      upLeft(row, col, content) === XMAS && count++;

      up(row, col, content) === SAMX && count++;
      right(row, col, content) === SAMX && count++;
      upRight(row, col, content) === SAMX && count++;
      upLeft(row, col, content) === SAMX && count++;
    }
  }
  expect(count).toBe(2297);
});
