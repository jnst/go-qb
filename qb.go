package qb

import (
	"image/color"
	"log"
	"os"
)

// Version represents Qubicle Binary file version.
type Version struct {
	Major   int8
	Minor   int8
	Release int8
	Build   int8
}

type ZAxisOrientation int

const (
	LeftHanded ZAxisOrientation = iota
	RightHanded
)

// Header represents Qubicle Binary file header.
type Header struct {
	Version          Version
	ColorFormat      color.RGBA
	ZAxisOrientation ZAxisOrientation
}

type QB struct {
	Version               uint32
	ColorFormat           uint32
	ZAxisOrientation      uint32
	Compressed            uint32
	VisibilityMaskEncoded uint32
	NumMatrices           uint32
	F                     *os.File
	I                     uint32
	J                     uint32
	X                     uint32
	Y                     uint32
	Z                     uint32
	SizeX                 uint32
	SizeY                 uint32
	SizeZ                 uint32
	PosX                  int32
	PosY                  int32
	PosZ                  int32
	//Matrix array
	//MatrixList array
	Index uint32
	Data  uint32
	Count uint32
}

const CODEFLAG = 2
const NEXTSLICEFLAG = 6

func Parse(filename string) {
	f, err := os.Open("~/Desktop/castle_lv1.qb")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := make([]byte, 4)
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	// version
	version := uint32(b)
	log.Println(version)

	//

}
