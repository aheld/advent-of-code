import org.junit.jupiter.api.DynamicTest.dynamicTest
import org.junit.jupiter.api.TestFactory
import kotlin.test.Test
import kotlin.test.assertEquals

internal class MainKtTest


internal class Day1Test {

    @Test
    fun testPart1() {
        val expected: Int = 24000
        assertEquals(part1("TestInput.txt"),expected)
    }

    @Test
    fun testPart2() {
        val expected: Int = 45000
        assertEquals(part2("TestInput.txt"),expected)
    }

    @TestFactory
    fun `part 1`() =
        listOf(
            "input.txt" to 66719,
            "TestInput.txt" to 24000,
        ).map { (filename, expected) ->
            dynamicTest(
                "given \"$filename\", " +
                        "when executing part1, " +
                        "then the total is $expected"
            ) {
                assertEquals(part1(filename),expected)
            }
        }

@TestFactory
fun `part 2`() =
    listOf(
        "input.txt" to 198551,
        "TestInput.txt" to 45000,

        ).map { (filename, expected) ->
        dynamicTest(
            "given \"$filename\", " +
                    "when executing part1, " +
                    "then the total is $expected"
        ) {
            assertEquals(part2(filename),expected)
        }
    }
}

