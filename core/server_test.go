package core

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	app := App{}
	app.HelloKoronet(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hi Koronet Team.\n"
	if string(body) != expected {
		t.Errorf("Unexpected body, got: %q but expected %q", string(body), expected)
	}
}
