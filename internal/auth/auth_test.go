package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		in          http.Header
		expectedKey string
		expectedErr error
	}

	tests := []test{}

	H1 := http.Header{}

	H1.Set("TestMe", "ApiKey 1234")

	tests = append(tests, test{
		in:          H1,
		expectedKey: "",
		expectedErr: ErrNoAuthHeaderIncluded,
	})

	H2 := http.Header{}

	H2.Set("Authorization", "ApiKey1234")

	tests = append(tests, test{
		in:          H2,
		expectedKey: "",
		expectedErr: ErrMalformedHeader,
	})

	H3 := http.Header{}

	H3.Set("Authorization", "ApiKey 1234")

	tests = append(tests, test{
		in:          H3,
		expectedKey: "1234",
		expectedErr: nil,
	})

	for i, curT := range tests {
		key, err := GetAPIKey(curT.in)
		if key != curT.expectedKey || err != curT.expectedErr {
			t.Fatalf("failed in test %v\nexpected key '%v' and err '%v'\ngot: key '%v' and err '%v'", i, curT.expectedKey, curT.expectedErr, key, err)
		}
	}
}
