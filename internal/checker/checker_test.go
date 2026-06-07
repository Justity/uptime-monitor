package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheck_ReturnsSuccessForHealthyServer(
	t *testing.T,
) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				w.WriteHeader(
					http.StatusOK,
				)
			},
		),
	)

	defer server.Close()

	result := Check(server.URL)

	if !result.IsUp {
		t.Fatal(
			"expected server to be up",
		)
	}

	if result.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected %d got %d",
			http.StatusOK,
			result.StatusCode,
		)
	}
}
