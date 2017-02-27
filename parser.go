package qb

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
)

// Parser presents Qubicle Binary file parser.
type Parser struct {
}

// New parser.
func New() *Parser {
	return &Parser{}
}

func (p Parser) Parse() {
	f, err := os.Open("/Users/jnst/Desktop/castle_lv1.qb")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := p.Next(f, 4)

	var major uint8
	err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, &major)
	if err != nil {
		log.Fatal(err)
	}

}

// Next returns next n bytes in file.
func (p Parser) Next(f *os.File, n int) []byte {
	buf := make([]byte, n)

	_, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	return buf
}
