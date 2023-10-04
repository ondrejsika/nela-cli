package nela

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math/rand"

	"github.com/qeesung/image2ascii/convert"

	_ "image/jpeg"
)

func PrintRandomNela() {
	values := []*[]byte{
		&NELA1_JPG,
		&NELA2_JPG,
		&NELA3_JPG,
		&NELA4_JPG,
		&NELA5_JPG,
	}
	randomIndex := rand.Intn(len(values))
	randomChoice := values[randomIndex]
	printRandomImage(*randomChoice)
}

func printRandomImage(data []byte) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}
	converter := convert.NewImageConverter()
	fmt.Print(converter.Image2ASCIIString(img, &convert.DefaultOptions))
}
