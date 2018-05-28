package main

import (
	"fmt"
	"time"

	"github.com/robjporter/go-utils/go/as"
)

func main() {
	time.Sleep(1 * time.Second)
	astostring()
	time.Sleep(1 * time.Second)
	astrimmed()
	time.Sleep(1 * time.Second)
	astofloat()
	time.Sleep(1 * time.Second)
	astorunelength()
	time.Sleep(1 * time.Second)
	astobool()
	time.Sleep(1 * time.Second)
	astobytes()
	time.Sleep(1 * time.Second)
	astoslice()
	time.Sleep(1 * time.Second)
	astoint()
	time.Sleep(1 * time.Second)
	astoip()
	time.Sleep(1 * time.Second)
	astobase64()
	time.Sleep(1 * time.Second)
	asfrombase64()
	time.Sleep(1 * time.Second)
	asisempty()
	time.Sleep(1 * time.Second)
	asiskind()
	time.Sleep(1 * time.Second)
	asofkind()
	time.Sleep(1 * time.Second)
	asoftype()
	time.Sleep(1 * time.Second)
	astotime()
	time.Sleep(1 * time.Second)
	astoduration()
	time.Sleep(1 * time.Second)
	astofixedlengthafter()
	time.Sleep(1 * time.Second)
	astofixedlengthbefore()
	time.Sleep(1 * time.Second)
	astofixedlengthcenter()
	time.Sleep(1 * time.Second)
	asisint()
	time.Sleep(1 * time.Second)
	asisbool()
	time.Sleep(1 * time.Second)
	asisfloat()
	time.Sleep(1 * time.Second)
	asisstring()
	time.Sleep(1 * time.Second)
	asistime()
	time.Sleep(1 * time.Second)
	asisnillable()
	time.Sleep(1 * time.Second)
	astoformattedbytes()
}

func astostring() {
	fmt.Println("")
	fmt.Println("AS TO STRING *******************************************************")
	fmt.Println("STRING: (32)                        >", `"`+as.ToString(32)+`"`)
	fmt.Println("STRING: (true)                      >", `"`+as.ToString(bool(true))+`"`)
	fmt.Println("STRING: ('mayonegg')                >", `"`+as.ToString("mayonegg")+`"`)         // "mayonegg"
	fmt.Println("STRING: (8)                         >", `"`+as.ToString(8)+`"`)                  // "8"
	fmt.Println("STRING: (8.31)                      >", `"`+as.ToString(8.31)+`"`)               // "8.31"
	fmt.Println("STRING: ([]byte('one time'))        >", `"`+as.ToString([]byte("one time"))+`"`) // "one time"
	fmt.Println("STRING: (nil)                       >", `"`+as.ToString(nil)+`"`)                // ""
	var foo interface{} = "one more time"
	fmt.Println("STRING: (interface{'one more time}) >", `"`+as.ToString(foo)) // "one more time"
}
func astrimmed() {
	fmt.Println("")
	fmt.Println("AS TRIMMED *******************************************************")
	fmt.Println("TRIMMED: ('    TEST      ')         >", `"`+as.Trimmed("    TEST      ")+`"`)
}
func astofloat() {
	fmt.Println("")
	fmt.Println("AS TO FLOAT *******************************************************")
	fmt.Println("FLOAT: (32.4400)                    >", as.ToFloat(32.4400))
	fmt.Println("FLOAT32: (32.4400)                  >", as.ToFloat32(32.4400))
}
func astorunelength() {
	fmt.Println("")
	fmt.Println("AS TO RUNE LENGTH*******************************************************")
	fmt.Println("RUNELENGTH: ('test')                >", as.ToRuneLength("test"))
	fmt.Println("RUNELENGTH: ('TEST')                >", as.ToRuneLength("TEST"))
	fmt.Println("RUNELENGTH: ('iiii')                >", as.ToRuneLength("iiii"))
	fmt.Println("RUNELENGTH: ('QQKK')                >", as.ToRuneLength("QQKK"))
	fmt.Println("RUNELENGTH: ('Lllm')                >", as.ToRuneLength("Lllm"))
}
func astobool() {
	fmt.Println("")
	fmt.Println("AS TO BOOL *******************************************************")
	fmt.Println("BOOL: (1)                           >", as.ToBool(1))
	fmt.Println("BOOL: (0)                           >", as.ToBool(0))
	fmt.Println("BOOL: ('1')                         >", as.ToBool("1"))
	fmt.Println("BOOL: ('true')                      >", as.ToBool("true"))
	fmt.Println("BOOL: ('down')                      >", as.ToBool("down"))
}
func astobytes() {
	fmt.Println("")
	fmt.Println("AS TO BYTES *******************************************************")
	fmt.Println("BYTES: ('Testing')                  >", as.ToBytes("Testing"))
}
func astoslice() {
	fmt.Println("")
	fmt.Println("AS TO SLICE *******************************************************")
	var foo2 []interface{}
	foo2 = append(foo2, "one") //more time"
	fmt.Println("SLICE: ('one')                      >", as.ToSlice(foo2))
}
func astoint() {
	fmt.Println("")
	fmt.Println("TO INT *******************************************************")
	fmt.Println("INT: ('1')                          >", as.ToInt("1"))
	fmt.Println("INT64: ('1')                        >", as.ToInt64("1"))
	fmt.Println("INT32: ('1')                        >", as.ToInt32("1"))
	fmt.Println("INT16: ('1')                        >", as.ToInt16("1"))
	fmt.Println("INT8: ('1')                         >", as.ToInt8("1"))
}

