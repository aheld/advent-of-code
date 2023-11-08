fn main() {
    println!("Part 1 {}", solve("iwrupvqb"));
    println!("Part 2 {}", solve_2("iwrupvqb"));
}

fn solve(input: &str) -> i32 {
    for n in 1..1609043 {
        let digest = md5::compute(format!("{}{}", input, n).as_bytes());
        //println!("{:x}", digest);
        if starts_with_five_zeroes(&format!("{:x}", digest)) {
            return n;
        }
    }
    0
}

fn solve_2(input: &str) -> i32 {
    for n in 9000000..10609043 {
        let digest = md5::compute(format!("{}{}", input, n).as_bytes());

        if format!("{:x}", digest).starts_with("000000") {
            return n;
        }
    }
    0
}

fn starts_with_five_zeroes(input: &str) -> bool {
    input.starts_with("00000")
}

#[test]
fn test_starts_with_five_zeroes() {
    assert_eq!(starts_with_five_zeroes("00000abcdef609043"), true);
    assert_eq!(starts_with_five_zeroes("0000pqrstuv1048970"), false);
}

#[test]
fn test_solve() {
    assert_eq!(solve("abcdef"), 609043);
    assert_eq!(solve("pqrstuv"), 1048970);
}
