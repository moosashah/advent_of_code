import { diffNumbers, sumNumberArray, readTextFile } from "../utils";

const path = process.argv.slice(2)[0];

const content = await readTextFile(path);

const col1: string[] = [];
const col2: string[] = [];
const differences: number[] = [];

content.split("\n").forEach((line) => {
  line.split("   ").forEach((col, idx) => {
    if (col) {
      col.trim();
      if (idx % 2) {
        col2.push(col);
      } else {
        col1.push(col);
      }
    }
  });
});

// console.log("col1", col1);
// console.log("col2", col2);

const col1Sorted = col1.toSorted();
const col2Sorted = col2.toSorted();

// console.log("col1Sorted", col1Sorted);
// console.log("col2Sorted", col2Sorted);

col1Sorted.forEach((num, idx) => {
  const num1 = +num;
  const num2 = +col2Sorted[idx];
  // console.log("num1", num1);
  // console.log("num2", num2);
  const diff = diffNumbers(num1, num2);
  // console.log(diff);
  differences.push(diff);
});

// console.log(differences);
console.log("part 1: ", sumNumberArray(differences));

//part 2
const similarities: number[] = [];

col1.forEach((str) => {
  let matches = 0;
  const num = +str;
  col2.forEach((comparison) => {
    if (num === +comparison) {
      matches++;
    }
  });
  const similarityScore = num * matches;
  if (similarityScore) {
    // console.log(similarityScore);
    similarities.push(similarityScore);
  }
});

// console.log(similarities);
console.log("part 2: ", sumNumberArray(similarities));
