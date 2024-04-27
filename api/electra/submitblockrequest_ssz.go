// Code generated by fastssz. DO NOT EDIT.
// Hash: 4beadbfbd95d5654c51275d7c365975e67d9b7d964b5d245c7ac5c64964f4c7f
// Version: 0.1.3
package electra

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SubmitBlockRequest object
func (s *SubmitBlockRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SubmitBlockRequest object to a target array
func (s *SubmitBlockRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(340)

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(v1.BidTrace)
	}
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (1) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if s.ExecutionPayload == nil {
		s.ExecutionPayload = new(electra.ExecutionPayload)
	}
	offset += s.ExecutionPayload.SizeSSZ()

	// Offset (2) 'BlobsBundle'
	dst = ssz.WriteOffset(dst, offset)
	if s.BlobsBundle == nil {
		s.BlobsBundle = new(deneb.BlobsBundle)
	}
	offset += s.BlobsBundle.SizeSSZ()

	// Field (3) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Field (1) 'ExecutionPayload'
	if dst, err = s.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'BlobsBundle'
	if dst, err = s.BlobsBundle.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SubmitBlockRequest object
func (s *SubmitBlockRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 340 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o2 uint64

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(v1.BidTrace)
	}
	if err = s.Message.UnmarshalSSZ(buf[0:236]); err != nil {
		return err
	}

	// Offset (1) 'ExecutionPayload'
	if o1 = ssz.ReadOffset(buf[236:240]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 340 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (2) 'BlobsBundle'
	if o2 = ssz.ReadOffset(buf[240:244]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (3) 'Signature'
	copy(s.Signature[:], buf[244:340])

	// Field (1) 'ExecutionPayload'
	{
		buf = tail[o1:o2]
		if s.ExecutionPayload == nil {
			s.ExecutionPayload = new(electra.ExecutionPayload)
		}
		if err = s.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'BlobsBundle'
	{
		buf = tail[o2:]
		if s.BlobsBundle == nil {
			s.BlobsBundle = new(deneb.BlobsBundle)
		}
		if err = s.BlobsBundle.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SubmitBlockRequest object
func (s *SubmitBlockRequest) SizeSSZ() (size int) {
	size = 340

	// Field (1) 'ExecutionPayload'
	if s.ExecutionPayload == nil {
		s.ExecutionPayload = new(electra.ExecutionPayload)
	}
	size += s.ExecutionPayload.SizeSSZ()

	// Field (2) 'BlobsBundle'
	if s.BlobsBundle == nil {
		s.BlobsBundle = new(deneb.BlobsBundle)
	}
	size += s.BlobsBundle.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the SubmitBlockRequest object
func (s *SubmitBlockRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SubmitBlockRequest object with a hasher
func (s *SubmitBlockRequest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(v1.BidTrace)
	}
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'ExecutionPayload'
	if err = s.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'BlobsBundle'
	if err = s.BlobsBundle.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SubmitBlockRequest object
func (s *SubmitBlockRequest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
