package main

func main() {
	type NoKTP string
	type Married bool

	var noKtpDicky NoKTP = "123456789"
	var marriedStatus Married = false
	println(noKtpDicky)
	println(NoKTP("1234231"))
	println(marriedStatus)
}