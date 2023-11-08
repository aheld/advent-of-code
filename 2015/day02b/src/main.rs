fn main() {
    let gifts = parse_input(include_str!("../input"));
    let areas_per_gift = gifts.iter().map(|gift| area_needed(gift));
    println!("Area Needed {}", areas_per_gift.sum::<u32>());
}

#[derive(PartialEq, Debug)]
struct Gift {
    l: u32,
    w: u32,
    h: u32,
}

fn area_needed(gift: &Gift) -> u32 {
    let mut s = vec![gift.w, gift.l, gift.h];
    s.sort();
    let ribbon = s[0] * 2 + s[1] * 2;
    let bow = gift.l * gift.w * gift.h;
    ribbon + bow
}

fn parse_input(input: &str) -> Vec<Gift> {
    input
        .lines()
        .map(|line| {
            let i = line
                .split("x")
                .map(|x| x.parse::<u32>().unwrap())
                .collect::<Vec<u32>>();
            Gift {
                l: i[0],
                w: i[1],
                h: i[2],
            }
        })
        .collect::<Vec<Gift>>()
}

#[test]
fn test_area() {
    let input = "2x3x4\n1x1x10";
    let gifts = parse_input(input);
    assert_eq!(gifts[0], Gift { l: 2, w: 3, h: 4 });
    assert_eq!(gifts[1], Gift { l: 1, w: 1, h: 10 });
    assert_eq!(area_needed(&gifts[0]), 34);
    assert_eq!(area_needed(&gifts[1]), 14);
}
