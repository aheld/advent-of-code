use std::cmp::Ordering;
use std::collections::HashMap;
fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(Debug)]
struct CamelHand<'a> {
    cards: &'a str,
    bid: i32,
}

#[derive(Debug, PartialEq, Eq)]
enum CamelHandType<'a> {
    FiveOfaKind(CamelHand<'a>),
    FourOfaKind(CamelHand<'a>),
    FullHouse(CamelHand<'a>),
    ThreeOfaKind(CamelHand<'a>),
    TwoPair(CamelHand<'a>),
    OnePair(CamelHand<'a>),
    HighCard(CamelHand<'a>),
}

impl CamelHand<'_> {
    pub fn new_hand(input: &str) -> CamelHand {
        let hand = if let Some((cards, bid)) = input.split_once(' ') {
            CamelHand {
                cards,
                bid: bid.parse().unwrap(),
            }
        } else {
            panic!("Can't parse card input {input}'");
        };
        hand
    }

    pub fn get_groups(ch: &CamelHand) -> HashMap<char, i32> {
        let cards: Vec<char> = ch.cards.chars().collect();
    
        let mut grouped_c: HashMap<char, i32> = HashMap::new();
    
        for c in cards {
            grouped_c.entry(c).and_modify(|cnt| *cnt += 1).or_insert(1);
        }
    
        grouped_c
    }

    fn consolidate_jokers(grouped_c: &mut HashMap<char, i32>) {
        let jokers = grouped_c.get(&'X');
        if let Some(j_count) = jokers {
            let max_key = &grouped_c
                .clone()
                .into_iter()
                .filter(|x| x.0 != 'X')
                .max_by(|a, b| a.1.cmp(&b.1))
                .map(|(k, _v)| k);
    
            match max_key {
                Some(k) => {
                    println!("Adding {:?} for X-{:?} for {:?}", k, j_count, max_key);
                    dbg!(&grouped_c);
                    *grouped_c.entry(*k).or_insert(0) += grouped_c[&'X'];
                    grouped_c.remove(&'X');
                    dbg!(&grouped_c);
                }
                None => (), // All Jokers!
            }
        }
    }
    
}


fn new_camel_hand(input: &str) -> CamelHandType {
    let ch = CamelHand::new_hand(input);
    let mut grouped_c = CamelHand::get_groups(&ch);

    CamelHand::consolidate_jokers(&mut grouped_c); // this is the enhancement for part 2

    if grouped_c.values().any(|v| *v == 5) {
        return CamelHandType::FiveOfaKind(ch);
    }
    if grouped_c.values().any(|v| *v == 4) {
        return CamelHandType::FourOfaKind(ch);
    }
    if grouped_c.values().any(|v| *v == 3) {
        if grouped_c.values().any(|v| *v == 2) {
            return CamelHandType::FullHouse(ch);
        } else {
            return CamelHandType::ThreeOfaKind(ch);
        }
    }
    if grouped_c.values().filter(|v| **v == 2).count() == 2 {
        return CamelHandType::TwoPair(ch);
    }
    if grouped_c.values().filter(|v| **v == 2).count() == 1 {
        return CamelHandType::OnePair(ch);
    }

    CamelHandType::HighCard(ch)
}

impl PartialOrd for CamelHandType<'_> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl CamelHandType<'_> {
    fn rank(&self) -> i16 {
        match &self {
            CamelHandType::FiveOfaKind(_) => 7,
            CamelHandType::FourOfaKind(_) => 6,
            CamelHandType::FullHouse(_) => 5,
            CamelHandType::ThreeOfaKind(_) => 4,
            CamelHandType::TwoPair(_) => 3,
            CamelHandType::OnePair(_) => 2,
            CamelHandType::HighCard(_) => 1,
        }
    }

    // This feels wrong
    fn cards(&self) -> &CamelHand {
        match &self {
            Self::FiveOfaKind(c)
            | Self::FourOfaKind(c)
            | Self::FullHouse(c)
            | Self::ThreeOfaKind(c)
            | Self::TwoPair(c)
            | Self::OnePair(c)
            | Self::HighCard(c) => c,
        }
    }
}

fn rank_cards(c: char) -> i32 {
    if c.is_ascii_digit() {
        return c.to_digit(10).unwrap() as i32;
    }
    match c {
        'A' => 14,
        'K' => 13,
        'Q' => 12,
        'J' => 11,
        'T' => 10,
        'X' => -1, //going to swap J for X for part 2 and I'm too lazy to abstract this part
        _ => panic!("unknown card value for {}", c),
    }
}

impl Ord for CamelHandType<'_> {
    fn cmp(&self, other: &Self) -> Ordering {
        // println!("Here {:?} vs {:?}", &self, &other);
        // println!("{:?} vs {:?}", &self.rank(), &other.rank());
        if self.rank().eq(&other.rank()) {
            for (s, o) in self.cards().cards.chars().zip(other.cards().cards.chars()) {
                // println!("{} : {}", s, o);
                // println!("{} : {}", rank_cards(s), rank_cards(o));
                if rank_cards(s) == rank_cards(o) {
                    continue;
                } else {
                    return rank_cards(s).cmp(&rank_cards(o));
                }
            }
            return Ordering::Equal;
        } else {
            return self.rank().cmp(&other.rank());
        }
    }
}

impl PartialOrd for CamelHand<'_> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for CamelHand<'_> {
    fn cmp(&self, other: &Self) -> Ordering {
        let me: &Vec<char> = &self.cards.chars().collect();
        let them: &Vec<char> = &other.cards.chars().collect();
        for (i, _) in me.iter().enumerate() {
            if me[i] == them[i] {
                continue;
            }
            return me[i].cmp(&them[i]);
        }
        return std::cmp::Ordering::Equal;
    }
}

impl PartialEq for CamelHand<'_> {
    fn eq(&self, other: &Self) -> bool {
        self.cards == other.cards
    }
}

impl Eq for CamelHand<'_> {}

fn solve(input: &str) -> usize {
    let mut hands = input
        .lines()
        .map(new_camel_hand)
        .collect::<Vec<CamelHandType>>();
    hands.sort();
    let mut total = 0;
    for (i, h) in hands.iter().enumerate() {
        // println!("{i} {:?}", h);
        total += (i + 1) * h.cards().bid as usize
    }
    total
}

fn solve_2(input: &str) -> usize {
    let input = input.replace("J", "X");
    solve(&input)
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;
    use std::collections::HashMap;

    // Note - I don't commit real inputs as per AOC rules.
    const INPUT: &str = include_str!("../input");
    const TEST_INPUT: &str = include_str!("../input_test");

    fn get_fixtures_hands() -> HashMap<&'static str, CamelHand<'static>> {
        let mut hands = HashMap::new();
        hands.insert("hand1", CamelHand::new_hand("12345 123"));
        hands.insert("hand1_same", CamelHand::new_hand("12345 124"));
        hands.insert("hand2", CamelHand::new_hand("23456 124"));
        hands.insert("hand122", CamelHand::new_hand("12234 124"));
        hands
    }

    #[test]
    fn test_suite_camel_hand() {
        let h = get_fixtures_hands();
        assert_eq!(h.get("hand1").unwrap(), h.get("hand1_same").unwrap());
        assert!(h.get("hand1").unwrap() < h.get("hand2").unwrap());
        assert!(h.get("hand1").unwrap() < h.get("hand2").unwrap());
        assert!(h.get("hand122").unwrap() < h.get("hand2").unwrap());
        assert!(h.get("hand122").unwrap() < h.get("hand1").unwrap());
    }

    #[test]
    fn test_suite_camel_hand_parse() {
        let hand = CamelHand::new_hand("55555 500");
        assert_eq!(
            hand,
            CamelHand {
                cards: "55555",
                bid: 500
            }
        );

        assert!(matches!(
            new_camel_hand("55555 500"),
            CamelHandType::FiveOfaKind(_)
        ));
        assert!(matches!(
            new_camel_hand("15555 500"),
            CamelHandType::FourOfaKind(_)
        ));
        assert!(matches!(
            new_camel_hand("15551 500"),
            CamelHandType::FullHouse(_)
        ));
        assert!(matches!(
            new_camel_hand("95551 500"),
            CamelHandType::ThreeOfaKind(_)
        ));
        assert!(matches!(
            new_camel_hand("91915 500"),
            CamelHandType::TwoPair(_)
        ));
        assert!(matches!(
            new_camel_hand("99123 500"),
            CamelHandType::OnePair(_)
        ));
        assert!(matches!(
            new_camel_hand("12345 500"),
            CamelHandType::HighCard(_)
        ));

        assert!(new_camel_hand("TJKA7 213") > new_camel_hand("TJ45K 434"));
    }

    #[test]
    fn test_solve() {
        assert_eq!(solve(TEST_INPUT), 6440);
    }

    #[test]
    fn test_solve_2() {
        assert_eq!(solve_2(TEST_INPUT), 5905);
    }
    #[test]
    fn test_real_inputs() {
        assert_eq!(solve(INPUT), 251029473);
        assert_eq!(solve_2(INPUT), 251003917);
    }
}
