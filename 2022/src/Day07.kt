import org.junit.Ignore
import org.junit.Test
import java.util.*
import kotlin.collections.ArrayDeque
import kotlin.test.assertEquals

class TestDay07 {
    private val testInput = readInput("Day07_test")
    private val input = readInput("Day07")

    data class File(val name: String, val size: Int)
    class Directory(
        name: String,
        parent: Directory?,
    ) {
        val name = name
        val parent = parent

        val subDirectories = mutableSetOf<Directory>()
        val files = mutableSetOf<File>()

        fun addFile(filename: String, size: Int) {
            files.add(File(filename, size))
        }

        fun addSubdirectory(dirname: String) {
            if (subDirectories.none { it.name === dirname }) {
                subDirectories.add(Directory(dirname, this))
            }
        }

        fun getParentOrRoot(): Directory {
            if (parent == null) return this
            return parent

        }

        fun getSize(): Int {
            return files.sumOf { it.size }  + subDirectories.sumOf { it.getSize() }
        }

        override fun toString(): String {
            return "$name(${getSize()})"
        }

        fun prettyString(): String {
            val stringBuilder = StringBuilder()
            print(stringBuilder, "", "")
            return stringBuilder.toString()
        }

        // https://github.com/AdrianKuta/Tree-Data-Structure/blob/master/treedatastructure/src/main/java/com/github/adriankuta/datastructure/tree/TreeNode.kt#L105
        private fun print(stringBuilder: StringBuilder, prefix: String, childrenPrefix: String) {
            stringBuilder.append(prefix)
            stringBuilder.append("$name <dir> ${getSize()}")
            stringBuilder.append('\n')
            val childIterator = subDirectories.iterator()
            while (childIterator.hasNext()) {
                val node = childIterator.next()
                if (childIterator.hasNext()) {
                    node.print(stringBuilder, "$childrenPrefix├── ", "$childrenPrefix│   ")
                } else {
                    node.print(stringBuilder, "$childrenPrefix└── ", "$childrenPrefix    ")
                }
            }
            val fileIterator = files.iterator()
            while (fileIterator.hasNext()) {
                val file = fileIterator.next()
                stringBuilder.append("├──  ")
                stringBuilder.append(file.name)
                stringBuilder.append("  ")
                stringBuilder.append(file.size)
                stringBuilder.append('\n')
            }
        }

        fun iterator(): Iterator<Directory> = PreOrderTreeIterator(this)

        //https://github.com/AdrianKuta/Tree-Data-Structure/blob/master/treedatastructure/src/main/java/com/github/adriankuta/datastructure/tree/PreOrderTreeIterator.kt
        class PreOrderTreeIterator(root: Directory) : Iterator<Directory> {
            private val stack = ArrayDeque<Directory>()
            init {
                stack.addFirst(root)
            }
            override fun hasNext(): Boolean = !stack.isEmpty()
            override fun next(): Directory {
                val node = stack.removeFirst()
                node.subDirectories
                    .forEach { stack.add(it) }
                return node
            }
        }
    }

    fun loadDirectories(input: List<String>): Directory {
        val fileSystem = Directory("/", null)
        var currentDir = fileSystem
        for (line in input) {
            if (line == "$ cd /") {
                currentDir = fileSystem
                continue
            }
            if (line == "$ ls") continue
            if (line[0].isDigit()) {
                val (size, filename) = line.split(" ")
                currentDir.addFile(filename, size.toInt())
                continue
            }
            if (line.startsWith("dir")) {
                val (_, dirName) = line.split(" ")
                currentDir.addSubdirectory(dirName)
                continue
            }
            if (line == "$ cd ..") {
                currentDir = currentDir.getParentOrRoot()
                continue
            }
            if (line.startsWith("$ cd")) {
                val (_, _, dirName) = line.split(" ")
                currentDir = currentDir.subDirectories.filter { it.name == dirName }.first()
                continue
            }
            println("No match for $line")
        }
        return fileSystem
    }

    fun testPart1(input: List<String>): Int{
        val fileSystem = loadDirectories(input)
        println(fileSystem.prettyString())

        var total = 0
        for ( d in fileSystem.iterator()) {
            if (d.getSize() < 100000) {
                total += d.getSize()
            }
        }
        return total
    }

    fun testPart2(input: List<String>): Int {
        val DiskTotal   = 70000000
        val NeededSpace = 30000000

        val fileSystem = loadDirectories(input)
        val CurrentFreeSpace = DiskTotal - fileSystem.getSize()
        val NeedToFind = NeededSpace - CurrentFreeSpace

        println("Need $NeedToFind")

        val directories = fileSystem.iterator()
            .asSequence()
            .sortedBy { it.getSize() }
            .toList()

        val suffix = if (directories.size > 5) { "..." } else { "]" }
        println("Sorted Dirs [${directories.take(5).joinToString()} $suffix")

        return directories.filter { it.getSize() >= NeedToFind }.first().getSize()
    }

    @Test
    fun testPart1() {
        assertEquals(95437, testPart1(testInput))
        println("Part 1: ${testPart1(input)}")
    }

    @Test
    fun testPart2() {
        assertEquals(24933642, testPart2(testInput))
        println("Part 2: ${testPart2(input)}")
    }
}