func astoip() {
	fmt.Println("")
	fmt.Println("TO IP *******************************************************")
	fmt.Println("IP ADDRESS: ('192.168.0.1')          >", as.ToIP("192.168.0.1"))   // "one more time"
	fmt.Println("IP ADDRESS: ('one more time')        >", as.ToIP("one more time")) //
	fmt.Println("IP ADDRESS: ('1')                    >", as.ToIP("1"))             // "one more time"
	fmt.Println("IP ADDRESS: ('1.0')                  >", as.ToIP("1.0"))           // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0')                >", as.ToIP("1.0.0"))         // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0.0/8')            >", as.ToIP("1.0.0.0/8"))     // "one more time"
}

func astobase64() {
	fmt.Println("")
	fmt.Println("TO BASE64 *******************************************************")
	fmt.Println("TOBASE64: ('This is a test')         >", as.ToBase64("This is a test"))
}

func asfrombase64() {
	fmt.Println("")
	fmt.Println("FROM BASE64 *******************************************************")
	fmt.Println("FROMBASE64: ('VGhpcyBpcyBhIHRlc3Q=') >", as.FromBase64("VGhpcyBpcyBhIHRlc3Q="))
}

func asisempty() {
	fmt.Println("")
	fmt.Println("AS IS EMPTY *******************************************************")
	fmt.Println("IP EMPTY: ('0')                      >", as.IsEmpty(0))
	fmt.Println("IP EMPTY: ('1')                      >", as.IsEmpty(1))
	fmt.Println("IP EMPTY: ('')                       >", as.IsEmpty(""))
	fmt.Println("IP EMPTY: ('sdasdass')               >", as.IsEmpty("sdasdass"))
	fmt.Println("IP EMPTY: ('[]string{}')             >", as.IsEmpty([]string{}))
}

func asiskind() {
	fmt.Println("")
	fmt.Println("AS IS KIND *******************************************************")
	fmt.Println("IS KIND: (string,0)                  >", as.IsKind("string", 0))
	fmt.Println("IS KIND: (string,'')                 >", as.IsKind("string", ""))
	fmt.Println("IS KIND: (int,0)                     >", as.IsKind("int", 0))
	fmt.Println("IS KIND: (int,'test')                >", as.IsKind("int", "test"))
}

func asofkind() {
	fmt.Println("")
	fmt.Println("AS OF KIND *******************************************************")
	fmt.Println("KIND OF: ('string')                  >", as.OfKind("string"))
	fmt.Println("KIND OF: ([]string{})                >", as.OfKind([]string{}))
	fmt.Println("KIND OF: (nil)                       >", as.OfKind(nil))
	fmt.Println("KIND OF: ([]byte('one time))         >", as.OfKind([]byte("one time")))
	fmt.Println("KIND OF: (bool(true))                >", as.OfKind(bool(true)))
	fmt.Println("KIND OF: (32)                        >", as.OfKind(32))
}
func asoftype() {
	fmt.Println("")
	fmt.Println("AS OF TYPE *******************************************************")
	fmt.Println("TYPE: (32)                           >", as.OfType(32))
	fmt.Println("TYPE: ('')                           >", as.OfType(""))
	fmt.Println("TYPE: ([]string{}])                  >", as.OfType([]string{}))
	fmt.Println("TYPE: (true)                         >", as.OfType(true))
	fmt.Println("TYPE: (1.0f)                         >", as.OfType(1.00))
	fmt.Println("TYPE: (int64(22))                    >", as.OfType(int64(22)))
}

