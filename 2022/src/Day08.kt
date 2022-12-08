import org.junit.Ignore
import org.junit.Test
import kotlin.test.assertEquals

class TestDay08 {
    private val testInput = readInput("Day08_test")
    private val input = readInput("Day08")

    data class Tree(val height: Int, var isVisible: Boolean)

    fun testPart1(input: List<String>): Int {
        val forrest = loadInput(input)

        // From top
        for (colNum in 0..forrest.lastIndex) {
            var sightHeight = -1
            for (rowNum in 0..forrest[0].lastIndex) {
                val tree = forrest[colNum][rowNum]
                if (tree.height > sightHeight) {
                    tree.isVisible = true
                    sightHeight = tree.height
                }
            }
        }

        // From bottom
        for (colNum in 0..forrest.lastIndex) {
            var sightHeight = -1
            for (rowNum in forrest[0].lastIndex downTo 0) {
                val tree = forrest[colNum][rowNum]
                if (tree.height > sightHeight) {
                    tree.isVisible = true
                    sightHeight = tree.height
                }
            }
        }

        // From left
        for (rowNum in 0..forrest[0].lastIndex) {
            var sightHeight = -1
            for (colNum in 0..forrest.lastIndex) {
                val tree = forrest[colNum][rowNum]
                if (tree.height > sightHeight) {
                    tree.isVisible = true
                    sightHeight = tree.height
                }
            }
        }

        // From right
        for (rowNum in 0..forrest[0].lastIndex) {
            var sightHeight = -1
            for (colNum in forrest.lastIndex downTo 0) {
                val tree = forrest[colNum][rowNum]
                if (tree.height > sightHeight) {
                    tree.isVisible = true
                    sightHeight = tree.height
                }
            }
        }

        var total = 0
        for (rowNum in 0..forrest[0].lastIndex) {
            for (colNum in 0..forrest.lastIndex) {
                if (forrest[colNum][rowNum].isVisible) {
                    print('V')
                    total += 1
                } else {
                    print('.')
                }
            }
            println()
        }

        return total

    }

    private fun loadInput(input: List<String>): Array<Array<Tree>> {
        val data = input.map {
            it.toCharArray()
                .map { Tree(it.toString().toInt(), false) }
        }

        //        val forrest: ArrayList<ArrayList<Tree>> = ArrayList()
        //transpose so its forrest[X],[Y]
        val cols = data[0].size
        val rows = data.size
        val forrest = Array(cols) { j ->
            Array(rows) { i ->
                data[i][j]
            }
        }
        return forrest
    }

    fun scenicScore(forrest: Array<Array<Tree>>, treeX: Int, treeY: Int): Int {
        data class Delta(val X: Int, val Y: Int)

        val views = listOf<Delta>(Delta(1, 0), Delta(-1, 0), Delta(0, 1), Delta(0, -1))

        var tree = forrest[treeX][treeY]
//        println("($treeX, $treeY): ${tree.height}")
        val viewScores = views.map {
            val view = it
            var viewedTrees = 0
            var x = treeX
            var y = treeY
            var tree = forrest[x][y]


            var viewHeight = tree.height
            while ((x in 0..forrest.lastIndex) && (y in 0..forrest[0].lastIndex)) {
                x += view.X
                y += view.Y
                var treeToSee: Tree
                try {
                    treeToSee = forrest[x][y]
                } catch (ex: ArrayIndexOutOfBoundsException) {
                    break
                }
                viewedTrees += 1
                if (treeToSee.height >= viewHeight) {
                    break
                }
            }
//            println("$view, $viewedTrees")
            viewedTrees
        }
        return viewScores[0] * viewScores[1] * viewScores[2] * viewScores[3]
    }

    fun testPart2(input: List<String>): Int {
        val forrest = loadInput(input)
        var bestScore = 0
        for (x in 0..forrest.lastIndex) {
            for (y in 0..forrest[0].lastIndex) {
                val score = scenicScore(forrest, x, y)
                if (score > bestScore) {
                    bestScore = score
//                    println("$x, $y : $bestScore")
                }
            }
        }

        return bestScore
    }
    
    @Test
    fun testPart1() {
        assertEquals(21, testPart1(testInput))
        assertEquals(1679, testPart1(input))
    }


    @Test
    fun testPart2() {
        assertEquals(8, testPart2(testInput))
        assertEquals(536625, testPart2(input))
    }
}