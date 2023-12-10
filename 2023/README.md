# Trying Rust this year

note: instructions for Macos/Linux, should be similar commands for windows
- Install [rustup](https://rustup.rs/)
 - `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`
- Install [cargo-binstall](https://github.com/cargo-bins/cargo-binstall)
 -  `curl -L --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/cargo-bins/cargo-binstall/main/install-from-binstall-release.sh | bash`
- install [aoc-cli](https://github.com/scarvalhojr/aoc-cli) 
  - `brew install scarvalhojr/tap/aoc-cli`
- Copy your AOC session cookie to a local var (see aoc-cli README)
    - Copy the session cookie value
    - `pbpaste > ~/.adventofcode.session`s
    

## Download a new day

    make newday day=01

    # run tests until you get the test input working
    cd day01
    make watchtest

    # run the main
    make run

    # submit the answer
    make submit part=1 answer=1234"

    # check your progress
    make calendar