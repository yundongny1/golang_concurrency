package imageprocessing

import (
	"image"
	"os"
	"testing"
)

func TestReadImage(t *testing.T) {
	path := "example_image_original.jpeg"
	//example image to test
	example_image := ReadImage(path)

	//if the image is not null, pass
	if example_image == nil {
		t.Error(example_image)
	}
}

func TestWriteImage(t *testing.T) {
	input_path := "example_image_original.jpeg"
	output_path := "example_image_test_write.jpeg"

	type Job struct {
		InputPath string
		Image     image.Image
		OutPath   string
	}
	job := Job{InputPath: input_path, OutPath: output_path}

	job.Image = ReadImage("example_image_original.jpeg")

	WriteImage(job.OutPath, job.Image)

	file, err := os.Open("example_image_test_write.jpeg")
	if err != nil {
		t.Error("Error opening image:", err)
	}
	defer file.Close()
}
