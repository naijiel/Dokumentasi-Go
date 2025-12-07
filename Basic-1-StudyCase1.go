// Seorang kasir di toko kecil harus menghitung total belanja pelanggan yang
// membeli berbagai barang dengan harga berbeda. Kadang pelanggan mendapat potongan
// harga berdasarkan total belanja atau jenis barang tertentu. Setelah transaksi selesai,
// kasir perlu menampilkan ringkasan pembelian beserta keterangan
// apakah pelanggan mendapatkan diskon atau tidak.
package main

import "fmt"

type dataBarang struct {
  nama, jenis string
  jumlah int
  harga float64
}

func main () {
  var i, n int
  var produk dataBarang

  fmt.Print("Jumlah barang yang ingin di input: ")
  fmt.Scan(&n)
  for i = 1; i <= n; i++ {
    fmt.Print("Nama barang: ")
    fmt.Scan(&produk.nama)
    fmt.Print("Harga: ")
    fmt.Scan(&produk.harga)
    fmt.Print("Jumlah: ")
    fmt.Scan(&produk.jumlah)
    fmt.Print("Jenis: ")
    fmt.Scan(&produk.jenis)
    sistem(&produk.harga, &produk.jumlah, &produk.jenis, &produk.nama)
  }
}

func sistem(*h, *sum, *jns, *nm) {
  var calculation float64
  if sum > 2 && jenis != "Promo"{
    calculation = discCountBySum(h, sum, jns)
    fmt.Printf("Harga total untuk %s\n", nm)
    fmt.Print("Rp.", sum * h)
    fmt.Print("Harga setelah diskon: Rp. ", calculation)
    fmt.Println()
    sumOfShopping(&calculation)
  } else if sum <= 2  && sum > 0 && jenis != "Promo"{
    fmt.Printf("Harga total untuk %s\n", nm)
    fmt.Print("Rp. ")
    calculation = normalCountBelanja(h, sum, jns)
    fmt.Println()
    sumOfShopping(&calculation)
  } else if jenis == "Promo" {
    calculation = disCountByJns(h, sum, jns)
    fmt.Printf("Harga total untuk %s\n", nm)
    fmt.Print("Rp.", sum * h)
    fmt.Print("Harga setelah diskon: Rp. ", calculation)
    fmt.Println()
    sumOfShopping(&calculation)
  }
}

func normalCountBelanja(harga, jumlah, jenis) float64 {
  return harga * jumlah
}

func dicCountBySum(harga, jumlah, jenis) float64 {
  var disc, total float64
  if jumlah == 3 {
    total = jumlah * harga
    disc = total * 0.25
    total = total - disc
  } else if jumlah > 3 {
    total = jumlah * harga
    disc = total * 0.6
    total = total - disc
  }

  return total
}

func dicCountByJns(harga, jumlah, jenis) float64 {
  var total, disc float64
  total = jumlah * harga
  disc = total * 0.45
  return total - disc
}

func sumOfShopping(*calcu) {
  var total float64
  total = total + calcu
  fmt.Println("Total belanja anda kali ini adalah")
  fmt.Println(calcu)
}
