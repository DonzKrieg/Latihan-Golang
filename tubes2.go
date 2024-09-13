package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Tipe Bentukan
type Flight struct {
	id            int
	aircraft      string
	origin        string
	destination   string
	price         float64
	departureDate time.Time
}

const maxFlights = 100

type tabFlight [maxFlights]Flight

// Menambahkan Data Penerbangan
func addFlight(flights *tabFlight, flightCount *int, id int, aircraft, origin, destination string, price float64, year, month, day int) {
	if *flightCount < maxFlights {
		flights[*flightCount] = Flight{
			id:            id,
			aircraft:      aircraft,
			origin:        origin,
			destination:   destination,
			price:         price,
			departureDate: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
		}
		*flightCount++
		fmt.Println("Penerbangan berhasil ditambahkan.")
	} else {
		fmt.Println("Tidak dapat menambah penerbangan lagi, kapasitas penuh.")
	}
}

// Pencarian Penerbangan berdasarkan Tujuan
func searchFlightsByDestination(flights *tabFlight, flightCount int, destination string) {
	found := false
	fmt.Println("=================================================================================================")
	fmt.Printf("%s %17s %14s %17s %15s %28s", "ID", "Pesawat", "Asal", "Tujuan", "Harga", "Tanggal Keberangkatan\n")
	fmt.Println("=================================================================================================")
	for i := 0; i < flightCount; i++ {
		if flights[i].destination == destination {
			printFlight(flights[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada penerbangan yang ditemukan dengan tujuan tersebut.")
	}
	fmt.Println("=================================================================================================")
}

// Pencarian Penerbangan berdasarkan Asal
func searchFlightsByOrigin(flights *tabFlight, flightCount int, origin string) {
	found := false
	fmt.Println("=================================================================================================")
	fmt.Printf("%s %17s %14s %17s %15s %28s", "ID", "Pesawat", "Asal", "Tujuan", "Harga", "Tanggal Keberangkatan\n")
	fmt.Println("=================================================================================================")
	for i := 0; i < flightCount; i++ {
		if flights[i].origin == origin {
			printFlight(flights[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada penerbangan yang ditemukan dengan asal tersebut.")
	}
	fmt.Println("=================================================================================================")
}

// Pencarian Penerbangan berdasarkan Asal, Tujuan, dan Tanggal
func searchFlightsOrder(flights *tabFlight, flightCount int, origin, destination string, year, month, day int, A *tabFlight, n *int) {
	found := false
	fmt.Println("=================================================================================================")
	fmt.Printf("%s %17s %14s %17s %15s %28s", "ID", "Pesawat", "Asal", "Tujuan", "Harga", "Tanggal Keberangkatan\n")
	fmt.Println("=================================================================================================")
	for i := 0; i < flightCount; i++ {
		if flights[i].origin == origin && flights[i].destination == destination && flights[i].departureDate == time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC) {
			printFlight(flights[i])
			found = true
			A[*n] = flights[i]
			*n++
		}
	}
	if !found {
		fmt.Println("Tidak ada penerbangan yang ditemukan dengan asal tersebut.")
	}
	fmt.Println("=================================================================================================")
}

// Ketersediaan Tiket
func findTicketById(flights *tabFlight, flightCount int, id int) int {
	for i := 0; i < flightCount; i++ {
		if id == flights[i].id {
			return i
		}
	}
	return -1
}

// Pengubahan Data Penerbangan
func editFlight(flights *tabFlight, flightCount int, id int, newAircraft, newOrigin, newDestination string, newPrice float64, year, month, day int) {
	for i := 0; i < flightCount; i++ {
		if flights[i].id == id {
			flights[i].aircraft = newAircraft
			flights[i].origin = newOrigin
			flights[i].destination = newDestination
			flights[i].price = newPrice
			flights[i].departureDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			fmt.Println("Penerbangan berhasil diubah.")
			return
		}
	}
	fmt.Println("Penerbangan tidak ditemukan.")
}

// Penghapusan Data Penerbangan
func deleteFlight(flights *tabFlight, flightCount *int, id int) {
	for i := 0; i < *flightCount; i++ {
		if flights[i].id == id {
			for j := i; j < *flightCount-1; j++ {
				flights[j] = flights[j+1]
			}
			*flightCount--
			fmt.Println("Penerbangan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Penerbangan tidak ditemukan.")
}

// Pengurutan Data Penerbangan Berdasarkan Tanggal Keberangkatan (Selection Sort)
func sortFlightsByDepartureDateAscending(flights *tabFlight, flightCount int) {
	var temp Flight
	for pass := 0; pass < flightCount-1; pass++ {
		minIndex := pass
		for i := pass + 1; i < flightCount; i++ {
			if flights[i].departureDate.Before(flights[minIndex].departureDate) {
				minIndex = i
			}
		}
		temp = flights[pass]
		flights[pass] = flights[minIndex]
		flights[minIndex] = temp
	}
}

// func sortFlightsByDepartureDateAscending(flights *tabFlight, flightCount int) {
// 	var pass int
// 	var temp Flight
// 	pass = 1
// 	for pass <= flightCount-1 {
// 		i := pass
// 		temp = flights[pass]
// 		for i > 0 && temp.departureDate.Before(flights[i-1].departureDate) {
// 			flights[i] = flights[i-1]
// 			i--
// 		}
// 		flights[i] = temp
// 		pass++

// 	}
// }

// Menampilkan Data Penerbangan
func displayFlights(flights *tabFlight, flightCount int) {
	if flightCount == 0 {
		fmt.Println("Tidak ada penerbangan yang tersedia.")
	} else {
		fmt.Println("=================================================================================================")
		fmt.Printf("%s %17s %14s %17s %15s %28s", "ID", "Pesawat", "Asal", "Tujuan", "Harga", "Tanggal Keberangkatan\n")
		fmt.Println("=================================================================================================")
		for i := 0; i < flightCount; i++ {
			printFlight(flights[i])
		}
		fmt.Println("=================================================================================================")
	}
}

// Fungsi Bantu
func printFlight(flight Flight) {
	fmt.Printf("%d %15s %14s %17s %15.2f %22s\n", flight.id, flight.aircraft, flight.origin, flight.destination, flight.price, flight.departureDate.Format("2006-01-02"))
}

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func inputFlightData(flights *tabFlight, flightCount *int) {
	var id int
	var aircraft, origin, destination string
	var price float64
	var year, month, day int

	fmt.Println("Masukkan ID Penerbangan:")
	fmt.Scan(&id)
	fmt.Println("Masukkan Nama Pesawat:")
	fmt.Scan(&aircraft)
	fmt.Println("Masukkan Asal:")
	fmt.Scan(&origin)
	fmt.Println("Masukkan Tujuan:")
	fmt.Scan(&destination)
	fmt.Println("Masukkan Harga:")
	fmt.Scan(&price)
	fmt.Println("Masukkan Tahun Keberangkatan:")
	fmt.Scan(&year)
	fmt.Println("Masukkan Bulan Keberangkatan:")
	fmt.Scan(&month)
	fmt.Println("Masukkan Hari Keberangkatan:")
	fmt.Scan(&day)

	addFlight(flights, flightCount, id, aircraft, origin, destination, price, year, month, day)
}

// Fungsi Main
func main() {
	var pilihan int
	var flights tabFlight
	var flightCount int

	for {
		fmt.Println("\n=======================================================================")
		fmt.Println(">>>>>>>>>>>>>>>> APLIKASI INFORMASI DATA PENERBANGAN <<<<<<<<<<<<<<<<<<")
		fmt.Println("=======================================================================")
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Data Penerbangan")
		fmt.Println("2. Cari Penerbangan Berdasarkan Tujuan")
		fmt.Println("3. Cari Penerbangan Berdasarkan Asal")
		fmt.Println("4. Ubah Data Penerbangan")
		fmt.Println("5. Hapus Data Penerbangan")
		fmt.Println("6. Tampilkan Semua Penerbangan Terurut Berdasarkan Tanggal Keberangkatan")
		fmt.Println("7. Pesan Tiket Pesawat")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			inputFlightData(&flights, &flightCount)
		case 2:
			var destination string
			fmt.Println("Masukkan Tujuan:")
			fmt.Scan(&destination)
			clearScreen()
			searchFlightsByDestination(&flights, flightCount, destination)
		case 3:
			var origin string
			fmt.Println("Masukkan Asal:")
			fmt.Scan(&origin)
			clearScreen()
			searchFlightsByOrigin(&flights, flightCount, origin)
		case 4:
			clearScreen()
			var id int
			var newAircraft, newOrigin, newDestination string
			var newPrice float64
			var year, month, day int

			fmt.Println("Masukkan ID Penerbangan yang Ingin Diubah:")
			fmt.Scan(&id)
			fmt.Println("Masukkan Nama Pesawat Baru:")
			fmt.Scan(&newAircraft)
			fmt.Println("Masukkan Asal Baru:")
			fmt.Scan(&newOrigin)
			fmt.Println("Masukkan Tujuan Baru:")
			fmt.Scan(&newDestination)
			fmt.Println("Masukkan Harga Baru:")
			fmt.Scan(&newPrice)
			fmt.Println("Masukkan Tahun Keberangkatan Baru:")
			fmt.Scan(&year)
			fmt.Println("Masukkan Bulan Keberangkatan Baru:")
			fmt.Scan(&month)
			fmt.Println("Masukkan Hari Keberangkatan Baru:")
			fmt.Scan(&day)

			editFlight(&flights, flightCount, id, newAircraft, newOrigin, newDestination, newPrice, year, month, day)
		case 5:
			clearScreen()
			var id int
			displayFlights(&flights, flightCount)
			fmt.Println("Masukkan ID Penerbangan yang Ingin Dihapus:")
			fmt.Scan(&id)
			deleteFlight(&flights, &flightCount, id)
		case 6:
			clearScreen()
			sortFlightsByDepartureDateAscending(&flights, flightCount)
			displayFlights(&flights, flightCount)
		case 7:
			clearScreen()
			var origin, destination string
			var year, month, day int
			var temp tabFlight
			var n int
			var id int
			fmt.Println("Masukkan Asal:")
			fmt.Scan(&origin)
			fmt.Println("Masukkan Tujuan:")
			fmt.Scan(&destination)
			fmt.Println("Masukkan Tahun:")
			fmt.Scan(&year)
			fmt.Println("Masukkan Bulan:")
			fmt.Scan(&month)
			fmt.Println("Masukkan Hari:")
			fmt.Scan(&day)
			clearScreen()
			searchFlightsOrder(&flights, flightCount, origin, destination, year, month, day, &temp, &n)
			fmt.Print("Pilih Tiket Penerbangan (Ketik ID): ")
			fmt.Scan(&id)
			idx := findTicketById(&temp, n, id)
			if idx != -1 {
				fmt.Printf("Tiket berhasil dipesan\n")
			}
			if idx == -1 {
				fmt.Println("Tiket tidak tersedia.\n")
			}
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
