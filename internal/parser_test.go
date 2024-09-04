package internal

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestGoVersion(t *testing.T) {
	// Определение тест-кейсов
	testCases := []struct {
		name     string
		fileName string
		want     string
		wantErr  bool
	}{
		{
			name:     "Valid HTML with version",
			fileName: "target_data.html",
			want:     "go1.21.5",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Чтение содержимого тестового файла
			testData, err := os.ReadFile("../tests/" + tc.fileName)
			if err != nil {
				t.Fatalf("Failed to read test data for %v: %v", tc.name, err)
			}

			// Создание мок-сервера
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err = w.Write(testData)
				assert.NoError(t, err)
			}))
			defer server.Close()

			// Вызов функции с URL мок-сервера
			got, err := GetLatestGoVersion(server.URL)
			if (err != nil) != tc.wantErr {
				t.Errorf("getLatestGoVersion error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// Проверка полученной версии
			if got != tc.want {
				t.Errorf("getLatestGoVersion = %q, want %q", got, tc.want)
			}
		})
	}
}
