package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"test/test"
)

func main() {
	msg := &test.Msg{
		Fint32:    90,
		Fint64:    -90,
		Fuint32:   50,
		Fuint64:   500,
		Fsint32:   -30, // 这里是int32
		Fsint64:   512, // 这里是int64
		Ffixed32:  60,  // uint32
		Ffixed64:  0,   // uint64
		Fdouble:   3.9,
		Ffloat:    4.2,
		Fbool:     true,
		Fenum:     test.DayOfWeek_TUESDAY,
		Fmessage:  &test.Child{Chile: 9},
		Fmap:      map[uint32]int32{1: 1, 2: 2},
		Frepeat:   []bool{true, false, true},
		Fstring:   "fjakfj",
		Fbytes:    []byte("jjiejf"),
		Fsfixed32: 90,
		Fsfixed64: 100,
	}
	b, _ := proto.Marshal(msg)
	for i, bb := range b {
		fmt.Printf("%08b ", bb)
		if i%10 == 9 {
			fmt.Println()
		}
	}
}
