
import json
import sys
from deep_eq import *


def _in_(l, lines):
    for line in lines:
        if l[0] != line[0] or l[1] != line[1]:
            continue
        if l[2] != line[2]:
            print l[1], '\n\t      Has:', l[2], '\n\tShould be:', line[2]
            return False
        if l[2] == 200 and len(l) == 4 and len(line) == 4:
            if not deep_eq(json.loads(l[3]), json.loads(line[3])):
                print l[1], '\n\t      Has:', l[3], '\n\tShould be:', line[3]
                return False
        return True


if len(sys.argv) != 3:
    print('format: ./test.py <file.answ> <file.answ>')
    exit(1)

f0 = open(sys.argv[1], 'r')
f0_lines = []
for line in f0:
    f0_lines.append(line.rstrip().split('\t', 4))

f1 = open(sys.argv[2], 'r')
f1_lines = []
for line in f1:
    if _in_(line.rstrip().split('\t', 4), f0_lines):
        continue

