package main

import (
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"
)

func textToBinary(text string) string {
	var binaryStrings []string
	for _, rune := range text {
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", rune))
	}
	return strings.Join(binaryStrings, " ")
}

func binaryToText(binary string) (string, error) {
	var text strings.Builder
	for i := 0; i < len(binary); i += 8 {
		if i+8 > len(binary) {
			break
		}
		bin := binary[i : i+8]
		num, err := strconv.ParseInt(bin, 2, 64)
		if err != nil {
			return "", err
		}
		text.WriteRune(rune(num))
	}
	return text.String(), nil
}

func binaryToImage(binary string) *image.Gray {
	binary = strings.ReplaceAll(binary, " ", "")
	width := len(binary)

	img := image.NewGray(image.Rect(0, 0, width, 1))
	for x, char := range binary {
		if char == '1' {
			img.SetGray(x, 0, color.Gray{Y: 0}) // Black
		} else {
			img.SetGray(x, 0, color.Gray{Y: 255}) // White
		}
	}
	return img
}

func imageToBinary(img *image.Gray) string {
	var sb strings.Builder
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if img.GrayAt(x, y).Y < 128 {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
	}
	return sb.String()
}
