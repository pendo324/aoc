import { mkdir, readFile, stat, writeFile } from 'fs/promises';
import { join } from 'path';
import { fileURLToPath } from 'url';
import { promisify } from 'util';

import got from 'got';
import * as dotenv from 'dotenv';
import { CookieJar } from 'tough-cookie';

import { capturePage } from './page-to-markdown';

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
} catch (e) {
  if (e.code === 'EEXIST') {
    console.log('Folder already exists.');
  } else {
    throw e;
  }
}

const packageFilePath = join(folderPath, 'package.json');
console.log(`Writing or overwriting package.json file at ${packageFilePath}`);

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

await writeFile(packageFilePath, packageJson);

const startingFile = `import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

`;

// don't overwrite index.js, there might be a solution in there
const indexPath = join(folderPath, 'index.ts');
let inputFileExists = false;
let wroteInputFile = true;
try {
  await writeFile(indexPath, startingFile, { flag: 'wx' });
} catch (e) {
  wroteInputFile = false;
  if (e.code === 'EEXIST') {
    inputFileExists = true;
  } else {
    throw e;
  }
}

if (!wroteInputFile) {
  if (inputFileExists) {
    console.log(
      `Skipped write to index.ts since it already exists at ${indexPath}`
    );
  } else {
    console.log(`Skipped write to index.ts because of unhandled error`);
  }
} else {
  console.log(`Wrote index.ts to ${indexPath}`);
}

// download input file from aoc
const baseUrl = `https://adventofcode.com/${year}/day/${day}`;
const inputUrl = `${baseUrl}/input`;

const cookieJar = new CookieJar();
const setCookie = promisify(cookieJar.setCookie.bind(cookieJar));
await setCookie(`session=${process.env.AOC_SESSION_COOKIE}`, baseUrl);

// only download the input file if it doesn't exist
const inputFilePath = join(folderPath, 'input');
let downloadedInputFile = false;
try {
  await stat(inputFilePath);
} catch (e) {
  if (e.code === 'ENOENT') {
    try {
      const response = await got(inputUrl, { cookieJar });
      const inputFile = response.body;

      await writeFile(inputFilePath, inputFile);
      downloadedInputFile = true;
    } catch (innerError) {
      if (innerError?.response?.statusCode === 404) {
        console.log(
          `The input file doesn't exist on the server. This usually means it was run before the prompt unlocked, or there was a problem with the input.`
        );
        const body = innerError?.response?.body;
        if (body) {
          console.log(`Error message from page: ${body}`);
        }
      } else if (innerError?.response?.statusCode === 403) {
        console.log(
          `Download forbidden. This usually means the token wasn't set properly. Re-check the .env file.`
        );
      } else {
        throw e;
      }
    }
  } else {
    throw e;
  }
}

if (downloadedInputFile) {
  console.log(`Downloaded input file to ${inputFilePath}`);
} else {
  console.log(
    `Skipped input file download since file exists at ${inputFilePath}`
  );
}

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
