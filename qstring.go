package qtcore

import (
	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtrt"
)

// 这个文件是非QObject的, QString, QVariant, QUrl

type QString struct {
	*qtrt.CObject
}

func QStringFromptr(ptr voidptr) *QString {
	me := &QString{qtrt.CObjectFromptr(ptr)}
	return me
}

func (me *QString) Dtor() {
	qtrt.Callany[int](me)
}

func QStringDtor(me *QString) { me.Dtor() }

// 静态方法，并且返回ROV!!!
func QString_FromUtf8(s string) *QString {
	clzsz := qtclzsz.Get("QString")
	rovp := cgopp.Malloc(clzsz) // todo when to free
	_ = rovp

	rv := qtrt.CallanyRov[int](rovp, nil, s, len(s))
	// log.Println(rv.Ptr(), s)
	_ = rv
	return QStringFromptr(rovp)
}
func (me *QString) ToUtf8() *QByteArray {
	rovp := cgopp.Malloc(qtclzsz.Get("QByteArray"))
	qtrt.CallanyRov[int](rovp, me)
	return QByteArrayFromptr(rovp)
}

func (me *QString) Length() int {
	rv := qtrt.Callany[int](nil, me)
	return rv
}

// ////
type QByteArray struct {
	*qtrt.CObject
}

func QByteArrayFromptr(ptr voidptr) *QByteArray {
	me := &QByteArray{qtrt.CObjectFromptr(ptr)}
	return me
}
func (me *QByteArray) Dtor() {
	qtrt.Callany[int](nil, me)
}
func (me *QByteArray) ConstData() string {
	rvo := qtrt.Callany[voidptr](me)
	rv := cgopp.GoString(rvo)
	return rv
}
func (me *QByteArray) Length() int {
	rvo := qtrt.Callany[int](me)
	return rvo
}

// ///
type QVariant struct {
	*qtrt.CObject
}

func QVariantFromptr(ptr voidptr) *QVariant {
	me := &QVariant{qtrt.CObjectFromptr(ptr)}
	return me
}

func (me *QVariant) Dtor() {
	qtrt.Callany0(me)
}

func QVariantDtor(me *QVariant) { me.Dtor() }

func NewQVariant[T any](vx T) *QVariant {
	return newQVariant(vx)
}

func newQVariant(vx any) *QVariant {
	rv := qtrt.Callany[voidptr](nil, vx)
	return QVariantFromptr(rv)
}

func (me *QVariant) ToInt() int {
	rv := qtrt.Callany[int](me, nil)
	return rv
}
func (me *QVariant) ToLongLong() int64 {
	rv := qtrt.Callany[int64](me, nil)
	return rv
}
func (me *QVariant) ToReal() float64 {
	rv := qtrt.Callany[float64](me, nil)
	return rv
}
func (me *QVariant) ToString() string {
	rovp := cgopp.Malloc(qtclzsz.Get("QString"))
	qtrt.CallanyRov[int](rovp, me)

	qstr := QStringFromptr(rovp)
	defer qstr.Dtor()
	// log.Println(qstr, qstr.Length())
	ba := qstr.ToUtf8()
	// log.Println(ba)
	// log.Println(ba.Length())
	rv := ba.ConstData()

	return rv
}

// 一般用作传递参数，而且是以结构体的形式传递，而不是引用或者指针传递
// 看了一没有引用和指针的用法，全是QAnyStringView传递，这里也可以不用指针了
// 注意字符串引用，不拷贝
// 难道和go的string结构体相同了？确认确实结构相同了!!!
type QAnyStringView struct {
	Ptr voidptr
	Len usize
}

// type QAnyStringViewITF interface {
// 	QAnyStringViewPTR() *QAnyStringView
// }

// func (me *QAnyStringView) QAnyStringViewITF() *QAnyStringView { return me }

// func QAnyStringViewFromptr(ptr voidptr) *QAnyStringView {
// 	return &QAnyStringView{ptr, 0}
// }

func NewQAnyStringView(s string) *QAnyStringView {
	// QAnyStringView::QAnyStringView<char, true>(char const*, long long)
	// symname := "__ZN14QAnyStringViewC2IcLb1EEEPKT_x" // template
	// fnsym := qtrt.GetQtSymAddr(symname)
	// cthis := cgopp.Malloc(qtclzsz.Get("QAnyStringView"))
	// s4c := cgopp.string
	// cgopp.FfiCall[int](cthis)
	ptr := (voidptr)(&s)
	// return QAnyStringViewFromptr(ptr)
	return (*QAnyStringView)(ptr)
}

// type QAnyStringView2 string

func (me *QAnyStringView) Size() int {
	// rv := qtrt.Callany[int](me)
	// return rv
	return int(me.Len)
}
func (me *QAnyStringView) Data() voidptr {
	// rv := qtrt.Callany[voidptr](me)
	// return rv
	return me.Ptr
}

type QUrl struct {
	*qtrt.CObject
}

func QUrlFromptr(ptr voidptr) *QUrl {
	return &QUrl{qtrt.CObjectFromptr(ptr)}
}

func (me *QUrl) Dtor() { qtrt.Callany0(me) }

func NewQUrl(u string, mode ...int) *QUrl {
	rv := qtrt.Callany[voidptr](nil, u, gopp.FirstofGv(mode))
	return QUrlFromptr(rv)
}
