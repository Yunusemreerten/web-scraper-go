# Web Scraper Go

Bu proje, Scraper Görevi kapsamýnda Go (Golang) kullanýlarak geliþtirilmiþtir.

## Projenin Amacý
Belirtilen bir URL üzerinden:
- Web sayfasýnýn HTML içeriðini çekmek
- Sayfadaki tüm baðlantýlarý analiz ederek `links.txt` dosyasýna kaydetmek
- Web sayfasýnýn tam ekran görüntüsünü (`screenshot.png`) almak

## Kullanýlan Teknolojiler
- Go (Golang)
- net/http
- chromedp
- goquery

## Kullaným þekli
```bash
go run main.go https://example.com
