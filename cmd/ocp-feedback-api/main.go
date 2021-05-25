package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	consumerTimeout, sleepTime := time.Second*2, time.Second*1
	sc := &SleepyConsumer{os.Stdout, consumerTimeout}
	if err := ObserveFile("some_file.txt", sc, sleepTime); err != nil {
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
func ObserveFile(fileName string, consumer *SleepyConsumer, timeout time.Duration) error {
	if consumer == nil {
		panic("consumer can't be nil")
	}

	readFile := func() error {
		return readFileContents(fileName, consumer, timeout)
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
// timeout is used to control amount of time given to consume the data
func readFileContents(fileName string, writer io.Writer, timeout time.Duration) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	totalSize := 0
	for leave, timeout_guard := false, time.After(timeout); !leave; {
		select {
		case <-timeout_guard:
			leave = true // consumer timed out
		default:
			chunkSize, err := writer.Write(data[totalSize:])
			if err != nil {
				return err
			}
			if totalSize += chunkSize; totalSize >= len(data) {
				return nil // consumer consumed all the data
			}
		}
	}
	return errors.New("time is up")
}
