fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(Debug)]
struct Race {
    time: usize,
    distance: usize,
}

fn parse_input(input: &str) -> Vec<Race> {
    let lines: Vec<&str> = input.lines().collect();
    let times = lines[0][5..].split_whitespace().map(|x| x.parse::<usize>().unwrap()).collect::<Vec<usize>>();
    let distances = lines[1][9..].split_whitespace().map(|x| x.parse::<usize>().unwrap()).collect::<Vec<usize>>();
    times.iter().zip(distances.iter())
        .map(|(t, d)| Race { time: *t, distance: *d })
        .collect()
}

fn solve(input: &str) -> usize {
    let races = parse_input(input);
    races.iter().map(|r| {
        let mut wins = 0;
        for i in 0..r.time {
            if (i * (r.time - i)) > r.distance {
                wins += 1;
            }
        }
        wins
    }).product()
}

fn solve_2(input: &str) -> usize {
    let races = parse_input(input);
    let time = races.iter().map(|r| format!("{}",r.time)).collect::<Vec<String>>().join("").parse::<usize>().unwrap();
    let distance = races.iter().map(|r| format!("{}",r.distance)).collect::<Vec<String>>().join("").parse::<usize>().unwrap();
    let mut wins = 0;
    for i in 0..time {
        if (i * (time - i)) > distance {
            wins += 1;
        }
    }
    wins

}

#[test]
fn test1() {
    let input = include_str!("../input_test");
    assert_eq!(solve(input), 288);
}

#[test]
fn test2() {
    let input = include_str!("../input_test");
    assert_eq!(solve_2(input), 71503);
}
