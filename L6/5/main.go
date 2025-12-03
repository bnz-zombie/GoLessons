package main
import "fmt"

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
    return "Beep! I am robot " + r.ID
}

func main() {
    greeters := []Greeter{
       Person{Name: "Фёдор"},
       Person{Name: "King"},
       Robot{ID: "RX-78"},
       Robot{ID: "Terminator"},
       Robot{ID: "AI"},
    }
    

    for _, g := range greeters {
       fmt.Println(g.Greet())
    }
}




