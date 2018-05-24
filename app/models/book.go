package models

import (
	//"reflect"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
    Title       string
    Author      string
    Date        string
    Isbn        string
    Image       string
}

type Books []Book

func CreateDB(){
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
        `create table book (
        title    varchar(20),
        author   varchar(20),
        publish  varchar(7),
        isbn     varchar(13) unique,
        image    varchar(2083)
        )`,
	)

	if err != nil {
		panic(err)
	}
}

func Insert(title string, author string, date string, isbn string, image string) (error){
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
	  panic(err)
	}
	// データの挿入
	_, err = db.Exec(
		`insert into book (title, author, publish, isbn, image) VALUES (?, ?, ?, ?, ?)`,
		title,
		author,
		date,
        isbn,
        image,
    )

    if err != nil {
        panic(err)
	}

	return err
	/*
	// 挿入処理の結果からIDを取得
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	*/
}

func Delete(isbn string) (error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
	  panic(err)
	}

    _, err = db.Exec(
        `delete from book where isbn = ?`,
        isbn,
    )

    if err != nil{
        panic(err)
    }

    return err
}

func Search(word string) Books{
	var titles []string
	var authors []string
	var dates []string
    var isbns []string
    var images []string

    var book Book

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
	  panic(err)
	}

    newQuery := "%"+word+"%"
	rows, err := db.Query(
		`select * from book where title like ? or author like ? or isbn like ?`,
        newQuery,
        newQuery,
        newQuery,
    )
    fmt.Println(rows)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var title string
		var author string
		var date string
        var isbn string
        var image string

		// カーソルから値を取得
		if err := rows.Scan(&title, &author, &date, &isbn, &image); err != nil {
			log.Fatal("rows.Scan()", err)
			//return
		}
		/*
		fmt.Println(name, " ", mac, " ", alive)
		fmt.Println(reflect.TypeOf(name))
		fmt.Println(reflect.TypeOf(mac))
		fmt.Println(reflect.TypeOf(alive))
		*/
		titles = append(titles, title)
		authors = append(authors, author)
		dates = append(dates, date)
        isbns = append(isbns, isbn)
        images = append(images, image)
	}

    fmt.Println(titles)
	var books Books

	for j:=0; j<len(titles); j++{
		book.Title = titles[j]
		book.Author = authors[j]
		book.Date = dates[j]
        book.Isbn = isbns[j]
        book.Image = images[j]
		temp := Book{
			Title: book.Title,
			Author: book.Author,
			Date: book.Date,
            Isbn: book.Isbn,
            Image: book.Image,
		}
		books = append(books, temp)
	}
    fmt.Println(books)

	return books
}

func Show() Books{
	var titles []string
	var authors []string
	var dates []string
    var isbns []string
    var images []string

    var book Book

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
	  panic(err)
	}

	rows, err := db.Query(
		`select * from book`,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var title string
		var author string
		var date string
        var isbn string
        var image string

		// カーソルから値を取得
		if err := rows.Scan(&title, &author, &date, &isbn, &image); err != nil {
			log.Fatal("rows.Scan()", err)
			//return
		}
		/*
		fmt.Println(name, " ", mac, " ", alive)
		fmt.Println(reflect.TypeOf(name))
		fmt.Println(reflect.TypeOf(mac))
		fmt.Println(reflect.TypeOf(alive))
		*/
		titles = append(titles, title)
		authors = append(authors, author)
		dates = append(dates, date)
        isbns = append(isbns, isbn)
        images = append(images, image)
	}

	var books Books

	for j:=0; j<len(titles); j++{
		book.Title = titles[j]
		book.Author = authors[j]
		book.Date = dates[j]
        book.Isbn = isbns[j]
        book.Image = images[j]
		temp := Book{
			Title: book.Title,
			Author: book.Author,
			Date: book.Date,
            Isbn: book.Isbn,
            Image: book.Image,
		}
		books = append(books, temp)
	}

	return books
}
