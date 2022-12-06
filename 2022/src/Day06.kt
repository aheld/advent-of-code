import org.junit.Test
import org.junit.jupiter.api.DynamicTest
import org.junit.jupiter.api.TestFactory
import kotlin.test.assertEquals

class TestDay06 {
    private val input = readInput("Day06")[0]

    fun process(input: String, lenToFind: Int): Int {
        val iarr = input.toCharArray()
        for (i in lenToFind-1..iarr.size - 1) {
            val subarray = iarr.copyOfRange(i - (lenToFind-1), i+1)
            if (subarray.distinct().size === lenToFind) return i + 1
        }
        return -1
    }

    fun part1(input: String): Int {
        return process(input, 4)
    }

    fun part2(input: String): Int {
        return process(input, 14)
    }

    data class TestCase(val signal: String, val expected: Int, val expected2: Int)

    val testCases: List<TestCase> = listOf(
        TestCase("bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23),
        TestCase("nppdvjthqldpwncqszvftbrmjlhg", 6,23),
        TestCase("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10,29),
        TestCase("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26)
    )
    @TestFactory
    fun `part 1 and 2`() = testCases.map { (signal: String, expected: Int, expected2: Int) ->
        DynamicTest.dynamicTest(
            signal
        ) {
            assertEquals(expected, part1(signal))
            assertEquals(expected2, part2(signal))
        }
    }

    @Test
    fun testPart1() {
        println("Part 1 ${part1(input)}")
    }
    @Test
    fun testPart2() {
        println("Part 2 ${part2(input)}")
    }

}