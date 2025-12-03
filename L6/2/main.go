package mainimport (
    "fmt")

type Point struct {
    X int    Y int}

func (p *Point) Move(dx, dy int) {
    p.X += dx
    p.Y += dy
}

func main() {
    p2 := Point{X: 1, Y: 2}
    p := &Point{X: 1, Y: 2}
    p.Move(3, -1)
    fmt.Println("Point:", p.X, p.Y) /
}


