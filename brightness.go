package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInt(file *os.File) (number int, err error) {
	// Array to store the data in...
	data := make([]byte, 100)

	bytesRead, err := file.Read(data)
	if err != nil {
		return 0, err
	}
	//parse the int and cut off the \n at the end
	number, err = strconv.Atoi(string(data[:bytesRead-1]))
	return number, err
}

func main() {
	// Filename to read/write brightness from
	filename := "/sys/class/backlight/intel_backlight/brightness"
	// How much to raise/lower brightness.
	factor := 1.5
	//Minimum brightness setting. Below here my display goes dark
	min := 6
	//path to maximum brightness file
	max_filename := "/sys/class/backlight/intel_backlight/max_brightness"

	max_file, err := os.Open(max_filename)
	if err != nil {
		log.Fatal(err)
	}
	max, err := readInt(max_file)
	defer max_file.Close()

	//make sure we have the right args.
	if len(os.Args) < 2 {
		log.Fatal("Must specify 'up' or 'down'")
	}

	//open the file in read/write mode
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	brightness, err := readInt(file)
	if err != nil {
		log.Fatal(err)
	}

	//compute the new brightness
	if os.Args[1] == "up" {
		brightness = int(float64(brightness)*factor + 1)
		if brightness > max {
			brightness = max
		}
	} else if os.Args[1] == "down" {
		brightness = int(float64(brightness) / factor)
		if brightness < min {
			brightness = min
		}
	} else if os.Args[1] == "print" {
		fmt.Println(brightness)
	} else {
		log.Fatal("Invalid argument")
	}
	strBright := fmt.Sprintf("%d\n", brightness)

	//write the new brightness back out
	file.Seek(0, 0)
	_, err = file.Write([]byte(strBright))
	if err != nil {
		log.Fatal(err)
	}

}
