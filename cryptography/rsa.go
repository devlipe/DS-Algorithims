package cryptography

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/devlipe/data_structures/calc"
)

// FIXME: Receive a byte array and return a byte array
func (k *RSAkey) Encrypt(bt []byte) []uint64 {

	eBuff := new(bytes.Buffer)
	temp := make([]uint64, 0)
	for _, b := range bt {

		temp = append(temp, calc.ModExp(uint64(b), k.e, k.n))
	}
	err := binary.Write(eBuff, binary.BigEndian, temp)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return temp
}

// FIXME: Receive a byte array and return a byte array
func (k *RSAkey) Decrypt(bt []uint64) []byte {
	// temp := make([]uint64, (len(bt)+7)/8)
	// buf := bytes.NewBuffer(bt)
	// fmt.Println("bt: ", bt)
	// binary.Read(buf, binary.BigEndian, &temp)
	// fmt.Println("temp: ", temp)
	// for _, b := range temp {
	//
	// 	fmt.Println("bt: ", b)
	// 	// a := calc.ModExp(b, d, n)
	// 	// decoded = append(decoded, byte(a))
	// }
	// return bt

	decoded := make([]byte, 0)

	for _, b := range bt {
		a := calc.ChineseRemainderRSA(k.p, k.q, k.dp, k.dq, k.qinv, b)
		decoded = append(decoded, byte(a))
	}

	return decoded
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
