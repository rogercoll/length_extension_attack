package merkledamgard

import (
	"encoding/hex"
	"github.com/rogercoll/merkledamgard/pkg/md5"
)

//length corresponts to the lenght of the original message
//hash => original hash H(secret || message)
//extend => extension that we want to add to the hash, for example: world

func LengthExtensionAttack(length int, hash, extend string) (string, error) {
	mockBytes := make([]byte, length)
	pb := md5.MakePaddingBlock(mockBytes, length)
	oldmd5, err := hex.DecodeString(hash)

	if err != nil {
		return "", err
	}
	d := md5.PreviousToSum(oldmd5[:])
	extendedBytes := []byte(extend)
	clen := md5.MakePaddingBlock(extendedBytes, len(extend)+len(pb))
	md5.BlockGeneric2(&d, clen)
	newForgedmd5 := d.ForgedCheckSum()
	inhash := hex.EncodeToString(newForgedmd5[:])
	return inhash, nil
}
