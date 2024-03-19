// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ouchihub/work/v1/work.proto

package workv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/work/v1"
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
	// WorkServiceName is the fully-qualified name of the WorkService service.
	WorkServiceName = "ouchihub.work.v1.WorkService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WorkServiceListWorksProcedure is the fully-qualified name of the WorkService's ListWorks RPC.
	WorkServiceListWorksProcedure = "/ouchihub.work.v1.WorkService/ListWorks"
	// WorkServiceCreateWorkProcedure is the fully-qualified name of the WorkService's CreateWork RPC.
	WorkServiceCreateWorkProcedure = "/ouchihub.work.v1.WorkService/CreateWork"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	workServiceServiceDescriptor          = v1.File_ouchihub_work_v1_work_proto.Services().ByName("WorkService")
	workServiceListWorksMethodDescriptor  = workServiceServiceDescriptor.Methods().ByName("ListWorks")
	workServiceCreateWorkMethodDescriptor = workServiceServiceDescriptor.Methods().ByName("CreateWork")
)

// WorkServiceClient is a client for the ouchihub.work.v1.WorkService service.
type WorkServiceClient interface {
	ListWorks(context.Context, *connect.Request[v1.ListWorksRequest]) (*connect.Response[v1.ListWorksResponse], error)
	CreateWork(context.Context, *connect.Request[v1.CreateWorkRequest]) (*connect.Response[v1.CreateWorkResponse], error)
}

// NewWorkServiceClient constructs a client for the ouchihub.work.v1.WorkService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WorkServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workServiceClient{
		listWorks: connect.NewClient[v1.ListWorksRequest, v1.ListWorksResponse](
			httpClient,
			baseURL+WorkServiceListWorksProcedure,
			connect.WithSchema(workServiceListWorksMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createWork: connect.NewClient[v1.CreateWorkRequest, v1.CreateWorkResponse](
			httpClient,
			baseURL+WorkServiceCreateWorkProcedure,
			connect.WithSchema(workServiceCreateWorkMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// workServiceClient implements WorkServiceClient.
type workServiceClient struct {
	listWorks  *connect.Client[v1.ListWorksRequest, v1.ListWorksResponse]
	createWork *connect.Client[v1.CreateWorkRequest, v1.CreateWorkResponse]
}

// ListWorks calls ouchihub.work.v1.WorkService.ListWorks.
func (c *workServiceClient) ListWorks(ctx context.Context, req *connect.Request[v1.ListWorksRequest]) (*connect.Response[v1.ListWorksResponse], error) {
	return c.listWorks.CallUnary(ctx, req)
}

// CreateWork calls ouchihub.work.v1.WorkService.CreateWork.
func (c *workServiceClient) CreateWork(ctx context.Context, req *connect.Request[v1.CreateWorkRequest]) (*connect.Response[v1.CreateWorkResponse], error) {
	return c.createWork.CallUnary(ctx, req)
}

// WorkServiceHandler is an implementation of the ouchihub.work.v1.WorkService service.
type WorkServiceHandler interface {
	ListWorks(context.Context, *connect.Request[v1.ListWorksRequest]) (*connect.Response[v1.ListWorksResponse], error)
	CreateWork(context.Context, *connect.Request[v1.CreateWorkRequest]) (*connect.Response[v1.CreateWorkResponse], error)
}

// NewWorkServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkServiceHandler(svc WorkServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workServiceListWorksHandler := connect.NewUnaryHandler(
		WorkServiceListWorksProcedure,
		svc.ListWorks,
		connect.WithSchema(workServiceListWorksMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workServiceCreateWorkHandler := connect.NewUnaryHandler(
		WorkServiceCreateWorkProcedure,
		svc.CreateWork,
		connect.WithSchema(workServiceCreateWorkMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/ouchihub.work.v1.WorkService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkServiceListWorksProcedure:
			workServiceListWorksHandler.ServeHTTP(w, r)
		case WorkServiceCreateWorkProcedure:
			workServiceCreateWorkHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkServiceHandler struct{}

func (UnimplementedWorkServiceHandler) ListWorks(context.Context, *connect.Request[v1.ListWorksRequest]) (*connect.Response[v1.ListWorksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchihub.work.v1.WorkService.ListWorks is not implemented"))
}

func (UnimplementedWorkServiceHandler) CreateWork(context.Context, *connect.Request[v1.CreateWorkRequest]) (*connect.Response[v1.CreateWorkResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchihub.work.v1.WorkService.CreateWork is not implemented"))
}
