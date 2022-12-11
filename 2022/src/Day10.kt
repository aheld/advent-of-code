import org.junit.Ignore
import org.junit.Test
import kotlin.test.assertEquals

class TestDay10 {
    private val testInput = readInput("Day10_test")
    private val input = readInput("Day10")

    data class Cmd(val cycles: Int, val x: Int)

    fun testPart1(input: List<String>): Int {
        var signals = loadSignals(input)
        val strengths = listOf<Int>(20, 60, 100, 140, 180, 220).map {
            val key = signals.keys.takeWhile { x -> x <= it }.max()
            it * signals[key]!!
        }
        return strengths.sum()
    }

    private fun loadSignals(input: List<String>): MutableMap<Int, Int> {
        val commands = input.map {
            if (it.startsWith("no")) Cmd(1, 0)
            else {
                Cmd(2, it.split(" ")[1].toInt())
            }
        }
        var cycle = 0
        var x = 1
        var signals = mutableMapOf<Int, Int>()
        for (cmd in commands) {
            signals[cycle + cmd.cycles + 1] = x + cmd.x
            cycle += cmd.cycles
            x += cmd.x
        }
        return signals
    }

    fun getSignal(signals: MutableMap<Int, Int>, cycle: Int): Int {
        val keys = signals.keys.takeWhile { it <= cycle }
        val k = when (keys.size) {
            0 -> 1
            else -> keys.max()
        }
        return signals[k]!!
    }

    fun testPart2(input: List<String>): String {
        var signals = loadSignals(input)
        signals[1] = 1
        var screen = StringBuilder()
        repeat(6) {
            val row = it
            for (cycle in 1..40) {
                val signal = getSignal(signals, cycle + row*40)

                if (signal in cycle - 2..cycle) screen.append("#")
                else screen.append(".")
//                println("$cycle --> $signal : $pixel ${(signal in cycle - 1..cycle + 1)}")
           }
            screen.append("\n")
        }
        return screen.toString()
    }

@Test
fun testPart1() {
    assertEquals(13140, testPart1(testInput))
    assertEquals(16020, testPart1(input))
//        println(testPart1(input))
}

@Test
fun testPart2() {
    val expected = """##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######....."""
    val result = testPart2(testInput)
    val resultsPairs = expected.split("\n").zip(result.split("\n"))
    println(resultsPairs)
    resultsPairs.mapIndexed { i, pair ->
        println("Row $i, ${pair.second.length}")
//            assertEquals(pair.first.length, pair.second.length)
        println(pair.first)
        println(pair.second)
        repeat(4) {
            for (i in 1..10) {
                if (i == 10) print(0)
                else print(i)
            }
        }
        println()
        assertEquals(pair.first, pair.second)
    }
        assertEquals(expected, testPart2(testInput).trimEnd())
    var solution = testPart2(input)
            .replace(".", " ")
    println(solution)
}
}