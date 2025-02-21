/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 */
package pbkdf2

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"
)

func TestPGP(t *testing.T) {
	//	b := make([]byte, 32)
	//	_, err := rand.Read(b)
	//	var obj Pbkdf2
	//	obj.Itr = 32
	//	obj.Cipher = []byte("V//Xn/LBA0cHTnwIpqkYQYHSsqLOU9f47FtGAEgJh402UCpVHo6u8mOI3TZ4zBVO")[32:]
	//	obj.Salt = []byte("V//Xn/LBA0cHTnwIpqkYQYHSsqLOU9f47FtGAEgJh402UCpVHo6u8mOI3TZ4zBVO")[:32]
	//	obj.KeyLen = 32
	//	obj.Plain = []byte("620664")
	//	resulterr = obj.Compare()
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	var tmp []byte
	//	tmp = append(tmp, obj.Salt...)
	//	tmp = append(tmp, obj.Cipher...)
	//	t.Log(len(tmp))
	//	out, err := base64.Encode(tmp)
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	t.Log(string(out))

	dat := "2017-04-05"
	layout := "2006-01-02"

	t1, _ := time.Parse(layout, dat)
	fmt.Println(t1)

	t1.Format("02-01-2006")
	fmt.Println(t1)

}

func BenchmarkPGP(t *testing.B) {
	for j := 0; j < t.N; j++ {
		b := make([]byte, 32)
		_, err := rand.Read(b)
		var obj Pbkdf2
		obj.Itr = 32
		obj.Plain = []byte("080328")
		obj.Salt = b
		obj.KeyLen = 32
		err = obj.Encrypt()
		if err != nil {
			t.Error(err)
			return
		}
	}
}
