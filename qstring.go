package qtcore

import (
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtrt"
)

type QString struct {
	*qtrt.CObject
}

func QStringFromptr(ptr voidptr) *QString {
	me := &QString{qtrt.CObjectFromptr(ptr)}
	return me
}

func (me *QString) Dtor() {
	qtrt.Callany[int](nil, me)
}

// 静态方法，并且返回ROV!!!
func QString_FromUtf8(s string) *QString {
	clzsz := qtclzsz.Get("QString")
	rovp := cgopp.Malloc(clzsz) // todo when to free
	_ = rovp

	rv := qtrt.Callany[int](rovp, nil, s, len(s))
	// log.Println(rv.Ptr(), s)
	_ = rv
	return QStringFromptr(rovp)
}
func (me *QString) ToUtf8() *QByteArray {
	rovp := cgopp.Malloc(qtclzsz.Get("QByteArray"))
	qtrt.Callany[int](rovp, me)
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
	rvo := qtrt.Callany[voidptr](nil, me)
	rv := cgopp.GoString(rvo)
	return rv
}
func (me *QByteArray) Length() int {
	rvo := qtrt.Callany[int](nil, me)
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
	qtrt.Callany[int](nil, me)
}

func NewQVariant[T any](vx T) *QVariant {
	return newQVariant(vx)
}

func newQVariant(vx any) *QVariant {
	rvo := qtrt.Callany[voidptr](nil, nil, vx)
	return QVariantFromptr(rvo)
}

func (me *QVariant) ToInt() int {
	rvo := qtrt.Callany[int](nil, me, nil)
	return rvo
}
func (me *QVariant) ToLongLong() int64 {
	rvo := qtrt.Callany[int64](nil, me, nil)
	return rvo
}
func (me *QVariant) ToReal() float64 {
	rvo := qtrt.Callany[float64](nil, me, nil)
	return rvo
}
func (me *QVariant) ToString() string {
	rovp := cgopp.Malloc(qtclzsz.Get("QString"))
	qtrt.Callany[int](rovp, me)

	qstr := QStringFromptr(rovp)
	defer qstr.Dtor()
	// log.Println(qstr, qstr.Length())
	ba := qstr.ToUtf8()
	// log.Println(ba)
	// log.Println(ba.Length())
	rv := ba.ConstData()

	return rv
}
