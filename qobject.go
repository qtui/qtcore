package qtcore

import "github.com/qtui/qtrt"

type QObject struct {
	*qtrt.CObject
}

type QObjectITF interface {
	QObjectPTR() *QObject
}

func (me *QObject) QObjectPTR() *QObject { return me }

func QObjectFromptr(ptr voidptr) *QObject {
	me := &QObject{qtrt.CObjectFromptr(ptr)}
	return me
}

func (me *QObject) Dtor() {
	qtrt.Callany0(me)
}
func (me *QObject) DeleteLater() {
	qtrt.Callany0(me)
}

func (me *QObject) SetCthis(ptr voidptr) {
	me.CObject = qtrt.CObjectFromptr(ptr)
}

func (me *QObject) SetObjectName(name string) {
	qtrt.Callany[int](me, name)
}
func (me *QObject) ObjectName() string {
	rv := qtrt.Callany[voidptr](me)
	return QStringFromptr(rv).ToUtf8().ConstData()
}

func (me *QObject) Parent() *QObject {
	rv := qtrt.Callany[voidptr](me)
	return QObjectFromptr(rv)
}

func (me *QObject) Inherits(classname string) bool {
	rv := qtrt.Callany[bool](me, classname)
	return rv
}

func (me *QObject) dynamicMetaObject() *QMetaObject {
	rv := qtrt.Callany[voidptr](me)
	return QMetaObjectFromptr(rv)
}

// ///
type QMetaObject struct {
	*qtrt.CObject
}

func QMetaObjectFromptr(ptr voidptr) *QMetaObject {
	return &QMetaObject{qtrt.CObjectFromptr(ptr)}
}
