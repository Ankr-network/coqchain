package structjson

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"testing"

	"github.com/Ankr-network/coqchain/rlp"
)

type Student struct {
	Name string
	Age  uint16
}

func (s *Student) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	fieldOne := len(s.Name)
	b16 := make([]byte, 2)
	binary.BigEndian.PutUint16(b16, uint16(fieldOne))
	buf.Write(b16)
	buf.WriteString(s.Name)
	binary.BigEndian.PutUint16(b16, 2)
	buf.Write(b16)
	binary.BigEndian.PutUint16(b16, s.Age)
	buf.Write(b16)
	return buf.Bytes(), nil
}

func (s *Student) UnmarshalJSON(d []byte) error {
	var (
		stx    uint16 = 0
		offset uint16 = 2
	)
	act := binary.BigEndian.Uint16(d[stx:offset])
	stx = offset
	offset = stx + act
	s.Name = string(d[stx:offset])
	stx = offset
	offset += 2
	act = binary.BigEndian.Uint16(d[stx:offset])
	stx = offset
	offset = stx + act
	s.Age = binary.BigEndian.Uint16(d[stx:offset])
	return nil
}

func TestStdJson(t *testing.T) {
	s := Student{"mobus", 32}

	bs, _ := json.Marshal(s)
	t.Log("std json size: ", len(bs))

	t.Log("test custom")
	bs, _ = s.MarshalJSON()
	t.Log("custom json size: ", len(bs))

	so := &Student{}
	so.UnmarshalJSON(bs)
	t.Logf("std name: %s age: %d", so.Name, so.Age)

	buf := bytes.NewBuffer([]byte{})
	err := rlp.Encode(buf, s)
	if err != nil {
		t.Logf("rlp error: %v \n", err)
		return
	}
	t.Logf("rlp size: %d \n", len(buf.Bytes()))
}
