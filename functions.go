package w32

var (
	getCurrentProcessId        = kernel32.NewProc("GetCurrentProcessId")
)

func GetCurrentProcessId() DWORD {
	id, _, _ := getCurrentProcessId.Call()
	return DWORD(id)
}