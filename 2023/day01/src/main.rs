use phf::phf_map;
use rstest::rstest;
use std::str::Chars;

fn main() {
    println!("Part 1:   {}", solve(include_str!("../input")));
    println!("Part 2:   {}", solve_2(include_str!("../input")));
    println!(
        "Part 2.1: {}",
        solve(&parse_input_2_1(include_str!("../input")))
    );
}

fn parse_input(input: &str) -> Vec<Chars<'_>> {
    input.lines().map(|line| line.chars()).collect()
}

fn parse_input_2(input: &str) -> Vec<&str> {
    input.lines().collect()
}

fn parse_input_2_1(input: &str) -> String {
    let mut inp = input.to_string();
    for (key, value) in STRING_NUM_STRING.entries() {
        inp = inp.replace(key, value);
    }
    inp
}

fn solve(input: &str) -> usize {
    let _calibration = parse_input(input);
    let line_totals = _calibration
        .iter()
        .map(|line| {
            let l: Vec<_> = line.clone().filter(|c| c.is_ascii_digit()).collect();
            let num_str = format!("{}{}", l[0], l[l.len() - 1]);
            num_str.parse::<usize>().unwrap()
        })
        .collect::<Vec<_>>();
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

static STRING_NUM_STRING: phf::Map<&'static str, &'static str> = phf_map! {
    "one" => "o1n",
    "two" =>  "t2o",
    "three" =>  "th3ree",
    "four" =>  "f4ur",
    "five" =>  "fi5ve",
    "six" =>  "s6x",
    "seven" =>  "se7en",
    "eight" =>  "ei8jt",
    "nine" =>  "n9ne",
};

fn parse_line(line: &&str) -> usize {
    let mut numbers = Vec::new();
    for i in 0..line.len() {
        let c = &line[i..i + 1];
        if c.parse::<usize>().is_ok() {
            numbers.push(c.parse::<usize>().unwrap());
            continue;
        }
        let substr = &line[i..line.len()];
        for (key, value) in STRING_NUM.entries() {
            if substr.starts_with(key) {
                numbers.push(**value);
            }
        }
    }
    let num_str = format!("{}{}", numbers[0], numbers[numbers.len() - 1]);
    num_str.parse::<usize>().unwrap()
}

fn solve_2(input: &str) -> usize {
    let calibrations_input = parse_input_2(input);
    let cals = calibrations_input
        .iter()
        .map(parse_line)
        .collect::<Vec<_>>();
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

#[test]
fn test_suite_2_1() {
    assert_eq!(
        solve(&parse_input_2_1(include_str!("../input_test_2"))),
        281
    );
}
#[rstest]
#[case("two1nine", 29)]
#[case("eightwothree", 83)]
#[case("abcone2threexyz", 13)]
#[case("xtwone3four", 24)]
#[case("4nineeightseven2", 42)]
#[case("zoneight234", 14)]
#[case("7pqrstsixteen", 76)]
fn part2_test(#[case] input: &str, #[case] expected: usize) {
    assert_eq!(expected, solve_2(input))
}
