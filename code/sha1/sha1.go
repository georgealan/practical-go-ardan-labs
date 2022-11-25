package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("sha1/http.log.gz") // I have to specify a directory folder for work.
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = sha1Sum("sha1/http.log.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

/*
if file names ends with .gz the: $ cat http.log.gz | gunzip | sha1sum
else: $ cat http.log.gz | sha1sum
*/
func sha1Sum(fileName string) (string, error) {
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close() // deferred are called in LIFO order
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gzFile, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gzFile.Close()
		r = gzFile
	}

	//io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
