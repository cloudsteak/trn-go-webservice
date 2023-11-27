# Go - Web Szolgáltatás (middleware)

Olyan Web szolgáltatást írunk itt, ahol a képernyőn megjelenítünk valamilyen adatot. Ez esetben a Mentork Klubnál elérhető Felhő képzések adatait írjuk ki a képernyőre JSON formátumban. 
Ezáltal már bármilyen webalkalmazás képes lesz ezt az adatot felolvasni és grafikusan megjeleníteni. 

## Előfeltételek

Az alábbi helyen megtalálod az előkészületeket a Go-ban való fejlesztéshez: https://github.com/cloudsteak/golang-basics

## Projekt létrehozás

1. Nyiss egy parancssort (CMD)
2. Navigálj abba a mappába ahol a kódod fogod tárolni a helyi gépeden.
3. Hozd létre a projekted mappáját. Pl.: `webszolgaltatas`

```bash
mkdir webszolgaltatas
```

4. Lépj be a mappába

```bash
cd webszolgaltatas
```

5. Készítsd el a projekted alap struktúráját

```bash
go mod init webszolgaltatas
```

6. indítsd el innen a Visual Studio Code-ot.

```bash
code .
```

## Go alkalmazás

1. Hozd létre a `main.go` fájlt a projekt gyökerében
2. A szerkesztőben a `main.go` fülön, ked el gépelni: `package`
3. A VS Code felajánl több lehetőséget is. Nekünk jelnleg a `package main` szükséges
4. Telepítsd a `gin-tonic` web keretrendszert

```bash
go install github.com/gin-gonic/gin@latest
```

5. Hozzuk létre a megjelenítendő adatokat struktúráját (ez lesz a képzés típusa). Illeszd az alábbi kódot a `package main` alá:

```go
type kepzes struct {
	ID     int     `json:"id"`     // Képzés azonosítója
	Kepzes string  `json:"kepzes"` // Képzés neve
	Felho  string  `json:"felho"`  // Melyik felhővel kaplcsolatos a képzés
	Szint  string  `json:"szint"`  // Milyen szintű a képzés
	Tipus  string  `json:"tipus"`  // Milyen formátumban elérhető
	Ora    float64 `json:"ora"`    // Milyen kiterjedésű
}
```


6. Ez alapján definiáljuk az adatot (jelenleg innen a kódból):

```go
// Képzések adatai egyben - Jelenleg helyben definiálva
var kepzesek = []kepzes{
	{ID: 1, Kepzes: "Cloud alapozó - AWS, Azure", Felho: "AWS, Azure", Szint: "alap", Tipus: "video", Ora: 3},
	{ID: 2, Kepzes: "Cloud alapozó - Azure", Felho: "Azure", Szint: "alap", Tipus: "egy napos", Ora: 7},
	{ID: 3, Kepzes: "Cloud alapozó - AWS", Felho: "AWS", Szint: "alap", Tipus: "egy napos", Ora: 7},
	{ID: 4, Kepzes: "Haladó Cloud - Azure 7 hetes képzés", Felho: "Azure", Szint: "haladó", Tipus: "7 hetes", Ora: 10.5},
	{ID: 5, Kepzes: "Azure haladó szinten", Felho: "Azure", Szint: "haladó", Tipus: "videó", Ora: 9},
}
```

7. Hozzunk létre egy függvényt, amely lekérdezi a képzéseket az adathalmazból


```go
// Függvény, amely visszadja a képzések listáját az elérhető adathalmazból
func kepzesLista(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, kepzesek)
}
```

8. Ez alá a sor alá pedig illeszd be az alábbi fő függvényt, ami már a web szolgáltatás része az alkalmazásunknak:

```go
// Fő függvény - az alkalmazás belépési pontja.
// Web alkalmazás definiálása, amely a /kepzesek ág meghívása esetén meghívja a képzések listázása függvényt
func main() {
	router := gin.Default()
	router.GET("/kepzesek", kepzesLista)
	// Az alkalmatás elérhető a 8080-as porton
	router.Run("localhost:8080")
}
```

9. Mentsd el a módosításokat

