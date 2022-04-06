package dpos

import (
	"bytes"
	"encoding/binary"
	_ "errors"
	_ "fmt"
	"io"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

func parseEpochExtra(header *types.Header) ([]common.Address, []*Proposal, [][]ElectedDelegator) {
	extras := unserialize(header.Extra)

	//extract signer
	signers := make([]common.Address, len(extras[1])/common.AddressLength)

	for i := 0; i < len(signers); i++ {
		copy(signers[i][:], extras[1][i*common.AddressLength:])
	}

	//extract proposal
	proposalCnt := len(extras[2]) / common.HashLength
	proposals := make([]*Proposal, 0)

	for i := 0; i < proposalCnt; i++ {
		proposal := &Proposal{}
		proposal.fromBytes(common.BytesToHash(extras[2][i*common.HashLength : (i+1)*common.HashLength]))
		proposals = append(proposals, proposal)
	}

	//extract delegator
	portionLen := 4
	segments := unserialize(extras[3])

	delegatorss := make([][]ElectedDelegator, 0)
	for _, segment := range segments {

		segmentLen := common.AddressLength + portionLen
		delegatorCnt := len(segment) / segmentLen
		delegators := make([]ElectedDelegator, 0)
		for i := 0; i < delegatorCnt; i++ {

			delegator := segment[i*segmentLen : (i+1)*segmentLen]
			address := common.BytesToAddress(delegator[:common.AddressLength])
			portion := math.Float32frombits(binary.BigEndian.Uint32(delegator[common.AddressLength:]))

			delegators = append(delegators, ElectedDelegator{address, portion})
		}

		delegatorss = append(delegatorss, delegators)

	}

	return signers, proposals, delegatorss

}

func RLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)

	return b.Bytes()
}

func unserialize(headerExtra []byte) [][]byte {

	var result [][]byte

	extra := make([]byte, len(headerExtra))
	copy(extra, headerExtra) //new slice

	start := uint64(0)
	limit := uint64(0)

	processLen := true
	extraLen := uint64(len(extra))

	for start < extraLen {
		if processLen {
			if extra[start] == 0xfd {
				limit = uint64(binary.BigEndian.Uint16(extra[start+1 : start+3]))
				start += 3
			} else {
				limit = uint64(extra[start])
				start += 1
			}

			processLen = false
		} else {
			result = append(result, extra[start:start+limit])

			start += limit
			processLen = true
		}
	}

	return result
}

func VarIntToBytes(item []byte) []byte {

	buf := new(bytes.Buffer)
	prefix := make([]byte, 0)

	switch {
	case len(item) <= 0xfc:
		binary.Write(buf, binary.BigEndian, uint8(len(item)))
	case len(item) <= 0xffff:
		binary.Write(buf, binary.BigEndian, uint16(len(item)))
		prefix = append(prefix, 0xfd)
	default:
		return []byte{}
	}

	result := buf.Bytes()

	if len(prefix) > 0 {
		result = append(prefix, result...)
	}

	return result
}

func BytesToVarInt() {

}

func encodeSigHeader(w io.Writer, header *types.Header) {
	//empty signature

	extras := unserialize(header.Extra)

	headerExtra := make([]byte, 0)
	copy(extras[0][:], bytes.Repeat([]byte{0x00}, crypto.SignatureLength))

	for _, extra := range extras {
		headerExtra = append(headerExtra, VarIntToBytes(extra)...)
		headerExtra = append(headerExtra, extra...)
	}

	toEncode := []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		headerExtra,
		header.MixDigest,
		header.Nonce,
	}

	err := rlp.Encode(w, toEncode)

	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data

	if len(header.Extra) < crypto.SignatureLength {
		return common.Address{}, errMissingSignature
	} else if header.Extra[0] != 0x41 {
		return common.Address{}, errMissingSignature
	}

	signature := header.Extra[1 : crypto.SignatureLength+1]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)

	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}
