import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const validPasswords = input
  .trim()
  .split('\n')
  .filter((i) => {
    const split = i.split(' ');
    const [lowerString, upperString] = split[0].split('-');
    const letter = split[1][0];
    const password = split[2];

    const lower = parseInt(lowerString);
    const upper = parseInt(upperString);

    let matches = 0;
    for (let passwordLetter of Array.from(password)) {
      if (passwordLetter === letter) {
        matches = matches + 1;
      }
    }

    return matches <= upper && matches >= lower;
  });

console.log(`Part two: ${validPasswords.length}`);

const validPasswordsTwo = input
  .trim()
  .split('\n')
  .filter((i) => {
    const split = i.split(' ');
    const [lowerString, upperString] = split[0].split('-');
    const letter = split[1][0];
    const password = split[2];

    const lowerIndex = parseInt(lowerString) - 1;
    const upperIndex = parseInt(upperString) - 1;

    return (
      password.charAt(lowerIndex) !== password.charAt(upperIndex) &&
      (password.charAt(lowerIndex) === letter ||
        password.charAt(upperIndex) === letter)
    );
  });

console.log(`Part two: ${validPasswordsTwo.length}`);
