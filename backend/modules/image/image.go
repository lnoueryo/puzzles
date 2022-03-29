package image

import (
	"backend/modules/crypto"
	"bytes"
	"encoding/base64"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)



func StoreImage(dir string, base64Data string) (string, error) {
	randStr, _ := crypto.MakeRandomStr(20)
	// サーバー側に保存するために空ファイルを作成
    coI := strings.Index(base64Data, ",")
    rawImage := base64Data[coI+1:]

    // Encoded Image DataUrl //
    unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))

    res := bytes.NewReader(unbased)
	var fileName string
	switch strings.TrimSuffix(base64Data[5:coI], ";base64") {
	case "image/png":
		extension := ".png"
		fileName = randStr + extension
		saveImage, err := os.Create("./upload/" + dir + "/" + fileName)
		if err != nil {
			message := "could't upload the file"
			err := errors.New(message)
			return fileName, err
		}
		defer saveImage.Close()
		pngI, _ := png.Decode(res)
        png.Encode(saveImage, pngI)
	case "image/jpeg":
		extension := ".jpeg"
		fileName = randStr + extension
		saveImage, err := os.Create("./upload/" + dir + "/" + fileName)
		if err != nil {
			message := "could't upload the file"
			err := errors.New(message)
			return fileName, err
		}
		defer saveImage.Close()
		jpgI, _ := jpeg.Decode(res)
        jpeg.Encode(saveImage, jpgI, &jpeg.Options{Quality: 75})
	}
	return fileName, nil
}