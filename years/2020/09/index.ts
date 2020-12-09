import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const numbers = input
  .trim()
  .split('\n')
  .map((n) => parseInt(n));

let n = 25;

let wrongNumber: number;
while (n < numbers.length) {
  let preamble = numbers.slice(n - 25, n);
  wrongNumber = numbers[n];

  let matched = false;
  for (let i = 0; i < preamble.length && !matched; i++) {
    let p1 = preamble[i];
    for (let j = 0; j < preamble.length && !matched; j++) {
      let p2 = preamble[preamble.length - j];

      if (wrongNumber === p1 + p2) {
        matched = true;
      }
    }
  }

  if (matched) {
    n++;
  } else {
    console.log(`Part one: ${wrongNumber}`);
    break;
  }
}

let candidates: number[] = [];
let matched = false;
for (let i = 0; i < numbers.length && !matched; i++) {
  const number1 = numbers[i];

  candidates = [];
  candidates.push(number1);
  for (let j = i + 1; j < numbers.length && !matched; j++) {
    const number2 = numbers[j];

    if (number1 + number2 < wrongNumber) {
      candidates.push(number2);

      const total = candidates.reduce((sum, c) => {
        sum = sum + c;
        return sum;
      }, 0);

      if (total > wrongNumber) {
        candidates = candidates.slice(0, 1);
      }

      if (total === wrongNumber) {
        matched = true;
      }
    }
  }
}

if (matched) {
  const sorted = candidates.sort((a, b) => b - a);
  const max = sorted[0];
  const min = sorted[sorted.length - 1];

  console.log(`Part two: ${max + min}`);
}
