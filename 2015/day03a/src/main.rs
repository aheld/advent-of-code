use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
    println!("{}", solve(include_str!("../input")));
}

#[derive(PartialEq, Debug, Eq, Hash, Clone, Copy)]
struct House {
    x: i32,
    y: i32,
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

fn count_houses(directions: &Vec<Move>) -> usize {
    let mut houses = HashSet::new();
    let mut current_house = House { x: 0, y: 0 };
    houses.insert(current_house);
    for direction in directions {
        current_house.x += direction.x;
        current_house.y += direction.y;
        houses.insert(current_house);
    }
    houses.len()
}

fn solve(input: &str) -> usize {
    let directions = parse_input(input);
    count_houses(&directions)
}

#[test]
fn test_single() {
    let input = ">";
    let directions = parse_input(input);
    assert_eq!(directions[0].x, 1);
    assert_eq!(directions[0].y, 0);
}

#[test]
fn test_moves() {
    let mut input = "^>v<";
    let mut count = solve(input);
    assert_eq!(count, 4);

    input = "^v^v^v^v^v";
    count = solve(input);
    assert_eq!(count, 2);
}
