package main

import (
	"fmt"
	"github.com/ebitengine/purego"
	"golang.org/x/sys/windows"
	"runtime"
	"syscall"
	"unsafe"
)

func main() {
	runtime.LockOSThread()
	a := NewApp(0, []string{"", ""})
	qWidget := NewWidget()
	b := NewButton("button", qWidget)
	//	   button.resize(200, 100);
	b.OnPressed(func() {
		fmt.Println("button pressed")
	})
	b.Show()
	qWidget.Show()
	a.Exec()
}

var (
	QApplication_newProc            *syscall.Proc
	QWidget_newProc                 *syscall.Proc
	QWidget_showProc                *syscall.Proc
	QPushButton_newProc             *syscall.Proc
	QPushButton_showProc            *syscall.Proc
	QPushButton_connect_pressedProc *syscall.Proc
	QApplication_execProc           *syscall.Proc
)

func init() {
	//check(windows.SetDllDirectory("plugins/platforms"))
	check(windows.SetDllDirectory("."))

	dll := syscall.MustLoadDLL("purego.dll")
	QApplication_newProc = dll.MustFindProc("QApplication_new")
	QWidget_newProc = dll.MustFindProc("QWidget_new")
	QWidget_showProc = dll.MustFindProc("QWidget_show")
	QPushButton_newProc = dll.MustFindProc("QPushButton_new")
	QPushButton_showProc = dll.MustFindProc("QPushButton_show")
	QPushButton_connect_pressedProc = dll.MustFindProc("QPushButton_connect_pressed")
	QApplication_execProc = dll.MustFindProc("QApplication_exec")
}

type App struct {
	handle uintptr //todo add more fields from C++ class members
}

func NewApp(argc int, argv []string) *App {
	argcPtr := uintptr(unsafe.Pointer(&argc))
	var argvPtr uintptr
	if argv != nil && len(argv) > 0 {
		argvPtr = uintptr(unsafe.Pointer(&argv[0])) // 获取argv的指针
	} else {
		argvPtr = 0 // 或者使用nil
	}
	r1, _, err := QApplication_newProc.Call(argcPtr, argvPtr)
	check(err)
	return &App{handle: r1}
}

func (a *App) Exec() {
	QApplication_execProc.Call(a.handle)
}
func check(err error) {
	if err != nil {
		//panic(err)
	}
}

type QWidget struct {
	handle uintptr //todo add more fields from C++ class members
}

func NewWidget() *QWidget {
	r1, _, err := QWidget_newProc.Call()
	check(err)
	return &QWidget{handle: r1}
}
func (w *QWidget) Show() {
	QWidget_showProc.Call(w.handle)
}

type Button struct {
	handle uintptr //todo add more fields from C++ class members
}

func NewButton(text string, parent *QWidget) *Button {
	title := [256]byte{}
	copy(title[:], []byte(text))
	r1, _, err := QPushButton_newProc.Call(
		uintptr(unsafe.Pointer(&title)),
		uintptr(unsafe.Pointer(parent.handle)),
	)
	check(err)
	return &Button{handle: r1}
}

func (b *Button) Show() {
	QPushButton_showProc.Call(b.handle)
}

func (b *Button) OnPressed(callback func()) {
	cb := func() uintptr {
		if callback != nil {
			callback()
		}
		return 0
	}
	newCallback := purego.NewCallback(cb)
	QPushButton_connect_pressedProc.Call(b.handle, newCallback)
}
