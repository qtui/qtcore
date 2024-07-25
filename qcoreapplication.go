package qtcore

import (
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
	return QCoreApplication_instance()
}
func (me *QCoreApplication) Instance() *QCoreApplication {
	return QCoreApplication_instance()
}
func QCoreApplication_instance() *QCoreApplication {
	rv := qtrt.Callany[voidptr](nil)
	return QCoreApplicationFromptr(rv)
}

type QGuiApplication struct {
	*QCoreApplication
}

func QGuiApplicationFromptr(ptr voidptr) *QGuiApplication {
	me := &QGuiApplication{QCoreApplicationFromptr(ptr)}
	return me
}

// func NewQGuiApplication()

func (me *QGuiApplication) ApplicationState() int {
	return QGuiApplication_applicationState()
}

func QGuiApplication_applicationState() int {
	return qtrt.Callany[int](nil)
}

func (me *QGuiApplication) platformName() string {
	return QGuiApplication_platformName()
}

func QGuiApplication_platformName() string {
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
