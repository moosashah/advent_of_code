import { expect, test } from "bun:test";
import { readTextFile } from "../utils";

function evaluateExpression(numbers: number[], operators: string[]): number {
  let results = numbers[0];
  for (let i = 0; i < operators.length; i++) {
    if (operators[i] === "+") {
      results += numbers[i + 1];
    } else if (operators[i] === "*") {
      results *= numbers[i + 1];
    }
  }
  return results;
}

const generateOperatorCombinations = (numOperators: number): string[][] => {
  const operators = ["+", "*"];
  const combinations: string[][] = [];

  const dfs = (curr: string[]) => {
    if (curr.length === numOperators) {
      combinations.push([...curr]);
      return;
    }
    for (const op of operators) {
      curr.push(op);
      dfs(curr);
      curr.pop();
    }
  };

  dfs([]);
  return combinations;
};

test("part 1 example", () => {
  const lines = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
    .split("\n")
    .filter(Boolean);

  const sum = lines.reduce((acc, curr) => {
    const [val, inputs] = curr.split(":");
    const targetValue = parseInt(val.trim());
    const numberInts = inputs
      .split(" ")
      .filter(Boolean)
      .map((num) => parseInt(num.trim()));

    const numOperators = numberInts.length - 1;
    const operatorCombinations = generateOperatorCombinations(numOperators);

    let foundSolution = false;

    for (const ops of operatorCombinations) {
      const results = evaluateExpression(numberInts, ops);
      if (results === targetValue) {
        foundSolution = true;
        break;
      }
    }

    if (foundSolution) return acc + targetValue;
    return acc;
  }, 0);
  expect(sum).toBe(3749);
});

test("part 1", async () => {
  const lines = (await readTextFile("./input.txt")).split("\n").filter(Boolean);
  const sum = lines.reduce((acc, curr) => {
    const [val, inputs] = curr.split(":");
    const targetValue = parseInt(val.trim());
    const numberInts = inputs
      .split(" ")
      .filter(Boolean)
      .map((num) => parseInt(num.trim()));

    const numOperators = numberInts.length - 1;
    const operatorCombinations = generateOperatorCombinations(numOperators);

    let foundSolution = false;

    for (const ops of operatorCombinations) {
      const results = evaluateExpression(numberInts, ops);
      if (results === targetValue) {
        foundSolution = true;
        break;
      }
    }

    if (foundSolution) return acc + targetValue;
    return acc;
  }, 0);
  expect(sum).toBe(1611660863222);
});
