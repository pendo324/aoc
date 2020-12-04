import { mkdir, readFile, writeFile } from 'fs/promises';
import { join } from 'path';
import { fileURLToPath } from 'url';
import { promisify } from 'util';

import got from 'got';
import * as dotenv from 'dotenv';
import { CookieJar } from 'tough-cookie';

dotenv.config();

let year = new Date().getFullYear() + '';
let day: string;
if (process.argv.length === 3) {
  day = process.argv[2];
} else if (process.argv.length === 4) {
  year = process.argv[2];
  day = process.argv[3];
} else if (process.argv.length === 2 || process.argv.length > 4) {
  throw 'Wrong number of arguments provided.';
}

const paddedDay = day.padStart(2, '0');
const folderPath = join(
  fileURLToPath(import.meta.url),
  `../../years/${year}/${paddedDay}`
);

try {
  await mkdir(folderPath);
} catch {
  console.log('Folder already exists.');
}

console.log('Writing or overwriting package.json file');

const packageJson = `{
  "name": "@pendo324/aoc-${year}-${paddedDay}",
  "version": "1.0.0",
  "description": "Advent of code solution for ${year}, day ${day}",
  "main": "index.js",
  "type": "module",
  "scripts": {
    "test": "node --no-warnings --loader ts-node/esm ./index.ts"
  },
  "devDependencies": {
    "@types/node": "^14.14.10"
  },
  "dependencies": {
    "ts-node": "^9.1.0"
  }
}
`;

const startingFile = `import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

`;

await writeFile(join(folderPath, 'package.json'), packageJson);
// don't overwrite index.js, there might be a solution in there
try {
  await writeFile(join(folderPath, 'index.ts'), startingFile, { flag: 'wx' });
} catch (e) {
  if (e.code !== 'EEXIST') {
    console.log('Skipping write to index.ts since it already exists');
    throw e;
  }
}

// download input file from aoc
const url = `https://adventofcode.com/${year}/day/${day}/input`;

const cookieJar = new CookieJar();
const setCookie = promisify(cookieJar.setCookie.bind(cookieJar));
await setCookie(`session=${process.env.AOC_SESSION_COOKIE}`, url);

const response = await got(url, { cookieJar });
const inputFile = response.body;

await writeFile(join(folderPath, 'input'), inputFile);

// add run script to root package file
const rootPackagePath = join(
  fileURLToPath(import.meta.url),
  '../../package.json'
);
const rootPackageFile = JSON.parse(await readFile(rootPackagePath, 'utf8'));

if (!!!rootPackageFile.scripts[day]) {
  rootPackageFile.scripts[
    day
  ] = `cd ./years/${year}/${paddedDay} && npm run test`;
}

await writeFile(rootPackagePath, `${JSON.stringify(rootPackageFile, null, 2)}\n`);