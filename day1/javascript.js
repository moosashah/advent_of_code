const { readFileSync } = require('fs')

const singleElf = readFileSync('./real_input', 'utf-8').trimEnd().split('\n\n')

const singleElfSum = singleElf.map((elf) =>
  elf
    .split('\n')
    .map((line) => Number(line))
    .reduce((sum, i) => sum + i, 0)
)

const [topElf, secondElf, thirdElf] = singleElfSum.sort((a, z) => z - a)
console.log('Top Elf: %s', topElf)

console.log('Top 3 Elves %s', topElf + secondElf + thirdElf)
