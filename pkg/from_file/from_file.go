package from_file

import (
	"fmt"

	"github.com/qeesung/image2ascii/convert"

	_ "image/jpeg"
	_ "image/png"
)

func PrintFromFile(filePath string) {
	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(filePath, &convert.DefaultOptions))
}
