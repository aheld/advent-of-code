fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

fn parse_input(input: &str) -> Vec<Vec<i64>> {
    input
        .lines()
        .map(|line| line.split(" ").map(|n| n.parse::<i64>().unwrap()).collect())
        .collect()
}

fn predict(mut diffs: Vec<i64>) -> i64 {
    let mut prediction = 0;

    while diffs.iter().any(|&n| n != 0) {
        prediction += get_diffs(&mut diffs);
    }

    prediction
}

fn get_diffs(num: &mut Vec<i64>) -> i64 {
    for i in 0..num.len() - 1 {
        num[i] = num[i + 1] - num[i];
    }
    num.pop().unwrap()
}

fn solve(input: &str) -> i64 {
    let oasis = parse_input(input);
    oasis.iter().map(|diffs| predict(diffs.clone())).sum()
}

fn solve_2(input: &str) -> i64 {
    let oasis = parse_input(input);
    oasis
        .iter()
        .map(|diffs| predict(diffs.iter().rev().copied().collect::<Vec<i64>>()))
        .sum()
}
#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
";

    #[test]
    fn test_solve() {
        assert_eq!(solve(INPUT), 114);
    }
    #[test]
    fn test_solve_2() {
        assert_eq!(solve_2(INPUT), 2);
    }
}
