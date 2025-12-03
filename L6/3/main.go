package mainimport "fmt"type Greeter interface {
    Greet() string}

type Person struct {
    Name string}

func (p Person) Greet() string {
    return "Hello " + p.Name
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    p := Person{Name: "Фёдор"}
    SayHello(p)
}

