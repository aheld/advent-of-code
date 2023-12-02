use criterion::{black_box, criterion_group, criterion_main, Criterion};

fn fibonacci(n: u64) -> u64 {
    match n {
        0 => 1,
        1 => 1,
        n => fibonacci(n-1) + fibonacci(n-2),
    }
}

#[path = "../src/main.rs"]
mod main;

fn criterion_benchmark(c: &mut Criterion) {
    let input = include_str!("../input");
    c.bench_function("aco day02", |b| b.iter(|| main::solve(black_box(input))));
    c.bench_function("aco day02", |b| b.iter(|| main::main_for_bench()));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);