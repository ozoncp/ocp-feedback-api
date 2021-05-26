package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	sleepTime := time.Second * 1
	sc := &SleepyConsumer{os.Stdout, sleepTime}
	if err := ObserveFile("some_file.txt", sc); err != nil {
		log.Fatal(err)
	}
}

// Sleeper is an interface that represents operation delay
type Sleeper interface {
	Sleep()
}

// SleepyConsumer consumes bytes with a defined delay
type SleepyConsumer struct {
	io.Writer
	sleepDuration time.Duration
}

// Sleep pauses execution for a certain period of time
func (c *SleepyConsumer) Sleep() {
	time.Sleep(c.sleepDuration)
}

// Write wraps around Writer's write method and adds some detailed output
func (c *SleepyConsumer) Write(p []byte) (n int, err error) {
	fmt.Printf("Consumed %v bytes from file. Contents are: \n", len(p))
	return c.Writer.Write(p)
}

// ObserveFile opens file and keeps reading its contents over and over
// Every time file contents are read, consumer will be notified
// Sleep method is used to add a delay between file reads
func ObserveFile(fileName string, consumer *SleepyConsumer) error {
	if consumer == nil {
		panic("consumer can't be nil")
	}

	readFile := func() error {
		return readFileContents(fileName, consumer)
	}

	for {
		if err := readFile(); err != nil {
			return fmt.Errorf("unable to read file %v: %v", fileName, err)
		}
		consumer.Sleep()
	}
}

// readFileContents opens file, reads its contents and passes the contents to the consumer
// if consumer is busy, function will try to wait until all chunks are consumed
func readFileContents(fileName string, writer io.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
