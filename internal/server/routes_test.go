package server_test

import (
	"github.com/maveonair/onepage/internal/server"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_Health(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}

	server, err := server.NewServer("")
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Code)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "ok") {
		t.Errorf("Expected body to contain 'ok', got '%s'", body)
	}
}

func Test_Index(t *testing.T) {
	file, err := os.CreateTemp("", "page.md")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	defer func() {
		if err := os.Remove(file.Name()); err != nil {
			t.Fatal(err)
		}
	}()

	if _, err := file.Write([]byte("# title")); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	server, err := server.NewServer(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Code)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "<h1 id=\"title\">title</h1>") {
		t.Errorf("Expected body to contain '<h1 id=\"title\">title</h1>', got '%s'", body)
	}
}

func Test_Edit(t *testing.T) {
	file, err := os.CreateTemp("", "page.md")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	defer func() {
		if err := os.Remove(file.Name()); err != nil {
			t.Fatal(err)
		}
	}()

	if _, err := file.Write([]byte("# title")); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, "/edit", nil)
	if err != nil {
		t.Fatal(err)
	}

	server, err := server.NewServer(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Code)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "# title") {
		t.Errorf("Expected body to contain '# title', got '%s'", body)
	}
}

func Test_Update(t *testing.T) {
	file, err := os.CreateTemp("", "page.md")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	defer func() {
		if err := os.Remove(file.Name()); err != nil {
			t.Fatal(err)
		}
	}()

	if _, err := file.Write([]byte("# title")); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/update", strings.NewReader("content=# updated"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	server, err := server.NewServer(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)

	if res.Code != http.StatusFound {
		t.Errorf("expected status Found (302); got %v", res.Code)
	}

	req, err = http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res = httptest.NewRecorder()
	server.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Code)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "<h1 id=\"updated\">updated</h1>") {
		t.Errorf("Expected body to contain '<h1 id=\"updated\">updated</h1>', got '%s'", body)
	}
}
