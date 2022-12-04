import org.junit.jupiter.api.DynamicTest.dynamicTest
import org.junit.jupiter.api.TestFactory
import kotlin.test.*

internal class MainKtTest


internal class Day1Test {

data class TestCase(val fileName: String, val part1: Int, val part2: Int)
val testCases: List<TestCase> = listOf(
    TestCase("Day01input.txt",  66719,198551),
    TestCase("Day01testinput.txt", 24000, 45000)
    )
        @TestFactory
fun `part 2`() = testCases.map { (filename:String, part1:Int, part2:Int) ->
        dynamicTest(
            filename
        ) {
            assertEquals(part1(filename),part1)
            assertEquals(part2(filename),part2)
        }
    }
}

