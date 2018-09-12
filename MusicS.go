package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

//Declaración de la estructura
type Songs struct {
    Id int	`gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Artist string `gorm:"not null" form:"artist" json:"artist"`
    Song  string `gorm:"not null" form:"song" json:"song"`
    Genre  int `gorm:"not null" form:"genre" json:"genre"`
    Length  int `gorm:"not null" form:"length" json:"length"`
}

func InitDb() *gorm.DB {
    // Openning file
    db, err := gorm.Open("sqlite3", "./jrdd.db")
    // Display SQL queries
    db.LogMode(true)

    // Error
    if err != nil {
   	 panic(err)
    }
    // Creating the table
    if !db.HasTable(&Songs{}) {
   	 db.CreateTable(&Songs{})
   	 db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Songs{})
    }

    return db
}


func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
   	 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
   	 c.Next()
    }
}

func main() {
    r := gin.Default()

    r.Use(Cors())

    v1 := r.Group("api/v1")
    {

   	 v1.GET("/songs", GetSongs)
   	 v1.GET("/songs/:artist", GetByArtist)
   	 //v1.GET("/songs/:genre", GetByGenre)
    }

    r.Run(":8080")
}

func GetSongs(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    var songs []Songs
    // SELECT * FROM Songs
    db.Find(&songs)

    // Display JSON result
    c.JSON(200, songs)

    // curl -i http://localhost:8080/api/v1/songs
}

func GetByArtist(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    artist := c.Params.ByName("artist")
    var song Songs
    // SELECT * FROM Songs WHERE artist = ?;
    //db.First(&song, artist)
    db.Where("artist = ?", artist).First(&song)

    //len(song.Artist) != 0
    //song.Artist != ""
    if len(song.Artist) != 0 {
   	 // Display JSON result
   	 c.JSON(200, song)
    } else {
   	 // Display JSON error
   	 c.JSON(404, gin.H{"error": "Artist not found"})
    }

    // curl -i http://localhost:8080/api/v1/songs/Artist
}

/*
func GetByGenre(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    genre := c.Params.ByName("genre")
    var song Songs
    // SELECT * FROM songs WHERE genre = 1;
    db.First(&song, genre)

    if song.Genre != 0 {
   	 // Display JSON result
   	 c.JSON(200, song)
    } else {
   	 // Display JSON error
   	 c.JSON(404, gin.H{"error": "Genre not found"})
    }

    // curl -i http://localhost:8080/api/v1/song/Genre
}
*/

func OptionsSongs(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    c.Next()
}
