import unittest
import sys
import pathlib

DEBUG = False

def increasing(levels):
    difs = set()
    prev=levels[0]
    for x in levels[1:]:
        if DEBUG:
            print(f'''{(prev+1)} <= {x} <= {(prev+3)}''')
        difs.add((x-prev))
        if (prev+1) <= x <= (prev+3):

            prev = x
            continue
        else:
            print(f'''Bad: {difs}''')
            return False
    print(f'''Good: {difs}''')
    return True

def decreasing(levels):
    difs = set()
    prev=levels[0]
    for x in levels[1:]:
        if DEBUG:
            print(f'''{(prev-1)} >= {x} >= {(prev-3)}''')
        difs.add((x-prev))
        if (prev-1) >= x >= (prev-3):
            prev = x
            continue
        else:
            print(f'''Bad: {difs}''')
            return False
    print(f'''Good: {difs}''')
    return True

def is_safe(levels, i=None):
    if i is None:
        l = levels
    else:
        l = levels[:i] + levels[i+1:]
    # print(l, i)
    if l[1]>l[0]:
        return increasing(l)
            
    if l[1]<l[0]:
        return decreasing(l)
    
    return False

def solve2(input):
    return solve(input, True)

def solve(input, damper=False):
    sum=0
    for report in input.split('\n'):
        direction=None
        if len(report) < 4:
            continue
        levels=[int(x.strip()) for x in report.split(' ')]
        safe = is_safe(levels)
        i = 0
        while damper and not safe and i < len(levels):
            safe = is_safe(levels, i)
            i+=1
        if safe:
            sum += 1
    return sum
        



class Day1Test(unittest.TestCase):
    def setUp(self):
        self.input = pathlib.Path("test_input.txt").read_text()
        self.real_input = pathlib.Path("input").read_text()

    def test_part1(self):
        self.assertEqual(solve(self.input), 2)
        self.assertEqual(solve(self.real_input), 299)

    def test_part2(self):
        self.assertEqual(solve2(self.input), 4)
        self.assertEqual(solve2(self.real_input), 364)


if __name__ == "__main__":
    if len(sys.argv) == 1:
        unittest.main()
    else:
        input = pathlib.Path("input").read_text()
        # print(solve(input))
        print(solve2(input))
