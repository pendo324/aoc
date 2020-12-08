import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const elements = input
  .trim()
  .split('\n')
  .map((i) => parseInt(i));

let ans1: number | null;
let ans2: number | null;
for (let [i, e] of elements.entries()) {
  if (i > elements.length / 2) {
    for (let [j, s] of elements.reverse().entries()) {
      const sum = e + s;
      if (sum === 2020) {
        ans1 = e;
        ans2 = elements[i];
      }
    }
  }
}

console.log(`Part two: ${ans1 * ans2}`);

const sums = {};
let part2ans1: number | null;
let part2ans2: number | null;
let part2ans3: number | null;

for (let [i, e] of elements.entries()) {
  if (i > elements.length / 2) {
    for (let [j, s] of elements.reverse().entries()) {
      const sum = e + s;
      if (sum < 2020) {
        sums[e] = {
          [s]: sum
        };
      }
    }
  }
}

for (let [first, s] of Object.entries(sums)) {
  for (let [second, twoSum] of Object.entries(s)) {
    for (let [i, e] of elements.entries()) {
      if (twoSum + e === 2020) {
        part2ans1 = Number(first);
        part2ans2 = Number(second);
        part2ans3 = e;
      }
    }
  }
}

console.log(`Part two: ${part2ans1 * part2ans2 * part2ans3}`);
