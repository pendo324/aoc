import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const joltages = input
  .trim()
  .split('\n')
  .map((n) => parseInt(n))
  .sort((a, b) => b - a);

const joltagesOne = joltages.slice();
joltagesOne.push(0);
joltagesOne.unshift(joltagesOne[0] + 3);

let oneDif = 0;
let threeDif = 0;
for (let i = 0; i < joltagesOne.length - 1; i++) {
  const dif = joltagesOne[i] - joltagesOne[i + 1];
  if (dif === 1) {
    oneDif = oneDif + 1;
  } else if (dif === 3) {
    threeDif = threeDif + 1;
  }
}

console.log(`Part one: ${oneDif * threeDif}`);

const joltagesTwo = joltages.slice();
joltagesTwo.sort((a, b) => a - b).unshift(0);
joltagesTwo.push(joltagesTwo[joltagesTwo.length - 1] + 3);

let prefixes = [];
prefixes.length = joltagesTwo.length;
prefixes.fill(0, 0, joltagesTwo.length);

prefixes[0] = 1;

for (let i = 1; i < joltagesTwo.length; i++) {
  const number = joltagesTwo[i];
  for (let j = 0; j < i; j++) {
    const secondNumber = joltagesTwo[j];

    const dif = number - secondNumber;

    if (dif <= 3) {
      prefixes[i] = prefixes[i] + prefixes[j];
    }
  }
}

console.log(`Part two: ${prefixes[prefixes.length - 1]}`);
