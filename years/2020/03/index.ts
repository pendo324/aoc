import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const map = input.split('\n');

const checkMap = (right: number, down: number) => {
  let xPos = 0;
  let yPos = 0;
  let trees = 0;

  // avoid overflow and ignore newline at EOF
  while (yPos < map.length - 2) {
    xPos = (xPos + right) % map[0].length;
    yPos = yPos + down;

    if (map[yPos][xPos] === '#') {
      trees = trees + 1;
    }
  }
  return trees;
};

const part1 = checkMap(3, 1);

const part2_1 = checkMap(1, 1);
const part2_2 = checkMap(5, 1);
const part2_3 = checkMap(7, 1);
const part2_4 = checkMap(1, 2);

console.log(`Part one: ${part1}`);
console.log(`Part two: ${part1 * part2_1 * part2_2 * part2_3 * part2_4}`);
