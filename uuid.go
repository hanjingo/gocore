package gocore

import (
	"sync"
	"sync/atomic"
	"time"
)

// StartTime ;you known it
var StartTime, _ = time.Parse("2006-01-02 15:04:05", "2018-01-01 00:00:00")

// DoLittleEndian ;use little endian?
var DoLittleEndian = false

var once sync.Once
var gen *UUID

// GetUUID ;return random generator
func GetUUID() *UUID {
	once.Do(func() {
		gen = &UUID{
			code:    []byte{},
			codeLen: 0,
			sec:     0,
			inc:     0,
			maxInc:  1<<32 - 1,
		}
	})
	return gen
}

// UUID ;random generator
// bits（64bit）:stable code(0~16bit) + total seconds(32bit) + increment(16~32bit)
type UUID struct {
	code    []byte // stable code
	codeLen uint32 // code length
	sec     uint32 // from 20180101 to now's second; 32bit, It can last more than 100 years
	inc     uint32 // increment number
	maxInc  uint32 // max increment number
}

func (u *UUID) SetStableCode(code []byte) {
	switch len(code) {
	case 1:
		atomic.SwapUint32(&u.maxInc, 1<<24-1)
		atomic.SwapUint32(&u.codeLen, 8)
		u.code = code
	case 2:
		atomic.SwapUint32(&u.maxInc, 1<<16-1)
		atomic.SwapUint32(&u.codeLen, 16)
		u.code = code
	default:
		return
	}
}

func (u *UUID) Reset() {
	atomic.SwapUint32(&u.codeLen, 8)
	atomic.SwapUint32(&u.maxInc, 1<<24-1)
	u.code = []byte{uint8(0)}
}

func (u *UUID) Generate() []byte {
	var back []byte
	copy(back, u.code)

RE_INC:
	if atomic.CompareAndSwapUint32(&u.inc, u.maxInc, u.inc) {
		span := uint32(time.Now().Sub(StartTime).Seconds())
		if atomic.CompareAndSwapUint32(&u.sec, u.sec, span) {
			goto RE_INC
		} else {
			atomic.SwapUint32(&u.inc, 0)
		}
	}
	atomic.SwapUint32(&u.inc, atomic.LoadUint32(&u.inc)+1) // inc

	back = append(back, Uint32ToBytes(atomic.LoadUint32(&u.sec), DoLittleEndian)...)
	switch atomic.LoadUint32(&u.codeLen) {
	case 16:
		back = append(back, Uint16ToBytes(uint16(u.inc), DoLittleEndian)...)
	default:
		back = append(back, Uint32ToBytes(u.inc, DoLittleEndian)...)
	}
	return back
}
