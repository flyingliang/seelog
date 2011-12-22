// Copyright 2011 Cloud Instruments Co. Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	log "github.com/cihub/sealog"
	"strings"
	"time"
)

var longMessage = strings.Repeat("A", 1024*100)

func main() {
	syncBehavior()
	fmt.Println()
	asyncLoopBehavior()
	fmt.Println()
	asyncTimerBehavior()
}

func syncBehavior() {
	fmt.Println("Sync test")

	testConfig := `
<sealog type="sync">
	<outputs>
		<filter levels="trace">
			<file path="log.log"/>
		</filter>
		<filter levels="debug">
			<console />
		</filter>
	</outputs>
</sealog>
`

	conf, _ := log.ConfigFromBytes([]byte(testConfig))
	log.UseConfig(conf)

	doTest()
}

func asyncLoopBehavior() {
	fmt.Println("Async loop test")

	testConfig := `
<sealog>
	<outputs>
		<filter levels="trace">
			<file path="log.log"/>
		</filter>
		<filter levels="debug">
			<console />
		</filter>
	</outputs>
</sealog>`

	conf, _ := log.ConfigFromBytes([]byte(testConfig))
	log.UseConfig(conf)

	doTest()

	time.Sleep(1e9)
}

func asyncTimerBehavior() {
	fmt.Println("Async timer test")

	testConfig := `
<sealog type="asynctimer" asyncinterval="500">
	<outputs>
		<filter levels="trace">
			<file path="log.log"/>
		</filter>
		<filter levels="debug">
			<console />
		</filter>
	</outputs>
</sealog>`

	conf, _ := log.ConfigFromBytes([]byte(testConfig))
	log.UseConfig(conf)

	doTest()

	time.Sleep(1e9)
}

func doTest() {
	start := time.Now()
	for i := 0; i < 30; i += 2 {
		fmt.Printf("%d\n", i)
		log.Trace(longMessage)
		log.Debug("%d", i+1)
	}
	end := time.Now()
	dur := end.Sub(start)
	fmt.Printf("Test took %d ns\n", dur)
}
