# Backend-Golang
- [Download](https://golang.org) dan install terlebih dahulu golang di komputer anda jika belum terinstall
- Minimal menginstall Golang versi v1.11
- Kita akan menggunakan go mod sebagai Versioning & Dependency management, sehingga kita tidak harus meletakan project go kita didalam direktori GOPATH lagi dan manual install dependency
- Buatkan satu file environment dengan nama `env.toml` dan letakan di root project. Sudah saya sertakan contohnya dengan nama env.toml.example
- Konfigurasi semua config didalam file tersebut sesuai environment kalian
- Jalankan perintah `go run main.go` (Perintah ini akan mendownload semua dependency jika baru pertama kali running) kemudian akan menjalankan project ini
- Semua `Log Query` dapat dilihat pada command line