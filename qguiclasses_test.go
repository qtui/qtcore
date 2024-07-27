package qtcore

import (
	"log"
	"os"
	"testing"
	"unsafe"

	"github.com/kitech/gopp"
	"github.com/qtui/qtsyms"
)

func init() {
	log.Println(gopp.Retn(os.Getwd()))
	qtsyms.LoadAllQtSymbols()
}

func TestColor1(t *testing.T) {

	v := QColor{}

	v2 := byte(0)
	log.Println(unsafe.Alignof(v2))
	log.Println(unsafe.Offsetof(v.Spec))
	log.Println(unsafe.Offsetof(v.Alpha))
	log.Println(unsafe.Offsetof(v.Red))
	log.Println(unsafe.Offsetof(v.Green))
	log.Println(unsafe.Offsetof(v.Blue))
	log.Println(unsafe.Offsetof(v.pad))
	log.Println(unsafe.Offsetof(v.padding))
	stsz := unsafe.Sizeof(v)
	if stsz != 16 {
		t.Error("struct size invalid", stsz, "want", 16)
	}
}

func TestColor2(t *testing.T) {
	c := NewQColor("darkblue")
	log.Println(c)

	c = NewQColorz2("darkgreen")
	log.Println(c)
}

// not crash is ok
func TestColor3(t *testing.T) {
	lst := QColor_ColorNames()
	// log.Println(lst.Size())
	// 测试的时候大概148
	if lst.Size() > 0 && lst.Size() < 12345 {
	} else {
		t.Error("size invalid", lst.Size())
	}
	for i := 0; i < lst.Size(); i++ {
		qstr := lst.QStringAt(i)
		_ = qstr.Length()
		_ = qstr.ToUtf8().ConstData()
		if false {
			log.Println(i, qstr.Length(), qstr.ToUtf8().ConstData())
		}
	}

}
