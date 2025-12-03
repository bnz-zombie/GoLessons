package mainimport (
    "fmt")

type Point struct {
    X int    Y int}

func (p Point) DistanceFromOrigin() int {
    if p.X < 0 {
       p.X = 0 - p.X
    }
    if p.Y < 0 {
       p.Y = 0 - p.Y
    }

    return p.X + p.Y
}

func main() {
    p := Point{X: -3, Y: 5}
    p2 := Point{}

    fmt.Println("Distance:", p.DistanceFromOrigin())
    fmt.Println("Point:", p2.X, p2.Y)
}
