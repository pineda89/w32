package w32

import "unsafe"

var (
	getCurrentProcessId        = kernel32.NewProc("GetCurrentProcessId")

	sendInput                        = moduser32.NewProc("SendInput")
)

func GetCurrentProcessId() DWORD {
	id, _, _ := getCurrentProcessId.Call()
	return DWORD(id)
}

func SendInput(inputs ...INPUT) uint32 {
	if len(inputs) == 0 {
		return 0
	}
	ret, _, _ := sendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		unsafe.Sizeof(inputs[0]),
	)
	return uint32(ret)
}