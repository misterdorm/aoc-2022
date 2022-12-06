#!/usr/bin/env python3

import sys

elf_calories = [0]
elf = 0

for line in map(str.rstrip, sys.stdin):
    print(line)
    try:
        elf_calories[elf] = elf_calories[elf] + int(line)
    except ValueError:
        print("Elf " + str(elf) + ": " + str(elf_calories[elf]))
        elf = elf + 1
        elf_calories.append(0)

elf_calories.sort(reverse=True)

print('---------------')
print(elf_calories[0])
print(elf_calories[0:3])
print(sum(elf_calories[0:3]))
