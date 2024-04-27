// Code generated by fastssz. DO NOT EDIT.
// Hash: 4beadbfbd95d5654c51275d7c365975e67d9b7d964b5d245c7ac5c64964f4c7f
// Version: 0.1.3
package electra

import (
	"github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ExecutionPayloadAndBlobsBundle object
func (e *ExecutionPayloadAndBlobsBundle) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadAndBlobsBundle object to a target array
func (e *ExecutionPayloadAndBlobsBundle) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if e.ExecutionPayload == nil {
		e.ExecutionPayload = new(electra.ExecutionPayload)
	}
	offset += e.ExecutionPayload.SizeSSZ()

	// Offset (1) 'BlobsBundle'
	dst = ssz.WriteOffset(dst, offset)
	if e.BlobsBundle == nil {
		e.BlobsBundle = new(deneb.BlobsBundle)
	}
	offset += e.BlobsBundle.SizeSSZ()

	// Field (0) 'ExecutionPayload'
	if dst, err = e.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'BlobsBundle'
	if dst, err = e.BlobsBundle.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadAndBlobsBundle object
func (e *ExecutionPayloadAndBlobsBundle) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'ExecutionPayload'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'BlobsBundle'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'ExecutionPayload'
	{
		buf = tail[o0:o1]
		if e.ExecutionPayload == nil {
			e.ExecutionPayload = new(electra.ExecutionPayload)
		}
		if err = e.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'BlobsBundle'
	{
		buf = tail[o1:]
		if e.BlobsBundle == nil {
			e.BlobsBundle = new(deneb.BlobsBundle)
		}
		if err = e.BlobsBundle.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadAndBlobsBundle object
func (e *ExecutionPayloadAndBlobsBundle) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'ExecutionPayload'
	if e.ExecutionPayload == nil {
		e.ExecutionPayload = new(electra.ExecutionPayload)
	}
	size += e.ExecutionPayload.SizeSSZ()

	// Field (1) 'BlobsBundle'
	if e.BlobsBundle == nil {
		e.BlobsBundle = new(deneb.BlobsBundle)
	}
	size += e.BlobsBundle.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadAndBlobsBundle object
func (e *ExecutionPayloadAndBlobsBundle) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadAndBlobsBundle object with a hasher
func (e *ExecutionPayloadAndBlobsBundle) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ExecutionPayload'
	if err = e.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'BlobsBundle'
	if err = e.BlobsBundle.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ExecutionPayloadAndBlobsBundle object
func (e *ExecutionPayloadAndBlobsBundle) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}
