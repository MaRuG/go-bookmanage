package controllers

import (
    "log"
	"github.com/revel/revel"
    "net/http"
    "io/ioutil"
    // "bufio"
    "os"
    "fmt"
    // "reflect"
    "unsafe"
    // "bytes"
    "strings"
    "encoding/json"
    "bookmanage/app/models"
)

type App struct {
	*revel.Controller
}

type Booksinfo struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		Etag       string `json:"etag"`
		SelfLink   string `json:"selfLink"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			ReadingModes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			PageCount        int      `json:"pageCount"`
			PrintType        string   `json:"printType"`
			Categories       []string `json:"categories"`
			AverageRating    float64  `json:"averageRating"`
			RatingsCount     int      `json:"ratingsCount"`
			MaturityRating   string   `json:"maturityRating"`
			AllowAnonLogging bool     `json:"allowAnonLogging"`
			ContentVersion   string   `json:"contentVersion"`
			ImageLinks       struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks"`
			Language            string `json:"language"`
			PreviewLink         string `json:"previewLink"`
			InfoLink            string `json:"infoLink"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink"`
		} `json:"volumeInfo"`
		SaleInfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			IsEbook     bool   `json:"isEbook"`
		} `json:"saleInfo"`
		AccessInfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			PublicDomain           bool   `json:"publicDomain"`
			TextToSpeechPermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				IsAvailable bool `json:"isAvailable"`
			} `json:"epub"`
			Pdf struct {
				IsAvailable bool `json:"isAvailable"`
			} `json:"pdf"`
			WebReaderLink       string `json:"webReaderLink"`
			AccessViewStatus    string `json:"accessViewStatus"`
			QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet"`
		} `json:"searchInfo"`
	} `json:"items"`
}

type Book struct {
    Title       string
    Image       string
    Info        string
}

type Books []Book

func (c App) Index() revel.Result {
    if b := Exists("./test.db"); b {
        fmt.Println("file exist!!")
        // db.CreateDB()
	} else {
        fmt.Println("file not exist!!")
        models.CreateDB()
	}

    m := "Welcome to Book Management Site"
    // m := Getinfo("9784621300251")
    // fmt.Println(m["title"], m["author"])    
    // if err := db.Insert(name, mac, active); err != nil {

	return c.Render(m)
}

func (c App) DbSearch(word string) revel.Result {
    books := models.Search(word)
    return c.Render(books)
}

func (c App) IsbnRegister(isbn string) revel.Result {

    isbn = strings.Replace(isbn, "-", "", -1)
    m := Getisbn(isbn)

    return c.Render(m)
}

func (c App) TitleRegister(title string) revel.Result {

    m := Gettitle(title)
    fmt.Println(m)

    return c.Render(m)
}

func (c App) RegisterPost(title string, author string, date string, isbn string, image string) revel.Result{
    models.Insert(title, author, date, isbn, image)
    return c.Render(title, author, date, isbn, image)
}
// func (c App) DeletePost(title string, author string, date string, isbn string, image string) revel.Result {
func (c App) DeletePost(isbn string) revel.Result {
    models.Delete(isbn)
    m := Getisbn(isbn)
    // return c.Render(title, author, date, isbn, image)
    return c.Render(m)
}

func (c App) Display() revel.Result{
    books := models.Show()
    // log.Fatalln(books)
    return c.Render(books)
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

func sbytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// func Gettitle(title string) ([]map[string]string){
func Gettitle(title string) Books{

    url := "https://www.googleapis.com/books/v1/volumes?q="
    // fmt.Println(url + isbn)
    response, _ := http.Get(url + title)
    body, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()

    // println(string(body))

    body_bytes := sbytes(string(body))
    var bi Booksinfo

    var ma []map[string]string
    var titles  []string
    var infos    []string
    var images  []string

    json.Unmarshal(body_bytes, &bi)
    // err := json.Unmarshal(body_bytes, &bi)
    // if err != nil {
		// fmt.Println("error:", err)
	// }
    for i:=0; i < len(bi.Items); i++ {
        title := bi.Items[i].VolumeInfo.Title
        info := bi.Items[i].VolumeInfo.InfoLink
        image := bi.Items[i].VolumeInfo.ImageLinks.SmallThumbnail
        m := map[string]string{"title": title, "info": info, "image": image}
        titles = append(titles, title)
        infos = append(infos, info)
        images = append(images, image)
        ma = append(ma,m)
    }

    var book Book
    var books Books

    for j:=0; j<len(titles); j++ {
        book.Title = titles[j]
        book.Info = infos[j]
        book.Image = images[j]
        temp := Book{
            Title:  book.Title,
            Info:   book.Info,
            Image:  book.Image,
        }
        books = append(books, temp)
    }

    fmt.Println(books)
    return books
}

func Getisbn(isbn string) (map[string]string){

    url := "https://www.googleapis.com/books/v1/volumes?q=isbn:"
    // fmt.Println(url + isbn)
    response, _ := http.Get(url + isbn)
    body, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()

    // println(string(body))

    body_bytes := sbytes(string(body))
    var bi Booksinfo

    json.Unmarshal(body_bytes, &bi)
    title := bi.Items[0].VolumeInfo.Title
    author := bi.Items[0].VolumeInfo.Authors
    date := bi.Items[0].VolumeInfo.PublishedDate
    // isbn10_type := bi.Items[0].VolumeInfo.IndustryIdentifiers[0].Type
    isbn10 := bi.Items[0].VolumeInfo.IndustryIdentifiers[0].Identifier
    // isbn13_type := bi.Items[0].VolumeInfo.IndustryIdentifiers[1].Type
    isbn13 := bi.Items[0].VolumeInfo.IndustryIdentifiers[1].Identifier
    image := bi.Items[0].VolumeInfo.ImageLinks.SmallThumbnail

    m := map[string]string{"title": title, "author": author[0], "date": date, "isbn10": isbn10, "isbn13": isbn13, "image": image}
    // fmt.Println(m["isbn13"])
    // fmt.Println(m)
    return m
}

