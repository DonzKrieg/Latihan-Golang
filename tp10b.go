package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0

	bacaData(&timnas, &nPemain)
	cetakPemain(timnas, nPemain)
	cetakNamaPemainTertinggi(timnas, nPemain)
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var p pemain

	for i := 0; i < NMAX; i++ {
		fmt.Scanln(&p.nama, &p.nomorPunggung, &p.posisi, &p.tinggi)
		if p.nama == "none" {
			break
		}
		A[i] = p
		*n++
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	fmt.Printf("Data Pemain: \n")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %d\n", A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var t pemain = A[0]
	for i := 0; i < n; i++ {
		if t.tinggi < A[i].tinggi {
			t = A[i]
		}
	}
	fmt.Printf("Pemain dengan badan tertinggi: %s\n", t.nama)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var t pemain = A[0]
	for i := 0; i < n; i++ {
		if t.tinggi > A[i].tinggi {
			t = A[i]
		}
	}
	fmt.Printf("Pemain dengan badan terendah: %s", t.nama)
}
