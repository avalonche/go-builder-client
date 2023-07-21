// Code generated by fastssz. DO NOT EDIT.
// Hash: c6b76f7f064aa399e33e95fc859f9dcc33fbe18976165fbb88abb8f54f1f9350
// Version: 0.1.3
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlobsBundle object
func (b *BlobsBundle) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlobsBundle object to a target array
func (b *BlobsBundle) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'Commitments'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Commitments) * 48

	// Offset (1) 'Proofs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Proofs) * 48

	// Offset (2) 'Blobs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Blobs) * 131072

	// Field (0) 'Commitments'
	if size := len(b.Commitments); size > 4 {
		err = ssz.ErrListTooBigFn("BlobsBundle.Commitments", size, 4)
		return
	}
	for ii := 0; ii < len(b.Commitments); ii++ {
		dst = append(dst, b.Commitments[ii][:]...)
	}

	// Field (1) 'Proofs'
	if size := len(b.Proofs); size > 4 {
		err = ssz.ErrListTooBigFn("BlobsBundle.Proofs", size, 4)
		return
	}
	for ii := 0; ii < len(b.Proofs); ii++ {
		dst = append(dst, b.Proofs[ii][:]...)
	}

	// Field (2) 'Blobs'
	if size := len(b.Blobs); size > 4 {
		err = ssz.ErrListTooBigFn("BlobsBundle.Blobs", size, 4)
		return
	}
	for ii := 0; ii < len(b.Blobs); ii++ {
		dst = append(dst, b.Blobs[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlobsBundle object
func (b *BlobsBundle) UnmarshalSSZ(buf []byte) error {
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

	// Offset (2) 'Blobs'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'Commitments'
	{
		buf = tail[o0:o1]
		num, err := ssz.DivideInt2(len(buf), 48, 4)
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
		num, err := ssz.DivideInt2(len(buf), 48, 4)
		if err != nil {
			return err
		}
		b.Proofs = make([]deneb.KzgProof, num)
		for ii := 0; ii < num; ii++ {
			copy(b.Proofs[ii][:], buf[ii*48:(ii+1)*48])
		}
	}

	// Field (2) 'Blobs'
	{
		buf = tail[o2:]
		num, err := ssz.DivideInt2(len(buf), 131072, 4)
		if err != nil {
			return err
		}
		b.Blobs = make([]deneb.Blob, num)
		for ii := 0; ii < num; ii++ {
			copy(b.Blobs[ii][:], buf[ii*131072:(ii+1)*131072])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobsBundle object
func (b *BlobsBundle) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'Commitments'
	size += len(b.Commitments) * 48

	// Field (1) 'Proofs'
	size += len(b.Proofs) * 48

	// Field (2) 'Blobs'
	size += len(b.Blobs) * 131072

	return
}

// HashTreeRoot ssz hashes the BlobsBundle object
func (b *BlobsBundle) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobsBundle object with a hasher
func (b *BlobsBundle) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Commitments'
	{
		if size := len(b.Commitments); size > 4 {
			err = ssz.ErrListTooBigFn("BlobsBundle.Commitments", size, 4)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Commitments {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.Commitments))
		hh.MerkleizeWithMixin(subIndx, numItems, 4)
	}

	// Field (1) 'Proofs'
	{
		if size := len(b.Proofs); size > 4 {
			err = ssz.ErrListTooBigFn("BlobsBundle.Proofs", size, 4)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Proofs {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.Proofs))
		hh.MerkleizeWithMixin(subIndx, numItems, 4)
	}

	// Field (2) 'Blobs'
	{
		if size := len(b.Blobs); size > 4 {
			err = ssz.ErrListTooBigFn("BlobsBundle.Blobs", size, 4)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Blobs {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.Blobs))
		hh.MerkleizeWithMixin(subIndx, numItems, 4)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlobsBundle object
func (b *BlobsBundle) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