func astotime() {
	fmt.Println("")
	fmt.Println("AS TO TIME *******************************************************")
	fmt.Println("TIME: ('2016-04-04')                 >", as.ToTime(false, "2016-04-04"))
	fmt.Println("TIME: ('04-04-2016')                 >", as.ToTime(false, "04-04-2016"))
	fmt.Println("TIME: ('2016-04-04 16:20:40')        >", as.ToTime(false, "2016-04-04 16:20:40"))
	fmt.Println("TIME: ('2016-04-04 16:20:40 +1 BST') >", as.ToTime(false, "2016-04-04 16:20:40 +1 BST"))
	t1 := time.Now()
	fmt.Println("TIME: NOW TO INT                     >", as.FromTime(t1))
	fmt.Println("TIME: INT TO TIME                    >", as.ToTime(true, as.FromTime(t1)))
}
func astoduration() {
	fmt.Println("")
	fmt.Println("AS TO DURATION *******************************************************")
	fmt.Println("DURATION: (1h44m)                    >", as.ToDuration("1h44m"))
	fmt.Println("DURATION: (44)                       >", as.ToDuration("44"))
	fmt.Println("DURATION: (44s)                      >", as.ToDuration("44s"))
	fmt.Println("DURATION: (444h)                     >", as.ToDuration("444h"))
	fmt.Println("DURATION: (88m)                      >", as.ToDuration("88m"))
}

func astofixedlengthafter() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH AFTER *******************************************************")
	fmt.Println("FIXED LENGTH AFTER (*,20):           >", as.ToFixedLengthAfter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH AFTER (-,50):           >", as.ToFixedLengthAfter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH AFTER (*,10):           >", as.ToFixedLengthAfter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH AFTER (*,8):            >", as.ToFixedLengthAfter("Test String", "*", 8))
}

func astofixedlengthbefore() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH BEFORE *******************************************************")
	fmt.Println("FIXED LENGTH BEFORE (*,20):          >", as.ToFixedLengthBefore("Test String", "*", 20))
	fmt.Println("FIXED LENGTH BEFORE (-,50):          >", as.ToFixedLengthBefore("Test String", "-", 50))
	fmt.Println("FIXED LENGTH BEFORE (*,10):          >", as.ToFixedLengthBefore("Test String", "*", 10))
	fmt.Println("FIXED LENGTH BEFORE (*,8):           >", as.ToFixedLengthBefore("Test String", "*", 8))
}

func astofixedlengthcenter() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH CENTER *******************************************************")
	fmt.Println("FIXED LENGTH CENTER (*,20):          >", as.ToFixedLengthCenter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH CENTER (-,50):          >", as.ToFixedLengthCenter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH CENTER (*,10):          >", as.ToFixedLengthCenter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH CENTER (*,8):           >", as.ToFixedLengthCenter("Test String", "*", 8))

}

func asisint() {
	fmt.Println("")
	fmt.Println("AS IS INT *******************************************************")
	fmt.Println("INT: (44)                           >", as.IsInt(44))
	fmt.Println("INT: (true)                         >", as.IsInt(true))
	fmt.Println("INT: (44.44)                        >", as.IsInt(44.44))
	fmt.Println("INT: ('test')                       >", as.IsInt("test"))
	fmt.Println("INT: ('14:14:14')                   >", as.IsInt(as.ToTime(false, "14:14:14")))
}

func asisbool() {
	fmt.Println("")
	fmt.Println("AS IS BOOL *******************************************************")
	fmt.Println("BOOL: (44)                           >", as.IsBool(44))
	fmt.Println("BOOL: (true)                         >", as.IsBool(true))
	fmt.Println("BOOL: (44.44)                        >", as.IsBool(44.44))
	fmt.Println("BOOL: ('test')                       >", as.IsBool("test"))
	fmt.Println("BOOL: ('14:14:14')                   >", as.IsBool(as.ToTime(false, "14:14:14")))
}

