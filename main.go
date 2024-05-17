package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func createUI() {
	a := app.New()
	w := a.NewWindow("Bit Display")
	w.Resize(fyne.NewSize(400, 200))

	textEntry := widget.NewEntry()
	textEntry.SetPlaceHolder("Enter text here")

	imageEntry := widget.NewEntry()
	imageEntry.SetPlaceHolder("Enter image file name here")

	outputButton := widget.NewButton("Output Image", func() {
		text := textEntry.Text
		if text == "" {
			dialog.ShowError(errors.New("No text entered"), w)
			return
		}

		img := binaryToImage(textToBinary(text))

		filePath := filepath.Join(".", "output.png")
		file, err := os.Create(filePath)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		dialog.ShowInformation("Success", "Image has been outputted to "+filePath, w)
	})
	inputButton := widget.NewButton("Decode Image", func() {
		imageName := imageEntry.Text
		if imageName == "" {
			dialog.ShowError(errors.New("no file name entered"), w)
			return
		}

		filePath := filepath.Join(".", imageName)
		file, err := os.Open(filePath)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		defer file.Close()

		// Decode the image file
		img, err := png.Decode(file)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		// Convert image.Image to *image.Gray
		grayImg, ok := img.(*image.Gray)
		if !ok {
			bounds := img.Bounds()
			grayImg = image.NewGray(bounds)
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					grayImg.Set(x, y, img.At(x, y))
				}
			}
		}

		binaryData, _ := binaryToText(imageToBinary(grayImg))
		textEntry.SetText(binaryData)
		dialog.ShowInformation("Success", "Copied to the text field!", w)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Text to encode:", Widget: textEntry},
			{Text: "Image to decode:", Widget: imageEntry},
			{Widget: container.NewHBox(outputButton, inputButton)},
		},
	}

	w.SetContent(form)
	w.ShowAndRun()
}

func main() {
	createUI()
}
