use std::cmp::Ordering;
use std::collections::HashMap;
fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive( Debug)]
struct CamelHand<'a> {
    cards: &'a str,
    bid: i32
}


#[derive( Debug, PartialEq, Eq)]
enum CamelHandType<'a> {
    FiveOfaKind(CamelHand<'a>),
    FourOfaKind(CamelHand<'a>),
    FullHouse(CamelHand<'a>),
    ThreeOfaKind(CamelHand<'a>),
    TwoPair(CamelHand<'a>),
    OnePair(CamelHand<'a>),
    HighCard(CamelHand<'a>),
}

fn new_hand(input: &str) -> CamelHand {
    let hand =  if let Some((cards, bid)) = input.split_once(' ') {
        CamelHand { cards, bid: bid.parse().unwrap() }
    }
    else {
        panic!("Can't parse card input {input}'");
    };
    hand
}


fn get_groups(ch: &CamelHand) -> HashMap<char, i32>{
    let cards: Vec<char> = ch.cards.chars().collect();

    let mut grouped_c: HashMap<char, i32> = HashMap::new();

    for c in cards {
        grouped_c.entry(c).and_modify(|cnt| { *cnt += 1 }).or_insert(1);
    }

    grouped_c
}

fn new_camel_hand(input: &str) -> CamelHandType {
    let ch = new_hand(input);
    let grouped_c = get_groups(&ch);    

    if grouped_c.len() == 1 {
        return CamelHandType::FiveOfaKind(ch);
    }

    let card_set_sizes = grouped_c.values().collect::<Vec<&i32>>();
    if grouped_c.len() == 2 {
        let first_group_len = *card_set_sizes[0];
        if first_group_len == 1 ||  first_group_len == 4 {
            return CamelHandType::FourOfaKind(ch);
        }
        return CamelHandType::FullHouse(ch);
    }
    
    if grouped_c.len() == 3 {
        let max_pairs = *card_set_sizes.iter().max().unwrap();
        if *max_pairs == 3 {
            return CamelHandType::ThreeOfaKind(ch);
        }
        if card_set_sizes.iter().filter(|x| ***x == 2).collect::<Vec<_>>().len() == 2 {
            return CamelHandType::TwoPair(ch);
        }
    }
    if grouped_c.len() == 4 {
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
            Self::FiveOfaKind(c) | 
                Self::FourOfaKind(c) | 
                Self::FullHouse(c) |
                Self::ThreeOfaKind(c) |
                Self::TwoPair(c) |
                Self::OnePair(c) |
                Self::HighCard(c)
                => c,
        }
    }
}

impl Ord for CamelHandType<'_> {
    fn cmp(&self, other: &Self) -> Ordering {
        println!("Here {:?} vs {:?}", &self, &other);
        println!("{:?} vs {:?}", &self.rank(), &other.rank());
        if self.rank().eq(&other.rank()) {
            return self.cards().cmp(&other.cards());
            }
        else {
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
        let me:&Vec<char> = &self.cards.chars().collect();
        let them: &Vec<char>= &other.cards.chars().collect();
        for (i, _) in me.iter().enumerate() {
            if me[i] == them[i] { continue; }
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

fn parse_input(input: &str) -> usize {
    for line in input.lines() {
        println!("{}", line);
        for c in line.chars() {
            println!("{}", c);
        }
    }
    1
}

fn solve(input: &str) -> usize {
    let _directions = parse_input(input);
    10
}

fn solve_2(input: &str) -> usize {
    let _directions = parse_input(input);
    10
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;
    use std::collections::HashMap;

    fn get_fixtures_hands() -> HashMap<&'static str, CamelHand<'static>> {
        let mut hands = HashMap::new();
        hands.insert("hand1", new_hand("12345 123"));
        hands.insert("hand1_same", new_hand("12345 124"));
        hands.insert("hand2", new_hand("23456 124"));
        hands.insert("hand122", new_hand("12234 124"));
        hands
    }
    
#[test]
fn test_suite_camel_hand() {
    let h = get_fixtures_hands();
    assert_eq!(h.get("hand1").unwrap(), h.get("hand1_same").unwrap());
    assert!(h.get("hand1").unwrap() <h.get("hand2").unwrap());
    assert!(h.get("hand1").unwrap() <h.get("hand2").unwrap());
    assert!(h.get("hand122").unwrap() <h.get("hand2").unwrap());
    assert!(h.get("hand122").unwrap() <h.get("hand1").unwrap());
}

#[test]
fn test_suite_camel_hand_parse() {
    let hand = new_hand("55555 500");
    assert_eq!(hand, CamelHand{cards: "55555", bid: 500});

    assert!( matches!(new_camel_hand("55555 500"),CamelHandType::FiveOfaKind(_)));
    assert!( matches!(new_camel_hand("15555 500"),CamelHandType::FourOfaKind(_)));
    assert!( matches!(new_camel_hand("15551 500"),CamelHandType::FullHouse(_)));
    assert!( matches!(new_camel_hand("95551 500"),CamelHandType::ThreeOfaKind(_)));
    assert!( matches!(new_camel_hand("91915 500"),CamelHandType::TwoPair(_)));
    assert!( matches!(new_camel_hand("99123 500"),CamelHandType::OnePair(_)));
    assert!( matches!(new_camel_hand("12345 500"),CamelHandType::HighCard(_))); 
}
}