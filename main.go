package main

import (
	"log"

	"github.com/MarcelArt/marcel-lib-pack/pkg/objects"
)

type StructType struct {
	Interface any
	Err       error
	Byte      byte
}

type AllType struct {
	I         int `json:"int" validate:"required"`
	I8        int8
	I16       int16
	I32       int32
	I64       int64
	U         uint
	U8        uint8
	U16       uint16
	U32       uint32
	U64       uint64
	String    string
	Str       *string `json:"str"`
	Float     float32
	Double    float64
	Boolean   bool
	Char      rune
	Object    StructType
	Interface any
	Err       error
	Byte      byte
	Binary    []byte
	Map       map[string]float64
}

func main() {
	var s AllType
	s.Float = 1.1

	fieldTypes, err := objects.Fields(s, "json", "validate")
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, fieldType := range fieldTypes {
		log.Println("Field Name:", fieldType.Field, "Field Type:", fieldType.Type)
		log.Println("Field Tags:", fieldType.Tags)
	}

	value, err := objects.Get(s, "Float")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Value:", value)
}
