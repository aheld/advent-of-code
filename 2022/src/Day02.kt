fun main() {
    operator fun String.component1() = this[0]
    operator fun String.component2() = this[1]
    operator fun String.component3() = this[2]


    /*1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).*/
    fun playScore(play: Char): Int {
        val score = when (play) {
            'X' -> 1
            'Y' -> 2
            'Z' -> 3
            else -> throw Exception("bad myPlay score")
        }
        return score
    }

    fun roundScore(myPlay: Char, oppPlay: Char): Int {
//        A for Rock, B for Paper, and C for Scissors.
//        X for Rock, Y for Paper, and Z for Scissors.

        return when ("$oppPlay-$myPlay") {
            "A-Z", "B-X", "C-Y" -> 0
            "A-X", "B-Y", "C-Z" -> 3
            "A-Y", "B-Z", "C-X" -> 6
            else -> {
                println("$oppPlay-$myPlay")
                throw Exception("Can't get here")
            }
        }
    }

    fun part1(input: List<String>): Int {
        return input.map {
            val (oppPlay, _, myPlay) = it
            playScore(myPlay) + roundScore(myPlay, oppPlay)
        }.sum()
    }

    //X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
    //A for Rock, B for Paper, and C for Scissors.
    // send back the letter of the play based on part 1
    // X for Rock, Y for Paper, and Z for Scissors.
    fun getMyPlay(oppPlay: Char, outcome: Char): Char {
        return when ("$outcome-$oppPlay") {
            "X-A" -> 'Z'
            "X-B" -> 'X'
            "X-C" -> 'Y'
            "Y-A" -> 'X'
            "Y-B" -> 'Y'
            "Y-C" -> 'Z'
            "Z-A" -> 'Y'
            "Z-B" -> 'Z'
            "Z-C" -> 'X'
            else -> throw Exception("should not get here")
        }
    }

    fun part2(input: List<String>): Int {
        return input.map {
            val (oppPlay, _, outcome) = it
            val myPlay = getMyPlay(oppPlay, outcome)
            playScore(myPlay) + roundScore(myPlay, oppPlay)
        }.sum()
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day02_test")
    check(part1(testInput) == 15)
    check(part2(testInput) == 12)

    val input = readInput("Day02")
    println(part1(input))
    println("Part 2 ${part2(input)}")
}