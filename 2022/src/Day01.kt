import java.io.File

fun main(args: Array<String>) {
    println(part1("Day01input.txt"))
    println(part2("Day01input.txt"))
}

fun loadInputAndTotal(filename: String):ArrayList<Int> {
    val inputs = File("src", filename).readLines() //.map { it.toInt() }

    var elfList = ArrayList<Int>()
    var calorieTotal = 0
    for (line in inputs) {
        if (line.isEmpty()) {
            elfList.add(calorieTotal)
            calorieTotal = 0
            continue
        }
        calorieTotal += line.toInt()
    }
    elfList.add(calorieTotal)
    return elfList
}

fun part1(filename:String): Int  {
    var elfList = loadInputAndTotal(filename)
    return (elfList.max())
}

fun part2(filename:String): Int  {
    var elfList = loadInputAndTotal(filename)
    elfList.sort()
    var total = 0
    for (i in elfList.lastIndex downTo (elfList.lastIndex-2)){
        total += elfList[i]
    }
    return (total)
}