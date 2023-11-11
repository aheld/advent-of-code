fn main() {
    println!("Hello, world!");
    println!("{}", solve(include_str!("../input")));
    println!("{}", solve2(include_str!("../input")));
}

struct SantaString {
    string: String,
}

fn parse_input(input: &str) -> Vec<SantaString> {
    input
        .lines()
        .map(|line| SantaString {
            string: line.to_string(),
        })
        .collect()
}

impl SantaString {
    fn how_many_vowels(&self) -> usize {
        let vowels = ['a', 'e', 'i', 'o', 'u'];
        let mut count = 0;
        for c in self.string.chars() {
            if vowels.contains(&c) {
                count += 1;
            }
        }
        return count;
    }
    fn has_double_letter(&self) -> bool {
        let mut prev = '-';
        for c in self.string.chars() {
            if c == prev {
                return true;
            }
            prev = c;
        }
        return false;
    }
    fn has_naughty_strings(&self) -> bool {
        let naughty_strings = ["ab", "cd", "pq", "xy"];
        for s in naughty_strings.iter() {
            if self.string.contains(s) {
                return true;
            }
        }
        return false;
    }
    fn is_nice(&self) -> bool {
        return self.how_many_vowels() >= 3
            && self.has_double_letter()
            && !self.has_naughty_strings();
    }

    fn non_overlapping_pairs(&self) -> bool {
        for i in 0..self.string.len() - 1 {
            let pair = self.string[i..i + 2].to_string();
            for j in 0..self.string.len() - 1 {
                if j == i || j == i + 1 || (i > 0 && j == i - 1) {
                    continue;
                }
                if pair.eq(&self.string[j..j + 2]) {
                    return true;
                }
            }
        }
        return false;
    }

    fn has_repeating_letter_with_gap(&self) -> bool {
        let mut prev = '-';
        let mut prev_prev = '-';
        for (i, c) in self.string.chars().enumerate() {
            if i > 1 && c == prev_prev {
                return true;
            }
            prev_prev = prev;
            prev = c;
        }
        return false;
    }
}

fn solve(input: &str) -> usize {
    let strings = parse_input(input);
    let mut count = 0;
    for s in strings {
        // println!("String: {}", s.string);
        if s.is_nice() {
            count += 1;
        }
    }
    count
}

fn solve2(input: &str) -> usize {
    let strings = parse_input(input);
    let mut count = 0;
    for s in strings {
        // println!("String: {}", s.string);
        if s.non_overlapping_pairs() && s.has_repeating_letter_with_gap() {
            count += 1;
        }
    }
    count
}

#[test]
fn test_how_many_vowels() {
    assert_eq!(
        SantaString {
            string: "aeiou".to_string()
        }
        .how_many_vowels(),
        5
    );
    assert_eq!(
        SantaString {
            string: "yyyyy".to_string()
        }
        .how_many_vowels(),
        0
    );
    assert_eq!(
        SantaString {
            string: "yyeyyy".to_string()
        }
        .how_many_vowels(),
        1
    );
}

#[test]
fn test_suite() {
    let cases = &[
        ("ugknbfddgicrmopn", 1, "nice"),
        ("ugknbfddgicrmopn\nugknbfddgicrmopn\nasdq", 2, "multiline"),
        (
            "jchzalrnumimnmhp",
            0,
            "is naughty because it has no double letter.",
        ),
        (
            "haegwjzuvuyypxyu",
            0,
            "is naughty because it contains the string `xy`.",
        ),
        (
            "dvszwmarrgswjxmb",
            0,
            "is naughty because it contains only one vowel.",
        ),
    ];

    for (input, expected, desc) in cases {
        assert_eq!(solve(*input), *expected, "{}", *desc);
    }
}

#[test]
fn test_suite_2() {
    let cases = &[
        ("qjhvhtzxzqqjkmpb", 1, "nice"),
        ("xxyxx", 1, "nice"),
        ("uurcxstgmygtbstg", 0, "is naughty"),
        ("ieodomkazucvgmuy", 0, "is naughty"),
    ];

    for (input, expected, _desc) in cases {
        assert_eq!(solve2(*input), *expected, "{}", *input);
    }
}

#[test]
fn test_suite_ieodomkazucvgmuy() {
    assert_eq!(solve2("ieodomkazucvgmuy"), 0, "ieodomkazucvgmuy");
}

#[test]
fn test_non_overlapping_pair() {
    let cases = &[
        // ("aabcdefgaa", true),
        // ("xyxy", true),
        ("aaa", false),
        // ("uurcxstgmygtbstg", true),
        // ("ieodomkazucvgmuy", false),
    ];
    for (input, expected) in cases {
        assert_eq!(
            SantaString {
                string: input.to_string()
            }
            .non_overlapping_pairs(),
            *expected,
            "{}",
            *input
        );
    }
}
