use nom::{
    bytes::complete::tag,
    character::complete::{self, multispace0, multispace1},
    sequence::tuple,
    IResult,
};

fn main() {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::DEBUG)
        .with_target(false)
        .init();
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

pub fn main_for_bench() {
    solve(include_str!("../input"));
    solve_2(include_str!("../input"));
}

#[derive(PartialEq, Debug)]
pub struct Card {
    pub id: i32,
    pub my_numbers: Vec<i32>,
    pub card_numbers: Vec<i32>,
    pub value: i32,
}

fn parse_card(input: &str) -> IResult<&str, Card> {
    let (input, _) = tuple((tag("Card"), multispace1))(input)?;
    let (input, id) = complete::i32(input)?;
    let (input, _) = tag(": ")(input)?;
    //parser to grab 41 48 83 86 17
    let win_numbers_parser =
        nom::multi::separated_list1(multispace1, tuple((multispace0, complete::i32)));
    let my_numbers_parser =
        nom::multi::separated_list1(multispace1, tuple((multispace0, complete::i32)));
    //parser to get [list] | [ list]
    let mut list_parser = tuple((win_numbers_parser, tag(" | "), my_numbers_parser));
    let (input, parsed_raw) = list_parser(input)?;
    let card_numbers: Vec<i32> = parsed_raw.0.iter().map(|(_, n)| *n).collect::<Vec<i32>>();
    let my_numbers: Vec<i32> = parsed_raw.2.iter().map(|(_, n)| *n).collect::<Vec<i32>>();
    // let my_numbers = vec![1, 2, 3, 4, 5, 6, 7, 8];
    // let card_numbers = vec![1, 2, 3, 4, 5, 6, 7, 8];
    Ok((
        input,
        Card {
            id,
            my_numbers,
            card_numbers,
            value: 1,
        },
    ))
}

fn parse_input(input: &str) -> Vec<Card> {
    input
        .lines()
        .map(|l| {
            let res = parse_card(l);
            res
        })
        .map(|result| match result {
            Ok((_, card)) => card,
            Err(e) => panic!("error parsing input: {:?}", e),
        })
        .collect()
}

impl Card {
    pub fn get_matching(&self) -> Vec<i32> {
        self.card_numbers
            .iter()
            .filter(|n| self.my_numbers.contains(n))
            .map(|n| *n)
            .collect()
    }
}
pub fn solve(input: &str) -> i32 {
    let mut total = 0;
    let cards = parse_input(input);
    for card in &cards {
        let matching_numbers = card.get_matching();
        dbg!(&matching_numbers);
        total += match matching_numbers.len() {
            0 => 0,
            1 => 1,
            x if x > 1 => {
                let base: u32 = 2;
                let power: u32 = x as u32;
                base.pow(power - 1)
            }
            _ => panic!("should not be here"),
        };
    }
    total as i32
}

pub fn solve_2(_input: &str) -> i32 {
    1
}

#[test]
fn test_single_line() {
    let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53";
    let parsed = parse_card(input).unwrap();
    // dbg!(&parsed);gcc
    assert_eq!(parsed.1.id, 1, "game id should be 1");
    assert_eq!(
        parsed.1.my_numbers,
        vec![83, 86, 6, 31, 17, 9, 48, 53],
        "my numbers should match"
    );
}

#[test]
fn test_single_line_real() {
    let input = "Card   1: 69 24 51 87  9 49 17 16 21 48 |  5 52 86 35 57 18 60 84 50 76 96 47 38 41 34 36 55 20 25 37  6 70 66 45  3";
    let parsed = parse_card(input).unwrap();
    // dbg!(&parsed);gcc
    assert_eq!(parsed.1.id, 1, "game id should be 1");
    assert_eq!(
        parsed.1.my_numbers,
        vec![
            5, 52, 86, 35, 57, 18, 60, 84, 50, 76, 96, 47, 38, 41, 34, 36, 55, 20, 25, 37, 6, 70,
            66, 45, 3
        ],
        "my numbers should match"
    );
}

#[test]
fn test_input_1() {
    let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53";
    let cards = parse_input(input);
    assert_eq!(cards.len(), 1, "should be 1 games");
    assert_eq!(solve(input), 8);
}
#[test]
fn test_solve_1() {
    let input = include_str!("../input_test");
    let total = solve(input);
    assert_eq!(total, 13);
}

#[test]
fn test_solve_2() {
    let input = include_str!("../input_test");
    let res = solve_2(input);
    assert_eq!(res, 2286);
}
