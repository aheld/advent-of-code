import unittest
import sys
import pathlib


def make_lists(input):
    left = []
    right = []
    for line in input.split("\n"):
        inputs = line.split("   ")
        if len(inputs) != 2:
            continue
        left.append(int(inputs[0]))
        right.append(int(inputs[1]))

    left.sort()
    right.sort()
    return left, right


def solve(input):
    left, right = make_lists(input)
    dual = zip(left, right)
    sum = 0
    for l, r in list(dual):
        sum += abs(l - r)
    return sum


def solve2(input):
    left, right = make_lists(input)
    sum = 0
    for i in left:
        matches = [x for x in right if x == i]
        print(f"""{i}: {matches}""")
        sum += i * len(matches)
    return sum


class Day1Test(unittest.TestCase):
    def setUp(self):
        self.input = pathlib.Path("test_input.txt").read_text()

    def test_part1(self):
        self.assertEqual(solve(self.input), 11)

    def test_part2(self):
        self.assertEqual(solve2(self.input), 31)


if __name__ == "__main__":
    if len(sys.argv) == 1:
        unittest.main()
    else:
        input = pathlib.Path("input").read_text()
        print(solve(input))
        print(solve2(input))
