package nsuser

/*
// Compiler flags.
#cgo CFLAGS: -Wall -x objective-c -arch x86_64 -std=gnu99 -fobjc-arc
// Linker flags.
#cgo LDFLAGS: -framework Foundation -framework Cocoa -arch x86_64

#import "nsuser_darwin.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type Notification struct {
	Title    string
	Subtitle string
	// InformativeText is the notification message.
	InformativeText string
	// ContentImage is the primary notification icon.
	ContentImage string
	// SoundName is the name of the sound that fires with a notification.
	SoundName string
}

func (n *Notification) Send() error {
	t := C.CString(n.Title)
	s := C.CString(n.Subtitle)
	i := C.CString(n.InformativeText)
	c := C.CString(n.ContentImage)
	sn := C.CString(n.SoundName)

	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(i))
	defer C.free(unsafe.Pointer(c))
	defer C.free(unsafe.Pointer(sn))

	runtime.LockOSThread()
	C.Send(t, s, i, c, sn)
	runtime.UnlockOSThread()

	return nil
}

func (n *Notification) SetMessage(m string) {
	n.InformativeText = m
}
