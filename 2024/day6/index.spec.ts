import { test, expect } from "bun:test";
import { readTextFile } from "../utils";

type Direction = {
  rowStep: number;
  colStep: number;
};

function isBarrier(row: number, col: number, input: string[]): Boolean {
  return input[row][col] === "#";
}

const direction: Direction[] = [
  { rowStep: -1, colStep: 0 }, //up
  { rowStep: 0, colStep: 1 }, //right
  { rowStep: 1, colStep: 0 }, //down
  { rowStep: 0, colStep: -1 }, //left
];

function changeDirection(currentDirection: Direction): Direction {
  const idx = direction.indexOf(currentDirection);
  if (idx === 3) return direction[0];
  return direction[idx + 1];
}

function step(
  currentDirection: Direction,
  currentLocation: [number, number],
  input: string[],
): { location: [number, number]; direction: Direction } {
  let newDirection = currentDirection;
  const [currRow, currCol] = currentLocation;
  const { rowStep, colStep } = newDirection;
  let nextRow = currRow + rowStep;
  let nextCol = currCol + colStep;
  if (isBarrier(nextRow, nextCol, input)) {
    newDirection = changeDirection(currentDirection);
    nextRow = currRow + newDirection.rowStep;
    nextCol = currCol + newDirection.colStep;
  }
  return { location: [nextRow, nextCol], direction: newDirection };
}

test("part 1 example", () => {
  const input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
    .split("\n")
    .filter(Boolean);

  const location: [number, number] = [0, 0];
  let currDirection = direction[0];
  input.forEach((line, row) => {
    const col = line.indexOf("^");
    if (col > -1) {
      location[0] = row;
      location[1] = col;
    }
  });

  const visited = new Set<string>();

  while (
    location[0] < input.length &&
    location[1] < input[location[0]].length
  ) {
    try {
      visited.add(location.join(","));
      const { location: newLocation, direction: newDirection } = step(
        currDirection,
        location,
        input,
      );
      location[0] = newLocation[0];
      location[1] = newLocation[1];
      currDirection = newDirection;
    } catch {
      break;
    }
  }

  const distinct = visited.size;
  expect(distinct).toBe(41);
});

test("part 1", async () => {
  const input = (await readTextFile("./input.txt")).split("\n").filter(Boolean);

  const location: [number, number] = [0, 0];
  let currDirection = direction[0];
  input.forEach((line, row) => {
    const col = line.indexOf("^");
    if (col > -1) {
      location[0] = row;
      location[1] = col;
    }
  });

  const visited = new Set<string>();

  while (
    location[0] < input.length &&
    location[1] < input[location[0]].length
  ) {
    try {
      visited.add(location.join(","));
      const { location: newLocation, direction: newDirection } = step(
        currDirection,
        location,
        input,
      );
      location[0] = newLocation[0];
      location[1] = newLocation[1];
      currDirection = newDirection;
    } catch {
      break;
    }
  }
  const distinct = visited.size;
  expect(distinct).toBe(4819);
});
