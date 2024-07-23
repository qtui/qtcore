package qtcore

import (
	"log"
	"testing"

	"github.com/qtui/qtsyms"
)

// 指定搜索 QtxxxInline 库的路径
// LD_LIBRARY_PATH=../../minqtgo/srcc go test -v

func TestQSVB1(t *testing.T) {

	qtsyms.LoadAllQtSymbols() // 这个初始化现在是必须的

	qstr := QString_FromUtf8("hehhee")
	log.Println(qstr.Length())
	if qstr.GetCthis() == nil {
		t.Error(qstr)
	}
	if qstr.Length() != 6 {
		t.Error(qstr)
	}
	qstr.Dtor()

	qvar := NewQVariant(12345)
	log.Println(qvar)
	log.Println(qvar.ToInt())
	log.Println(qvar.ToLongLong())
	if qvar.GetCthis() == nil {
		t.Error(qvar)
	}
	if qvar.ToInt() != 12345 {
		t.Error(qvar)
	}
	if qvar.ToLongLong() != 12345 {
		t.Error(qvar)
	}
	qvar.Dtor()

	qvar = NewQVariant("abcdefg")
	log.Println(qvar)
	log.Println(qvar.ToString())
	if qvar.GetCthis() == nil {
		t.Error(qvar)
	}
	if qvar.ToString() != "abcdefg" {
		t.Error(qvar)
	}
	qvar.Dtor()

	qvar = NewQVariant(123.456)
	log.Println(qvar.ToReal())
	if qvar.GetCthis() == nil {
		t.Error(qvar)
	}
	if qvar.ToReal() != 123.456 {
		t.Error()
	}
}
