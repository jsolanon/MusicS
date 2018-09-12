# MusicS

MusicS is a simple REST app, using golang and sqlite

# REST API

-GetSongs

Shows all the songs in the database
http://localhost:8080/api/v1/songs

-GetByArtist

Search a song by artist
http://localhost:8080/api/v1/songs/NombreDelArtista

-GetBySong

Search a song by song tittle
http://localhost:8080/api/v1/songs/NombreDeLaCancion

-GetGenres

Shows all the genres in the database
http://localhost:8080/api/v1/genres

-GetByGenre

Search a song by genre name
http://localhost:8080/api/v1/genres/NombreGenero


# Ejecutar

go run MusicS.go

y seleccionar uno de los métodos antes mencionados en el navegador.