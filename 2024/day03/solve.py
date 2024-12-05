import unittest
import sys
import pathlib
import re
from pprint import pprint as pp


def solve(input):
    pattern = re.compile(r"mul\((\d{1,3}),(\d{1,3})\)")

    matches = pattern.findall(input)
    sum = 0
    for match in matches:
        sum += int(match[0]) * int(match[1])

    return sum


def solve2(input):
    does = ""
    buffer = input
    while True:
        if not buffer or len(buffer) < 2:
            break

        found_pos = buffer.find("don't()")
        if found_pos != -1:
            does += buffer[:found_pos]
            buffer = buffer[found_pos + len("don't()") :]
        else:
            does += buffer
            break

        found_pos = buffer.find("do()")
        if found_pos != -1:
            buffer = buffer[found_pos + len("do()") :]
        else:
            break

    return solve(does)


class Day1Test(unittest.TestCase):
    def setUp(self):
        self.input = (
            "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
        )

    def test_part1(self):
        self.assertEqual(solve(self.input), 161)

    def test_part2(self):
        self.assertEqual(
            solve2(
                "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
            ),
            48,
        )


if __name__ == "__main__":
    if len(sys.argv) == 1:
        unittest.main()
    else:
        input = pathlib.Path("input").read_text()
        print(solve(input))
        print(solve2(input))
