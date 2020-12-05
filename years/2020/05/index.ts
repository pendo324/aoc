import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const split = input.split('\n');

const seatIds = split
  .map((l) => {
    const row = l.slice(0, 7);
    const noFs = row.replaceAll('F', '0');
    const noBs = noFs.replaceAll('B', '1');

    const rowNumber = parseInt(noBs, 2);

    const column = l.slice(7);
    const noRs = column.replaceAll('R', '1');
    const noLs = noRs.replaceAll('L', '0');
    const columnNumber = parseInt(noLs, 2);

    return rowNumber * 8 + columnNumber;
  })
  .filter((n) => !Number.isNaN(n));

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
