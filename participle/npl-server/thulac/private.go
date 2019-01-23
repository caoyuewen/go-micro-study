package thulac

import (
	"sync"
	"syscall"
	"unsafe"
)

var (
	dll              *syscall.DLL
	thulacInit       *syscall.Proc
	thulacDestory    *syscall.Proc
	thulacGetCtx     *syscall.Proc
	thulacFreeCtx    *syscall.Proc
	thulacCut        *syscall.Proc
	thulacFreeResult *syscall.Proc
)

var (
	ctxCh  chan uintptr
	doneWG sync.WaitGroup
)

// Init Init
func Init(workers int) {
	dll = syscall.MustLoadDLL("libthulac.dll")

	/*
	   __declspec(dllexport) void thulac_init(const char *path, int maxWorkers);
	   __declspec(dllexport) void thulac_destory();

	   __declspec(dllexport) void *thulac_get_ctx();
	   __declspec(dllexport) void thulac_free_ctx(void *ctx);

	     // nullptr: failed
	     // !null: result
	   __declspec(dllexport) const char *thulac_cut(void *ctx, const char *seg, int len, int32_t *outBufferSize);
	   __declspec(dllexport) void  thulac_free_result(const char *res);
	*/

	thulacInit = dll.MustFindProc("thulac_init")
	thulacDestory = dll.MustFindProc("thulac_destory")
	thulacGetCtx = dll.MustFindProc("thulac_get_ctx")
	thulacFreeCtx = dll.MustFindProc("thulac_free_ctx")
	thulacCut = dll.MustFindProc("thulac_cut")
	thulacFreeResult = dll.MustFindProc("thulac_free_result")

	path := ""
	if workers < 1 {
		workers = 1
	}
	thulacInit.Call(uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(workers))

	ctxCh = make(chan uintptr, workers)

	for len(ctxCh) < workers {
		ctx, _, _ := thulacGetCtx.Call()
		if ctx != 0 {
			ctxCh <- ctx
		}
	}
}

// Destory Destory
func Destory() {

	doneWG.Wait()

	for len(ctxCh) > 0 {
		ctx := <-ctxCh
		thulacFreeCtx.Call(ctx)
	}

	thulacDestory.Call()

	dll.Release()
}
