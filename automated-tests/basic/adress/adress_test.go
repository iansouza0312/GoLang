package adress

import "testing"

func TestAdressType(t *testing.T) {
	testAdress := "Highway New Orleans"
	// testAdressError := "Highway New Orleans"
	expectedAdressType := "Highway"

	recievedAdressType := AdressType(testAdress)
	// recievedAdressType := AdressType(testAdressError)

	if recievedAdressType != expectedAdressType {
		t.Errorf("The type recieved is different of expected ! Expected %s and Recieved %s", expectedAdressType, recievedAdressType)
	}

}
