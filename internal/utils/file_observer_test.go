package utils

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
	c.bytes = p
	return len(p), nil
}

// brokenComsumer is a consumer stub that returns an error
type brokenConsumer struct{}

func (c *brokenConsumer) Write(p []byte) (n int, err error) {
	return 1, errors.New("i'm broken :(")
}

func TestReadFile(t *testing.T) {
	t.Run("empty file name", func(t *testing.T) {
		sleepTime := time.Second
		sc := &SleepyConsumer{os.Stdout, sleepTime}
		err := readFileContents("", sc)

		assertNonNilError(t, err)
	})

	t.Run("wrong file name", func(t *testing.T) {
		sleepTime := time.Second
		sc := &SleepyConsumer{os.Stdout, sleepTime}
		err := readFileContents("invalid_filename", sc)

		assertNonNilError(t, err)
	})

	t.Run("nil consumer", func(t *testing.T) {
		fileName, deleteFile := createTempFile(t, "some_file", "")
		defer assertPanic(t)
		defer deleteFile()

		_ = ObserveFile(fileName, nil)
		t.Error("goroutine must enter panic state")
	})

	t.Run("read file contents", func(t *testing.T) {
		consumer := &spyConsumer{}
		file_contents := "file_contents"
		fileName, deleteFile := createTempFile(t, "file_name", file_contents)
		defer deleteFile()
		err := readFileContents(fileName, consumer)

		assertNilError(t, err)

		if !reflect.DeepEqual(consumer.bytes, []byte(file_contents)) {
			t.Errorf("got %v want %v", consumer.bytes, []byte(file_contents))
		}
	})

	t.Run("consumer error", func(t *testing.T) {
		consumer := &brokenConsumer{}
		file_contents := "file_contents"
		fileName, deleteFile := createTempFile(t, "file_name", file_contents)
		defer deleteFile()
		err := readFileContents(fileName, consumer)

		assertNonNilError(t, err)
	})
}

func createTempFile(t *testing.T, fileName string, initialData string) (string, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", fileName)
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	if _, err = tmpfile.Write([]byte(initialData)); err != nil {
		t.Fatalf("unsable to write into temp file: %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile.Name(), removeFile
}
