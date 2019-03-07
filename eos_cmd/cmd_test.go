package eos_cmd

import "testing"

func TestMd5(t *testing.T) {
	md5, err := md5Sum("pct")
	checkError(err, t)
	t.Log(md5)
	if md5 != "973333fd6d3ab217522f3785d5d3e6ea" {
		t.Fatal("pct's md5 must equal 973333fd6d3ab217522f3785d5d3e6ea")
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		if t != nil {
			t.Fatal(err)
		} else {
			panic(err)
		}
	}
}