10. A Go érzékeli, hogy van csomag függőség és pár másodperc múlva módosítja is a kódot ennek megfeleően. Ezt lesz a `main.go` tartalma

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type kepzes struct {
	ID     int     `json:"id"`     // Képzés azonosítója
	Kepzes string  `json:"kepzes"` // Képzés neve
	Felho  string  `json:"felho"`  // Melyik felhővel kaplcsolatos a képzés
	Szint  string  `json:"szint"`  // Milyen szintű a képzés
	Tipus  string  `json:"tipus"`  // Milyen formátumban elérhető
	Ora    float64 `json:"ora"`    // Milyen kiterjedésű
}

// Képzések adatai egyben - Jelenleg helyben definiálva
var kepzesek = []kepzes{
	{ID: 1, Kepzes: "Cloud alapozó - AWS, Azure", Felho: "AWS, Azure", Szint: "alap", Tipus: "video", Ora: 3},
	{ID: 2, Kepzes: "Cloud alapozó - Azure", Felho: "Azure", Szint: "alap", Tipus: "egy napos", Ora: 7},
	{ID: 3, Kepzes: "Cloud alapozó - AWS", Felho: "AWS", Szint: "alap", Tipus: "egy napos", Ora: 7},
	{ID: 4, Kepzes: "Haladó Cloud - Azure 7 hetes képzés", Felho: "Azure", Szint: "haladó", Tipus: "7 hetes", Ora: 10.5},
	{ID: 5, Kepzes: "Azure haladó szinten", Felho: "Azure", Szint: "haladó", Tipus: "videó", Ora: 9},
}

// Függvény, amely visszadja a képzések listáját az elérhető adathalmazból
func kepzesLista(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, kepzesek)
}

// Fő függvény - az alkalmazás belépési pontja.
// Web alkalmazás definiálása, amely a /kepzesek ág meghívása esetén meghívja a képzések listázása függvényt
func main() {
	router := gin.Default()
	router.GET("/kepzesek", kepzesLista)
	// Az alkalmatás elérhető a 8080-as porton
	router.Run("localhost:8080")
}

```

Megjegyzés: Nem minden csomagot tud automatikusan telepíteni a Go. Tehát van amit majd külön parancsok futtatásával kell telepítneni.


## Alkalmazás futtatása

1. A megírt kódot az alábbi paranccsal tudjuk futtani a project fő mappájából: `go run .` vagy `go run main.go`
2. Ha megjelenik a terminal-ban a `[GIN-debug] Listening and serving HTTP on localhost:8080` az alkalmazás készen áll a használatra
3. Egy böngésző ablakban nyisd meg a http://localhost:8080

Eredmény:
```html
404 page not found
```
4. Ez normális, hiszen a képzések listája az alábbi linken érhető el: http://localhost:8080/kepzesek
Eredmény:
```html
[
    {
        "id": 1,
        "kepzes": "Cloud alapozó - AWS, Azure",
        "felho": "AWS, Azure",
        "szint": "alap",
        "tipus": "video",
        "ora": 3
    },
    {
        "id": 2,
        "kepzes": "Cloud alapozó - Azure",
        "felho": "Azure",
        "szint": "alap",
        "tipus": "egy napos",
        "ora": 7
    },
    {
        "id": 3,
        "kepzes": "Cloud alapozó - AWS",
        "felho": "AWS",
        "szint": "alap",
        "tipus": "egy napos",
        "ora": 7
    },
    {
        "id": 4,
        "kepzes": "Haladó Cloud - Azure 7 hetes képzés",
        "felho": "Azure",
        "szint": "haladó",
        "tipus": "7 hetes",
        "ora": 10.5
    },
    {
        "id": 5,
        "kepzes": "Azure haladó szinten",
        "felho": "Azure",
        "szint": "haladó",
        "tipus": "videó",
        "ora": 9
    }
]
```


## Alkalmazás fordítása (build)

Ha szeretnénk az alkalmazásunkat máshol is futtatni, anélkül, hogy minden fejlesztői eszközt és függőséget telepíyteni kellene, akkor azt egy csomagba le is tudjuk fordítani (buld). Ehhez az alábbi parancsot kell futtatni: `go build`

Eredményképpen Windows-on egy exe fájlt kapunk, amit futtathatunk a Go fejlesztői környezewten kívül is.



