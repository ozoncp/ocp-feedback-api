package main

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

// spyConsumer is a consumer stub that allows to examine consumed data
type spyConsumer struct {
	bytes []byte
}

func (c *spyConsumer) Write(p []byte) (n int, err error) {
	consumedElements := 1 // consume 1 element at a time
	c.bytes = append(c.bytes, p[:consumedElements]...)
	return consumedElements, nil
}

// busyConsumer is a consumer stub that imitates non-responding consumer
type busyConsumer struct{}

func (c *busyConsumer) Write(p []byte) (n int, err error) {
	return 0, nil
}

// brokenComsumer is a consumer stub that returns an error
type brokenConsumer struct{}

func (c *brokenConsumer) Write(p []byte) (n int, err error) {
	return len(p), errors.New("i'm broken :(")
}

func TestReadFile(t *testing.T) {
	t.Run("empty file name", func(t *testing.T) {
		sleepTime, consumerTimeout := time.Second, time.Second
		sc := &SleepyConsumer{os.Stdout, sleepTime}
		err := readFileContents("", sc, consumerTimeout)

		assertNonNilError(t, err)
	})

	t.Run("wrong file name", func(t *testing.T) {
		sleepTime, consumerTimeout := time.Second, time.Second
		sc := &SleepyConsumer{os.Stdout, sleepTime}
		err := readFileContents("invalid_filename", sc, consumerTimeout)

		assertNonNilError(t, err)
	})

	t.Run("nil consumer", func(t *testing.T) {
		fileName, deleteFile := createTempFile(t, "some_file", "")
		defer assertPanic(t)
		defer deleteFile()

		_ = ObserveFile(fileName, nil, time.Second)
		t.Error("goroutine must enter panic state")
	})

	t.Run("read file contents", func(t *testing.T) {
		consumer := &spyConsumer{}
		file_contents := "file_contents"
		fileName, deleteFile := createTempFile(t, "file_name", file_contents)
		defer deleteFile()
		err := readFileContents(fileName, consumer, time.Second*10)

		assertNilError(t, err)

		if !reflect.DeepEqual(consumer.bytes, []byte(file_contents)) {
			t.Errorf("got %v want %v", consumer.bytes, []byte(file_contents))
		}
	})

	t.Run("consumer timeout", func(t *testing.T) {
		consumer := &busyConsumer{}
		file_contents := "file_contents"
		fileName, deleteFile := createTempFile(t, "file_name", file_contents)
		defer deleteFile()
		err := readFileContents(fileName, consumer, time.Millisecond)

		assertNonNilError(t, err)
	})

	t.Run("consumer error", func(t *testing.T) {
		consumer := &brokenConsumer{}
		file_contents := "file_contents"
		fileName, deleteFile := createTempFile(t, "file_name", file_contents)
		defer deleteFile()
		err := readFileContents(fileName, consumer, time.Second*10)

		assertNonNilError(t, err)
	})
}

func createTempFile(t *testing.T, fileName string, initialData string) (string, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", fileName)
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile.Name(), removeFile
}

func assertNonNilError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("error must be returned")
	}
}

func assertNilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("error must not be returned")
	}
}

func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("goroutine must enter panic state")
	}
}
