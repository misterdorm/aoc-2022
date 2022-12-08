#!/usr/bin/env python3

import re
import sys

stacks = []

for line in map(str.rstrip, sys.stdin):
    if not line:
        break

    print(line)
    stack = 0
    for crate in [line[i:i+4] for i in range(0, len(line), 4)]:
        if len(stacks) < stack + 1:
            stacks.append([])

        if crate[0:1] == '[':
            stacks[stack].append(crate[1:2])

        stack = stack + 1

# Fix the ordering, so bottom is at index 0, top is at high index
for i in range(0, len(stacks)):
    stacks[i] = stacks[i][::-1]

print(stacks)

for line in map(str.rstrip, sys.stdin):
    m = re.search("move (\d+) from (\d+) to (\d+)", line)
    count = int(m.group(1))
    src = int(m.group(2)) - 1
    dst = int(m.group(3)) - 1

    stacks[dst].extend(stacks[src][-count:])
    del(stacks[src][-count:])
    print(line)
    print(stacks)
    print("------")

tops = ""
for i in range(0, len(stacks)):
    tops = tops + stacks[i][-1]

print(tops)
