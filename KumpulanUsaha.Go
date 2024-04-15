//Usaha dan Berdoa//
//BISMILLAH AJA DULU//



package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	sum := 0

	for i := 1; i <= N; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	fmt.Printf("Hasil penjumlahan bilangan ganjil antara 1 hingga %d adalah %d\n", N, sum)
}

func main() {
	var bil, sesudah, sebelum int
	var naik bool
	
	for bil <= 0 {
		fmt.scan(&bil)
	}
	digit = bil % 10
	bil= bil /10
	naik = true

	for bil !=0 && naik {
		sesudah = bil % 10
		naik = (sesudah+1 == digit)
		digit = bil % 10
		bil = bil / 10
	}
	fmt.Print(naik)
}	
package main

import (
	"fmt"
)

func main() {
	days := []int{52, 46, 65, 65, 69, 52, 82, 89, 75, 111}
	mean, decreases := averageAndDecreases(days)

	fmt.Println("Rata-rata pengunjung:", mean)
	fmt.Println("Penurunan pengunjung:", decreases)

	if mean < 25 || decreases > 5 {
		fmt.Println("Kafe akan ditutup")
	} else {
		fmt.Println("Kafe akan dibuka permanen")
	}
}

func averageAndDecreases(days []int) (float64, int) {
	sum := 0
	decreases := 0

	for i := 1; i < len(days); i++ {
		if days[i] < days[i-1] {
			decreases++
		}
		sum += days[i]
	}

	mean := float64(sum) / float64(len(days))
	return mean, decreases
}

func dapatSim(umur int, nilai int, golang bool) bool {
	if umur >= 10 && nilai == 100 && golang {
		return true
	} else {
		return false
	}
}

func main() {
	var umur, nilai int
	var golang bool
	fmt.Scan(&umur, &golang, &nilai)
	fmt.Println(dapatSim(umur, nilai, golang))
}
package main // luas tabung

import "fmt"

const pi float64 = 3.14

func main() {
	var v float64
	fmt.Printf("%.1f\n", v)
	fmt.Scan(&v)
}

func volume(r, t float64) float64 {
	var v float64
	v = pi * r * r * t
	return v
}
package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Println("60:")
	fmt.Scanln(&60)

	// Call function to check number
	checkBilangan(60)
}

