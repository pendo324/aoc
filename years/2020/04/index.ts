import { readFile } from 'fs/promises';

const input = await readFile('./input', 'utf8');

const passports = input.split('\n\n');

const requiredFields = [
  'byr',
  'iyr',
  'eyr',
  'hgt',
  'hcl',
  'ecl',
  'pid',
] as const;

const validFields = [...requiredFields, 'cid'] as const;

export type PassportRequiredFields = {
  [key in typeof requiredFields[number]]: string;
};

interface Passport extends PassportRequiredFields {
  cid?: string;
}

interface MaybePassport {
  [key: string]: string;
}

const parsedPassports = passports.map((p) => {
  const splitSpaces = p.split(' ');
  const splitLines = splitSpaces.map((f) => f.split('\n'));
  const flat = splitLines.flat();
  const split = flat.map((f) => {
    return f.split(':');
  });

  return split.reduce((accum, i) => {
    accum[i[0]] = i[1];
    return accum;
  }, {});
}) as MaybePassport[];

const validPassports = (parsedPassports.filter((p) => {
  const keys = Object.keys(p);
  return validFields.reduce((accum, f) => {
    if (f === 'cid') {
      return accum && true;
    }
    return accum && keys.includes(f);
  }, true);
}) as unknown) as Passport[];

console.log(`Part one: ${validPassports.length}`);

const eyeColors = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'];

const checkedPassports = validPassports.filter((v) => {
  return ((Object.keys(v) as unknown) as typeof validFields).reduce(
    (accum, k) => {
      const value = v[k];
      if (k === 'byr') {
        const valueAsNum = Number(value);
        return (
          accum &&
          value.length === 4 &&
          valueAsNum >= 1920 &&
          valueAsNum <= 2002
        );
      }
      if (k === 'iyr') {
        const valueAsNum = Number(value);
        return (
          accum &&
          value.length === 4 &&
          valueAsNum >= 2010 &&
          valueAsNum <= 2020
        );
      }
      if (k === 'eyr') {
        const valueAsNum = Number(value);
        return (
          accum &&
          value.length === 4 &&
          valueAsNum >= 2020 &&
          valueAsNum <= 2030
        );
      }
      if (k === 'hgt') {
        if (value.endsWith('in')) {
          const height = Number(value.split('in')[0]);
          return accum && height >= 59 && height <= 76;
        } else if (value.endsWith('cm')) {
          const height = Number(value.split('cm')[0]);
          return accum && height >= 150 && height <= 193;
        }
      }
      if (k === 'hcl') {
        const color = value.split('#')[1];
        return accum && /^[0-9a-f]{6}$/.test(color);
      }
      if (k === 'ecl') {
        return accum && eyeColors.includes(value);
      }
      if (k === 'pid') {
        return accum && /^[0-9]{9}$/.test(value);
      }
      if (k === 'cid') {
        return accum && true;
      }
      return accum && false;
    },
    true
  );
});

console.log(`Part two: ${checkedPassports.length}`);
