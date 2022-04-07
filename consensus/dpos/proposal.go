package dpos

import (
	"bytes"
	"encoding/binary"
	"errors"
	_ "fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	AnkrProposal uint8 = iota + 1
)

type Proposal struct {
	Id               uint8
	Values           []interface{}
	Description      string
	ValidateValuesFn func(uint8, []interface{}) error
	ValidateBytesFn  func(common.Hash) error
	ToBytesFn        func([]interface{}) []byte
	FromBytesFn      func(common.Hash) []interface{}
}

var Proposals map[uint8]*Proposal = map[uint8]*Proposal{
	AnkrProposal: {
		Id:          AnkrProposal,
		Values:      make([]interface{}, 0),
		Description: "proposal by dpos",

		ValidateValuesFn: func(id uint8, values []interface{}) error {
			value := values[0].(uint8)

			if !(value > 0 && value <= 255) {
				return errors.New("Invalid proposal#" + string(id))
			}

			return nil
		},

		ValidateBytesFn: func(_bytes common.Hash) error {
			value := _bytes[1]

			if !(value > 0 && value <= 255) {
				return errors.New("Invalid proposal#" + string(_bytes[0]))
			}

			if !bytes.Equal(_bytes[2:], bytes.Repeat([]byte{0x00}, common.HashLength-2)) {
				return errors.New("Invalid proposal#" + string(_bytes[0]))
			}

			return nil
		},

		ToBytesFn: func(values []interface{}) []byte {
			value := values[0].(uint8)

			buf := bytesBufferPool.Get().(*bytes.Buffer)
			defer bytesBufferPool.Put(buf)

			binary.Write(buf, binary.BigEndian, value)

			return buf.Bytes()
		},

		FromBytesFn: func(bytes common.Hash) []interface{} {
			return []interface{}{bytes[1]}
		},
	},
}

func getProposal(id uint8) (*Proposal, error) {
	proposal, ok := Proposals[id]

	if ok {
		return &(*proposal), nil //new
	} else {
		return &Proposal{}, errors.New("Proposal not found")
	}
}

func (this *Proposal) toBytes() (common.Hash, error) {

	if err := this.ValidateValuesFn(this.Id, this.Values); err != nil {
		return common.Hash{}, err
	}

	result := []byte{uint8(this.Id)}
	result = append(result, this.ToBytesFn(this.Values)...)
	result = append(result, bytes.Repeat([]byte{0x00}, common.HashLength-len(result))...)

	return common.BytesToHash(result), nil
}

func (this *Proposal) fromBytes(proposalBytes common.Hash) error {

	id := proposalBytes[0]

	proposal, err := getProposal(id)

	if err == nil {

		this.ValidateBytesFn = proposal.ValidateBytesFn

		if err := this.ValidateBytesFn(proposalBytes); err != nil {
			return err
		}

		this.Id = proposal.Id
		this.Description = proposal.Description
		this.ValidateValuesFn = proposal.ValidateValuesFn
		this.ToBytesFn = proposal.ToBytesFn
		this.FromBytesFn = proposal.FromBytesFn
		this.Values = this.FromBytesFn(proposalBytes)

	} else {
		return errors.New("Proposal not found")
	}

	return nil
}
