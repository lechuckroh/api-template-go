package app

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestAdminIndex(t *testing.T) {
	Convey("should response 404 if not admin", t, func() {
		server := NewServer()
		server.routes()

		_, w, _ := server.get("/admin", nil, nil)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("should response 200 if admin", t, func() {
		server := NewServer()
		server.routes()

		headers := map[string]string{"admin": "true"}
		_, w, _ := server.get("/admin", nil, headers)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
}

func TestHealthcheck(t *testing.T) {
	Convey("should response 200", t, func() {
		server := NewServer()
		server.routes()

		_, w, _ := server.get("/api/healthcheck", nil, nil)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
}

func TestAddItem(t *testing.T) {
	type request struct {
		Name string `json:"name"`
	}
	type response struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	Convey("should response 400 if name is invalid type", t, func() {
		server := NewServer()
		server.routes()

		_, w, _ := server.postJSON("/api/items", struct {
			Name int `json:"name"`
		}{Name: 1}, nil)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("should response 400 if name is empty", t, func() {
		server := NewServer()
		server.routes()

		_, w, _ := server.postJSON("/api/items", request{Name: ""}, nil)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("should response 201 if name is set", t, func() {
		server := NewServer()
		server.routes()

		req := request{Name: "foo"}
		_, w, _ := server.postJSON("/api/items", req, nil)
		So(w.Code, ShouldEqual, http.StatusCreated)

		res := response{}
		_ = json.Unmarshal(w.Body.Bytes(), &res)
		So(res, ShouldResemble, response{ID: 1, Name: req.Name})
	})
}
