import { readFile } from 'fs/promises';
import { parse } from 'ts-node';
import { isConstructorTypeNode } from 'typescript';

const input = await readFile('./input', 'utf8');
const instructions = input.trim().split('\n');

enum INSTRUCTION {
  'ACC' = 'acc',
  'JMP' = 'jmp',
  'NOP' = 'nop'
}

let acc = 0;

interface ProgramInstruction {
  index: number;
  operation: string;
}

let executedInstructions: ProgramInstruction[] = [];

let pc = 0;
while (pc < instructions.length) {
  let ins = instructions[pc];
  const [op, valString] = ins.split(' ');
  const val = parseInt(valString);

  const l = executedInstructions.find(
    (ei) => ei.index === pc && ei.operation === ins
  );
  if (l) {
    console.log(`Part one: ${acc}`);
    break;
  }

  executedInstructions.push({
    index: pc,
    operation: ins
  });

  if (op === INSTRUCTION.ACC) {
    acc = acc + val;
    pc = pc + 1;
  } else if (op === INSTRUCTION.JMP) {
    pc = pc + val;
  } else if (op === INSTRUCTION.NOP) {
    pc = pc + 1;
  }
}

let pc2 = 0;
let acc2 = 0;

let executedInstructions2: ProgramInstruction[] = [];
const tryExecute = (inputInstructions: string[]) => {
  while (pc2 < inputInstructions.length) {
    let ins = inputInstructions[pc2];
    const [op, valString] = ins.split(' ');
    const val = parseInt(valString);

    const l = executedInstructions2.find(
      (ei) => ei.index === pc2 && ei.operation === ins
    );
    if (l) {
      return false;
    }

    executedInstructions2.push({
      index: pc2,
      operation: ins
    });

    if (op === INSTRUCTION.ACC) {
      acc2 = acc2 + val;
      pc2 = pc2 + 1;
    } else if (op === INSTRUCTION.JMP) {
      pc2 = pc2 + val;
    } else if (op === INSTRUCTION.NOP) {
      pc2 = pc2 + 1;
    }
  }

  return true;
};

const exhaustivelySwitchInstructions = (
  original: INSTRUCTION,
  replacement: INSTRUCTION
): number | null => {
  let lastInstructionIndex = 0;

  const finalItemIndex =
    instructions.length -
    1 -
    instructions
      .slice()
      .reverse()
      .findIndex((v) => {
        return v.split(' ')[0] === original;
      });

  while (lastInstructionIndex < finalItemIndex) {
    const jpmToSwitch = instructions.findIndex((v, i) => {
      return v.split(' ')[0] === original && i > lastInstructionIndex;
    });

    lastInstructionIndex = jpmToSwitch;

    const inputInstructions = instructions.slice();
    const newCommandSplit = inputInstructions[jpmToSwitch].split(' ');
    const newCommand = `${newCommandSplit[0].replace(
      INSTRUCTION.JMP,
      replacement
    )} ${newCommandSplit[1]}`;

    inputInstructions.splice(jpmToSwitch, 1, newCommand);

    const res = tryExecute(inputInstructions);

    if (pc2 === instructions.length) {
      return acc2;
    }

    if (res === false) {
      // reset state
      pc2 = 0;
      acc2 = 0;
      executedInstructions2 = [];
    }
  }

  return null;
};

const switchedJmpToNop = exhaustivelySwitchInstructions(
  INSTRUCTION.JMP,
  INSTRUCTION.NOP
);
const switchedNopToJmp = exhaustivelySwitchInstructions(
  INSTRUCTION.NOP,
  INSTRUCTION.JMP
);

console.log(`Part two: ${switchedJmpToNop ?? switchedNopToJmp}`);
