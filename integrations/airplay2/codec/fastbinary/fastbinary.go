package fastbinary

import (
	"encoding/binary"
	"io"
)

const (
	Int16DataSize = 2
)

func ReadInt16(r io.Reader, data *int16) error {
	bs := make([]byte, Int16DataSize)
	if _, err := io.ReadFull(r, bs); err != nil {
		return err
	}
	*data = int16(binary.LittleEndian.Uint16(bs))
	return nil
}

func WriteInt16(w io.Writer, data int16) error {
	bs := make([]byte, Int16DataSize)
	binary.LittleEndian.PutUint16(bs, uint16(data))
	_, err := w.Write(bs)
	return err
}
