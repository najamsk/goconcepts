// Code generated by goa v3.5.3, DO NOT EDIT.
//
// users HTTP server encoders and decoders
//
// Command:
// $ goa gen mario/goa/design

package server

import (
	"context"
	"io"
	users "mario/goa/gen/users"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the users
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*users.User)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the users create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateUser(&body)

		return payload, nil
	}
}

// EncodeAllResponse returns an encoder for responses returned by the users all
// endpoint.
func EncodeAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*users.User)
		enc := encoder(ctx, w)
		body := NewAllResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalUsersUserToUserResponse builds a value of type *UserResponse from a
// value of type *users.User.
func marshalUsersUserToUserResponse(v *users.User) *UserResponse {
	res := &UserResponse{
		Name:      v.Name,
		BirthYear: v.BirthYear,
	}

	return res
}
