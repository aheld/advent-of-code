use nom::{
    bytes::complete::tag,
    character::complete::{self, alpha1, multispace1},
    multi::separated_list1,
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

#[derive(PartialEq, Debug)]
pub struct Draw {
    pub red: i32,
    pub green: i32,
    pub blue: i32,
}

#[derive(PartialEq, Debug)]
pub struct Game {
    pub draws: Vec<Draw>,
    pub id: i32,
}

// Nom parser for a single line that returns a Game struct
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// 1. Parse the game id
// 2. write a sub parser for each draw
// 3. wrap that parser in a separated list by ';'
fn parse_game(input: &str) -> IResult<&str, Game> {
    let (input, _) = tag("Game ")(input)?;
    let (input, id) = complete::i32(input)?;
    let (input, _) = tag(":")(input)?;
    let count_parser = nom::multi::separated_list1(
        tag(","),
        nom::sequence::tuple((
            multispace1,
            complete::i32,
            nom::character::complete::char(' '),
            alpha1,
        )),
    );
    let mut draw_parser  =  separated_list1(tag(";"), count_parser);
    let (input, mut draws_raw) = draw_parser(input)?;
    let draws:Vec<Draw> = draws_raw
        .iter_mut()
        .map(|draw| {
            let mut d = Draw {
                red: 0,
                green: 0,
                blue: 0,
            };
            for (_, count, _, color) in draw.iter_mut() {
                match *color {
                    "red" => d.red = *count,
                    "green" => d.green = *count,
                    "blue" => d.blue = *count,
                    _ => panic!("unknown color"),
                }
            };
            d
        }).collect();
    Ok((
        input,
        Game {
            draws,
            id,
        },
    ))
}

fn parse_input(input: &str) -> Vec<Game> {
    input
        .lines()
        .map(|l| {
            let res = parse_game(l);
            res})
        .map(|result|{
             match result {
            Ok((_, game)) => game,
            Err(e) => panic!("error parsing input: {:?}", e),
            }
        })
        .collect()
}

fn solve(input: &str) -> i32 {
    let games = parse_input(input);
    let pg = games.iter().filter(|g| {
        let mut possible = true;
            for draw in &g.draws {
                // println!("Game: {:?}, red: {:?}", g.id, draw.red);
            if draw.red > 12 || draw.green > 13 || draw.blue > 14 {
                possible = false;
                break;
            }
        }
        // println!("Game: {:?}:=> {:?}", g.id, possible);        
        possible
        })
        .map(|g| g.id)
        .collect::<Vec<i32>>();
        pg.iter().sum()
}

fn solve_2(input: &str) -> i32 {
    let games = parse_input(input);
    let min_cubes:Vec<Draw> = games.iter()
    .map(|g| 
        {
            // println!("Game: {:?}: {:?}", g.id, g.draws.iter().filter(|d| d.red > 0 ).map(|d| d.red).collect::<Vec<i32>>());
            // println!("Game: {:?}: {:?}", g.id, g.draws.iter().filter(|d| d.red > 0 ).map(|d| d.red).max().unwrap());
            Draw {
                red:g.draws.iter().filter(|d| d.red > 0 ).map(|d| d.red).max().unwrap(),
                green: g.draws.iter().filter(|d| d.green > 0 ).map(|d| d.green).max().unwrap(),
                blue: g.draws.iter().filter(|d| d.blue > 0 ).map(|d| d.blue).max().unwrap(),
            }
        }).collect();
        //sum the powers of the min number of cubes
        min_cubes.iter().map(|d| d.red * d.green * d.blue).sum()
}

#[test]
fn test_single_line() {
    let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";
    let parsed = parse_game(input).unwrap();
    assert_eq!(parsed.1.id, 1, "game id should be 1");
}
#[test]
fn test_input_1() {
    let input = include_str!("../input_test");
    let games = parse_input(input);
    assert_eq!(games.len(), 5, "should be 5 games");
}
#[test]
fn test_solve_1() {
    let input = include_str!("../input_test");
    let possible_games = solve(input);
    assert_eq!(possible_games, 8);
}

#[test]
fn test_solve_2() {
    let input = include_str!("../input_test");
    let res = solve_2(input);
    assert_eq!(res, 2286);
}
