const filename = "input";
let [time, distance] = (await Bun.file(filename).text())
  .split("\n")
  .map((f) => f.split(":")[1]);

const clean = (i: string): number[] => {
  return [
    parseInt(
      i
        .split(" ")
        .filter((x) => x)
        .join(""),
    ),
  ];
};

const tm = clean(time);
const dst = clean(distance);

const winnings = [];
for (let i = 0; i < tm.length; ++i) {
  const t = tm[i];
  const distanceToBeat = dst[i];

  let count = 0;
  for (let j = 0; j < t; ++j) {
    const myDist = j * (tm[i] - j);
    if (myDist > distanceToBeat) {
      count++;
    }
  }
  winnings.push(count);
}

console.log(
  winnings.reduce((acc, curr) => {
    return acc * curr;
  }, 1),
);
