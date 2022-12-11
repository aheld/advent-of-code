import org.junit.Ignore
import org.junit.Test
import kotlin.test.assertEquals

class TestDay11 {
    private val testInput = readInput("Day11_test")
    private val input = readInput("Day11")

    class ThrowTest(val divisible: Long, val trueMonkey: Int, val falseMonkey: Int)
    class Monkey(
        private val number: Int,
        private val items: List<Long>,
        val operation: (Long) -> Long,
        val throwTest: ThrowTest
    ) {
        var inspected = 0

        private val itemQueue = ArrayDeque(items)

        fun catchItem(item: Long) {
            itemQueue.add(item)
        }

        fun inspect(divisor: Long = 3): List<Pair<Int, Long>> {

            val res = itemQueue.map {
                inspected += 1
                var new = operation(it)
                new %= divisor
                if (new.mod(throwTest.divisible) == 0.toLong()) {
//                    println("Monkey ${this.number}: $it -> $new  to: ${throwTest.trueMonkey} (T)")
                    Pair(throwTest.trueMonkey, new)
                } else {
//                    println("Monkey ${this.number}: $it -> $new to: ${throwTest.falseMonkey} (F)")
                    Pair(throwTest.falseMonkey, new)
                }
            }
            itemQueue.clear()
            return res
        }

        override fun toString(): String {
            return "Monkey ($number) inspected $inspected and has  ${itemQueue.joinToString()}"
        }
    }

    private fun loadMonkeys(input: List<String>): List<Monkey> {
        return input.chunked(7) { cmd ->
            val line = cmd[0]
            val number = line["Monkey ".length].toString().toInt()
            val startingItems = cmd[1].substring("  Starting items: ".length)
                .split(", ").map { it.toLong() }
            val operator = cmd[2]["  Operation: new = old ".length].toString()
            val x = cmd[2].substring("  Operation: new = old * ".length)


            val op: (Long) -> Long = if (operator == "+" && x == "old") { old: Long -> old + old }
            else if (operator == "+" && x != "old") { old: Long -> old + x.toLong() }
            else if (operator == "*" && x == "old") { old: Long -> old * old }
            else if (operator == "*" && x != "old") { old: Long -> old * x.toLong() }
            else {
                println("Error $operator, $x")
                throw Exception("Can't get here")
            }


            val throwTest = ThrowTest(
                cmd[3].substring("  Test: divisible by ".length).toLong(),
                cmd[4].substring("    If true: throw to monkey ".length).toInt(),
                cmd[5].substring("    If false: throw to monkey ".length).toInt()
            )
            val monkey = Monkey(number, startingItems, op, throwTest)
//            println(monkey)
            monkey
        }
    }

    private fun testPart1(input: List<String>): Int {
        val monkeys = loadMonkeys(input)
        repeat(20) {
            for (monkey in monkeys) {
                val throws = monkey.inspect()
                for (toss in throws) {
                    monkeys[toss.first].catchItem(toss.second)
                }
            }

        }
        println(monkeys)

        val sortMonkys = monkeys.map { it.inspected }.sorted()
//        println(sortMonkys)
        return sortMonkys[sortMonkys.lastIndex] * sortMonkys[sortMonkys.lastIndex - 1]
    }

    fun testPart2(input: List<String>): Long {
        val monkeys = loadMonkeys(input)
        val commonDivisor: Long = monkeys.map { it.throwTest.divisible }.reduce { a, b -> a * b }

        repeat(10000) {
//            println("Round $it")
            for (monkey in monkeys) {
                val throws = monkey.inspect(commonDivisor)
                for (toss in throws) {
                    monkeys[toss.first].catchItem(toss.second)
                }
            }

        }
        println(monkeys)

        val sortMonkys = monkeys.map { it.inspected }.sorted()
        println(sortMonkys)
        return sortMonkys[sortMonkys.lastIndex].toLong() * sortMonkys[sortMonkys.lastIndex - 1].toLong()
    }

    @Ignore
    @Test
    fun testPart1() {
        assertEquals(10605, testPart1(testInput))
        println(testPart1(input))
    }

    @Test
    fun testPart2() {
        assertEquals(expected = 2713310158L, actual = testPart2(testInput))
        println("res ${testPart2(testInput)}")
        println("res ${testPart2(input)}")
    }

}
