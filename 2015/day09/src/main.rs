fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

fn parse_input(input: &str) -> Vec<&str> {
    return input.split("\n").collect();
}
/*
* solve Advent of Code Day 09, year 2015
* (See http://adventofcode.com/day/9.)
*/
fn solve(input: &str) -> usize {
    let _directions = parse_input(input);
    return 1;
}

fn solve_2(input: &str) -> usize {
    let _directions = parse_input(input);
    return 1;
}

#[test]
fn test_suite() {
    let cases = &[("^v", 3), ("^>v<", 3), ("^v^v^v^v^v", 11)];

    for (input, expected) in cases {
        assert_eq!(solve(*input), *expected);
    }
}
