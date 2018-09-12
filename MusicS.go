package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

//Declaración de la estructura
type Songs struct {
    //Id int	`gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Artist string `gorm:"not null" form:"artist" json:"artist"`
    Song  string `gorm:"not null" form:"song" json:"song"`
    //Genre  int `gorm:"not null" form:"genre" json:"genre"`
    Genre  string `gorm:"not null" form:"genre" json:"genre"`
    Length  int `gorm:"not null" form:"length" json:"length"`
}

type Genres struct{
    //ID int	`gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Name string `gorm:"not null" form:"name" json:"name"`
}

//Estructura para la segunda consulta extra
type ExGenres struct {
    Genre string `gorm:"not null" form:"genre" json:"genre"`
    Qty int `gorm:"not null" form:"qty" json:"qty"`
    Total int `gorm:"not null" form:"total" json:"total"`
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
   	 v1.GET("/genres", GetGenres)
   	 //v1.GET("/songs/:artist", GetByArtist)
   	 v1.GET("/genres/:genre", GetByGenre)
   	 v1.GET("/songs/:cancion", GetByTitle)

    }

    r.Run(":8080")
}

//Queries:
//SELECT S.artist, S.song, G.name, S.length FROM Songs S JOIN Genres G ON S.genre = G.id WHERE S.artist = ?

//SELECT G.name, COUNT(S.song) AS [TotalSongs], SUM(S.length) AS [TotalLenght] FROM Songs S JOIN Genres G ON S.genre = G.id GROUP BY G.name

//Mostrar la info de todas las canciones
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

//Mostrar todos los generos
func GetGenres(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    var genres []Genres
    // SELECT * FROM Genres
    db.Find(&genres)

    // Display JSON result
    c.JSON(200, genres)

    // curl -i http://localhost:8080/api/v1/genres
}

/*
func GetByArtist(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    artist := c.Params.ByName("artist")
    var song Songs
    // SELECT * FROM Songs WHERE artist = ?;
    //db.First(&song, artist)
    //db.Where("artist = ?", artist).First(&song)
    //db.Raw("SELECT S.artist, S.song, G.name, S.length FROM Songs S INNER JOIN Genres G ON S.genre = G.id WHERE S.artist = ?", artist).Scan(&song)
    db.Table("Songs").Select("Songs.Artist, Songs.Song, Genres.Name, Songs.Length").Joins("inner join Genres on Songs.Genre = Genres.ID").Where("Songs.Artist = ?", artist).Scan(&song)

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
*/

func GetByTitle(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    cancion := c.Params.ByName("cancion")
    var song Songs
    // SELECT * FROM Songs WHERE artist = ?;
    //db.First(&song, artist)
    //db.Where("artist = ?", artist).First(&song)
    //db.Raw("SELECT S.artist, S.song, G.name, S.length FROM Songs S INNER JOIN Genres G ON S.genre = G.id WHERE S.artist = ?", artist).Scan(&song)
    db.Table("Songs").Select("Songs.Artist, Songs.Song, Genres.Name, Songs.Length").Joins("inner join Genres on Songs.Genre = Genres.ID").Where("Songs.Song = ?", cancion).Scan(&song)

    //len(song.Song) != 0
    //song.Song != ""
    if len(song.Song) != 0 {
   	 // Display JSON result
   	 c.JSON(200, song)
    } else {
   	 // Display JSON error
   	 c.JSON(404, gin.H{"error": "Song not found"})
    }

    // curl -i http://localhost:8080/api/v1/songs/Title
}


func GetByGenre(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    genre := c.Params.ByName("genre")
    var song Songs
    // SELECT * FROM songs WHERE genre = 1;
    //db.First(&song, genre)
    //db.Where("genre = ?", genre).First(&song)
    db.Table("Songs").Select("Songs.Artist, Songs.Song, Genres.Name, Songs.Length").Joins("inner join Genres on Genres.ID = Songs.Genre").Where("Genre.Name = ?", genre).Scan(&song)

    //song.Genre != 0 -> len(song.Artist) != 0
    if len(song.Genre) != 0 {
   	 // Display JSON result
   	 c.JSON(200, song)
    } else {
   	 // Display JSON error
   	 c.JSON(404, gin.H{"error": "Genre not found"})
    }

    // curl -i http://localhost:8080/api/v1/song/Genre
}


func OptionsSongs(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    c.Next()
}


