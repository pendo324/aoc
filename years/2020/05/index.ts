import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const split = input.split('\n');
// remove trailing newline
const trimmed = split.slice(0, split.length - 1);

const seatIds = trimmed.map((l) => {
  const ones = l.replaceAll(/[FL]/g, '0');
  const zeros = ones.replaceAll(/[BR]/g, '1');

  const row = parseInt(zeros.slice(0, 7), 2);
  const column = parseInt(zeros.slice(7), 2);

  return row * 8 + column;
});

const sortedIds = seatIds.sort((a, b) => b - a);

const highestId = sortedIds[0];

console.log(`Part one: ${highestId}`);

const min = sortedIds[sortedIds.length - 1];
const max = sortedIds[0];

// create a new array with all of the values between min and max, and check if each one is in the original array
// the one that's not, is my seat
const mySeat = Array.from(Array(max - min), (_, i) => i + min).filter(
  (i) => !sortedIds.includes(i)
);

console.log(`Part two: ${mySeat}`);
