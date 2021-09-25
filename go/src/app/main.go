package main

import (
    "fmt"
    "time"
    "net/http"
    // "io/ioutil"
    // "encoding/json"
    "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


// album represents data about a record album.
type Album struct {
    gorm.Model
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var (
    db  *gorm.DB
    err error
)

func init() {
	db,err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
        panic(err.Error())
    }
    fmt.Println("db connected: ", &db)
    db.Set("gorm:table_options", "ENGINE=InnoDB")
    db.AutoMigrate(&Album{})
    db.LogMode(true)
}

// albums slice to seed record album data.
// var albums = []Album{
//     {Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
    defer db.Close()
	router := gin.Default()
	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
    router.PUT("/albums/:id", updateAlbumByID)
    router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("0.0.0.0:8080")
}


// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    newAlbum := Album{}
    now := time.Now()
    newAlbum.CreatedAt = now
    newAlbum.UpdatedAt = now

    // Call BindJSON to bind the received JSON to newAlbum.
    err := c.BindJSON(&newAlbum)
    if err != nil {
        c.String(http.StatusBadRequest, "Request failed: " + err.Error())
    }
    
    db.NewRecord(newAlbum)
    db.Create(&newAlbum)
    if !db.NewRecord(newAlbum) {
        c.JSON(http.StatusOK, newAlbum)
    }
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    albums := []Album{}
    db.Find(&albums)
    c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    album := Album{}

    db.Where("ID = ?", id).First(&album)
    c.IndentedJSON(http.StatusOK, album)

    // c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func updateAlbumByID(c *gin.Context) {
    album := Album{}
    id := c.Param("id")

    data := Album{}
    err := c.BindJSON(&data)
    if err != nil {
        c.String(http.StatusBadRequest, "Request failed" + err.Error())
    }

    db.Where("ID = ?", id).First(&album).Updates(&data)
}


func deleteAlbumByID(c *gin.Context) {
    album := Album{}
    id := c.Param("id")

    db.Where("ID = ?", id).Delete(&album)
}