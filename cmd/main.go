package main

import (
	"github.com/rafaeljesus/composing-go-example/httpclient"
	"github.com/rafaeljesus/composing-go-example/user"
)

func main() {
	reqter := new(httpclient.Requester)
	poster := httpclient.NewHTTPPoster(reqter)
	client := httpclient.New(
		poster,
		httpclient.NewHTTPGetter(reqter),
	)

	_ = user.NewUserSyncer(poster)
	_ = user.NewUserStorer(client)

	// pass syncer and storer downstream to
	// http handler, messaging handler, etc...
}
