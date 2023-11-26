use std::cell::RefCell;
use std::collections::HashMap;
use std::sync::atomic::{AtomicUsize, Ordering};

fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug, Clone)]
enum Gate {
    SignalWire(String),
    And(String, String),
    Or(String, String),
    LShift(String, u16),
    RShift(String, u16),
    Not(String),
}

struct Context {
    gates: HashMap<String, Gate>,
    wires: RefCell<HashMap<String, u16>>,
    counter: AtomicUsize,
}

fn parse_input(input: &str) -> Vec<(Gate, String)> {
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
                Some(_) => panic!("Unknown op"),
                None => (Gate::SignalWire(first.to_string()), wire.to_string()),
            }
        })
        .collect()
}

impl Context {
    pub fn new(gates: HashMap<String, Gate>) -> Context {
        let counter = AtomicUsize::new(0);
        let wires = RefCell::new(HashMap::new());
        Context {
            gates,
            counter,
            wires,
        }
    }

    pub fn get(&self, wire: &str) -> u16 {
        // let wire = wire.to_string();
        self.counter.fetch_add(1, Ordering::SeqCst);

        if let Some(value) = self.wires.borrow().get(wire) {
            return *value;
        }

        let value = wire.parse::<u16>();
        let ret_value = match value {
            Ok(n) => n,
            Err(_) => {
                let gate = self.gates.get(wire).unwrap();
                // if self.counter.load(Ordering::SeqCst) % 1000000 == 0 {
                // println!("{:?}: gate: {}=> {:?}", self.counter, wire, gate);
                // }
                match gate {
                    Gate::SignalWire(w) => self.get(w),
                    Gate::And(a, b) => self.get(a) & self.get(b),
                    Gate::Or(a, b) => self.get(a) | self.get(b),
                    Gate::LShift(a, b) => self.get(a) << b,
                    Gate::RShift(a, b) => self.get(a) >> b,
                    Gate::Not(a) => {
                        let v = self.get(a);
                        !v
                    }
                }
            }
        };
        self.wires.borrow_mut().insert(wire.to_string(), ret_value);
        ret_value
    }

    pub fn override_b(&self, signal: u16) {
        self.wires.borrow_mut().clear();
        self.wires.borrow_mut().insert("b".to_string(), signal);
    }
}

fn solve(input: &str) -> u16 {
    let input_gates = parse_input(input);
    let mut gates = HashMap::new();
    for g in input_gates.iter() {
        gates.insert(g.1.to_string(), g.0.clone());
    }
    let c = Context::new(gates);
    c.get("a")
}

fn solve_2(input: &str) -> u16 {
    let input_gates = parse_input(input);
    let mut gates = HashMap::new();
    for g in input_gates.iter() {
        gates.insert(g.1.to_string(), g.0.clone());
    }
    let c = Context::new(gates);
    c.override_b(c.get("a"));
    c.get("a")
}

#[test]
fn test_suite() {
    let gates = parse_input(include_str!("../input_test"));
    print!("{:?}", gates);
    assert_eq!(gates.len(), 9);
    assert_eq!(gates[2], (Gate::SignalWire("456"), "y".to_string()));
    assert_eq!(gates[3], (Gate::Not("y".to_string()), "a".to_string()));
}

#[test]
fn test_suite_part1_i() {
    assert_eq!(solve(include_str!("../input_test")), 65079);
}
