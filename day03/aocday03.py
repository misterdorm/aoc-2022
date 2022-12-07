#!/usr/bin/env python3

import sys

def chunker(seq, size):
    return (seq[pos:pos + size] for pos in range(0, len(seq), size))

sum = 0

for lines in chunker(list(map(str.rstrip, sys.stdin)), 3):
    print(lines)
    item = ''.join(set(lines[0]).intersection(lines[1]).intersection(lines[2]))
    print(item)

    if item.islower():
        p = ord(item) - 96
    else:
        p = ord(item) - 38

    print(item + ": " + str(ord(item)) + " --> " + str(p))
    sum = sum + p

print(sum)
