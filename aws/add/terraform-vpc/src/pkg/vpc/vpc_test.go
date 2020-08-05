package vpc

import "testing"

func TestParseAZS(t *testing.T) {
	want := "\"us-east-1a\",\"us-east-1b\""
	in := Inputs{
		VPCAZS: "us-east-1a,us-east-1b",
	}

	in.parseAZS()

	if in.VPCAZS != want {
		t.Errorf("parseAZS got (%v), want (%v)", in.VPCAZS, want)
	}
}
