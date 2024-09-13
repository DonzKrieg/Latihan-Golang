package main

import "fmt"

type time struct {
	hour, min, sec int
}

func main() {
	var in time
	fmt.Scan(&in.hour, &in.min, &in.sec)
	var out time
	fmt.Scan(&out.hour, &out.min, &out.sec)
	lamaPark := hitungLama(in, out)
	fmt.Println(lamaPark.hour, lamaPark.min, lamaPark.sec)
}

func hitungLama(in, out time) time {
	var lama time
	detikLama := (out.hour*3600 + out.min*60 + out.sec) - (in.hour*3600 + in.min*60 + in.sec)

	lama.hour = detikLama / 3600
	lama.min = (detikLama % 3600) / 60
	lama.sec = (detikLama % 3600) % 60

	return lama
}
