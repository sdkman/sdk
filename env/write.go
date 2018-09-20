package env

import "os"

func WriteVersion(version string, file string) {

	//open existing file in write-only append mode
	f, err := os.OpenFile(file, os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	//clobber default version with `version`
	_, err = f.WriteAt([]byte(version), 0)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
