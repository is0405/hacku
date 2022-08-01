package util

import(
	"net/http"
	"strings"
	"strconv"
)

// http://localhost/aaaa/{id} のidを取り出す
func URLToInt(r *http.Request) (int, error) {
	url := r.URL.Path
	str := url[strings.LastIndex(url, "/")+1:]
	id, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}

// http://localhost/aaaa/{id}/ssssのidを取り出す
func URLToSecondInt(r *http.Request) (int, error) {
	url := r.URL.Path
	url = url[:strings.LastIndex(url, "/")]
	str := url[strings.LastIndex(url, "/")+1:]

	id, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}
