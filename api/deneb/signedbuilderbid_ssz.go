// Code generated by fastssz. DO NOT EDIT.
// Hash: 9eb249d7f5f08129f3d3a9bd85af740731ca3ae3150677f16209e6b728acd4f5
// Version: 0.1.3
package deneb

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedBuilderBid object
func (s *SignedBuilderBid) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedBuilderBid object to a target array
func (s *SignedBuilderBid) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(100)

	// Offset (0) 'Message'
	dst = ssz.WriteOffset(dst, offset)
	if s.Message == nil {
		s.Message = new(BuilderBid)
	}
	offset += s.Message.SizeSSZ()

	// Field (1) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Field (0) 'Message'
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedBuilderBid object
func (s *SignedBuilderBid) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 100 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Message'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 100 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Signature'
	copy(s.Signature[:], buf[4:100])

	// Field (0) 'Message'
	{
		buf = tail[o0:]
		if s.Message == nil {
			s.Message = new(BuilderBid)
		}
		if err = s.Message.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBuilderBid object
func (s *SignedBuilderBid) SizeSSZ() (size int) {
	size = 100

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(BuilderBid)
	}
	size += s.Message.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the SignedBuilderBid object
func (s *SignedBuilderBid) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedBuilderBid object with a hasher
func (s *SignedBuilderBid) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedBuilderBid object
func (s *SignedBuilderBid) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
