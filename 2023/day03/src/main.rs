use std::collections::{HashMap, HashSet};

fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug, Hash, Eq, Clone)]
struct Point {
    x: usize,
    y: usize,
}

#[derive(PartialEq, Debug, Hash, Eq, Clone)]
enum Part {
    Symbol(char),
    Number { value: usize, len: usize },
}

#[derive(Debug)]
struct Schemantic {
    numbers: HashMap<Point, Part>,
    symbols: HashMap<Point, Part>,
}

fn parse_input(input: &str) -> Schemantic {
    let mut numbers: HashMap<Point, Part> = HashMap::new();
    let mut symbols: HashMap<Point, Part> = HashMap::new();
    for (y, line) in input.lines().enumerate() {
        //padd the puzzle with an extra column to make parsing easier
        let eline = format!("{line}.");
        let mut in_digit = false;
        let mut digit = String::from("");
        for (x, c) in eline.chars().enumerate() {
            if c.is_ascii_digit() {
                in_digit = true;
                digit.push(c);
            } else {
                if in_digit {
                    numbers.insert(
                        Point {
                            x: x - digit.len(),
                            y,
                        },
                        Part::Number {
                            value: digit.parse().unwrap(),
                            len: digit.len(),
                        },
                    );
                    digit = "".to_string();
                }
                in_digit = false;
                match c {
                    '.' => continue,
                    _ => {
                        symbols.insert(Point { x, y }, Part::Symbol(c));
                    }
                }
            }
        }
    }
    // dbg!(&numbers);
    // dbg!(&symbols);
    Schemantic { numbers, symbols }
}

impl Schemantic {
    fn has_surrounding_symbols(&self, point: &Point) -> bool {
        let symbols = self.get_surrounding_symbols(point);
        !symbols.is_empty()
    }
    fn get_surrounding_symbols(&self, point: &Point) -> Vec<char> {
        let mut symbols = Vec::new();
        let part = self.get_number(point);
        // dbg!(&part);
        let num_len = match part {
            Part::Number { len, .. } => len,
            _ => panic!("Not a number"),
        };
        let start_x = match point.x {
            0 => 0,
            _ => point.x - 1,
        };

        let start_y = match point.y {
            0 => 0,
            _ => point.y - 1,
        };
        for x in start_x..point.x + num_len + 1 {
            for y in start_y..point.y + 2 {
                // println!("Checking {:?}--{:?}", x, y);
                if x == point.x && y == point.y {
                    continue;
                }
                let p = Point { x, y };
                let sym = self.get_symbol(&p);
                match sym {
                    Some(Part::Symbol(c)) => {
                        symbols.push(*c);
                    }
                    _ => continue,
                }
            }
        }
        symbols
    }
    fn get_surrounding_numbers(&self, point: &Point) -> HashSet<usize> {
        let mut numbers = HashSet::new();
        let start_x = match point.x {
            0 => 0,
            _ => point.x - 1,
        };
        let end_x = point.x + 1;
        let start_y = match point.y {
            0 => 0,
            _ => point.y - 1,
        };
        for y in start_y..point.y + 2 {
            let cells_to_check = self.get_numbers_by_row(y);
            // dbg!(&cells_to_check);
            for (cell_point, part) in cells_to_check.iter() {
                let (value, num_len) = match part {
                    Part::Number { len, value } => (value, len),
                    _ => panic!("Not a number"),
                };
                for x in cell_point.x..cell_point.x + num_len {
                    if x >= start_x && x <= end_x {
                        numbers.insert(*value);
                        continue;
                    }
                }
            }
        }
        numbers
    }
    fn get_numbers_by_row(&self, y: usize) -> HashMap<Point, Part> {
        let mut cells: HashMap<Point, Part> = HashMap::new();
        for (point, part) in self.numbers.iter() {
            if point.y == y {
                cells.insert(point.clone(), part.clone());
            }
        }
        cells
    }
    fn get_number(&self, point: &Point) -> Part {
        let p = self.numbers.get(point);
        match p {
            Some(Part::Number { value, len }) => {
                //self.remove_number(point);
                Part::Number {
                    value: *value,
                    len: *len,
                }
            }
            _ => panic!("No number found at {:?}", point),
        }
    }
    fn get_symbol(&self, point: &Point) -> Option<&Part> {
        self.symbols.get(point)
    }
}

fn solve(input: &str) -> usize {
    let schematic = parse_input(input); //parse the input
    let mut total = 0;
    for digit in schematic.numbers.iter() {
        let surrounding_symbols = schematic.has_surrounding_symbols(digit.0);
        // println!("{:?}--{:?}", digit, surrounding_symbols);
        if surrounding_symbols {
            total += match digit.1 {
                Part::Number { value, len: _ } => value,
                _ => panic!("can't get here"),
            };
        }
    }

    total
}

fn solve_2(input: &str) -> usize {
    let schematic = parse_input(input);
    let mut total = 0;
    for sym in schematic
        .symbols
        .iter()
        .filter(|s| s.1 == &Part::Symbol('*'))
    {
        let surrounding_numbers = schematic.get_surrounding_numbers(sym.0);
        println!("\n\n{:?}--{:?}", sym, surrounding_numbers);
        if surrounding_numbers.len() == 2 {
            total += surrounding_numbers.into_iter().product::<usize>();
        }
    }

    total
}

#[test]
fn test_suite_parser() {
    let input = include_str!("../input_test");
    let schematic = parse_input(input);
    assert_eq!(schematic.numbers.len(), 10);
    assert_eq!(schematic.symbols.len(), 6, "symbol len's don't match");
}
#[test]
fn test_suite_1() {
    let input = include_str!("../input_test");
    let res = solve(input);
    assert_eq!(res, 4361);
}
#[test]
fn test_suite_2() {
    let input = include_str!("../input_test");
    let res = solve_2(input);
    assert_eq!(res, 467835);
}
