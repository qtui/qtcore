package qtcore

import (
	"fmt"
	"runtime"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtclzsz"
	"github.com/qtui/qtrt"
)

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

func (me *QObject) SetProperty(name string, valx any) bool {
	// valx => QVariant&
	qvar := newQVariant(valx)
	defer qvar.Dtor()
	rv := qtrt.Callany[bool](me, name, qvar)
	return rv
}
func (me *QObject) Property(name string) *QVariant {
	rovp := cgopp.Mallocpg(qtclzsz.Get("QVariant"))
	qtrt.CallanyRov[int](rovp, me, name)
	qvar := QVariantFromptr(rovp)

	runtime.SetFinalizer(qvar, QVariantDtor)
	return qvar
}

// 使用QAnyStringView优化版本
// 不需要new QString, delete QString
func (me *QObject) SetObjectName(name string) {
	//  symname := "__ZN7QObject13setObjectNameE14QAnyStringView"
	name4purego := Fatptrof(&name)
	qtrt.Callany[int](me, name4purego)
}
func (me *QObject) SetObjectNamez2(name string) {
	qtrt.Callany[int](me, name)
}
func (me *QObject) ObjectName() string {
	cthis := cgopp.Malloc(qtclzsz.Get("QString"))
	qtrt.CallanyRov[int](cthis, me)
	qstr := QStringFromptr(cthis)
	defer qstr.Dtor()
	return qstr.ToUtf8().ConstData()
}

func (me *QObject) Parent() *QObject {
	rv := qtrt.Callany[voidptr](me)
	return QObjectFromptr(rv)
}

func (me *QObject) Inherits(classname string) bool {
	rv := qtrt.Callany[bool](me, classname)
	return rv
}

func (me *QObject) MetaObject() *QMetaObject {
	rv := qtrt.Callany[voidptr](me)
	return QMetaObjectFromptr(rv)
}

func (me *QObject) Dbgstr() string {
	// goptr: %p, ptr: %p, class: %s, name: %s
	if me.GetCthis() == nil {
		return fmt.Sprintf("QObject(nil)")
	}
	mo := me.MetaObject()
	clzname := mo.ClassName()
	objname := me.ObjectName()
	s := fmt.Sprintf("QObject(%p){ptr: %v, class: %v, name: %v}", me, me.GetCthis(), clzname, objname)
	return s
}

type Fatptr64 struct {
	H int64
	L int64
}
type Fatptr32 struct {
	H int32
	L int32
}

// purego传递结构做，变长的类型不行
func Fatptrof[T any](ptr *T) any {
	var rv any
	if gopp.UintptrTySz == 4 {
		rv = *((*Fatptr32)(voidptr(ptr)))
	} else {
		rv = *((*Fatptr64)(voidptr(ptr)))
	}
	return rv
}

func (me *QObject) FindChild(objname string) *QObject {
	// QObject* QObject::findChild<QObject*>(QAnyStringView, QFlags<Qt::FindChildOption>) const
	symname := "__ZNK7QObject9findChildIPS_EET_14QAnyStringView6QFlagsIN2Qt15FindChildOptionEE"
	fnsym := qtrt.GetQtSymAddr(symname)
	// objname4c := NewQAnyStringView(objname)
	var objname4c = Fatptrof(&objname)
	rv := cgopp.FfiCall[voidptr](fnsym, me.GetCthis(), objname4c, Qt__FindChildrenRecursively)
	return QObjectFromptr(rv)
}

// ///
type QMetaObject struct {
	*qtrt.CObject
}

func QMetaObjectFromptr(ptr voidptr) *QMetaObject {
	return &QMetaObject{qtrt.CObjectFromptr(ptr)}
}

func (me *QMetaObject) ClassName() string {
	rv := qtrt.Callany[voidptr](me)
	return cgopp.GoString(rv)
}

func (me *QMetaObject) SuperClass() *QMetaObject {
	rv := qtrt.Callany[voidptr](me)
	return QMetaObjectFromptr(rv)
}
