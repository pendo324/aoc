import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const groups = input.split('\n\n');
const groupAns = groups.map((group) => {
  const answers = new Set();
  for (let [_, ans] of group.split('\n').entries()) {
    for (let letter of ans) {
      answers.add(letter);
    }
  }
  return answers;
});

const lengths = groupAns.map((ans) => {
  return Array.from(ans).length;
});

const sum = lengths.reduce((s, l) => {
  const sum = s + l;
  return sum;
}, 0);

console.log(`Part one: ${sum}`);

const groupEveryoneAns = groups
  .map((group, i, arr) => {
    const answers = {};
    const split = group.split('\n');
    let members = split.length;
    // remove the trailing newline lol.
    if (i === arr.length - 1) {
      members = members - 1;
    }

    for (let [_, ans] of split.entries()) {
      for (let letter of ans) {
        if (answers[letter]) {
          answers[letter] = answers[letter] + 1;
        } else {
          answers[letter] = 1;
        }
      }
    }
    for (let [ans] of Object.entries<string>(answers)) {
      if (answers[ans] !== members) {
        delete answers[ans];
      }
    }
    return answers;
  })
  .filter((g) => Object.keys(g).length);

const lengthsEveryone = groupEveryoneAns.map((ans) => {
  return Object.keys(ans).length;
});

const sumEveryone = lengthsEveryone.reduce((s, l) => {
  const sum = s + l;
  return sum;
}, 0);

console.log(`Part two: ${sumEveryone}`);
