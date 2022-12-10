import org.junit.Test
import kotlin.math.pow
import kotlin.test.assertEquals

class TestDay09 {
    private val testInput = readInput("Day09_test") //1
    private val testInput2 = readInput("Day09_test2") // 36
    private val input = readInput("Day09")

    data class Move(val dir: String, val dist: Int)
    data class Point(var X: Int, var Y: Int)

    data class Board(val Knots: ArrayList<Point>, val visited: MutableList<Point>)

    fun solve(input: List<String>, numKnots: Int = 2): Int {
        val moves = input.map {
            val (dir, dist) = it.split(" ")
            Move(dir, dist.toInt())
        }
        val knots = MutableList<Point>(numKnots) { Point(0, 0) }
        val visited = mutableSetOf<Point>()
        val deltas = mapOf<String, Point>(
            "R" to Point(1, 0),
            "U" to Point(0, 1),
            "L" to Point(-1, 0),
            "D" to Point(0, -1)
        )

        for (move in moves) {
            val delta = deltas[move.dir]!!
//            println(move)
            repeat(move.dist) {
//                println("$move, $it")
//                println("Head: ${knots[0]}")
                for (i in 1..knots.lastIndex) {


                    if (i == 1) {
                        knots[i - 1] = Point(knots[i - 1].X + (delta.X), knots[i - 1].Y + (delta.Y))
                    }

                    var head = knots[i - 1]
                    var tail = knots[i]
                    var distance = (tail.X - head.X).toDouble().pow(2) +
                            (tail.Y - head.Y).toDouble().pow(2)
//                    println("++ ${knots[i - 1]}[$i-1] ${knots[i]}[$i]  -- $distance")

                    if (distance >= 3) {
                        knots[i] =
                            Point(
                                head.X + (tail.X - head.X) / 2,
                                head.Y + (tail.Y - head.Y) / 2
                            )
//                        println("++ MOVED ${knots[i - 1]}[$i-1] >${knots[i]}[$i]")
                    }
                }
                visited.add(knots[knots.lastIndex])
            }
//            println(visited)
        }
        return visited.size
    }

    @Test
    fun testPart1() {
        assertEquals(13, solve(testInput))
        println(solve(input))
    }


    @Test
    fun testPart2() {
        assertEquals(1, solve(testInput, 10))
        assertEquals(36, solve(testInput2, 10))
        println(solve(input, 10))
    }
}

//check out https://gist.github.com/demidko/e6fdafd02e66a20cd325006e7afc6f70