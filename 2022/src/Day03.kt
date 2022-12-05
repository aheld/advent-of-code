import org.junit.Test
import kotlin.test.assertEquals


fun part1(input: List<String>): Int {
    return input.map { line: String ->
        line.substring(0..line.length / 2 - 1) to line.substring((line.length / 2))
    }
        .also { println(it) }
        .flatMap { (first, second) -> first.toSet() intersect second.toSet() }
        .also { println(it) }
        .sumOf { it -> charToScore(it) }
}

fun part2(input: List<String>): Int {
    val total = input.chunked(3)
        .also { println(it) }
        .map { group ->
            group.zipWithNext()
         .also {println(it)}
                .map { (first, second) -> first.toCharArray() intersect second.toSet() }
        }
        .also { println(it) }
        .flatMap { (first, second) -> first intersect second }
        .also { println(it) }
        .sumOf { it -> charToScore(it) }

    return total
}

fun charToScore(letter: Char) =
    if (letter.isUpperCase()) {
        letter - 'A' + 27
    } else {
        letter - 'a' + 1
    }

class TestDay3 {
    val testInput = readInput("Day03_test")
    val input = readInput("Day03")

    @Test
    fun testPart1() {
        assertEquals(157, part1(testInput))
        println(part1(input))
    }

    @Test
    fun testPart2() {
        assertEquals(70, part2(testInput))
        println(part2(input))
    }
}

