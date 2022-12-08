#!/usr/bin/env python3

import sys

count = 0
for line in map(str.rstrip, sys.stdin):
    elves = line.split(',')
    elves[0] = list(map(int, elves[0].split('-')))
    elves[1] = list(map(int, elves[1].split('-')))

    print(str(elves))
# Overlap completely (part 1)
#    if ( elves[0][0] <= elves[1][0] and elves[0][1] >= elves[1][1] ) or ( elves[0][0] >= elves[1][0] and elves[0][1] <= elves[1][1] ):

# Overlap at all (part 2)
    if not ( elves[0][1] < elves[1][0] or elves[1][1] < elves[0][0] ):
       count = count + 1
       print("match")
    print(count)
