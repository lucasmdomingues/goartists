# Installation

```go 
go get github.com/lucasmdomingues/goartists
```

# Examples

### Search artists by name.

```go
import (
	"log"

	"github.com./lucasmdomingues/goartists"
)

func main() {
	service := NewService("API_KEY")

	_, err := service.Search("Megadeth")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Search artist events.

```go
import (
	"log"
	
	"github.com./lucasmdomingues/goartists"
)

func main() {
	service := NewService("API_KEY")

	artist, err := service.Search("Megadeth")
	if err != nil {
		log.Fatal(err)
	}

	if err = service.GetEvents(artist); err != nil {
		log.Fatal(err)
	}
}
```
### Bandsintown
https://manager.bandsintown.com/support/bandsintown-api
