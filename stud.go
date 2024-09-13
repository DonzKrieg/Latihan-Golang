Program Student
Kamus
    type Student <
        Name : string
        Age : integer
        Grade : array of integer with size 3
    >
    student1 : Student
Algoritma
    Output ("Masukkan nama mahasiswa:")
    Input (student1.Name)
    Output ("Masukkan usia mahasiswa:")
    Input (student1.Age)
    Output ("Masukkan nilai-nilai Grade (sebanyak 6):")
    For i <- 0 to 2 do
        Output ("Nilai Grade ", i+1, ": ")
        Input (student1.Grade[i])
    End For
    Output ("Tabel Data Mahasiswa:")
    Output ("Nama:", student1.Name)
    Output ("Umur:", student1.Age)
    For i from 0 to 2 do
        Output ("Grade ", i+1, ": ", student1.Grade[i])
    End For
Endprogram

package main

import "fmt"

// Definisikan struct untuk student
type Student struct {
	Name  string
	Age   int
	Grade [3]int
}

func main() {
	var student1 Student

	// Mengisi data mahasiswa
	fmt.Println("Masukkan nama mahasiswa:")
	fmt.Scanln(&student1.Name)
	fmt.Println("Masukkan usia mahasiswa:")
	fmt.Scanln(&student1.Age)
	fmt.Println("Masukkan nilai-nilai Grade (sebanyak 6):")
	for i := 0; i < len(student1.Grade); i++ {
		fmt.Printf("Nilai Grade %d: ", i+1)
		fmt.Scanln(&student1.Grade[i])
	}
	// Menampilkan data mahasiswa
	fmt.Println("Tabel Data Mahasiswa:")
	fmt.Println("Nama:", student1.Name)
	fmt.Println("Umur:", student1.Age)
	for i := 0; i < len(student1.Grade); i++ {
		fmt.Printf("Grade %d: %d\n", i+1, student1.Grade[i])
	}
}
