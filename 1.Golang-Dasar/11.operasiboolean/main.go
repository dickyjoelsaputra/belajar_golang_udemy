package main

func main() {
	var nilaiAkhir = 70
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusNilaiAkhir || lulusAbsensi
	println(lulus)
}