package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func main() {
	var fileUrl string
	fmt.Print("Enter url:")
	fmt.Fscan(os.Stdin, &fileUrl)
	//fileUrl := "https://lapkins.ru/upload/uf/b50/b50d4643a4f0e08fe678a227fbd49bdb.jpg"
	err := DownloadFile("profile.webp", fileUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Downloaded: " + fileUrl)

	fmt.Println("https://yandex.ru/images/search?rpt=imageview&url="+fileUrl)
}
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}