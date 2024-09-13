package main

func main() {

	profile := map[string]string{

		"Nama":   "Muhammad Ihsan Romdhon",
		"Umur":   "19",
		"Asrama": "Gedung 9 No. 218",
	}

	// fmt.Println(profile)
	// fmt.Println(profile["Nama"])

	variable := make(map[string]string)

	variable["p"] = "Andi seorang mahasiswa Tel-U"
	variable["q"] = "Andi mengerjakan tugasnya dengan baik"
	variable["r"] = "Andi anak yang rajin"

	// fmt.Println(variable["q"])

	noKtp := make(map[string]int)

	noKtp["ferdi"] = 10302233546
	noKtp["bena"] = 10302233435
	noKtp["bijer"] = 10302233123

	// fmt.Println(noKtp)
	// fmt.Println(noKtp["bena"])

}
