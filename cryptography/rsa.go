package cryptography

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"

	"github.com/devlipe/data_structures/calc"
)

func (k *RSAkey) Encrypt(bt []byte) []byte {

	eBuff := new(bytes.Buffer)
	enc := gob.NewEncoder(eBuff)
	temp := make([]big.Int, 0)

	for _, b := range bt {

		a := new(big.Int).Exp(big.NewInt(int64(b)), k.e, k.n)
		temp = append(temp, *a)

	}
	// err := binary.Write(eBuff, binary.BigEndian, temp)
	err := enc.Encode(temp)

	if err != nil {
		log.Fatalln("enc.Encode failed:", err)
	}

	return eBuff.Bytes()
}

func (k *RSAkey) Decrypt(bt []byte) []byte {
	temp := make([]big.Int, (len(bt)+7)/8)

	buf := bytes.NewBuffer(bt)
	dec := gob.NewDecoder(buf)
	// binary.Read(buf, binary.BigEndian, &temp)
	err := dec.Decode(&temp)
	if err != nil {
		log.Fatalln("dec.Decode failed:", err)
	}
	dBuff := new(bytes.Buffer)

	for _, b := range temp {

		a := calc.ChineseRemainderRSA(k.p, k.q, k.dp, k.dq, k.qinv, &b)
		dBuff.WriteByte(byte(a))
	}
	return dBuff.Bytes()
}

// Print a key information in a human readable format
func (k *RSAkey) Print() {
	fmt.Println("n: ", k.n)
	fmt.Println("e: ", k.e)
	fmt.Println("d: ", k.d)
	fmt.Println("p: ", k.p)
	fmt.Println("q: ", k.q)
	fmt.Println("dp: ", k.dp)
	fmt.Println("dq: ", k.dq)
	fmt.Println("qinv: ", k.qinv)
}
