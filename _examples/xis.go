package main

import (
	"fmt"
	"time"

	"../xis"
)

func main() {
	time.Sleep(1 * time.Second)
	isempty()
	time.Sleep(1 * time.Second)
	isemail()
	time.Sleep(1 * time.Second)
	isint()
	time.Sleep(1 * time.Second)
	isfloat()
	time.Sleep(1 * time.Second)
	isbool()
	time.Sleep(1 * time.Second)
	isstring()
	time.Sleep(1 * time.Second)
	istime()
	time.Sleep(1 * time.Second)
	isinintslice()
	time.Sleep(1 * time.Second)
	isinstringslice()
	time.Sleep(1 * time.Second)
	isinslice()
	time.Sleep(1 * time.Second)
	isipaddress()
	time.Sleep(1 * time.Second)
	ishttp()
	time.Sleep(1 * time.Second)
	isuuid()
}

func isuuid() {
	fmt.Println("")
	fmt.Println("IS UUID *******************************************************")
	fmt.Println("isuuid(''):                  >", is.IsUUID(""))
}

func ishttp() {
	fmt.Println("")
	fmt.Println("IS HTTP *******************************************************")
	fmt.Println("IsURL(''):                  >", is.IsURL(""))
	fmt.Println("IsURL(''):                  >", is.IsURL("www.google.co.uk"))
	fmt.Println("IsURL(''):                  >", is.IsURL("http://www.google.co.uk"))
	fmt.Println("IsURL(''):                  >", is.IsURL("https://www.google.co.uk"))
}

func isempty() {
	fmt.Println("")
	fmt.Println("IS EMPTY *******************************************************")
	fmt.Println("IsEmpty(''):                  >", is.IsEmpty(""))
	fmt.Println("IsEmpty('test'):              >", is.IsEmpty("test"))
	fmt.Println("IsEmpty('  '):                >", is.IsEmpty("  "))
	fmt.Println("IsNotEmpty(''):               >", is.IsNotEmpty(""))
	fmt.Println("IsNotEmpty('test'):           >", is.IsNotEmpty("test"))
	fmt.Println("IsNotEmpty('  '):             >", is.IsNotEmpty("  "))
	fmt.Println("IsBlank(''):                  >", is.IsBlank(""))
	fmt.Println("IsBlank('test'):              >", is.IsBlank("test"))
	fmt.Println("IsUppercase('a'):             >", is.IsStringUppercase("a"))
	fmt.Println("IsUppercase('A'):             >", is.IsStringUppercase("A"))
	fmt.Println("IsStringAlpha('ab!'):         >", is.IsStringAllAlpha("ab!"))
	fmt.Println("IsStringAlpha('Ab'):          >", is.IsStringAllAlpha("Ab"))
	fmt.Println("IsStringAlpha('ab4'):         >", is.IsStringAllAlpha("ab4"))
	fmt.Println("IsStringContainAlpha('ab4'):  >", is.IsStringContainAlpha("ab4"))
	fmt.Println("IsStringContainAlpha('ab'):   >", is.IsStringContainAlpha("ab"))
	fmt.Println("IsStringContainAlpha('AB4'):  >", is.IsStringContainAlpha("AB4"))
	fmt.Println("IsStringContainNumber('ab4'): >", is.IsStringContainNumber("ab4"))
	fmt.Println("IsStringContainNumber('ab'):  >", is.IsStringContainNumber("ab"))
	fmt.Println("IsStringContainNumber('AB4'): >", is.IsStringContainNumber("AB4"))
	fmt.Println("IsAlphaNumeric('AB4'):        >", is.IsAlphaNumeric("AB4"))
	fmt.Println("IsAlphaNumeric('AB'):         >", is.IsAlphaNumeric("AB"))
	fmt.Println("IsAlphaNumeric('ab4'):        >", is.IsAlphaNumeric("ab4"))
	fmt.Println("IsAlphaNumeric('44'):         >", is.IsAlphaNumeric("44"))
	fmt.Println("IsBlank(''):                  >", is.IsBlank(""))
	fmt.Println("IsBlank('  '):                >", is.IsBlank("  "))
	fmt.Println("IsNotBlank(''):               >", is.IsNotBlank(""))
	fmt.Println("IsNotBlank('test'):           >", is.IsNotBlank("test"))
	fmt.Println("IsNotBlank('  '):             >", is.IsNotBlank("  "))
	fmt.Println("Reverse('test'):              >", is.Reverse("test"))
}

func isemail() {
	fmt.Println("")
	fmt.Println("IS EMAIL *******************************************************")

	fmt.Println("IsEmail(''):                  >", is.IsEmail(""))
	fmt.Println("IsEmail('test'):              >", is.IsEmail("test"))
	fmt.Println("IsEmail('test.test'):         >", is.IsEmail("test.test"))
	fmt.Println("IsEmail('test@test'):         >", is.IsEmail("test@test"))
	fmt.Println("IsEmail('test@test.test'):    >", is.IsEmail("test@test.com"))
	fmt.Println("IsEmail2(''):                 >", is.IsEmail2(""))
	fmt.Println("IsEmail2('test'):             >", is.IsEmail2("test"))
	fmt.Println("IsEmail2('test.test'):        >", is.IsEmail2("test.test"))
	fmt.Println("IsEmail2('test@test'):        >", is.IsEmail2("test@test"))
	fmt.Println("IsEmail2('test@test.test'):   >", is.IsEmail2("test@test.com"))
}

