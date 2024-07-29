package qtcore

import (
	"runtime"

	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtrt"
)

// 正好16字节 = gopp.Fatptr64
// 拆开之后go做改变了，变成18了
// 只好这样了
type QColor struct {
	Spec byte // 不管用 `binary:"spec" align:"4"`

	// union in cpp
	Alpha  byte // uint16
	Alpha1 byte
	Red    byte // uint16
	Red1   byte
	Green  byte // uint16
	Green1 byte
	Blue   byte // uint16
	Blue1  byte
	pad    byte // uint16
	pad1   byte

	// padding
	padding [5]byte
}

func (me *QColor) GetCthis() voidptr { return voidptr(me) }

func (me *QColor) Dtor() {
	// cgopp.Cfreepg((voidptr)(me))
	qtrt.Callany0(me)
}

func NewQColorz2(name string) QColor {
	rv := qtrt.Callany[voidptr](nil, name)
	tv := (*QColor)(rv)
	defer tv.Dtor()
	me := *tv
	return me
}

// QStringView
func NewQColorz3(name string) QColor {
	sv := NewQAnyStringView(name)
	rv := qtrt.Callany[voidptr](nil, sv)
	tv := (*QColor)(rv)
	defer tv.Dtor()
	me := *tv
	return me
}

// QStringView no ctor no malloc
func NewQColor(name string) QColor {
	return QColor0.FromString(name)
}

var QColor0 *QColor

func (me *QColor) New(name string) QColor {
	return QColor_FromString(name)
}
func (me *QColor) FromString(name string) QColor {
	return QColor_FromString(name)
}
func (me *QColor) ColorNames() *QList {
	return QColor_ColorNames()
}

func QColor_FromString(name string) QColor {
	fp := Fatptrof(NewQAnyStringView(name))
	rv := qtrt.Callany[QColor](nil, fp)
	return rv
}

func QColor_ColorNames() *QList {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QList"))
	rvx := qtrt.CallanyRov[voidptr](rovp, nil)

	rv := QListFromptr(rvx)
	runtime.SetFinalizer(rv, QStringListDtor)
	return rv
}

type QPoint struct {
}
type QPointF struct {
}

type QWindow struct {
	*QObject
}
type QWindowITF interface {
	QWindowPTR() *QWindow
}

func (me *QWindow) QWindowPTR() *QWindow { return me }

func QWindowFromptr(ptr voidptr) *QWindow {
	return &QWindow{QObjectFromptr(ptr)}
}

func (me *QWindow) Show() { qtrt.Callany0(me) }
