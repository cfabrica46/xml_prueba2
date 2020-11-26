package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type instrument struct {
	Name     string   `json:"title"`
	Brand    string   `json:"Brand"`
	Model    string   `json:"Model"`
	Electric bool     `json:"Electric"`
	Color    string   `json:"Color"`
	Strings  int      `json:"Strings"`
	Notes    []string `json:"Notes"`
}

func main() {

	instruments := []instrument{
		instrument{
			Name:     "Guitarra",
			Brand:    "Fender",
			Model:    "Affinity",
			Electric: true,
			Color:    "Negro",
			Strings:  6,
			Notes:    []string{"E", "B", "G", "D", "A", "E"},
		},
		instrument{
			Name:     "Ukulele",
			Brand:    "Fender",
			Model:    "SeaSide",
			Electric: false,
			Color:    "Cedro",
			Strings:  4,
			Notes:    []string{"A", "E", "C", "G"},
		},
		instrument{
			Name:     "Bass",
			Brand:    "Gio",
			Model:    "Ibanez",
			Electric: true,
			Color:    "Azul Electrico",
			Strings:  4,
			Notes:    []string{"G", "D", "A", "E"},
		},
	}

	data, err := xml.MarshalIndent(instruments, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	archivo, err := os.OpenFile("databases.xml", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	archivo.Write(data)

	instruments = append(instruments, instruments[0])

	data, err = xml.MarshalIndent(instruments, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	archivo.Seek(0, 0)

	_, err = archivo.Write(data)

	if err != nil {
		log.Fatal(err)
	}

	contenido := []byte{}

	buf := make([]byte, 50)

	nuevo := []instrument{}

	fmt.Println(instruments)

	_, err = archivo.Seek(0, 0)

	if err != nil {
		log.Fatal(err)
	}

	for {

		n, err := archivo.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		contenido = append(contenido, buf[:n]...)

	}

	err = xml.Unmarshal(contenido, &nuevo)

	fmt.Println(string(contenido))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(instruments[1])
}