func isint() {
	fmt.Println("")
	fmt.Println("IS INT *******************************************************")
	fmt.Println("IsInt('4'):                   >", is.IsInt("4"))
	fmt.Println("IsInt(4):                     >", is.IsInt(4))
	fmt.Println("IsInt(4.4):                   >", is.IsInt(4.4))
	fmt.Println("IsInt(true):                  >", is.IsInt(true))
	fmt.Println("IsInt(nil):                   >", is.IsInt(nil))
	fmt.Println("IsInt('07-06-2017'):          >", is.IsInt(time.Now()))
}

func isfloat() {
	fmt.Println("")
	fmt.Println("IS FLOAT *******************************************************")
	fmt.Println("IsFloat('4'):                 >", is.IsFloat("4"))
	fmt.Println("IsFloat(4):                   >", is.IsFloat(4))
	fmt.Println("IsFloat(4.4):                 >", is.IsFloat(4.4))
	fmt.Println("IsFloat(true):                >", is.IsFloat(true))
	fmt.Println("IsFloat(nil):                 >", is.IsFloat(nil))
	fmt.Println("IsInt('07-06-2017'):          >", is.IsFloat(time.Now()))
}

func isbool() {
	fmt.Println("")
	fmt.Println("IS BOOL *******************************************************")
	fmt.Println("IsBool('4'):                  >", is.IsBool("4"))
	fmt.Println("IsBool(4):                    >", is.IsBool(4))
	fmt.Println("IsBool(4.4):                  >", is.IsBool(4.4))
	fmt.Println("IsBool(true):                 >", is.IsBool(true))
	fmt.Println("IsBool(nil):                  >", is.IsBool(nil))
	fmt.Println("IsInt('07-06-2017'):          >", is.IsBool(time.Now()))
}

func isstring() {
	fmt.Println("")
	fmt.Println("IS STRING *******************************************************")
	fmt.Println("IsString('4'):                >", is.IsString("4"))
	fmt.Println("IsString(4):                  >", is.IsString(4))
	fmt.Println("IsString(4.4):                >", is.IsString(4.4))
	fmt.Println("IsString(true):               >", is.IsString(true))
	fmt.Println("IsString(nil):                >", is.IsString(nil))
	fmt.Println("IsString('07-06-2017'):       >", is.IsString(time.Now()))
}

func istime() {
	fmt.Println("")
	fmt.Println("IS TIME *******************************************************")
	fmt.Println("IsTime('4'):                  >", is.IsTime("4"))
	fmt.Println("IsTime(4):                    >", is.IsTime(4))
	fmt.Println("IsTime(4.4):                  >", is.IsTime(4.4))
	fmt.Println("IsTime(true):                 >", is.IsTime(true))
	fmt.Println("IsTime(nil):                  >", is.IsTime(nil))
	fmt.Println("IsTime(now()):                >", is.IsTime(time.Now()))
}

func isinintslice() {
	fmt.Println("")
	fmt.Println("IS IN INT SLICE *******************************************************")
	fmt.Println("IsInIntSlice():               >", is.IsInIntSlice(4, []int{1, 2, 3, 4}))
	fmt.Println("IsInIntSlice():               >", is.IsInIntSlice(4, []int{1, 2, 3, 5}))
}

func isinstringslice() {
	fmt.Println("")
	fmt.Println("IS IN STRING SLICE *******************************************************")
	fmt.Println("IsInStringSlice():            >", is.IsInStringSlice("a", []string{"a", "b", "c", "d"}))
	fmt.Println("IsInStringSlice():            >", is.IsInStringSlice("a", []string{"b"}))
}

func isinslice() {
	fmt.Println("")
	fmt.Println("IS IN SLICE *******************************************************")
	/*a := []interface{"1", "2", "3", "4"}
	fmt.Println("IsInSlice():                  >", is.IsInSlice(4, a))
	a = []interface{1, 2, 3, 5}
	fmt.Println("IsInSlice():                  >", is.IsInSlice(4, a))
	fmt.Println("IsInSlice():                  >", is.IsInSlice("a", []string{"a", "b", "c", "d"}))
	fmt.Println("IsInSlice():                  >", is.IsInSlice("a", []string{"b"}))
	*/
}

func isipaddress() {
	fmt.Println("")
	fmt.Println("IS IP ADDRESS *******************************************************")
	fmt.Println("IS IP (1111.1111.1111.1111):  >", is.IsIPAddress("1111.1111.1111.1111"))
	fmt.Println("IS IP (111.111.111.111):      >", is.IsIPAddress("111.111.111.111"))
	fmt.Println("IS IP (11.11.11.11):          >", is.IsIPAddress("11.11.11.11"))
	fmt.Println("IS IP (1.1.1.1):              >", is.IsIPAddress("1.1.1.1"))
}
