// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ouchihub/category/v1/category.proto

package categoryv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/category/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CategoryServiceName is the fully-qualified name of the CategoryService service.
	CategoryServiceName = "ouchihub.category.v1.CategoryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CategoryServiceListCategoriesProcedure is the fully-qualified name of the CategoryService's
	// ListCategories RPC.
	CategoryServiceListCategoriesProcedure = "/ouchihub.category.v1.CategoryService/ListCategories"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	categoryServiceServiceDescriptor              = v1.File_ouchihub_category_v1_category_proto.Services().ByName("CategoryService")
	categoryServiceListCategoriesMethodDescriptor = categoryServiceServiceDescriptor.Methods().ByName("ListCategories")
)

// CategoryServiceClient is a client for the ouchihub.category.v1.CategoryService service.
type CategoryServiceClient interface {
	ListCategories(context.Context, *connect.Request[v1.ListCategoriesRequest]) (*connect.Response[v1.ListCategoriesResponse], error)
}

// NewCategoryServiceClient constructs a client for the ouchihub.category.v1.CategoryService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCategoryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CategoryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &categoryServiceClient{
		listCategories: connect.NewClient[v1.ListCategoriesRequest, v1.ListCategoriesResponse](
			httpClient,
			baseURL+CategoryServiceListCategoriesProcedure,
			connect.WithSchema(categoryServiceListCategoriesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// categoryServiceClient implements CategoryServiceClient.
type categoryServiceClient struct {
	listCategories *connect.Client[v1.ListCategoriesRequest, v1.ListCategoriesResponse]
}

// ListCategories calls ouchihub.category.v1.CategoryService.ListCategories.
func (c *categoryServiceClient) ListCategories(ctx context.Context, req *connect.Request[v1.ListCategoriesRequest]) (*connect.Response[v1.ListCategoriesResponse], error) {
	return c.listCategories.CallUnary(ctx, req)
}

// CategoryServiceHandler is an implementation of the ouchihub.category.v1.CategoryService service.
type CategoryServiceHandler interface {
	ListCategories(context.Context, *connect.Request[v1.ListCategoriesRequest]) (*connect.Response[v1.ListCategoriesResponse], error)
}

// NewCategoryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCategoryServiceHandler(svc CategoryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	categoryServiceListCategoriesHandler := connect.NewUnaryHandler(
		CategoryServiceListCategoriesProcedure,
		svc.ListCategories,
		connect.WithSchema(categoryServiceListCategoriesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/ouchihub.category.v1.CategoryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CategoryServiceListCategoriesProcedure:
			categoryServiceListCategoriesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCategoryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCategoryServiceHandler struct{}

func (UnimplementedCategoryServiceHandler) ListCategories(context.Context, *connect.Request[v1.ListCategoriesRequest]) (*connect.Response[v1.ListCategoriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchihub.category.v1.CategoryService.ListCategories is not implemented"))
}
