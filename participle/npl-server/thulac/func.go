package thulac

import (
	"encoding/binary"
	"reflect"
	"syscall"
	"unsafe"
)

// CutItem CutItem
type CutItem struct {
	Word string
	Typ  WordType
}

// CutResult CutResult
type CutResult struct {
	Seqs []CutItem
}

// Cut Cut
func Cut(s string, goId int) *CutResult {

	doneWG.Add(1)
	defer doneWG.Done()

	//1 获取上下文
	ctx := <-ctxCh
	defer func(c uintptr) {
		ctxCh <- c
	}(ctx)

	inlen := len(s)

	var outBufferSize int32

	//2 分词
	outBuffer, _, _ := thulacCut.Call(
		ctx,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(s))),
		uintptr(inlen),
		uintptr(unsafe.Pointer(&outBufferSize)))
	defer thulacFreeResult.Call(uintptr(outBuffer))

	var bb []byte
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bb)) // case 1
	hdr.Data = outBuffer                               // case 6 (this case)
	hdr.Len = int(outBufferSize)
	hdr.Cap = int(outBufferSize)

	// 分词结果数量
	cnt := binary.LittleEndian.Uint32(bb)

	// 结果
	results := &CutResult{
		Seqs: make([]CutItem, 0, cnt),
	}

	// 结果 buffer 的开始
	buf := bb[4:]
	for cnt > 0 {
		cnt--

		i := 0
		for i < len(buf) && buf[i] != 0 {
			i++
		}

		s := string(buf[:i])
		i++

		t := buf[i]
		i++

		results.Seqs = append(results.Seqs, CutItem{
			Word: s,
			Typ:  WordType(t),
		})

		buf = buf[i:]
	}

	return results
}
