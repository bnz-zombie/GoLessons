package mainimport "fmt"

type Greeter interface {
    Greet() string}

type Person struct {
    Name string}

type Robot struct {
    ID string}

func (p Person) Greet() string {
    return "Hello " + p.Name
}

func (r Robot) Greet() string {
    return "Beep. I am robot " + r.ID
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    p := Person{Name: "Фёдор"}
    r := Robot{ID: "RX-78"}

    SayHello(p)
    SayHello(r)
}


