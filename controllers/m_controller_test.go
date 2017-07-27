package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVolumeStatus(t *testing.T) {
	req, err := http.NewRequest("Get", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVolumeStatus)
	handler.ServeHTTP(requestRecorder, req)

	if requestRecorder.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Got %d , Want %d", requestRecorder.Code, http.StatusOK)
	} else {
		t.Logf("Handler returned correct status code %d", requestRecorder.Code)
	}

	expected := `{"alive" : true}`

	if expected != requestRecorder.Body.String() {
		t.Errorf("Handler returned wrong body, Got %s Want %s", requestRecorder.Body.String(), expected)
	} else {
		t.Logf("Handler returned correct body %s", requestRecorder.Body.String())
	}
}
