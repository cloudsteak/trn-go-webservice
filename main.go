package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

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
var kepzesek = []kepzes{}

// Adatok olvasása fájlból
func adatOlvasas(dataFajl string) {
	// Adatfájl megnyitása
	jsonFajl, err := os.Open(dataFajl)
	// Hibakezelés, ha valamiért nem sikerül megnyitni a fájlt
	if err != nil {
		log.Fatalln(err)
	}

	// Értékek kiolvasása byte tipusba
	byteErtekek, _ := io.ReadAll(jsonFajl)

	// Átalakítás JSON formátumra
	json.Unmarshal(byteErtekek, &kepzesek)

	log.Printf("%s sikeresen megnyitva\n", dataFajl)
	// Fájl lezárása
	defer jsonFajl.Close()
}

// Függvény, amely visszadja a képzések listáját az elérhető adathalmazból
func kepzesLista(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, kepzesek)
}

// kepzesUj segítségével egy új elemet adhatunk a képzési listánkhoz
func kepzesUj(c *gin.Context) {
	var ujKepzes kepzes

	// ujKepzes
	if err := c.BindJSON(&ujKepzes); err != nil {
		// Hibakezelés
		return
	}

	// Adjuk hozzá az új képzést a már meglévőekhez
	kepzesek = append(kepzesek, ujKepzes)
	c.IndentedJSON(http.StatusCreated, ujKepzes)
}

// Fő függvény - az alkalmazás belépési pontja.
// Web alkalmazás definiálása, amely a /kepzesek ág meghívása esetén meghívja a képzések listázása függvényt
func main() {
	adatOlvasas("data.json")
	router := gin.Default()
	// Képzések lekérdezése
	router.GET("/kepzesek", kepzesLista)
	// Új képzés
	//router.POST("/kepzesek", kepzesUj)
	// Az alkalmatás elérhető a 8080-as porton
	router.Run("localhost:8080")
}
