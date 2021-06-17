/*
Copyright (C) GMO GlobalSign, Inc. 2019 - All Rights Reserved.

Unauthorized copying of this file, via any medium is strictly prohibited.
No distribution/modification of whole or part thereof is allowed.

Proprietary and confidential.
*/

package hvclient

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// headerFromResponse retrieves the value of a header from an HTTP response. If there
// is more than one header value, only the first is returned.
func headerFromResponse(r *http.Response, name string) (string, error) {
	if len(r.Header[name]) == 0 {
		return "", fmt.Errorf("no values in response for header %q", name)
	}

	return r.Header[name][0], nil
}

// basePathHeaderFromResponse retrieves the base part of the path value contained in a
// header in an HTTP response. If there is more than one header value, only
// the first is returned.
func basePathHeaderFromResponse(r *http.Response, name string) (string, error) {
	var location, err = headerFromResponse(r, name)
	if err != nil {
		return "", err
	}

	return filepath.Base(location), nil
}

// intHeaderFromResponse retrieves the integer value of a header from an HTTP
// response. If there is more than one header value, only the first is
// returned.
func intHeaderFromResponse(r *http.Response, name string) (int64, error) {
	var s, err = headerFromResponse(r, name)
	if err != nil {
		return 0, err
	}

	var n int64
	n, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// paginationString builds a query string for paginated API requests.
// perPage, from and to are optional.
func paginationString(
	page, perPage int,
	from, to time.Time,
) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("?page=%d", page))

	if perPage > 0 {
		builder.WriteString(fmt.Sprintf("&per_page=%d", perPage))
	}

	if !from.IsZero() {
		builder.WriteString(fmt.Sprintf("&from=%d", from.Unix()))
	}

	if !to.IsZero() {
		builder.WriteString(fmt.Sprintf("&to=%d", to.Unix()))
	}

	return builder.String()
}