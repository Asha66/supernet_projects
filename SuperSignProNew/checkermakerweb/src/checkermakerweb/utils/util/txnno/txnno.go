/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 */
package txnno

import (
	"fmt"
	"math/rand"

	//	"strconv"
	"time"
)

const (
	MAX_TRANS       = 20
	MAX_NANO        = 10000
	MAX_RAND        = 9000
	MAX_NANO_DIGIT  = 10000000000
	MAX_TRANS_DIGIT = 5
)

func getSeed(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func Generate() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS, et)
}

func Generate13Digit() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO_DIGIT
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS_DIGIT, et)
}

func Generate20Digit() string {
	lt := (time.Now().UnixNano())
	lt = lt / MAX_NANO_DIGIT
	epochTrans := fmt.Sprintf("%d", lt)
	ln := getSeed(MAX_RAND)
	ri := fmt.Sprintf("%d", ln)
	et := ri + epochTrans
	return fmt.Sprintf("%0*s", MAX_TRANS, et)
}
