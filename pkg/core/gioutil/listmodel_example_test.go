package gioutil_test

import (
	"fmt"

	"github.com/diamondburned/gotk4/pkg/core/gioutil"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

var peopleListModel = gioutil.NewListModelType[Person]()

func ExampleListModel() {
	list := peopleListModel.New()
	list.Append(Person{"Alice", 20})
	list.Append(Person{"Bob", 30})
	list.Append(Person{"Charlie", 40})

	// AllItems() can be iterated over if rangefunc is supported.
	all := list.AllItems()
	all(func(p Person) bool {
		fmt.Println(p)
		return true
	})

	// Output:
	// Alice: 20
	// Bob: 30
	// Charlie: 40
}
