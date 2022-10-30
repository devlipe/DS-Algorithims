package cryptography

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/devlipe/data_structures/calc"
)

func (k *RSAkey) Encrypt(bt []byte) []byte {

	eBuff := new(bytes.Buffer)
	temp := make([]uint64, 0)
	for _, b := range bt {

		temp = append(temp, calc.ModExp(uint64(b), k.e, k.n))
	}
	err := binary.Write(eBuff, binary.BigEndian, temp)

	if err != nil {
		log.Fatalln("binary.Write failed:", err)
	}

	return eBuff.Bytes()
}

func (k *RSAkey) Decrypt(bt []byte) []byte {
	temp := make([]uint64, (len(bt)+7)/8)
	buf := bytes.NewBuffer(bt)
	binary.Read(buf, binary.BigEndian, &temp)
	dBuff := new(bytes.Buffer)

	for _, b := range temp {

		a := calc.ChineseRemainderRSA(k.p, k.q, k.dp, k.dq, k.qinv, b)
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
