// Code generated by fastssz. DO NOT EDIT.
// Hash: 8580ab2c0eff14bc106c09e71bc3832ef122f8469bceca0a3b29a5437a53e207
// Version: 0.1.3
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
	"github.com/holiman/uint256"
)

// MarshalSSZ ssz marshals the BuilderBid object
func (b *BuilderBid) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BuilderBid object to a target array
func (b *BuilderBid) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(312)

	// Offset (0) 'Header'
	dst = ssz.WriteOffset(dst, offset)
	if b.Header == nil {
		b.Header = new(deneb.ExecutionPayloadHeader)
	}
	offset += b.Header.SizeSSZ()

	// Offset (1) 'BlindedBlobsBundle'
	dst = ssz.WriteOffset(dst, offset)
	if b.BlindedBlobsBundle == nil {
		b.BlindedBlobsBundle = new(BlindedBlobsBundle)
	}
	offset += b.BlindedBlobsBundle.SizeSSZ()

	// Field (2) 'Value'
	value := b.Value.Bytes32()
	for i := 0; i < 32; i++ {
		dst = append(dst, value[31-i])
	}

	// Field (3) 'Pubkey'
	dst = append(dst, b.Pubkey[:]...)

	// Field (0) 'Header'
	if dst, err = b.Header.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'BlindedBlobsBundle'
	if dst, err = b.BlindedBlobsBundle.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BuilderBid object
func (b *BuilderBid) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 312 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Header'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 312 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'BlindedBlobsBundle'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'Value'
	value := make([]byte, 32)
	for i := 0; i < 32; i++ {
		value[i] = buf[39-i]
	}
	if b.Value == nil {
		b.Value = new(uint256.Int)
	}
	b.Value.SetBytes32(value)

	// Field (3) 'Pubkey'
	copy(b.Pubkey[:], buf[264:312])

	// Field (0) 'Header'
	{
		buf = tail[o0:o1]
		if b.Header == nil {
			b.Header = new(deneb.ExecutionPayloadHeader)
		}
		if err = b.Header.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'BlindedBlobsBundle'
	{
		buf = tail[o1:]
		if b.BlindedBlobsBundle == nil {
			b.BlindedBlobsBundle = new(BlindedBlobsBundle)
		}
		if err = b.BlindedBlobsBundle.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BuilderBid object
func (b *BuilderBid) SizeSSZ() (size int) {
	size = 312

	// Field (0) 'Header'
	if b.Header == nil {
		b.Header = new(deneb.ExecutionPayloadHeader)
	}
	size += b.Header.SizeSSZ()

	// Field (1) 'BlindedBlobsBundle'
	if b.BlindedBlobsBundle == nil {
		b.BlindedBlobsBundle = new(BlindedBlobsBundle)
	}
	size += b.BlindedBlobsBundle.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BuilderBid object
func (b *BuilderBid) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BuilderBid object with a hasher
func (b *BuilderBid) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Header'
	if err = b.Header.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'BlindedBlobsBundle'
	if err = b.BlindedBlobsBundle.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Value'
	value := b.Value.Bytes32()
	for i, j := 0, 31; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	hh.PutBytes(value[:])

	// Field (3) 'Pubkey'
	hh.PutBytes(b.Pubkey[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BuilderBid object
func (b *BuilderBid) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
