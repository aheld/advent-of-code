fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug)]
struct Move {
    x: i32,
    y: i32,
}

fn parse_input(input: &str) -> Vec<Move> {
    input
        .chars()
        .map(|c| match c {
            '^' => Move { x: 0, y: 1 },
            'v' => Move { x: 0, y: -1 },
            '>' => Move { x: 1, y: 0 },
            '<' => Move { x: -1, y: 0 },
            _ => Move { x: 0, y: 0 },
        })
        .collect()
}

fn solve(input: &str) -> usize {
    let directions = parse_input(input);
    return 10;
}

fn solve_2(input: &str) -> usize {
    let directions = parse_input(input);
    return 10;
}

#[test]
fn test_suite() {
    let cases = &[("^v", 3), ("^>v<", 3), ("^v^v^v^v^v", 11)];

    for (input, expected) in cases {
        assert_eq!(solve(*input), *expected);
    }
}
