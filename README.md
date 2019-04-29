# Instalação

```go 
go get github.com/lucasmdomingues/goartists
```

# Exemplos

### Busca de artista ou banda por nome.

```go
import (
	"fmt"
	"goartists"
)

func main() {

	artist, err := SearchArtist(appID, "Megadeth")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", artist.Name)
}
```

### Busca de eventos por artista ou banda.

```go
import (
	"fmt"
	"goartists"
)

func main() {
  
  const appID = ""

	artist, err := SearchArtist(appID, "Megadeth")
	if err != nil {
		t.Error(err)
		return
	}

	events, err := artist.GetEvents(appID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", events[0].Venue)
}
```
### Bandsintown
https://manager.bandsintown.com/support/bandsintown-api
