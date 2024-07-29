package qtcore

import (
	"runtime"

	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtrt"
)

type QCoreApplication struct {
	*QObject
}

func QCoreApplicationFromptr(ptr voidptr) *QCoreApplication {
	me := &QCoreApplication{QObjectFromptr(ptr)}
	return me
}
func QApp() *QCoreApplication {
	return QCoreApplication_Instance()
}
func (me *QCoreApplication) Instance() *QCoreApplication {
	return QCoreApplication_Instance()
}
func QCoreApplication_Instance() *QCoreApplication {
	rv := qtrt.Callany[voidptr](nil)
	return QCoreApplicationFromptr(rv)
}
func (me *QCoreApplication) Exec() int {
	rv := qtrt.Callany[int](me)
	return rv
}
func (me *QCoreApplication) Arguments() *QList {
	return QCoreApplication_Arguments()
}
func QCoreApplication_Arguments() *QList {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QList"))
	qtrt.CallanyRov[int](rovp, nil)

	rv := QListFromptr(rovp)
	runtime.SetFinalizer(rv, QStringListDtor)
	return rv
}

type QGuiApplication struct {
	*QCoreApplication
}

func QGuiApplicationFromptr(ptr voidptr) *QGuiApplication {
	me := &QGuiApplication{QCoreApplicationFromptr(ptr)}
	return me
}
func NewQGuiApplication(argv []string, f int) *QGuiApplication {
	rv := qtrt.Callany[voidptr](nil, len(argv), argv, f)
	return QGuiApplicationFromptr(rv)
}

// func NewQGuiApplication()

func (me *QGuiApplication) ApplicationState() int {
	return QGuiApplication_ApplicationState()
}

func QGuiApplication_ApplicationState() int {
	return qtrt.Callany[int](nil)
}

func (me *QGuiApplication) PlatformName() string {
	return QGuiApplication_PlatformName()
}

func QGuiApplication_PlatformName() string {
	rovp := cgopp.Malloc(qtclzsz.Get("QString"))
	qtrt.CallanyRov[int](rovp, nil)
	qstr := QStringFromptr(rovp)
	defer qstr.Dtor()
	return qstr.ToUtf8().ConstData()
}

// todo
var QDesktopServices = &QDesktopServices_{}

type QDesktopServices_ struct{}

func (me *QDesktopServices_) OpenUrl(u string) {
	QDesktopServices_openUrl(u)
}
func QDesktopServices_openUrl(u string) {
	uo := NewQUrl(u)
	qtrt.Callany0(nil, uo)
}
