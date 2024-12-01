import { readFile } from "fs/promises";

export async function readTextFile(filePath: string): Promise<string> {
  try {
    return await readFile(filePath, "utf8");
  } catch (error) {
    console.error(`Error reading file from disk: ${error}`);
    throw error;
  }
}

export const diffNumbers = (number1: number, number2: number): number =>
  Math.abs(number1 - number2);

export const sumNumberArray = (numbers: number[]): number =>
  numbers.reduce((acc, curr) => acc + curr, 0);
