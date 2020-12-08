import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

interface ColorSize {
  name: string;
  size: number;
}

interface Item {
  color: string;
  contains?: ColorSize[];
}

let bags: Item[] = [];
input
  .trim()
  .split('\n')
  .forEach((r) => {
    const getColor = /(\S+\s\S+)(.*)/g;

    const [_, color] = [...r.matchAll(getColor)][0];

    const getContains = /(\d) (\S+\s\S+)/g;

    const containedColors = r.matchAll(getContains);

    const bag = {
      color,
      contains: []
    };

    for (let match of containedColors) {
      bag.contains.push({
        size: parseInt(match[1]),
        name: match[2]
      });
    }

    bags.push(bag);
  });

let total = 0;

const checkBag = (outer: Item, inner: ColorSize) => {
  if (outer?.contains?.find((o) => o.name === inner.name)) {
    return true;
  }

  let hasBag = false;
  if (outer?.contains) {
    for (let b of outer.contains) {
      const bag = bags?.find((o) => o.color === b.name);
      if (!bag) {
        hasBag = false;
      } else {
        hasBag = hasBag || checkBag(bag, inner);
      }
    }
  }

  return hasBag;
};

for (let bag of bags) {
  const shinyGoldBag: ColorSize = {
    name: 'shiny gold',
    size: 0
  };

  if (checkBag(bag, shinyGoldBag)) {
    total = total + 1;
  }
}

console.log(`Part one: ${total}`);

const shinyGoldBag = bags.find((b) => b.color === 'shiny gold');

const countBags = (rootBag: Item) =>
  rootBag.contains.reduce((accum, insideBag) => {
    accum =
      accum +
      insideBag.size +
      countBags(bags.find((b) => b.color === insideBag.name)) * insideBag.size;

    return accum;
  }, 0);

console.log(`Part two: ${countBags(shinyGoldBag)}`);
