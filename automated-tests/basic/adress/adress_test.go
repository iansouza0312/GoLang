package adress

import "testing"

type testCase struct {
	adressInsert   string
	expectedReturn string
}

func TestAdressType(t *testing.T) {
	testCases := []testCase{
		{"Highway New Orleans", "Highway"},
		{"Street blablalba", "Street"},
		{"Avenue something", "Avenue"},
		{"Street elder", "Invalid Type"},
		{"STREET FISHBOY", "Street"},
		{"Highway newest", "Invalid Type"},
	}

	for _, tc := range testCases {
		adressTypeRecieved := AdressType(tc.adressInsert)
		if adressTypeRecieved != tc.expectedReturn {
			t.Errorf("Type Recieved %s is different of Expected %s", adressTypeRecieved, tc.expectedReturn)
		}
	}
}
