package structtool

import (
	"testing"
)

func TestStructFloatFormat(t *testing.T) {
	type Example struct {
		A float64
		B float64
		C int
		D string
	}
	expStruct := &Example{
		A: 1.4444444,
		B: 0.0222222222,
		C: 5,
		D: "ddd",
	}
	expStruct = StructFloatFormat(expStruct).(*Example)
	t.Logf("expStruct:%v", expStruct) // expStruct:&{1.44 0.02 5 ddd}
}

func TestStructCopy(t *testing.T) {
	type Common struct {
		S1 int
		S2 string
	}
	type Example1 struct {
		AA float64
		B  float64
		C  int
		SS *Common
		D  string
	}
	type Example2 struct {
		A  float64
		B  float64
		C  int
		SS *Common
		D  float64
	}
	exp1 := &Example1{
		AA: 1.4444444,
		B:  0.0222222222,
		C:  5,
		SS: &Common{
			S1: 2,
			S2: "aaa",
		},
		D: "ddd",
	}
	exp2 := &Example2{}
	t.Logf("exp1:%v   exp2:%v", exp1, exp2) // exp1:&{1.4444444 0.0222222222 5 0xc000084060 ddd}   exp2:&{0 0 0 <nil> 0}
	StructCopy(exp2, exp1)
	t.Logf("exp1:%v   exp2:%v", exp1, exp2) // exp1:&{1.4444444 0.0222222222 5 0xc00008a080 ddd}   exp2:&{0 0.0222222222 5 0xc00008a080 0}
}

func TestStructURLDecode(t *testing.T) {
	type Request struct {
		AppID       string
		Count       int
		PageNo      int
		PlateNumber string
		Signature   string
	}
	req := &Request{
		AppID:       "aaaa",
		Count:       10,
		PageNo:      1,
		PlateNumber: "%E8%B5%A3A84897",
		Signature:   "sh%2FuYfecIceZ%2FWKLRdmEkQ%3D%3D",
	}
	t.Logf("req:%v\n", req) // req:&{aaaa 10 1 %E8%B5%A3A84897 sh%2FuYfecIceZ%2FWKLRdmEkQ%3D%3D}
	req = StructURLDecode(req).(*Request)
	t.Logf("decode req:%v\n", req) // decode req:&{aaaa 10 1 赣A84897 sh/uYfecIceZ/WKLRdmEkQ==}
}
