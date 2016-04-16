package restful

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnmodifiableContainer(t *testing.T) {
	wc := NewContainer()
	ws1 := new(WebService).Path("/users")
	ws1.Route(ws1.GET("{i}").To(dummy))
	ws1.Route(ws1.POST("{i}").To(dummy))
	wc.Add(ws1)

	um := wc.Unmodifiable()

	httpRequest, _ := http.NewRequest("GET", "http://api.his.com/users/1", nil)
	recorder := httptest.NewRecorder()
	um.ServeHTTP(recorder, httpRequest)
	if recorder.Code != http.StatusOK {
		t.Errorf("unexpected code %d", recorder.Code)
	}
}