func checkBilangan(number int) {
	// Check if number is a prime number
	if isPrime(60) {
		fmt.Printf("%d is a prime number\n", number)
	} else {
		fmt.Printf("%d is not a prime number\n", number)
	}

	// Check if number is a perfect square
	if isPerfectSquare(number) {
		fmt.Printf("%d is a perfect square\n", number)
	} else {
		fmt.Printf("%d is not a perfect square\n", number)
	}
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func isPerfectSquare(number int) bool {
	root := int(math.Sqrt(float64(number)))
	return number == root*root
}
func cekBilangan2(bil int) bool {
	if bil <= 1 {
		return false
	}
	for i := 2; i*i <= bil; i++ {
		if bil%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var bil int
	fmt.Println("Masukkan bilangan:")
	fmt.Scan(&bil)

	result := cekBilangan2(bil)
	fmt.Println("Bilangan tersebut adalah bilangan berikutnya?", result)
}
package main// CLO 4 SEMESTER 1

import (
	"fmt"
	"time"
)

func addDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func main() {
	var day, month, year int
	fmt.Println("Masukkan tanggal peminjaman: ")
	fmt.Scanln(&day)
	fmt.Scanln(&month)
	fmt.Scanln(&year)

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	returnDate := addDays(date, 7)

	fmt.Println("Tanggal pengembalian buku:", returnDate.Day(), int(returnDate.Month()), returnDate.Year())
}
package main

import (
	"fmt"
)

func kategoriUsia(usia int) string {
	if usia < 0 {
		return "Tidak valid"
	} else if usia < 19 {
		return "Anak-anak"
	} else if usia >= 65 {
	 	return "Orang tua"
	}
	return "Dewasa"
}

func hitungProsentase(usiaList []int) map[string]float64 {
	kategoriCount := map[string]int{"Anak-anak": 0, "Dewasa": 0, "Orang tua": 0}
	totalOrang := len(usiaList)

	for _, usia := range usiaList {
		kategori := kategoriUsia(usia)
		if kategori != "Tidak valid" {
			kategoriCount[kategori]++
		}
	}

	prosentase := map[string]float64{}
	for kategori, count := range kategoriCount {
		prosentase[kategori] = (float64(count) / float64(totalOrang)) * 100
	}

	return prosentase
}

func main() {
	usiaList := []int{25, 27, 8, 55, 65, 78, 39, 51}
	usiaList = append(usiaList, 999)

	prosentase := hitungProsentase(usiaList)

	fmt.Println("Anak-anak:", prosentase["Anak-anak"], "%")
	fmt.Println("Dewasa:", prosentase["Dewasa"], "%")
	fmt.Println("Orang tua:", prosentase["Orang tua"], "%")
}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		highestIPK, lowestIPK float64
		highestCount          int
	)

	for {
		var NIM string
		fmt.Print("Masukkan NIM (atau '00' untuk keluar): ")
		fmt.Scanln(&NIM)

		if NIM == "00" {
			break
		}

		fmt.Print("Masukkan nilai IPK mahasiswa: ")
		var IPKInput string
		fmt.Scanln(&IPKInput)
		IPK, err := strconv.ParseFloat(IPKInput, 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if IPK > highestIPK {
			highestIPK = IPK
			highestCount = 1
		} else if IPK == highestIPK {
			highestCount++
		}

		if IPK < lowestIPK || lowestIPK == 0 {
			lowestIPK = IPK
		}
	}

	fmt.Printf("IPK tertinggi: %.2f\n", highestIPK)
	fmt.Printf("Jumlah mahasiswa dengan IPK tertinggi: %d\n", highestCount)
	fmt.Printf("IPK terendah: %.2f\n", lowestIPK)
}
package main

import (
	"fmt"
	"time"
)

func addDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func main() {
	var day, month, year int
	fmt.Println("Masukkan tanggal peminjaman: ")
	fmt.Scanln(&day)
	fmt.Scanln(&month)
	fmt.Scanln(&year)	

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	returnDate := addDays(date, 7)

	fmt.Println("Tanggal pengembalian buku:", returnDate.Day(), int(returnDate.Month()), returnDate.Year())
}
package main

import "fmt"

func main() {
	var x, y, z int

	// Input
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("Masukkan nilai z: ")
	fmt.Scan(&z)

	x = 0

	// Output header
	fmt.Printf("%-5s%-5s%-5s\n", "x", "y", "z")
	fmt.Printf("%-5d%-5d%-5d\n", x, y, z)

	for z > 0 {
		if z%2 == 1 {
			x += y
		}
		y *= 2
		z /= 2

		// Output baris tabel
		fmt.Printf("%-5d%-5d%-5d\n", x, y, z)
	}

	// Output hasil akhir
	fmt.Println("Hasil: ", x)
}
package main

import (
	"fmt"
	"os"
	"strconv"
)

func isDigitsAscending(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("n lebih dari 1000 dan n kuranf dari:")
	fmt.Scan(&n)
	if n >= 1000 && n <= 9999 {
		result := isDigitsAscending(n)
		var a []any = []any{n, func() string {
			if result {
				return ""
			}
			return "not "
		}()}
		fmt.Fprintf(os.Stdout, "The digits in %d are %sascending.\n", a...)
	} else {
		fmt.Println("Invalid input. Please enter a number between 1000 and 9999.")
	}
}
package main//Mencari Posisi

import "fmt"

func main() {
	var posisi, posisi0, kecepatan, waktu int
	fmt.Scan(&posisi0, &kecepatan, &waktu)
	posisi = posisi0 + (kecepatan * waktu)
	fmt.Println(posisi)
}
package main

import (
	"fmt"
)

func main() {
	var celcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)

	reamur := celcius * (4.0 / 5.0)
	fahrenheit := (celcius * 9.0 / 5.0) + 32.0
	kelvin := celcius + 273.15

	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}


package main

import (
	"fmt"
	"os"~
)

func hitungProsentase(data []int, umur int) float64 {
	count := 0
	for _, value := range data {
		if value >= umur {
			count++
		}
	}
	return float64(count) / float64(len(data)) * 100
}

