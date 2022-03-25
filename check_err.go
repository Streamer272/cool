package main

func CheckErr(err interface{}) {
	if err != nil {
		panic(err)
	}
}
