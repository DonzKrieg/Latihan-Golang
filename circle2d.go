package main

import "fmt"

func main() {
	var r, s, ll, lp, kl, kp, tl, tk float64
	var ra, sa []float64
	for {
		fmt.Scanln(&r, &s)
		if r == 0 && s == 0 {
			break
		}
		ra = append(ra, r)
		sa = append(sa, s)
	}
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	for i := 0; i < len(ra); i++ {
		hitungLuasKelilingLingkaran(ra[i], &ll, &kl)
		hitungLuasKelilingPersegi(sa[i], &lp, &kp)
		hitungTotal(ll, lp, kl, kp, &tl, &tk)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", ra[i], sa[i], ll, lp, kl, kp, tl, tk)
	}
}

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	*l = 3.14 * r * r
	*k = 2 * 3.14 * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(ll, lp, kl, kp float64, totLuas, totKel *float64) {
	*totLuas = ll + lp
	*totKel = kl + kp
}
