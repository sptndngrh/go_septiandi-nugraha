# Materi 18 Middleware

Middleware adalah entitas yang terhubung ke pemrosesan permintaan/respons server untuk menerapkan berbagai fungsi seputar Permintaan HTTP masuk ke server Anda dan respons keluar.

Tiga poin tentang middleware:
> CORS Middleware:
    CORS (Cross-Origin Resource Sharing) memungkinkan sumber daya untuk dibagikan antara berbagai domain/origin.
Konfigurasi umum CORS mencakup Access-Control-Allow-Origin (menentukan domain yang dapat mengirimkan permintaan), Access-Control-Allow-Headers (menentukan header yang dapat digunakan dalam permintaan), Access-Control-Allow-Methods (menentukan metode HTTP yang diizinkan dalam permintaan), dan Access-Control-Max-Age (menentukan waktu maksimum untuk caching preflight request).
Contoh Implementasi:

===============================================
go
Copy code
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.CORS())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "hello")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
===============================================

> Rate Limiter Middleware:
    Rate Limiter digunakan untuk membatasi jumlah permintaan yang dapat diterima oleh server.
Ini membantu mencegah serangan DDoS, Brute Force, dan menjaga kinerja server.
Contoh Implementasi:

===============================================
go
Copy code
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "hello")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
===============================================

> Logger Middleware:
    Middleware logger digunakan untuk mencatat informasi tentang permintaan HTTP.
Ini digunakan sebagai jejak dan sumber data untuk analitika.
Dalam tambahan, ada juga pembahasan tentang Auth Middleware yang dapat digunakan untuk mengidentifikasi pengguna dan melindungi data. Ada dua jenis auth middleware yang disebutkan yaitu Basic Authentication dan JSON Web Token (JWT) Authentication. Basic Authentication mengharuskan pengguna memasukkan username dan password dalam header permintaan HTTP dalam format yang dienkripsi.