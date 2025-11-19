package mainimport "fmt"func Sumarr(arr []int) int {
    sum := 0    for _, v := range arr {
       sum += v
    }
    return sum
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 30, 5}
    sum := Sumarr(arr)
    fmt.Println(sum)
}
