package main

import "fmt"

// Tipe bentukan pegawai dengan komponen (field) nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama       string
	gajiPokok  int
	masaKerja  int
	besarBonus int
}

func main() {
	// Deklarasi variabel bertipe pegawai
	var p pegawai

	// Baca data pegawai
	fmt.Println("Masukkan nama, gaji pokok, dan masa kerja pegawai:")
	fmt.Scan(&p.nama, &p.gajiPokok, &p.masaKerja)

	// Hitung bonus dengan memanggil prosedur hitung_bonus
	hitungBonus(&p)

	// Cetak besar bonus
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %d\n", p.nama, p.besarBonus)
}

func hitungBonus(p *pegawai) {
	/* IS: p.nama, p.gaji_pokok, p.masa_kerja terdefinisi
	   Proses: Besar bonus dihitung dengan mengalikan masa kerja dengan angka pengali
	           Jika masa kerja minimal 10 tahun, angka pengalinya 1.5
	           jika masa kerja di bawah 10 tahun hingga 5 tahun, angka pengalinya 0.75
			   di bawah 5 tahun, angka pengalinya 0.5
	   FS: p.besar_bonus berisi nilai
	*/
	switch {
	case p.masaKerja >= 10:
		p.besarBonus = int(1.5 * float64(p.gajiPokok))
	case p.masaKerja >= 5:
		p.besarBonus = int(0.75 * float64(p.gajiPokok))
	default:
		p.besarBonus = int(0.5 * float64(p.gajiPokok))
	}
}
