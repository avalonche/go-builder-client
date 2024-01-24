// Code generated by fastssz. DO NOT EDIT.
// Hash: 7892d1b8a34d12e16ce7b5e0af8aef325f91e09b3a36208f80b706f3ffed4b78
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
	offset := int(88)

	// Offset (0) 'Header'
	dst = ssz.WriteOffset(dst, offset)
	if b.Header == nil {
		b.Header = new(deneb.ExecutionPayloadHeader)
	}
	offset += b.Header.SizeSSZ()

	// Offset (1) 'BlobKZGCommitments'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlobKZGCommitments) * 48

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

	// Field (1) 'BlobKZGCommitments'
	if size := len(b.BlobKZGCommitments); size > 4096 {
		err = ssz.ErrListTooBigFn("BuilderBid.BlobKZGCommitments", size, 4096)
		return
	}
	for ii := 0; ii < len(b.BlobKZGCommitments); ii++ {
		dst = append(dst, b.BlobKZGCommitments[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BuilderBid object
func (b *BuilderBid) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 88 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Header'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 88 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'BlobKZGCommitments'
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
	copy(b.Pubkey[:], buf[40:88])

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

	// Field (1) 'BlobKZGCommitments'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 48, 4096)
		if err != nil {
			return err
		}
		b.BlobKZGCommitments = make([]deneb.KZGCommitment, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobKZGCommitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BuilderBid object
func (b *BuilderBid) SizeSSZ() (size int) {
	size = 88

	// Field (0) 'Header'
	if b.Header == nil {
		b.Header = new(deneb.ExecutionPayloadHeader)
	}
	size += b.Header.SizeSSZ()

	// Field (1) 'BlobKZGCommitments'
	size += len(b.BlobKZGCommitments) * 48

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

	// Field (1) 'BlobKZGCommitments'
	{
		if size := len(b.BlobKZGCommitments); size > 4096 {
			err = ssz.ErrListTooBigFn("BuilderBid.BlobKZGCommitments", size, 4096)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobKZGCommitments {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.BlobKZGCommitments))
		hh.MerkleizeWithMixin(subIndx, numItems, 4096)
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
