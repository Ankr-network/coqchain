package dpos

import (
	"bytes"
	"encoding/binary"
	"errors"
	_ "fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	//action type
	becomeCandidate uint8 = iota + 1
	becomeDelegator
	quitCandidate
	quitDelegator
)

// dpos specify contract address for become cadidate/delegator tx
var (
	contractAddress = common.HexToAddress("0x00000000000000000000000000000000416e6b72")
)

type Action struct {
	Id               uint8
	Values           []interface{}
	Description      string
	ValidateValuesFn func(uint8, []interface{}) error
	ValidateBytesFn  func([]byte) error
	ToBytesFn        func([]interface{}) []byte
	FromBytesFn      func([]byte) []interface{}
}

var Actions map[uint8]*Action = map[uint8]*Action{
	becomeCandidate: {
		Id:          becomeCandidate,
		Values:      make([]interface{}, 0),
		Description: "Register to become a candidate",

		ValidateValuesFn: func(id uint8, values []interface{}) error {
			return nil
		},

		ValidateBytesFn: func(_bytes []byte) error {
			if len(_bytes) != 1 {
				return errors.New("Invalid action#" + string(_bytes[0]))
			}
			return nil
		},

		ToBytesFn: func(values []interface{}) []byte {
			value := values[0].(uint8)

			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, value)

			return buf.Bytes()
		},

		FromBytesFn: func(bytes []byte) []interface{} {
			return []interface{}{}
		},
	},

	becomeDelegator: {
		Id:          becomeDelegator,
		Values:      make([]interface{}, 0),
		Description: "Register to become a delegator",

		ValidateValuesFn: func(id uint8, values []interface{}) error {

			address := values[0].([]byte)

			if len(address) != common.AddressLength {
				return errors.New("Invalid action#" + string(id))
			}

			return nil
		},

		ValidateBytesFn: func(_bytes []byte) error {

			if len(_bytes) != common.AddressLength+1 {
				return errors.New("Invalid action#" + string(_bytes[0]))
			}

			return nil
		},

		ToBytesFn: func(values []interface{}) []byte {
			value := values[0].(uint8)

			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, value)

			return buf.Bytes()
		},

		FromBytesFn: func(bytes []byte) []interface{} {
			return []interface{}{common.BytesToAddress(bytes[1:])}
		},
	},

	quitCandidate: {
		Id:          quitCandidate,
		Values:      make([]interface{}, 0),
		Description: "To quit as candidate",

		ValidateValuesFn: func(id uint8, values []interface{}) error {
			return nil
		},

		ValidateBytesFn: func(_bytes []byte) error {
			if len(_bytes) != 1 {
				return errors.New("Invalid action#" + string(_bytes[0]))
			}
			return nil
		},

		ToBytesFn: func(values []interface{}) []byte {
			value := values[0].(uint8)

			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, value)

			return buf.Bytes()
		},

		FromBytesFn: func(bytes []byte) []interface{} {
			return []interface{}{}
		},
	},

	quitDelegator: {
		Id:          quitCandidate,
		Values:      make([]interface{}, 0),
		Description: "To quit as candidate",

		ValidateValuesFn: func(id uint8, values []interface{}) error {
			return nil
		},

		ValidateBytesFn: func(_bytes []byte) error {
			if len(_bytes) != 1 {
				return errors.New("Invalid action#" + string(_bytes[0]))
			}
			return nil
		},

		ToBytesFn: func(values []interface{}) []byte {
			value := values[0].(uint8)

			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, value)

			return buf.Bytes()
		},

		FromBytesFn: func(bytes []byte) []interface{} {
			return []interface{}{}
		},
	},
}

func getAction(id uint8) (*Action, error) {
	action, ok := Actions[id]

	if ok {
		return &(*action), nil //new
	} else {
		return &Action{}, errors.New("Action not found")
	}
}

func (self *Action) toBytes() ([]byte, error) {

	if err := self.ValidateValuesFn(self.Id, self.Values); err != nil {
		return []byte{}, err
	}

	result := []byte{uint8(self.Id)}
	result = append(result, self.ToBytesFn(self.Values)...)

	return result, nil
}

func (self *Action) fromBytes(actionBytes []byte) error {

	id := actionBytes[0]

	action, err := getAction(id)

	if err == nil {

		self.ValidateBytesFn = action.ValidateBytesFn

		if err := self.ValidateBytesFn(actionBytes); err != nil {
			return err
		}

		self.Id = action.Id
		self.Description = action.Description
		self.ValidateValuesFn = action.ValidateValuesFn
		self.ToBytesFn = action.ToBytesFn
		self.FromBytesFn = action.FromBytesFn
		self.Values = self.FromBytesFn(actionBytes)

	} else {
		return errors.New("Action not found")
	}

	return nil
}
