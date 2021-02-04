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

func SendInput(inputs ...INPUT2) uint32 {
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

// INPUT is used in SendInput. To create a concrete INPUT type, use the helper
// functions MouseInput, KeyboardInput and HardwareInput. These are necessary
// because the C API uses a union here, which Go does not provide.
type INPUT2 struct {
	Type uint32
	// use MOUSEINPUT for the union because it is the largest of all allowed
	// structures
	mouse MOUSEINPUT2
}

type MOUSEINPUT2 struct {
	Dx        int32
	Dy        int32
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type KEYBDINPUT2 struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type HARDWAREINPUT2 struct {
	Msg    uint32
	ParamL uint16
	ParamH uint16
}

func MouseInput2(input MOUSEINPUT2) INPUT2 {
	return INPUT2{
		Type:  INPUT_MOUSE,
		mouse: input,
	}
}

func KeyboardInput2(input KEYBDINPUT2) INPUT2 {
	return INPUT2{
		Type:  INPUT_KEYBOARD,
		mouse: *((*MOUSEINPUT2)(unsafe.Pointer(&input))),
	}
}

func HardwareInput2(input HARDWAREINPUT2) INPUT2 {
	return INPUT2{
		Type:  INPUT_HARDWARE,
		mouse: *((*MOUSEINPUT2)(unsafe.Pointer(&input))),
	}
}