package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	
	if len(os.Args) < 2 {
		fmt.Println("HATA: Lütfen bir URL girin.")
		fmt.Println("Kullanım şekli: go run main.go https://example.com")
		return
	}

	url := os.Args[1]
	fmt.Println("Girilen URL:", url)


	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP isteği sırasında hata oluştu:")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Status Code:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("HTML okunamadı:", err)
		return
	}

	htmlFile, err := os.Create("site_data.html")
	if err != nil {
		fmt.Println("Dosya oluşturulamadı:", err)
		return
	}
	defer htmlFile.Close()

	htmlFile.Write(body)
	fmt.Println("HTML içeriği site_data.html dosyasına kaydedildi.")


	file, err := os.Open("site_data.html")
	if err != nil {
		fmt.Println("HTML dosyası açılamadı:", err)
		return
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("HTML parse edilemedi:", err)
		return
	}

	linkFile, err := os.Create("links.txt")
	if err != nil {
		fmt.Println("links.txt oluşturulamadı:", err)
		return
	}
	defer linkFile.Close()

	fmt.Println("Sayfadaki linkler:")

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			fmt.Println(link)
			linkFile.WriteString(link + "\n")
		}
	})

	fmt.Println("Linkler links.txt dosyasına kaydedildi.")


	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var screenshotBuf []byte

	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(3*time.Second),
		chromedp.FullScreenshot(&screenshotBuf, 90),
	)

	if err != nil {
		fmt.Println("Screenshot alınırken hata oluştu:")
		fmt.Println(err)
		return
	}

	err = os.WriteFile("screenshot.png", screenshotBuf, 0644)
	if err != nil {
		fmt.Println("Screenshot dosyaya yazılamadı:", err)
		return
	}

	fmt.Println("Ekran görüntüsü screenshot.png olarak kaydedildi.")
}
