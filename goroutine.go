package main

func sum(values [] int, resultChan chan int) {
    println("start sum:", resultChan)
    sum := 0
    for _,value := range(values) {
        sum += value
    }
    resultChan <- sum
    println("end sum : ", resultChan)
}
func main() {
    values := [] int {1,2,3,4,5,6,7,8,9,10}

    resultChan := make(chan int, 1)
    go sum(values[:len(values)/2], resultChan)
    go sum(values[len(values)/2:], resultChan)
    sum1 := <- resultChan
    sum2 := <- resultChan


    //println("Result:", sum1)
    println("Result:", sum1, sum2, sum1+sum2)
}
