// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/Astemirdum/transactions/proto/user/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	SignUp(context.Context, *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error)
	SignIn(context.Context, *connect_go.Request[v1.SignInRequest]) (*connect_go.Response[v1.SignInResponse], error)
	SignOut(context.Context, *connect_go.Request[v1.SignOutRequest]) (*connect_go.Response[v1.SignOutResponse], error)
	Auth(context.Context, *connect_go.Request[v1.AuthRequest]) (*connect_go.Response[v1.AuthResponse], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		signUp: connect_go.NewClient[v1.SignUpRequest, v1.SignUpResponse](
			httpClient,
			baseURL+"/user.v1.UserService/SignUp",
			opts...,
		),
		signIn: connect_go.NewClient[v1.SignInRequest, v1.SignInResponse](
			httpClient,
			baseURL+"/user.v1.UserService/SignIn",
			opts...,
		),
		signOut: connect_go.NewClient[v1.SignOutRequest, v1.SignOutResponse](
			httpClient,
			baseURL+"/user.v1.UserService/SignOut",
			opts...,
		),
		auth: connect_go.NewClient[v1.AuthRequest, v1.AuthResponse](
			httpClient,
			baseURL+"/user.v1.UserService/Auth",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	signUp  *connect_go.Client[v1.SignUpRequest, v1.SignUpResponse]
	signIn  *connect_go.Client[v1.SignInRequest, v1.SignInResponse]
	signOut *connect_go.Client[v1.SignOutRequest, v1.SignOutResponse]
	auth    *connect_go.Client[v1.AuthRequest, v1.AuthResponse]
}

// SignUp calls user.v1.UserService.SignUp.
func (c *userServiceClient) SignUp(ctx context.Context, req *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error) {
	return c.signUp.CallUnary(ctx, req)
}

// SignIn calls user.v1.UserService.SignIn.
func (c *userServiceClient) SignIn(ctx context.Context, req *connect_go.Request[v1.SignInRequest]) (*connect_go.Response[v1.SignInResponse], error) {
	return c.signIn.CallUnary(ctx, req)
}

// SignOut calls user.v1.UserService.SignOut.
func (c *userServiceClient) SignOut(ctx context.Context, req *connect_go.Request[v1.SignOutRequest]) (*connect_go.Response[v1.SignOutResponse], error) {
	return c.signOut.CallUnary(ctx, req)
}

// Auth calls user.v1.UserService.Auth.
func (c *userServiceClient) Auth(ctx context.Context, req *connect_go.Request[v1.AuthRequest]) (*connect_go.Response[v1.AuthResponse], error) {
	return c.auth.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	SignUp(context.Context, *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error)
	SignIn(context.Context, *connect_go.Request[v1.SignInRequest]) (*connect_go.Response[v1.SignInResponse], error)
	SignOut(context.Context, *connect_go.Request[v1.SignOutRequest]) (*connect_go.Response[v1.SignOutResponse], error)
	Auth(context.Context, *connect_go.Request[v1.AuthRequest]) (*connect_go.Response[v1.AuthResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/user.v1.UserService/SignUp", connect_go.NewUnaryHandler(
		"/user.v1.UserService/SignUp",
		svc.SignUp,
		opts...,
	))
	mux.Handle("/user.v1.UserService/SignIn", connect_go.NewUnaryHandler(
		"/user.v1.UserService/SignIn",
		svc.SignIn,
		opts...,
	))
	mux.Handle("/user.v1.UserService/SignOut", connect_go.NewUnaryHandler(
		"/user.v1.UserService/SignOut",
		svc.SignOut,
		opts...,
	))
	mux.Handle("/user.v1.UserService/Auth", connect_go.NewUnaryHandler(
		"/user.v1.UserService/Auth",
		svc.Auth,
		opts...,
	))
	return "/user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) SignUp(context.Context, *connect_go.Request[v1.SignUpRequest]) (*connect_go.Response[v1.SignUpResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.SignUp is not implemented"))
}

func (UnimplementedUserServiceHandler) SignIn(context.Context, *connect_go.Request[v1.SignInRequest]) (*connect_go.Response[v1.SignInResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.SignIn is not implemented"))
}

func (UnimplementedUserServiceHandler) SignOut(context.Context, *connect_go.Request[v1.SignOutRequest]) (*connect_go.Response[v1.SignOutResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.SignOut is not implemented"))
}

func (UnimplementedUserServiceHandler) Auth(context.Context, *connect_go.Request[v1.AuthRequest]) (*connect_go.Response[v1.AuthResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.Auth is not implemented"))
}