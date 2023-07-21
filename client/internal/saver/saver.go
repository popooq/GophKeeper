package saver

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Saver struct {
	file       *os.File
	readwriter *bufio.ReadWriter
}

// New функция создает новый Saver
func New(storefile string) *Saver {
	file, err := os.OpenFile(storefile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o777)
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)

	return &Saver{
		file:       file,
		readwriter: bufio.NewReadWriter(reader, writer),
	}
}

func (s *Saver) SaveJWT(jwt []byte) error {
	_, err := s.readwriter.Write(jwt)
	if err != nil {
		return err
	}
	return s.readwriter.Flush()
}

func (s *Saver) LoadJWT() ([]byte, error) {
	if s.readwriter == nil {
		return nil, nil
	}
	data, err := io.ReadAll(s.readwriter)
	if err != nil {
		log.Printf("read err : %s", err)
		return nil, err
	}
	return data, nil
}
