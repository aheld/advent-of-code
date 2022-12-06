import org.junit.Test
import kotlin.test.assertEquals

class TestDay00 {
    private val testInput = readInput("Day04_test")
    private val input = readInput("Day04")
    fun testPart1(input: List<String>): Int{
        return -1
    }
    fun testPart2(input: List<String>): Int{
        return -1
    }

    @Test
    fun testPart1(){
        assertEquals(2, testPart1(testInput))
        println(testPart1(input))
    }

    @Test
    fun testPart2(){
        assertEquals(4, testPart2(testInput))
        println(testPart2(input))
    }
}