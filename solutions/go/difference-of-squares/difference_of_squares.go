package differenceofsquares

func SquareOfSum(n int) int {
    acc := 0
    
    for i := range n + 1 {
        acc = acc + i
    } 

    return (acc * acc)
}

func SumOfSquares(n int) int {
    acc := 0
    
    for i := range n + 1 {
        acc = acc + (i * i)
    } 

    return acc
}

func Difference(n int) int {
    tmp := SumOfSquares(n) - SquareOfSum(n)
    if (tmp < 0) {
        return -tmp
    }
    return tmp
}
