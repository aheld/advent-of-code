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
    let mut robo_current_house = House { x: 0, y: 0 };
    houses.insert(current_house);

    let mut santa = true;

    for direction in directions {
        if santa {
            current_house.x += direction.x;
            current_house.y += direction.y;
            houses.insert(current_house);
        } else {
            robo_current_house.x += direction.x;
            robo_current_house.y += direction.y;
            houses.insert(robo_current_house);
        }

        santa = !santa;
    }
    houses.len()
}

fn solve(input: &str) -> usize {
    let directions = parse_input(input);
    count_houses(&directions)
}

// #[test]
// fn test_moves() {
//     let mut input = "^v";
//     let mut count = solve(input);
//     assert_eq!(count, 3);

//     input = "^>v<";
//     count = solve(input);
//     assert_eq!(count, 3);

//     input = "^v^v^v^v^v";
//     count = solve(input);
//     assert_eq!(count, 11);
// }

#[test]
fn test_suite() {
    let cases = &[("^v", 3), ("^>v<", 3), ("^v^v^v^v^v", 11)];

    for (input, expected) in cases {
        assert_eq!(solve(*input), *expected);
    }
}
