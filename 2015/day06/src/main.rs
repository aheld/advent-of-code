use core::panic;
use std::collections::HashMap;
use std::collections::HashSet;

fn main() {
    // println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug)]
enum Action {
    On,
    Off,
    Toggle,
}

#[derive(PartialEq, Hash, Eq, Debug)]
struct Point {
    x: i32,
    y: i32,
}
#[derive(PartialEq, Debug)]
struct Move {
    start: Point,
    end: Point,
    action: Action,
}

fn make_move(action: Action, start: &str, end: &str) -> Move {
    let mut start_point = start.split(",").map(|x| x.parse::<i32>().unwrap());
    let sp = Point {
        x: start_point.next().unwrap(),
        y: start_point.next().unwrap(),
    };
    let mut end_point = end.split(",").map(|x| x.parse::<i32>().unwrap());
    let ep = Point {
        x: end_point.next().unwrap(),
        y: end_point.next().unwrap(),
    };
    return Move {
        start: sp,
        end: ep,
        action,
    };
}

fn parse_on(line: &str) -> Move {
    let split: Vec<&str> = line.split(" ").collect();
    // println!("{:?}", split);
    match split[1] {
        "on" => {
            return make_move(Action::On, split[2], split[4]);
        }
        "off" => {
            return make_move(Action::Off, split[2], split[4]);
        }
        _ => match split[0] {
            "toggle" => {
                return make_move(Action::Toggle, split[1], split[3]);
            }
            _ => panic!("Invalid action"),
        },
    }
}

fn parse_input(input: &str) -> Vec<Move> {
    return input.lines().map(|line| parse_on(line)).collect();
}

fn solve(input: &str) -> usize {
    let moves = parse_input(input);
    let mut lights = HashSet::new();
    for m in moves {
        println!("{:?}", m);
        for i in m.start.x..=m.end.x {
            for j in m.start.y..=m.end.y {
                let p = Point { x: i, y: j };
                match m.action {
                    Action::On => {
                        lights.insert(p);
                    }
                    Action::Off => {
                        lights.remove(&p);
                    }
                    Action::Toggle => {
                        if lights.contains(&p) {
                            lights.remove(&p);
                        } else {
                            lights.insert(p);
                        }
                    }
                }
            }
        }
    }
    return lights.len();
}

fn solve_2(input: &str) -> usize {
    let moves = parse_input(input);
    let mut lights = HashMap::new();
    for m in moves {
        // println!("{:?}", m);
        for i in m.start.x..=m.end.x {
            for j in m.start.y..=m.end.y {
                let p = Point { x: i, y: j };
                match m.action {
                    Action::On => {
                        *lights.entry(p).or_insert(0) += 1;
                    }
                    Action::Off => {
                        if lights.contains_key(&p) {
                            lights.entry(p).and_modify(|x| {
                                if *x > 0 {
                                    *x -= 1;
                                }
                            });
                        }
                    }
                    Action::Toggle => {
                        *lights.entry(p).or_insert(0) += 2;
                    }
                }
            }
        }
    }
    // let vec2 = Vec::from_iter(lights.values());
    // for l in vec2 {
    //     println!("{:?}", l);
    // }

    return lights.values().sum();
}

#[test]
fn test_suite() {
    let input = include_str!("../input_test");
    assert_eq!(solve(input), 1000000 - 1000 - 4);
}
