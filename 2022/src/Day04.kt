import org.junit.Test
import kotlin.test.assertEquals

class TestDay04 {
    private val testInput = readInput("Day04_test")
    private val input = readInput("Day04")
    fun testPart1(input: List<String>): Int{
       return  input.map {
            line -> line.split(',')
        }
//            .also{println(it)}
            .map { pair -> {
                val (lower, upper) = pair[0].split("-").map { it.toInt()}
                val (lower2, upper2) = pair[1].split("-").map { it.toInt()}
                if ((lower <= lower2 && upper >= upper2 ) || (lower2 <= lower && upper2 >=upper)) {
                     1
                } else {
                    0
                }
                //overlap
            }}.map { it.invoke()}
//            .also { println(it)}
            .sum()

    }
    fun testPart2(input: List<String>): Int{
        return  input.map {
                line -> line.split(',')
        }
//            .also{println(it)}
            .map { pair -> {
                val (lower, upper) = pair[0].split("-").map { it.toInt()}
                val (lower2, upper2) = pair[1].split("-").map { it.toInt()}
                if ((upper in lower2..upper2)
                    || (lower in lower2..upper2)
                    || (upper2 in lower .. upper)
                    || (lower2 in lower .. upper)) {
                    1
                } else {
                    0
                }
                //overlap
            }}.map { it.invoke()}
//            .also { println(it)}
            .sum()

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