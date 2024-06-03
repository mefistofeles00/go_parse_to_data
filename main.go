package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func createSlug(text string, wordLimit int) string {
	words := strings.Fields(text)
	if len(words) > wordLimit {
		words = words[:wordLimit]
	}
	slug := strings.Join(words, "-")
	return strings.ToLower(slug)
}

func fetchAndProcessPage(url string, db *sql.DB) {
	// SAYFA VERILERINI ALIYORUM
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Failed to fetch page: %d %s", res.StatusCode, res.Status)
	}

	// BURADA  HTMLI PARSSE EDIYORUZ
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// CLASS ADINA GORE VERILERI BULMA VE VERITABANINA EKLIYORUM
	doc.Find("span.profile-description").Each(func(index int, item *goquery.Selection) {
		profileDescription := item.Text()
		fmt.Println("Profile Description:", profileDescription)

		// Slug oluşturma
		slug := createSlug(profileDescription, 4) // 4 kelime sınırlaması ile slug oluşturma

		// VERIYI CATEGORIES TABLOSUNEA EKLIYORUM
		query := "INSERT INTO categories (name, slug) VALUES (?, ?)"
		_, err = db.Exec(query, profileDescription, slug)
		if err != nil {
			log.Println("Error inserting data into categories:", err)
		} else {
			fmt.Println("Data inserted successfully into categories table")
		}
	})
}

func main() {
	baseURL := "https://www.sahibinden.com/marangoz-hizmetleri?currentPageIndex="
	totalPages := 55 // toplam sayfa sayısını buraya girin

	// MySQL veritabanına bağlanma
	dsn := "username:password@tcp(orkneksunucuadi.com)/database_name"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
		url := baseURL + strconv.Itoa(pageIndex)
		fmt.Println("Fetching URL:", url)
		fetchAndProcessPage(url, db)
	}
}
