import org.junit.Test
import kotlin.test.assertEquals

data class Move(val qty: Int, val source: Int, val target: Int)

class TestDay05 {
    private val testInput = readInput("Day05_test")
    private val input = readInput("Day05")

    fun numStacks(input: List<String>): Int = input
        .dropWhile { it.contains("]") }
        .first()
        .split(" ")
        .filter { it.isNotBlank() }
        .last().toInt()

    fun getStacks(input: List<String>): List<ArrayDeque<Char>> {
        val stacks = List(numStacks(input)) { ArrayDeque<Char>()}

        val rows = input.filter { it.contains("[") }
            .map { it.substring(1).toCharArray() }


        for (row in rows ) {
            for (i in 0..row.size step 4) {
                val box = row[i]
                if (box.isLetter()){
                    stacks[i/4].addLast(box)
                }
            }
        }
        return stacks
    }

    fun getMoves(input: List<String>): List<Move> {
        return input.filter { it.contains("move") }
            .map {it.split(" ") }
            .map { l -> l.map{ it.toIntOrNull()}.filterNotNull() }
            .map { Move(it[0], it[1]-1, it[2]-1)}
    }

    fun testPart1(input: List<String>): String{
        val stacks = getStacks(input)
        val moves = getMoves(input)

        moves.forEach { step ->
            repeat(step.qty) { stacks[step.target].addFirst(stacks[step.source].removeFirst()) }
        }
        return stacks.map { it.first()}.joinToString("")
    }

    fun testPart2(input: List<String>): String {
        val stacks = getStacks(input)
        val moves = getMoves(input)

        moves.forEach { step ->
            stacks[step.source].subList(0,step.qty).asReversed()
                .also { println(it)}
                .map { stacks[step.target].addFirst(it)}
                .map {stacks[step.source].removeFirst()}
        }
        return stacks.map { it.first()}.joinToString("")
    }

    @Test
    fun testPart1(){
        assertEquals("CMZ", testPart1(testInput))
        println(testPart1(input))
    }

    @Test
    fun testPart2(){
        assertEquals("MCD", testPart2(testInput))
        println(testPart2(input))
    }
}