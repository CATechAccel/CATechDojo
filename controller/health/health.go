package health

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "Hello,World!"); err != nil {
		http.Error(w, "ヘルスチェックに失敗しました", http.StatusInternalServerError)
	}
}
