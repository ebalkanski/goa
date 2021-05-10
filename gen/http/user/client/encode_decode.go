// Code generated by goa v3.3.1, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/ebalkanski/goa/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	user "github.com/ebalkanski/goa/gen/user"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "user" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		name string
	)
	{
		p, ok := v.(*user.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "get", "*user.GetPayload", v)
		}
		name = p.Name
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserPath(name)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the user get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "get", err)
			}
			res := NewGetUserOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "user" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the user create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.User)
		if !ok {
			return goahttp.ErrInvalidType("user", "create", "*user.User", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the user
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "create", resp.StatusCode, string(body))
		}
	}
}
