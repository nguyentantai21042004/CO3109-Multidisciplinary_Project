package recognize

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const shopID = "c4596303-de42-424b-afcb-ea5be63ab060"

// Helper functions
func DownloadFile(url string) (*os.File, error) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "image-*.jpg")
	if err != nil {
		return nil, err
	}

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Copy the body to the file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	// Seek to the beginning of the file
	_, err = tempFile.Seek(0, 0)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	return tempFile, nil
}

func CreateSaveImageRequest(file *os.File, userID string) (*http.Request, error) {
	// Create a new file buffer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form file field
	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create the request
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:6000/ai/save/%s/%s", shopID, userID), &requestBody)
	if err != nil {
		return nil, err
	}

	// Set the content type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func SendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func ReadResponseBody(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}
	return body
}

// FindImages sends a request to find images for a specific shop
func CreateFindImagesRequest(file *multipart.FileHeader, shopID string) (*http.Request, error) {
	// Create a new file buffer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form file field
	part, err := writer.CreateFormFile("image", filepath.Base(file.Filename))
	if err != nil {
		return nil, err
	}

	// Copy the file content to the form field
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	_, err = io.Copy(part, fileContent)
	if err != nil {
		return nil, err
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create the request
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:6000/ai/find/%s", shopID), &requestBody)
	if err != nil {
		return nil, err
	}

	// Set the content type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func SendFindImagesRequest(req *http.Request) (string, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrUserNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
