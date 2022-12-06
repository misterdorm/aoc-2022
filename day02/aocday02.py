#!/usr/bin/env python3

import sys

# A = X = ROCK
# B = Y = PAPER
# C = Z = SCISSORS

score_map = { 'X': 1, 'Y': 2, 'Z': 3 }
match_map = { 'X': 'A', 'Y': 'B', 'Z': 'C' }
beats_map = { 'X': 'C', 'Y': 'A', 'Z': 'B' }
score = 0

for line in map(str.rstrip, sys.stdin):
    shapes = line.split()
    print(shapes)
    score = score + score_map[shapes[1]]
    print(score)

    if shapes[0] == match_map[shapes[1]]:
        print("draw")
        score = score + 3
    elif shapes[0] == beats_map[shapes[1]]:
        print("won")
        score = score + 6

    print(score)
    print()