func main() {
	dataUsia := []int{}
	for {
		fmt.Print("Masukkan usia penduduk (999 untuk berhenti): ")
		var inputUsia int
		fmt.Scan(&inputUsia)
		if inputUsia == 999 {
			break
		}
		dataUsia = append(dataUsia, inputUsia)
	}

	if len(dataUsia) == 0 {
		fmt.Println("Tidak ada data")
		os.Exit(0)
	}

	anak := hitungProsentase(dataUsia, 0)
	dewasa := hitungProsentase(dataUsia, 19)
	orangTua := hitungProsentase(dataUsia, 65)

	fmt.Printf("Anak-anak: %.1f%%\n", anak)
	fmt.Printf("Dewasa: %.1f%%\n", dewasa)
	fmt.Printf("Orang tua: %.1f%%\n", orangTua)
}
package main

import (
	"fmt"
)

func main() {
	days := []int{52, 46, 65, 65, 69, 52, 82, 89, 75, 111}
	mean, decreases := averageAndDecreases(days)

	fmt.Println("Rata-rata pengunjung:", mean)
	fmt.Println("Penurunan pengunjung:", decreases)

	if mean < 25 || decreases > 5 {
		fmt.Println("Kafe akan ditutup")
	} else {
		fmt.Println("Kafe akan dibuka permanen")
	}
}

func averageAndDecreases(days []int) (float64, int) {
	sum := 0
	decreases := 0

	for i := 1; i < len(days); i++ {
		if days[i] < days[i-1] {
			decreases++
		}
		sum += days[i]
	}

	mean := float64(sum) / float64(len(days))
	return mean, decreases
}
package main

import (
	"fmt"
)

func kategoriUsia(usia int) string {~
	if usia < 0 {
		return "Tidak valid"
	} else if usia < 19 {
		return "Anak-anak"
	} else if usia < 65 {
		return "Dewasa"
	} else {
		return "Orang tua"
	}
}

func hitungProsentase(usiaList []int) map[string]float64 {
	kategoriCount := map[string]int{"Anak-anak": 0, "Dewasa": 0, "Orang tua": 0}
	totalOrang := len(usiaList)

	for _, usia := range usiaList {
		kategori := kategoriUsia(usia)
		if kategori != "Tidak valid" {
			kategoriCount[kategori]++
		}
	}

	prosentase := map[string]float64{}
	for kategori, count := range kategoriCount {
		prosentase[kategori] = (float64(count) / float64(totalOrang)) * 100
	}

	return prosentase
}

func main() {
	usiaList := []int{25, 27, 8, 55, 65, 78, 39, 51}
	usiaList = append(usiaList, 999)

	prosentase := hitungProsentase(usiaList)

	fmt.Println("Anak-anak:", prosentase["Anak-anak"], "%")
	fmt.Println("Dewasa:", prosentase["Dewasa"], "%")
	fmt.Println("Orang tua:", prosentase["Orang tua"], "%")
}
package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Println("jejari:", radius)
	fmt.Scanln(&radius)
	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	luas := 4.0 * math.Pi * math.Pow(radius, 2)
	fmt.Println("volume:", volume)
	fmt.Println("luasBola:", luas)
}

func main() {
	var x int
	fmt.Scan(&x)
	x = 20
	if x > 9 {
		fmt.Println("KON")
	} else {
		fmt.Println("kon")
	}
}
package main

import "fmt"

func main() {
	age := 18
	name := "Farid"

	fmt.Printf("%v %v\n", name, age)
}

/*"%v": Ini adalah placeholder format. Ketika Anda menggunakan "%v" dalam string format,
itu menunjukkan bahwa nilai variabel akan dimasukkan di tempat tersebut. Jika Anda menggunakan Printf() atau fungsi serupa,
nilai variabel akan dimasukkan sesuai dengan urutan argumen yang diberikan setelah string format*/

func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Masukkan Tahun kabisat: ")
	fmt.Scan(&tahun)

	kabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Println(kabisat)
}
func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	var pita string
	for i := 0; i < n; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i+1)
		fmt.Scan(&bunga)
		if i > 0 {
			pita += " - " // agar outpunya menghasil seperti ini (merah "-" , putih "-", kertas"-")
		}
		pita = pita + bunga
	}

	fmt.Println("pita:", pita)
}

