package main
import "fmt"
func main(){
  var publisher, writer, artist, title, genre string
  title = "Mr. GoToSleep"
  publisher = "DizzyBooks Publishing Inc."
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  genre = "Fantasy"
  var year, pageNumber int
  year = 1997
  pageNumber = 14
  var grade float32
  grade = 6.5
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "Genre", genre, "Year", year, "Pages", pageNumber, "Condition", grade )
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn."
  artist = "Phoebe Paperclips"
  genre = "Adventure"
  year = 2013
  pageNumber = 160
  grade = 9.0
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "Genre", genre, "Year", year, "Pages", pageNumber, "Condition", grade  )
}
