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
	qtrt.Callany(nil, me.GetCthis())
}

// 静态方法，并且返回ROV!!!
func QString_FromUtf8(s string) *QString {
	clzsz := qtclzsz.Get("QString")
	rovp := cgopp.Malloc(clzsz) // todo when to free
	_ = rovp

	rv := qtrt.Callany(rovp, nil, s, len(s))
	// log.Println(rv.Ptr(), s)
	_ = rv
	return QStringFromptr(rovp)
}

func (me *QString) Length() int {
	rvp := qtrt.Callany(nil, me.GetCthis())
	return rvp.Int()
}
