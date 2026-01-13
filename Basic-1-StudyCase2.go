// Sebuah aplikasi tiket bus ingin membantu penumpang memilih tujuan perjalanan,
// waktu keberangkatan, dan tipe kursi. Harga tiket berbeda tergantung rute dan
// waktu keberangkatan. Penumpang juga bisa mendapatkan promo tertentu jika
// memenuhi syarat tertentu. Setelah semua pilihan dibuat, sistem menampilkan
// total harga dan status tiket.
package main

import "fmt"

type bus struct {
  tujuan string
  waktu string
}
func main () {
  var pilih bus
  fmt.Println("--Pilih tiket yang anda butuhkan--")
  tujuan()
}

func tujuan() {
  var pilih int
  fmt.Println("-Pilih tujuan anda-")
  fmt.Println("1. Kota A")
  fmt.Println("2. Kota B")
  fmt.Println("3. Kota C")
  fmt.Println("4. Kota D")
  fmt.Println("5. Kota E")
  fmt.Print("\nPilih: ")
  fmt.Scan(&pilih)
  waktu(&pilih)
}

func waktu(tujuan *int) {
  var pilih bus
  fmt.Println("-kapan waktu keberangkatan anda-")
  fmt.Scan(&pilih.waktu) //terdapat 3 jenis waktu keberangkatan yaitu pagi, siang, malam
  Diskon(&tujuan, &pilih)
}

func Diskon(tujuan *int, pilih *string) { // simbol * dapat berarti pointer atau nilai yang diberikan akan di keluarkan lagi
  var disc float64
  if (pilih == "siang" || pilih == "pagi") && tujuan == 3{
    disc = 0.15
  }  else if pilih == "malam" {
    disc = 0.4
  } else {
    disc = 0
  }
  kalkulasi(&disc, &tujuan, &pilih)
}

func kalkulasi(D *float64, tujuan *int, pilih *string) {
  var total float64
  total = 320000 - D
  if tujuan == 1 {
    fmt.Println("selamat menikmati perjalanan anda menuju kota A")
    fmt.Printf("Total pembayaran untuk satu tiket anda %f\n", total)
  } else if tujuan == 2 {
    fmt.Println("selamat menikmati perjalanan anda menuju kota B")
    fmt.Printf("Total pembayaran untuk satu tiket anda %f\n", total)
  } else if tujuan == 3 {
    fmt.Println("selamat menikmati perjalanan anda menuju kota C")
    fmt.Printf("Total pembayaran untuk satu tiket anda %f\n", total)
  } else if tujuan == 4 {
    fmt.Println("selamat menikmati perjalanan anda menuju kota D")
    fmt.Printf("Total pembayaran untuk satu tiket anda %f\n", total)
  } else if tujuan == 5 {
    fmt.Println("selamat menikmati perjalanan anda menuju kota E")
    fmt.Printf("Total pembayaran untuk satu tiket anda %f\n", total)
  }
}
