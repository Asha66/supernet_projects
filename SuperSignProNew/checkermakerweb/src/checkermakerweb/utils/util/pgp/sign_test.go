/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package pgp_test

import (
	"fmt"
	"testing"

	"github.com/jchavannes/go-pgp/pgp"
)

func TestSignature(t *testing.T) {
	fmt.Println("Signature test: START")
	entity, err := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created private key entity.")

	signature, err := pgp.Sign(entity, []byte(TestMessage))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created signature of test message with private key entity.")

	publicKeyEntity, err := pgp.GetEntity([]byte(TestPublicKey), []byte{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created public key entity.")

	err = pgp.Verify(publicKeyEntity, []byte(TestMessage), signature)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Signature verified using public key entity.")
	fmt.Println("Signature test: END")
}
