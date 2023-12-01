use std::str::Chars;
use phf::{phf_map};

fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}


fn parse_input(input: &str) -> Vec<Chars<'_>> {
    input
        .lines()
        .map(|line| line.chars())
        .collect()
}

fn parse_input_2(input: &str)-> Vec<&str> {
    input.lines().collect()
}

fn solve(input: &str) -> usize {
    let _calibration = parse_input(input);
    // println!("{:?}", _calibration);
    let line_totals = _calibration.iter().map(|line| {
        let l: Vec<_> = line.clone().filter(|c| c.is_digit(10)).collect();
        let num_str = format!("{}{}", l[0], l[l.len()-1]);
        num_str.parse::<usize>().unwrap()
    }).collect::<Vec<_>>();
    // println!("{:?}", line_totals);
    return line_totals.iter().sum();
}

static STRING_NUM: phf::Map<&'static str, &'static usize> = phf_map! {
    "one" => &1,
    "two" =>  &2,
    "three" =>  &3,
    "four" =>  &4,
    "five" =>  &5,
    "six" =>  &6,
    "seven" =>  &7,
    "eight" =>  &8,
    "nine" =>  &9,
};

fn solve_2(input: &str) -> usize {
    let calibrations_input = parse_input_2(input);
    let cals = calibrations_input.iter().map(|line| {
        let mut numbers = Vec::new();
        for i in 0..line.len() {
            let c = &line[i..i+1];
            if c.parse::<usize>().is_ok() {
                numbers.push(c.parse::<usize>().unwrap().clone());
                println!("********{}", c);
                continue;
            }
            let substr = &line[i..line.len()];
            println!("{}", substr);
            for (key, value) in STRING_NUM.entries() {
                if substr.starts_with(key) {
                    println!("\n******\n{} {}", key, value);
                    numbers.push(**value);
                }
            }
        }
        let num_str = format!("{}{}", numbers[0], numbers[numbers.len()-1]);
        num_str.parse::<usize>().unwrap()
    }).collect::<Vec<_>>();
    println!("CALS\n{:?}", cals);
    return cals.iter().sum();
}

#[test]
fn test_suite() {
    let input = include_str!("../input_test");
    assert_eq!(solve(input), 142);
}

#[test]
fn test_suite_2() {
    let input = include_str!("../input_test_2");
    assert_eq!(solve_2(input), 281);
}
