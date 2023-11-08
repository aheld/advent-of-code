pub fn main() {
 let moves: Vec<i32> = include_str!("../input")
        .chars()
        .map(|c| -> i32 {
            match c {
                '(' => 1,
                ')' => -1,
                _ => 0,
            }
        }).collect();
    print!("{}", moves.iter().sum::<i32>());
}
