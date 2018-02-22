package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	uncrypto, err := rsa.NewReader(conn)
	uncompress, err := gzip.NewReader(uncrypto)
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(uncompress)
	for {
		hdr, err := tr.Next()
		if err != nil {
			return
		}
		fmt.Println(hdr.Name)
		io.Copy(ioutil.Discard, tr)
	}

}
