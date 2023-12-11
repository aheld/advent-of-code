fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug, Hash, Eq)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug)]
enum Part {
    Symbol(char),
    Number { vaule: i32, len: usize },
}

fn parse_input(input: &str) -> usize {
    for line in input.lines() {
        println!("{}", line);
        for c in line.chars() {
            println!("{}", c);
        }
    }
    1
}

fn solve(input: &str) -> usize {
    let _directions = parse_input(input);
    10
}

fn solve_2(input: &str) -> usize {
    let _directions = parse_input(input);
    10
}

#[test]
fn test_suite_parser() {
    let input = include_str!("../input_test");
    assert_eq!(parse_input(input), 3);
}
