package qtcore

import "github.com/qtui/qtrt"

type QObject struct {
	*qtrt.CObject
}

func QObjectFromptr(ptr voidptr) *QObject {
	me := &QObject{qtrt.CObjectFromptr(ptr)}
	return me
}

func (me *QObject) SetCthis(ptr voidptr) {
	me.CObject = qtrt.CObjectFromptr(ptr)
}

func (me *QObject) ObjectName() string {

	return ""
}
