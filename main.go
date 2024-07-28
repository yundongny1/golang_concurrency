package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strings"
	"time"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func loadImageIndependent(path string) {
	// For each input path create a job and add it to
	// the out channel
	job := Job{InputPath: path,
		OutPath: strings.Replace(path, "images/", "images/output/", 1)}
	job.Image = imageprocessing.ReadImage(path)
	job.Image = imageprocessing.Resize(job.Image)
	job.Image = imageprocessing.Grayscale(job.Image)
	imageprocessing.WriteImage(job.OutPath, job.Image)
	fmt.Println("Success")
}

func main() {
	fmt.Println("Input 1 for concurrency, 2 for non-concurrency.")
	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	// var then variable name then variable type
	var input int16

	// Taking input from user
	fmt.Scanln(&input)

	// Print function is used to
	// display output in the same line
	if input == 1 {
		start := time.Now()

		fmt.Println("Process is running with concurrency.")

		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if success {
				fmt.Println("Success!")
			} else {
				fmt.Println("Failed!")
			}
		}

		// End time
		end := time.Now()
		// Calculate the duration
		duration := end.Sub(start)
		// Print the throughput time
		fmt.Printf("Throughput time: %v\n", duration)

	} else if input == 2 {
		start := time.Now()

		fmt.Println("Process is running without concurrency.")

		for _, i := range imagePaths {
			loadImageIndependent(i)
		}

		// End time
		end := time.Now()
		// Calculate the duration
		duration := end.Sub(start)
		// Print the throughput time
		fmt.Printf("Throughput time: %v\n", duration)
	} else {
		fmt.Println("There are only two input options: 1 for concurrency, 2 for independent.")
	}
}
