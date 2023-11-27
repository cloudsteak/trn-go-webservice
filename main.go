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
	router := gin.Default()
	// Képzések lekérdezése
	router.GET("/kepzesek", kepzesLista)
	// Új képzés
	router.POST("/kepzesek", kepzesUj)
	// Az alkalmatás elérhető a 8080-as porton
	router.Run("localhost:8080")
}
