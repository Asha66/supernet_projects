/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package pbkdf2

import (
	"bytes"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"

	//"tk.com/crypto/aes"
	"checkermakerweb/utils/crypto/aes"
	//"krishivala.com/crypto/aes"
)

type Pbkdf2 struct {
	Key    []byte
	Salt   []byte
	Plain  []byte
	Cipher []byte
	Itr    int
	KeyLen int
}

func (obj *Pbkdf2) Encrypt() (err error) {
	obj.Key = pbkdf2.Key(obj.Plain, obj.Salt, obj.Itr, obj.KeyLen, sha256.New)
	obj.Cipher, err = aes.Encrypt(obj.Plain, obj.Key)
	if err != nil {
		return
	}
	return
}

func (obj *Pbkdf2) Compare() (result bool, err error) {
	obj.Key = pbkdf2.Key(obj.Plain, obj.Salt, obj.Itr, obj.KeyLen, sha256.New)

	cp, err := aes.Encrypt(obj.Plain, obj.Key)
	if err != nil {
		return
	}

	if !bytes.Equal(cp, obj.Cipher) {
		result = false
		return
	}
	result = true
	return
}
