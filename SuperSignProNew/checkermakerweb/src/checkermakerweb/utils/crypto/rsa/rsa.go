/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 */
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	//	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func ReadPrivateKey(filename string) (key *rsa.PrivateKey, err error) {
	fd, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : private key not found")
		return
	}

	blk, _ := pem.Decode(fd)
	if blk == nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : private key decode fail")
		return
	}
	if got, want := blk.Type, "RSA PRIVATE KEY"; got != want {
		err = errors.New(filename + " : invalid private key")
		return
	}

	key, err = x509.ParsePKCS1PrivateKey(blk.Bytes)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : private key parsing fail")
		return
	}
	return
}

func ReadPublicKey(filename string) (key *rsa.PublicKey, err error) {

	fd, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : public key not found")
		return
	}

	blk, _ := pem.Decode(fd)
	if blk == nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : public key decode fail")
		return
	}

	if got, want := blk.Type, "CERTIFICATE"; got != want {
		err = errors.New(filename + " : invalid public key")
		return
	}

	cer, err := x509.ParseCertificate(blk.Bytes)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : certificate parsing fail")
		return
	}
	if cer == nil {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : certificate data not found")
		return
	}

	key, ok := cer.PublicKey.(*rsa.PublicKey)
	if !ok {
		log.Println("Error", "Error", err)
		err = errors.New(filename + " : certificate invalid")
		return
	}
	return
}

func Encrypt(in []byte, key *rsa.PublicKey) (out []byte, err error) {

	out, err = rsa.EncryptPKCS1v15(rand.Reader, key, in)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New("rsa encrypt fail")
		return
	}
	return
}

func Decrypt(in []byte, key *rsa.PrivateKey) (out []byte, err error) {

	out, err = rsa.DecryptPKCS1v15(nil, key, in)
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New("rsa decrypt fail")
		return
	}
	return
}
