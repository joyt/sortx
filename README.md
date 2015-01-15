# sortx

Sortx is a simple golang sort extension that provides a method for sorting an array of structs by a field.

### Installation:
`go get github.com/joyt/sortx`

### Example:
```go
import (
	"time"
	"github.com/joyt/sortx"
)
type Person struct {
	Name string
	Age int
	BirthDate time.Time
	Married bool
	Popularity float64
}

func main() {

	people := []Person{
		{Name: "Santa Claus", Age: 424, BirthDate: time.Date(1590, 12, 25, 0,0,0,0, time.Local()), Married: true, Popularity: 95.7},
		{Name: "Easter Bunny", Age: 231, BirthDate:  time.Date(1783, 4, 5, 0,0,0,0, time.Local()), Married: false, Popularity: 20.4},
		{Name: "Cupid", Age: 1250, time.Date(764, 2, 14, 0,0,0,0, time.Local()), Married: false, Popularity: 55.5},
		{Name: "Tooth Fairy", Age: 104, BirthDate: time.Date(1910, 2, 14, 0,0,0,0, time.Local()), Married: true, Popularity: 32.1},
	}

	// Sort by name alphabectially
	sortx.SortByField(people, "Name", sortx.Ascending)

	// Sort by birth date, oldest first.
	sortx.SortByField(people, "BirthDate", sortx.Ascending)

	// Sort by age, oldest first
	sortx.SortByField(people, "Age", sortx.Descending)

	// Sort by popularity
	sortx.SortByField(people, "Popularity", sortx.Descending)
}
```
