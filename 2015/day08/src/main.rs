use std::str::Chars;

fn main() {
    println!("Part 1: {}", solve(include_str!("../input")));
    println!("Part 2: {}", solve_2(include_str!("../input")));
}

#[derive(PartialEq, Debug)]
struct Counts {
    code: usize,
    chars: usize,
    encoded: usize,
}

fn parse_input(input: &str) -> Vec<Counts> {
    input
        .lines()
        .map(|l| {
            println!("{:?}", l);
            let chars = l.chars();
            println!("{:?}", chars);
            let count  = Counts {
                code: chars.clone().count(),
                chars: get_character_len(chars),
                encoded: get_encoded_len(l.chars()),
            };
            println!("{:?}: {:?}",l, count);
            count
        })
        .collect()
}

fn get_character_len(chars_input: Chars) -> usize {
    let mut chars: Vec<char> = chars_input.collect();
    let code_len = chars.len();
    if code_len == 0 {
        return 0;
    }
    let mut char_len = 0;

    chars = chars[1..code_len - 1].to_vec();
    while !chars.is_empty() {
        char_len += 1;
        let (ch, rest) = chars.split_first().unwrap();
        if let '\\' = ch {
            match rest.split_first().unwrap() {
                ('"' | '\\', _) => {
                    chars = rest[1..].to_vec();
                }
                ('x', _) => {
                    chars = rest[3..].to_vec();
                }
                _ => panic!("Should not happen")
            }
        } else {
            chars = rest.to_vec();
        }
    }
    return char_len;
}

fn get_encoded_len(chars_input: Chars) -> usize {
    let mut chars: Vec<char> = chars_input.collect();
    let code_len = chars.len();
    
    code_len
            + 2
            + chars
                .iter()
                .filter(|&ch| ch.eq(&'"') || ch.eq(&'\\'))
                .count()
}

fn solve(input: &str) -> usize {
    let counts = parse_input(input);
    let mut total_chars = 0;
    let mut total_code = 0;
    counts.into_iter().for_each(|c| {
        total_chars += c.chars;
        total_code += c.code;
    });
    return total_code - total_chars;
}

fn solve_2(input: &str) -> usize {
    let counts = parse_input(input);
    let mut total_encode = 0;
    let mut total_code = 0;
    counts.into_iter().for_each(|c| {
        total_encode += c.encoded;
        total_code += c.code;
    });
    return total_encode - total_code;
}

#[test]
fn test_suite() {
    let input = r#"""
"abc"
"aaa\"aaa"
"\x27"
"\"njro\\x68qgbx\\xe4af\\\"\\\\suan\""
"#;

    let counts = parse_input(input);
    println!("{:?}", counts);
    assert_eq!(counts[0].code, 2);
    assert_eq!(counts[0].chars, 0);

    assert_eq!(counts[1].code, 5);
    assert_eq!(counts[1].chars, 3);

    assert_eq!(counts[2].code, 10);
    assert_eq!(counts[2].chars, 7);

    assert_eq!(counts[3].code, 6);
    assert_eq!(counts[3].chars, 1);

    assert_eq!(counts[4].chars, 28);
    assert_eq!(counts[4].code, 38);

}

#[test]
fn test_suite_2() {
    let input = r#"""
"abc"
"aaa\"aaa"
"\x27"
"#;

    let counts = parse_input(input);
    println!("{:?}", counts);
    assert_eq!(counts[0].code, 2);
    assert_eq!(counts[0].encoded, 6);

    assert_eq!(counts[1].code, 5);
    assert_eq!(counts[1].encoded, 9);

    assert_eq!(counts[2].code, 10);
    assert_eq!(counts[2].encoded, 16);

    assert_eq!(counts[3].code, 6);
    assert_eq!(counts[3].encoded, 11);

    assert_eq!(counts[4].chars, 28);
    assert_eq!(counts[4].code, 38);

}
