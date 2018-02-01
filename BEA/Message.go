package BEA

import (
	"CreditCard/ISO8583"
	_ "crypto/cipher"
	"crypto/des"
	"fmt"
)

//加密ISO8583消息
func encryptISO8583Message(msg []byte) []byte {
	keyV1 := ISO8583.Base16Decode("ABCDEF0123456789EEEEEEEEEEEEEEEE")
	keyV2 := ISO8583.Base16Decode("FFFFFFFFFFFFFFFF9876543210FEDCBA")
	var key []byte

	for i := 0; i < len(keyV1); i++ {
		key = append(key, keyV1[i]^keyV2[i])
	}

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, key[:16]...)
	tripleDESKey = append(tripleDESKey, key[:8]...)

	tdesCipher, err := des.NewTripleDESCipher(tripleDESKey)

	if err != nil {
		panic(err)
	}

	round := (len(msg) + 7) / 8
	dstLen := round * 8
	fmt.Printf("msg len:%d,round:%d,dstLen:%d\r\n", len(msg), round, dstLen)
	plainMsg := make([]byte, dstLen)
	copy(plainMsg, msg)
	encryptedMsg := make([]byte, 0)
	encBlock := make([]byte, 8)
	var plainBlock []byte

	for i := 0; i < round; i++ {
		plainBlock = plainMsg[i*8 : i*8+8]
		tdesCipher.Encrypt(encBlock, plainBlock)
		encryptedMsg = append(encryptedMsg, encBlock...)
	}

	return encryptedMsg
}

/**
	根据给定的交易数据和需要打包的位图信息生成ISO8583 报文

**/

func createIISO8583Message(fieldsMap map[uint8]string, config *Config) ([]byte, error) {
	msg, err := ISO8583.PrepareISO8583Message(fieldsMap)
	if err != nil {
		return nil, fmt.Errorf("ISO8583::PrepareISO8583Message error: %s", err.Error())
	}

	// msg = encryptISO8583Message(msg)

	dstMsg := make([]byte, 0)
	dstMsg = append(dstMsg, 0x00, 0x00) // len
	dstMsg = append(dstMsg, ISO8583.Base16Decode(config.TPDU)...)
	dstMsg = append(dstMsg, ISO8583.Base16Decode(config.EDS)...)
	dstMsg = append(dstMsg, msg...)
	dstMsg[2+5+4] = byte(len(msg) >> 8)
	dstMsg[2+5+5] = byte(len(msg) & 0x00FF)

	dstMsg[0] = byte((len(dstMsg) - 2) >> 8)
	dstMsg[1] = byte((len(dstMsg) - 2) & 0x00FF)

	return dstMsg, nil
}
