use std::collections::{HashMap, VecDeque};

fn main() {
    println!("Running main");
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("DONE");
    // println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug)]
enum Gate {
    Signal(u16),
    SignalWire(String),
    And(String, String),
    Or(String, String),
    LShift(String, u16),
    RShift(String, u16),
    Not(String),
}
fn parse_input(input: &str) -> Vec<(Gate, String)> {
    println!("Parse Input");
    input
        .lines()
        .map(|line| {
            let mut parts = line.split(" -> ");
            let op_part = parts.next().unwrap();
            let wire = parts.next().unwrap();
            let mut op_parts = op_part.split(" ");
            let first = op_parts.next().unwrap();
            if first == "NOT" {
                let signal = op_parts.next().unwrap();
                println!("In NOT");
                return (Gate::Not(signal.to_string()), wire.to_string());
            }
            let maybe_second = op_parts.next();
            match maybe_second {
                Some("AND") => {
                    let third = op_parts.next().unwrap();
                    (
                        Gate::And(first.to_string(), third.to_string()),
                        wire.to_string(),
                    )
                }
                Some("OR") => {
                    let third = op_parts.next().unwrap();
                    (
                        Gate::Or(first.to_string(), third.to_string()),
                        wire.to_string(),
                    )
                }
                Some("LSHIFT") => {
                    let third = op_parts.next().unwrap();
                    (
                        Gate::LShift(first.to_string(), third.parse::<u16>().unwrap()),
                        wire.to_string(),
                    )
                }
                Some("RSHIFT") => {
                    let third = op_parts.next().unwrap();
                    (
                        Gate::RShift(first.to_string(), third.parse::<u16>().unwrap()),
                        wire.to_string(),
                    )
                }
                Some(_) => {
                    println!("{:?}", maybe_second);
                    (
                        Gate::Signal(first.parse::<u16>().unwrap()),
                        wire.to_string(),
                    )
                }
                None => {
                    let sig = first.parse();
                    match sig {
                        Ok(s) => (Gate::Signal(s), wire.to_string()),
                        Err(_) => (Gate::SignalWire(first.to_string()), wire.to_string()),
                    }
                }
            }
        })
        .collect()
}

fn solve(input: &str) -> u16 {
    let mut circuit: HashMap<String, u16> = HashMap::new();
    let gates = parse_input(input);
    for g in gates {
        println!("{:?}", g);
    }
    return 0;
    let mut queue = VecDeque::from_iter(gates.iter());
    println!("Process Queue");
    while let Some(g) = queue.pop_front() {
        //println!("{:?}", queue);
        // println!("{:?}", queue.len());
        // println!("{:?}", circuit.len());
        // println!("{:?}", g);
        match g {
            (Gate::Signal(n), wire) => {
                circuit.insert(wire.to_string(), *n);
            }
            (Gate::SignalWire(a), wire) => {
                let a_val = circuit.get(a);
                if a_val.is_none() {
                    queue.push_back(g);
                    continue;
                }
                circuit.insert(wire.to_string(), *a_val.unwrap());
            }
            (Gate::And(a, b), wire) => {
                let a_val = circuit.get(a);
                let b_val = circuit.get(b);
                match (a_val, b_val) {
                    (Some(a), Some(b)) => {
                        circuit.insert(wire.to_string(), a & b);
                    }
                    _ => {
                        queue.push_back(g);
                    }
                }
            }
            (Gate::Or(a, b), wire) => {
                let a_val = circuit.get(a);
                let b_val = circuit.get(b);
                match (a_val, b_val) {
                    (Some(a), Some(b)) => {
                        circuit.insert(wire.to_string(), a | b);
                    }
                    _ => {
                        queue.push_back(g);
                    }
                }
            }
            (Gate::LShift(a, b), wire) => {
                let a_val = circuit.get(a);
                if a_val.is_none() {
                    queue.push_back(g);
                    continue;
                }
                circuit.insert(wire.to_string(), a_val.unwrap() << b);
            }
            (Gate::RShift(a, b), wire) => {
                let a_val = circuit.get(a);
                if a.contains("f") && a.len() == 1 {
                    println!("RSHIFT {:?}: {:?}", a_val, a);
                }
                if a_val.is_none() {
                    queue.push_back(g);
                    continue;
                }
                circuit.insert(wire.to_string(), a_val.unwrap() >> b);
            }
            (Gate::Not(a), wire) => {
                let a_val = circuit.get(a);
                if a_val.is_none() {
                    queue.push_back(g);
                    continue;
                }
                circuit.insert(wire.to_string(), !a_val.unwrap());
            }
        }
        // println!("{:?} post", queue.len());
    }
    // for c in circuit {
    //     println!("{:?}", c);
    // }
    return *circuit.get("a").unwrap();
}

// fn solve_2(input: &str) -> u16 {
//     let directions = parse_input(input);
//     return 10;
// }

#[test]
fn test_suite() {
    let gates = parse_input(include_str!("../input_test"));
    print!("{:?}", gates);
    assert_eq!(gates.len(), 8);
    assert_eq!(gates[0], (Gate::Signal(123), "x".to_string()));
    assert_eq!(gates[7], (Gate::Not("y".to_string()), "a".to_string()));
}

#[test]
fn test_suite_part1_i() {
    assert_eq!(solve(include_str!("../input_test")), 65079);
}
