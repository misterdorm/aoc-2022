#!/usr/bin/env python3

import sys

sum = 0

for line in map(str.rstrip, sys.stdin):
    c1 = line[:len(line)//2]
    c2 = line[len(line)//2:]
    print(c1 + ", " + c2)

    type = ''.join(
        set(c1).intersection(c2)
        )

    if type.islower():
        p = ord(type) - 96
    else:
        p = ord(type) - 38

    print(type + ": " + str(ord(type)) + " --> " + str(p))
    sum = sum + p

print(sum)