func asisfloat() {
	fmt.Println("")
	fmt.Println("AS IS FLOAT *******************************************************")
	fmt.Println("FLOAT: (44)                           >", as.IsFloat(44))
	fmt.Println("FLOAT: (true)                         >", as.IsFloat(true))
	fmt.Println("FLOAT: (44.44)                        >", as.IsFloat(44.44))
	fmt.Println("FLOAT: ('test')                       >", as.IsFloat("test"))
	fmt.Println("FLOAT: ('14:14:14')                   >", as.IsFloat(as.ToTime(false, "14:14:14")))
}

func asisstring() {
	fmt.Println("")
	fmt.Println("AS IS STRING *******************************************************")
	fmt.Println("STRING: (44)                         >", as.IsString(44))
	fmt.Println("STRING: (true)                       >", as.IsString(true))
	fmt.Println("STRING: (44.44)                      >", as.IsString(44.44))
	fmt.Println("STRING: ('test')                     >", as.IsString("test"))
	fmt.Println("STRING: ('14:14:14')                 >", as.IsString(as.ToTime(false, "14:14:14")))
}

func asistime() {
	fmt.Println("")
	fmt.Println("AS IS TIME *******************************************************")
	fmt.Println("TIME: (44)                           >", as.IsTime(44))
	fmt.Println("TIME: (true)                         >", as.IsTime(true))
	fmt.Println("TIME: (44.44)                        >", as.IsTime(44.44))
	fmt.Println("TIME: ('test')                       >", as.IsTime("test"))
	fmt.Println("TIME: ('14:14:14')                   >", as.IsTime(as.ToTime(false, "14:14:14")))
}

func asisnillable() {
	fmt.Println("")
	fmt.Println("AS IS NILLABLE *******************************************************")
	fmt.Println("NILLABLE: ('')                       >", as.IsNillable(""))
	fmt.Println("NILLABLE: ([]string{})               >", as.IsNillable([]string{}))
}

func astoformattedbytes() {
	fmt.Println("")
	fmt.Println("AS TO FORMATTED BYTES *******************************************************")
	fmt.Println("FORMAT: (44)                       >", as.FormatIntToByte(44))
	fmt.Println("FORMAT: (444)                      >", as.FormatIntToByte(444))
	fmt.Println("FORMAT: (4444)                     >", as.FormatIntToByte(4444))
	fmt.Println("FORMAT: (44444)                    >", as.FormatIntToByte(44444))
	fmt.Println("FORMAT: (444444)                   >", as.FormatIntToByte(444444))
	fmt.Println("FORMAT: (4444444)                  >", as.FormatIntToByte(4444444))
	fmt.Println("FORMAT: (44444444)                 >", as.FormatIntToByte(44444444))
	fmt.Println("FORMAT: (444444444)                >", as.FormatIntToByte(444444444))
	fmt.Println("FORMAT: (4444444444)               >", as.FormatIntToByte(4444444444))
	fmt.Println("FORMAT: (44444444444)              >", as.FormatIntToByte(44444444444))
	fmt.Println("FORMAT: (444444444444)             >", as.FormatIntToByte(444444444444))
	fmt.Println("FORMAT: (4444444444444)            >", as.FormatIntToByte(4444444444444))
	fmt.Println("FORMAT: (44444444444444)           >", as.FormatIntToByte(44444444444444))
	fmt.Println("FORMAT: (444444444444444)          >", as.FormatIntToByte(444444444444444))
	fmt.Println("FORMAT: (4444444444444444)         >", as.FormatIntToByte(4444444444444444))
	fmt.Println("FORMAT: (44444444444444444)        >", as.FormatIntToByte(44444444444444444))
	fmt.Println("FORMAT: (444444444444444444)       >", as.FormatIntToByte(444444444444444444))
	fmt.Println("FORMAT: (999999999999999999)       >", as.FormatIntToByte(999999999999999999))
	fmt.Println("FORMAT: (1000000000000000000)      >", as.FormatIntToByte(1152921504606846976))
}
