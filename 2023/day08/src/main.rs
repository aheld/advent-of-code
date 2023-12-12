use std::collections::HashMap;

fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(Debug, Clone)]
struct Direction {
    l: String,
    r: String,
}
#[derive(Debug,Clone)]
struct Directions {
    steps: Vec<char>,
    mapping: HashMap<String, Direction>
}

fn parse_input(input: &str) -> Directions {
    let mut mapping = HashMap::new();
    let (steps, map_in)=input.split_at(input.find("\n").unwrap());
    for line in map_in.lines() {
        if line.is_empty() {
            continue;
        }
        let (key, value) = line.split_at(line.find("=").unwrap());
        let (l, r) = value.split_once(", ").unwrap();
        mapping.insert(key.trim().to_string(), Direction { l: l[3..].to_string(), r: r[0..3].to_string() });
    }
    Directions {
        steps: steps.chars().collect(),
        mapping
    }
}

fn solve(input: &str) -> usize {
    let directions = parse_input(input);
    navigate_map(directions, "AAA", "ZZZ")
}

fn navigate_map(directions :Directions, start: &str, end: &str) -> usize {
    let mut location = start.to_string();
    let mut i = 0;
    let mut j = 0;
    while !location.ends_with(end) {
        j+=1;
        let next_step = directions.steps[i];
        if i == directions.steps.len() - 1 {
            i = 0;
        } else {
            i += 1;
        }
        let direction = directions.mapping.get(&location).unwrap();
        if next_step == 'R' {
            location = direction.r.clone();
        } else {
            location = direction.l.clone();
        }
    }
    j
}

fn solve_2(input: &str) -> usize {
    let directions = parse_input(input);
    let locations = directions.mapping.keys().filter(|k| k.ends_with("A")).collect::<Vec<_>>();
    let mut counts = HashMap::<String, usize>::new();
    
    for location in locations.iter() {
        let count = navigate_map(directions.clone(), location, "Z");
        counts.insert(location.to_string(), count);
    }

    dbg!(&counts);
    least_common_multiple(&counts.values().cloned().collect::<Vec<_>>())
}

fn least_common_multiple(nums: &[usize]) -> usize {
    let mut result = 1;
    for &num in nums {
        result = num * result / gcd(num, result);
    }
    result
}

fn gcd(a: usize, b: usize) -> usize {
    if b == 0 {
        return a;
    }

    gcd(b, a % b)
}

#[cfg(test)]
mod tests {
    use super::*;

    const TEST_INPUT: &str="RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
";

    const TEST_INPUT_2: &str="LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
";
    
    const TEST_INPUT_3: &str="LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
";
    
    #[test]
    fn test_solve() {
        assert_eq!(solve(TEST_INPUT), 2);
        assert_eq!(solve(TEST_INPUT_2), 6);
    }

    #[test]
    fn test_solve_2() {
        assert_eq!(solve_2(TEST_INPUT_3), 6);
    }
    #[test]
    fn test_solve_lcm() {
        assert_eq!(least_common_multiple(&[3,6,27]), 54);
    }
}
