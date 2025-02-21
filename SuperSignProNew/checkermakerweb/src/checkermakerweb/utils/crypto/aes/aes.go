/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"errors"
	"fmt"

	"checkermakerweb/utils/encoding/base64"

	log "github.com/sirupsen/logrus"
	// "krishivala.com/encoding/base64"
	// "krishivala.com/util/log"
)

func Generate(len int) (key []byte, err error) {
	key = make([]byte, len)
	n, err := rand.Read(key)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New("AES key generation fail")
	}
	if n != len {
		err = errors.New("AES key legnth invalid")
	}
	return
}

func Encrypt(in []byte, key []byte) (out []byte, err error) {

	obj, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}

	plainPad, err := DoPKCS7Pad(in, obj.BlockSize())
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}
	tmp := make([]byte, len(plainPad))

	if len(plainPad)%obj.BlockSize() != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		log.Println("Error", "Error", err)
		return
	}
	if len(tmp) < len(plainPad) {
		err = errors.New("crypto/cipher: output smaller than input")
		log.Println("Error", "Error", err)
		return
	}
	for len(plainPad) > 0 {
		obj.Encrypt(tmp, plainPad[:obj.BlockSize()])
		plainPad = plainPad[obj.BlockSize():]
		tmp = tmp[:obj.BlockSize()]
		out = append(out, tmp...)
	}
	return
}

func Decrypt(in []byte, key []byte) (out []byte, err error) {

	obj, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}

	tmp_arr := make([]byte, len(in))
	var cipherPad []byte
	for len(in) > 0 {
		obj.Decrypt(tmp_arr, in[:obj.BlockSize()])
		in = in[obj.BlockSize():]
		tmp_arr = tmp_arr[:obj.BlockSize()]
		cipherPad = append(cipherPad, tmp_arr...)
	}
	if len(cipherPad)%obj.BlockSize() != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		log.Println("Error", "Error", err)
		return
	}
	fmt.Println(cipherPad)
	out, err = DoPKCS7Unpad(cipherPad, obj.BlockSize())
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}

	return
}

func DoPKCS7Pad(data []byte, blocklen int) ([]byte, error) {

	if blocklen <= 0 {
		return nil, fmt.Errorf("invalid blocklen %d", blocklen)
	}
	padlen := 1
	for ((len(data) + padlen) % blocklen) != 0 {
		padlen = padlen + 1
	}

	pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
	return append(data, pad...), nil
}

func DoPKCS7Unpad(data []byte, blocklen int) ([]byte, error) {

	if blocklen <= 0 {
		return nil, fmt.Errorf("invalid blocklen %d", blocklen)
	}
	if len(data)%blocklen != 0 || len(data) == 0 {
		return nil, fmt.Errorf("invalid data len %d", len(data))
	}
	padlen := int(data[len(data)-1])
	if padlen > blocklen || padlen == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	// check padding
	pad := data[len(data)-padlen:]
	for i := 0; i < padlen; i++ {
		if pad[i] != byte(padlen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return data[:len(data)-padlen], nil
}

func Base64Encrypt(in []byte, key []byte) (out []byte, err error) {

	tmp1 := []byte("00000000000000000000000000000000")
	if len(key) != 32 {
		key = append(key, tmp1[:32-len(key)]...)
	}

	obj, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}

	plainPad, err := DoPKCS7Pad(in, obj.BlockSize())
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}
	tmp := make([]byte, len(plainPad))

	if len(plainPad)%obj.BlockSize() != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		log.Println("Error", "Error", err)
		return
	}
	if len(tmp) < len(plainPad) {
		err = errors.New("crypto/cipher: output smaller than input")
		log.Println("Error", "Error", err)
		return
	}
	var t []byte

	for len(plainPad) > 0 {
		obj.Encrypt(tmp, plainPad[:obj.BlockSize()])
		plainPad = plainPad[obj.BlockSize():]
		tmp = tmp[:obj.BlockSize()]
		t = append(t, tmp...)
	}

	out, err = base64.Encode(t)
	if err != nil {
		return
	}
	return
}

func Base64Decrypt(in []byte, key []byte) (out []byte, err error) {
	tmp1 := []byte("00000000000000000000000000000000")
	if len(key) != 32 {
		key = append(key, tmp1[:32-len(key)]...)
	}

	obj, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}

	ind, err := base64.Decode(in)
	if err != nil {
		return
	}

	tmp := make([]byte, len(ind))
	var cipherPad []byte
	for len(ind) > 0 {
		obj.Decrypt(tmp, ind[:obj.BlockSize()])
		ind = ind[obj.BlockSize():]
		tmp = tmp[:obj.BlockSize()]
		cipherPad = append(cipherPad, tmp...)
	}
	if len(cipherPad)%obj.BlockSize() != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		log.Println("Error", "Error", err)
		return
	}
	fmt.Println(string(cipherPad))
	out, err = DoPKCS7Unpad(cipherPad, obj.BlockSize())
	if err != nil {
		log.Println("Error", "Error", err)
		return
	}
	return
}
