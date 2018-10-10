package cmd

func check(e error) {
	if e != nil {
		panic(e)
	}
}
