package fetch

import (
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// create a new HTTP client
	c := urlfetch.Client(ctx)

	// and use it to request the Google homepage
	res, err := c.Get("https://google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// we need to close the body at the end of this function
	defer res.Body.Close()

	// then we can dump the whole webpage onto our output
	_, err = io.Copy(w, res.Body)
	if err != nil {
		log.Errorf(ctx, "could not copy the response: %v", err)
	}
}

func init() {
	http.HandleFunc("/", handler)
}
