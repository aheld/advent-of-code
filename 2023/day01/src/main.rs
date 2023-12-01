use std::str::Chars;

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

fn solve(input: &str) -> usize {
    let _calibration = parse_input(input);
    // println!("{:?}", _calibration);
    let line_totals = _calibration.iter().map(|line| {
        let l: Vec<_> = line.clone().filter(|c| c.is_digit(10)).collect();
        let num_str = format!("{}{}", l[0], l[l.len()-1]);
        num_str.parse::<usize>().unwrap()
    }).collect::<Vec<_>>();
    println!("{:?}", line_totals);
    return line_totals.iter().sum();
}

fn solve_2(input: &str) -> usize {
    let _directions = parse_input(input);
    return 10;
}

#[test]
fn test_suite() {
    let input = include_str!("../input_test");
    assert_eq!(solve(input), 142);
}
