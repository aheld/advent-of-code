use std::collections::HashMap;

use nom::{
    bytes::complete::{tag, take_until1},
    character::complete::line_ending,
    character::streaming::i64,
    multi::separated_list1,
    sequence::tuple,
    IResult,
};

fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

pub fn main_for_bench() {
    solve(include_str!("../input"));
    solve_2(include_str!("../input"));
}

#[derive(Debug)]
struct Almanac {
    seeds: Vec<i64>,
    mappings: HashMap<String, Vec<AlmanacMap>>,
}
#[derive(Debug)]
struct AlmanacMap {
    dest_start: i64,
    source_start: i64,
    range: i64,
}

// note - added an extra line ending to the end of the input file
fn parse_input(input: &str) -> IResult<&str, Almanac> {
    let (input, seeds) = tuple((tag("seeds: "), separated_list1(tag(" "), i64)))(input)?;

    let (input, _) = tuple((line_ending, line_ending))(input)?;
    let (_, blocks) =
        separated_list1(tuple((line_ending, line_ending)), take_until1("\n\n"))(input)?;

        let mut mappings = HashMap::new();
    for block in blocks {
        let (input, s) = take_until1(" map:\n")(block)?;
        let (input, _) = tuple((take_until1("\n"), tag("\n")))(input)?;
        let maps = input.split('\n').map(|m|{
            let inputs = m
                .split(' ')
                .map(|m| m.parse::<i64>().unwrap())
                .collect::<Vec<i64>>();
            AlmanacMap {
                dest_start: inputs[0],
                source_start: inputs[1],
                range: inputs[2],
            }
        }).collect();
        mappings.insert(s.to_string(), maps);
    }
    let a = Almanac {
        seeds: seeds.1,
        mappings
    };
    Ok((input, a))
}

fn get_input(input: &str) -> Almanac{
    let a = parse_input(input);
    match a {
        Ok((_, a)) => {
            a
        }
        Err(e) => {
            println!("Error: {:?}", e);
            panic!("Error - parse input");
        }
    }

}

fn find_mapped_location(seed: i64, mappings: &Vec<AlmanacMap>) -> i64 {
    for map in mappings {
        if seed >= map.source_start && seed < map.source_start + map.range {
            return map.dest_start + seed - map.source_start;
        }
    }
    seed
}

pub fn solve(input: &str) -> i64 {
    let a = get_input(input);
    let mut lowest  = i64::MAX;
    let map_list = "seed-to-soil soil-to-fertilizer fertilizer-to-water water-to-light light-to-temperature temperature-to-humidity humidity-to-location".split(' ').collect::<Vec<&str>>();
    for seed in a.seeds {
        let mut location = seed;
        for map in &map_list {
            location = find_mapped_location(location, &a.mappings[*map]);
            println!("Seed {} maps to location {} for {}", seed, location, *map);
        }
        lowest = lowest.min(location);
    }
    lowest
}

pub fn solve_2(input: &str) -> i64 {
    let a = get_input(input);
    let mut lowest  = i64::MAX;
    let map_list = "seed-to-soil soil-to-fertilizer fertilizer-to-water water-to-light light-to-temperature temperature-to-humidity humidity-to-location".split(' ').collect::<Vec<&str>>();
    for seedlist in a.seeds.chunks(2) {
        println!("Seedlist {:?}", seedlist );
        for seed in seedlist[0]..(seedlist[0]+seedlist[1]) {
            let mut location = seed;
            for map in &map_list {
                location = find_mapped_location(location, &a.mappings[*map]);
                // println!("Seed {} maps to location {} for {}", seed, location, *map);
            }
        lowest = lowest.min(location);
        }
    }
    lowest
}

#[test]
fn test_input_1() {
    let input = include_str!("../input_test");
    let res = parse_input(input);
    match res {
        Ok((_, a)) => {
            // dbg!(&a);
            assert_eq!(a.seeds, vec![79, 14, 55, 13]);
            assert_eq!(a.mappings["seed-to-soil"][1].range, 48);
            assert_eq!(a.mappings["seed-to-soil"][0].source_start, 98, "Source start");
            assert_eq!(a.mappings["seed-to-soil"][1].dest_start, 52);
            assert_eq!(a.mappings["seed-to-soil"][0].dest_start, 50);
            assert_eq!(a.mappings["humidity-to-location"][1].dest_start, 56);
        }
        Err(e) => {
            println!("Error: {:?}", e);
            assert!(false);
        }
    }
}
#[test]
fn test_solve_1() {
    let input = include_str!("../input_test");
    let location = solve(input);
    assert_eq!(location, 35);
}

#[test]
fn test_solve_2() {
    let input = include_str!("../input_test");
    let res = solve_2(input);
    assert_eq!(res, 46);
}
