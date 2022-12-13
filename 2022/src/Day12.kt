import org.junit.Test
import kotlin.test.assertEquals

class TestDay12 {
    private val testInput = readInput("Day12_text")
    private val input = readInput("Day12")
    private val forrestTest = loadInput(testInput)
    private val forrest = loadInput(input)

    private fun testPart1(forrest: Forrest): Int {
        //        val forrest = loadInput(input)
        fun getPoints(c: Char): List<Point> = forrest.indices.flatMap { x ->
            forrest[x].mapIndexed { y, elevation -> Pair(Point(x, y), elevation) }
        }.filter { p -> p.second == c }.map { it.first }

        val initialPoint: Point = getPoints('S').first()

        forrest[initialPoint.x][initialPoint.y] = 'a'

        val endPoint: Point = forrest.indices.flatMap { x ->
            forrest[x].mapIndexed { y, elevation -> Pair(Point(x, y), elevation) }
        }.first { p -> p.second == 'E' }.first
        forrest[endPoint.x][endPoint.y] = 'z'

        return getShortestPath(initialPoint, forrest, endPoint)
    }

    private fun getShortestPath(initialPoint: Point, forrest: Forrest, endPoint: Point): Int {
        val costs = mutableMapOf(initialPoint to 0)
        val toVisit = mutableSetOf(initialPoint)

        while (toVisit.isNotEmpty()) {
            val cell = toVisit.first()
            toVisit.remove(cell)
            val cost = costs[cell]!!

            fun validatePoint(p: Point): Point? {
                return if (p.x < 0 || p.y < 0 || p.x >= forrest.size || p.y >= forrest[0].size) null
                else p
            }

            fun Point.getElevation() = forrest[x][y].code

            val adjacentCells = listOfNotNull(
                validatePoint(Point(cell.x + 1, cell.y)),
                validatePoint(Point(cell.x - 1, cell.y)),
                validatePoint(Point(cell.x, cell.y + 1)),
                validatePoint(Point(cell.x, cell.y - 1)),
            ).filter {
                //                println(cell)
                //                println(cell.getElevation())
                //                println(it)
                //                println(it.getElevation() - 1)
                (cell.getElevation() >= it.getElevation() - 1)
            }
            adjacentCells.forEach { neighbor ->

                val neighborCost = costs.getOrElse(neighbor) { Int.MAX_VALUE }
                if (neighborCost >= cost + 1) {
                    costs[neighbor] = cost + 1
                    toVisit += (neighbor)
                }
            }
        }
        //        println(costs)

        return costs[endPoint] ?: 0
    }

    private fun part2(forrest: Forrest): Int {

        fun getPoints(c: Char): List<Point> = forrest.indices.flatMap { x ->
            forrest[x].mapIndexed { y, elevation -> Pair(Point(x, y), elevation) }
        }.filter { p -> p.second == c }.map { it.first }

        val initialPoints: List<Point> = getPoints('a')

        val endPoint: Point = getPoints('E').first()
        forrest[endPoint.x][endPoint.y] = 'z'

        val res = initialPoints.map { s ->
            getShortestPath(s, forrest, endPoint)
        }
//        println(res)
        return res.filter { (it > 0) }.min()
    }

    private fun loadInput(input: List<String>): Forrest {
        val data = input.map {
            it.toCharArray()
        }
        //transpose so its forrest[X],[Y]
        val rows = data[0].size
        val cols = data.size
        return Array(cols) { j ->
            Array(rows) { i ->
                data[j][i]
            }
        }
    }

    @Test
    fun testPart1() {
        assertEquals(31, testPart1(forrestTest))
        val part1 = testPart1(forrest)
        assertEquals(447, part1)
        println(part1)
    }

    @Test
    fun testPart2() {
        assertEquals(29, part2(forrestTest))
        val part2 = part2(forrest)
        assertEquals(446, part2)
        println(part2)
    }
}

data class Point(val x: Int, val y: Int)

typealias Forrest = Array<Array<Char>>