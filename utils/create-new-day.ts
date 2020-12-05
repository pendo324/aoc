import { mkdir, readFile, writeFile } from 'fs/promises';
import { join } from 'path';
import { fileURLToPath } from 'url';
import { promisify } from 'util';

import got from 'got';
import * as dotenv from 'dotenv';
import { CookieJar } from 'tough-cookie';

// TODO: figure out a better way to get TypeScript to like local imports with extensions
// @ts-ignore
import { capturePage } from './page-to-markdown.ts';

dotenv.config();

let year = new Date().getFullYear().toString(10);
let day: string;
if (process.argv.length === 3) {
  day = process.argv[2];
} else if (process.argv.length === 4) {
  year = process.argv[2];
  day = parseInt(process.argv[3]).toString(10);
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
const baseUrl = `https://adventofcode.com/${year}/day/${day}`;
const inputUrl = `${baseUrl}/input`;

const cookieJar = new CookieJar();
const setCookie = promisify(cookieJar.setCookie.bind(cookieJar));
await setCookie(`session=${process.env.AOC_SESSION_COOKIE}`, baseUrl);

const response = await got(inputUrl, { cookieJar });
const inputFile = response.body;

await writeFile(join(folderPath, 'input'), inputFile);

const pageMarkdown = capturePage(
  (await got(baseUrl, { cookieJar })).body,
  baseUrl
);

await writeFile(join(folderPath, `README.md`), pageMarkdown);

// add run script to root package file
const rootPackagePath = join(
  fileURLToPath(import.meta.url),
  '../../package.json'
);
const rootPackageFile = JSON.parse(await readFile(rootPackagePath, 'utf8'));

const dayKey = `${year}-${day}`;
if (!!!rootPackageFile.scripts[dayKey]) {
  rootPackageFile.scripts[
    dayKey
  ] = `cd ./years/${year}/${paddedDay} && exitzero npm --loglevel silent run test`;
}

await writeFile(
  rootPackagePath,
  `${JSON.stringify(rootPackageFile, null, 2)}\n`
);
