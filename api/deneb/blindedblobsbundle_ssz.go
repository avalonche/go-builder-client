// Code generated by fastssz. DO NOT EDIT.
// Hash: a670c8936ed78dd3a7f9257ba32b0b7638c7af80cb190629e6da3f1672e5d9f3
// Version: 0.1.3
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlindedBlobsBundle object
func (b *BlindedBlobsBundle) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlindedBlobsBundle object to a target array
func (b *BlindedBlobsBundle) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'Commitments'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Commitments) * 48

	// Offset (1) 'Proofs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Proofs) * 48

	// Offset (2) 'BlobRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlobRoots) * 32

	// Field (0) 'Commitments'
	if size := len(b.Commitments); size > 4096 {
		err = ssz.ErrListTooBigFn("BlindedBlobsBundle.Commitments", size, 4096)
		return
	}
	for ii := 0; ii < len(b.Commitments); ii++ {
		dst = append(dst, b.Commitments[ii][:]...)
	}

	// Field (1) 'Proofs'
	if size := len(b.Proofs); size > 4096 {
		err = ssz.ErrListTooBigFn("BlindedBlobsBundle.Proofs", size, 4096)
		return
	}
	for ii := 0; ii < len(b.Proofs); ii++ {
		dst = append(dst, b.Proofs[ii][:]...)
	}

	// Field (2) 'BlobRoots'
	if size := len(b.BlobRoots); size > 4096 {
		err = ssz.ErrListTooBigFn("BlindedBlobsBundle.BlobRoots", size, 4096)
		return
	}
	for ii := 0; ii < len(b.BlobRoots); ii++ {
		dst = append(dst, b.BlobRoots[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlindedBlobsBundle object
func (b *BlindedBlobsBundle) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'Commitments'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Proofs'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'BlobRoots'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'Commitments'
	{
		buf = tail[o0:o1]
		num, err := ssz.DivideInt2(len(buf), 48, 4096)
		if err != nil {
			return err
		}
		b.Commitments = make([]deneb.KzgCommitment, num)
		for ii := 0; ii < num; ii++ {
			copy(b.Commitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}

	// Field (1) 'Proofs'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 48, 4096)
		if err != nil {
			return err
		}
		b.Proofs = make([]deneb.KzgProof, num)
		for ii := 0; ii < num; ii++ {
			copy(b.Proofs[ii][:], buf[ii*48:(ii+1)*48])
		}
	}

	// Field (2) 'BlobRoots'
	{
		buf = tail[o2:]
		num, err := ssz.DivideInt2(len(buf), 32, 4096)
		if err != nil {
			return err
		}
		b.BlobRoots = make([]phase0.Root, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlindedBlobsBundle object
func (b *BlindedBlobsBundle) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'Commitments'
	size += len(b.Commitments) * 48

	// Field (1) 'Proofs'
	size += len(b.Proofs) * 48

	// Field (2) 'BlobRoots'
	size += len(b.BlobRoots) * 32

	return
}

// HashTreeRoot ssz hashes the BlindedBlobsBundle object
func (b *BlindedBlobsBundle) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlindedBlobsBundle object with a hasher
func (b *BlindedBlobsBundle) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Commitments'
	{
		if size := len(b.Commitments); size > 4096 {
			err = ssz.ErrListTooBigFn("BlindedBlobsBundle.Commitments", size, 4096)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Commitments {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.Commitments))
		hh.MerkleizeWithMixin(subIndx, numItems, 4096)
	}

	// Field (1) 'Proofs'
	{
		if size := len(b.Proofs); size > 4096 {
			err = ssz.ErrListTooBigFn("BlindedBlobsBundle.Proofs", size, 4096)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Proofs {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.Proofs))
		hh.MerkleizeWithMixin(subIndx, numItems, 4096)
	}

	// Field (2) 'BlobRoots'
	{
		if size := len(b.BlobRoots); size > 4096 {
			err = ssz.ErrListTooBigFn("BlindedBlobsBundle.BlobRoots", size, 4096)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.BlobRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, 4096)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlindedBlobsBundle object
func (b *BlindedBlobsBundle) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
