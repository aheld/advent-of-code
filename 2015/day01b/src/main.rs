pub fn main() {
    let moves = include_str!("../input").chars().map(|c| -> i32 {
        match c {
            '(' => 1,
            ')' => -1,
            _ => 0,
        }
    });
    let mut floor = 0;
    for (i, m) in moves.enumerate() {
        // println!("Item {} = {}", i, m);
        floor += m;
        println!("Item {} = {} {}", i + 1, m, floor);
        if floor < 0 {
            break;
        };
    }
}
