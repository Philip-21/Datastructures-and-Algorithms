package main

import (
	"bytes"
	"fmt"
)

//interfaces specifies what method a type should have and the type decides how to implement these methods
//interfaces dont describe data, interfaces  describe behaviours
//only structs cant be used to implement interfaces

//composing interfaces together
//this is one of the keys to scalability,

func Main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("hello all coders ,this is a test")) //converting the string to a byte slice,the result will be printed in 8 character chunks
	wc.Close()                                           // this method is required to print out the characters in 8 chunks
}

type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}

//composing the writer and closer interface(embedding interfaces)
type WriterCloser interface { // the WriterCloser interface will be implemented based on methods in the Writer & Closer interface
	Writer
	Closer
}

//creating a Struct  related to a concept
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) { // implementing the Write method
	n, err := bwc.buffer.Write(data) //when data is passed unto the write method it will store it in the BufferedWriterCloser struct
	if err != nil {
		return 0, err
	} //writing out what gets sent into the bufferCloser and print it out to the console in 8 character increment
	v := make([]byte, 8)
	// if buffer has more than 8 characters the conditions will be executed below
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8) //pulling the next 8characters out writing it to the console
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// a constructor function that is returning a pointer to a bufferedWriterCloser
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}), //initializing the internal buffer to a NewBuffer
	}
}
