package grpc

import (
	"github.com/go-chi/jwtauth"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

// Add middleware to validate that all gRPC requests are authorized with Auth0.
func validationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		// TODO: rework ordering to require all endpoints beyond a certain point to
		// be authenticated
		if wrappedServer.IsGrpcWebRequest(req) || isGatewayGrpcRequest(req) {
			// Perform auth validation only on gRPC requests.
			// TODO: perform extra validation on the iss/aud/exp fields.
			verifier := jwtauth.Verifier(jwtauth.New("RS256", nil, authKey))
			verifier(jwtauth.Authenticator(next)).ServeHTTP(resp, req)
		} else {
			next.ServeHTTP(resp, req)
		}
	})
}

func grpcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		grpclog.Infof("Received request with C-T: %s\n", req.Header.Get("content-type"))
		if wrappedServer.IsAcceptableGrpcCorsRequest(req) {
			grpclog.Infof("Incoming grpc cors-preflight request")
			wrappedServer.ServeHTTP(resp, req)
		} else if wrappedServer.IsGrpcWebRequest(req) || isGatewayGrpcRequest(req) {
			grpclog.Infof("Incoming grpc request")
			wrappedServer.ServeHTTP(resp, req)
		} else {
			grpclog.Infof("Incoming non-grpc request")
			next.ServeHTTP(resp, req)
		}
	})
}
