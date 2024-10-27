package utils

import "os"

func WriteHTMLToFile(html string, filename string) error {
	// Create or overwrite the HTML file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the HTML content to the file
	_, err = file.WriteString(html)
	return err
}
