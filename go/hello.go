package main

import (
	"fmt"
	"hash"
	"math/big"

	// "encoding/base64"
	"golang.org/x/crypto/sha3"
)

// ECDSASignature represents an ECDSA signature
type ECDSASignature struct {
	R, S *big.Int
}

func getHashSHA3(bitsize int) (hash.Hash, error) {
	switch bitsize {
	case 224:
		return sha3.New224(), nil
	case 256:
		return sha3.New256(), nil
	case 384:
		return sha3.New384(), nil
	case 512:
		return sha3.New512(), nil
	case 521:
		return sha3.New512(), nil
	default:
		return nil, fmt.Errorf("Invalid bitsize. It was [%d]. Expected [224, 256, 384, 512, 521]", bitsize)
	}
}

func computeHash(msg []byte, bitsize int) ([]byte, error) {
	hash, err := getHashSHA3(bitsize)
	if err != nil {
		return nil, err
	}

	// base64Msg := base64.StdEncoding.EncodeToString(msg)
	// myLogger.Debugf("base64 Msg -----------: %s\n", base64Msg)
	// hash.Write([]byte(base64Msg))
	hash.Write(msg)
	return hash.Sum(nil), nil
}

func getPublicKeyHash(key []byte) string {
	hashPub, err := computeHash(key, 224)
	if err != nil {
		// myLogger.Error("Invalid signature")
	}

	tmp := fmt.Sprintf("%x", hashPub)
	fmt.Printf("hash -------------: %s\n", tmp)

	return tmp
}
//func main() {
//	str := "308201bb30820161a003020102020101300a06082a8648ce3d0403033031310b300906035504061302555331143012060355040a130b48797065726c6564676572310c300a06035504031303656361301e170d3137303232303036313230345a170d3137303532313036313230345a304a310b300906035504061302555331143012060355040a130b48797065726c65646765723125302306035504030c1c6f6e655f636861696e5f61646d696e5c696e737469747574696f6e733059301306072a8648ce3d020106082a8648ce3d03010703420004b08cf2a408d5088e755323a48f8c1715d8c4df84a29071ae462a7d73ce001fd416c5ee7e77d04285ffb2c48558329be25ef5749d7b3f0cc56abde3b7fceebaf7a351304f300e0603551d0f0101ff040403020780300c0603551d130101ff04023000300d0603551d0e0406040401020304300f0603551d2304083006800401020304300f06065103040506070101ff04023136300a06082a8648ce3d040303034800304502201c04afc9ed8b843b2b781d12052e2c723a8ecbba84e169dfbb55d054edcf1892022100d459f507a8f7ae0beb4f49479e37e8acdf08429cc173c116c81533818828de1e"
//	fmt.Printf("$$$$$$$$$123:", getPublicKeyHash([]byte(str)))
//}
