package app

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type validateable interface {
	Validate() error
}

func GrpcInterceptor() grpc.ServerOption {
	grpcServerOption := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if v, ok := req.(validateable); ok {
			if err := v.Validate(); err != nil {
				return nil, err
			}
		}
		// resp, err = handler(ctx, req)
		return handler(ctx, req)
	})
	return grpcServerOption
}

// Custom Interceptor
func CustomInterceptor() runtime.ServeMuxOption {
	httpServerOptions := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		return nil
	})

	return httpServerOptions
}

func MyFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	// hashcode := w.Header().Get("Hashcode") + " "
	// apiName := w.Header().Get("Apiname")
	// logrus.Debug(hashcode, "Response ", apiName, ": ", util.Serialize(resp))
	delete(w.Header(), "Hashcode") // remove header Hashcode
	delete(w.Header(), "Apiname")  // remove header Apiname
	// respModel := model.ResponseModel{}
	// json.Unmarshal([]byte(util.Serialize(resp)), &respModel)
	// logrus.Debug(hashcode, "Response ", apiName, ": ", respModel.Status)
	// w.WriteHeader(respModel.Status)
	return nil
}

func allowedOrigin(origin string) bool {
	return strings.Contains(os.Getenv("ALLOWED_ORIGIN"), origin)

}

func HttpInterceptor(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// hashcode := "(" + util.GenerateHashcode() + ") "
		// apiname := util.GetAPIName(r.URL.Path, r.Method)
		// start := util.GetTimeNowMillis()
		// r.Header.Add("Hashcode", hashcode)

		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, sessionid")
		}

		if r.Method == "OPTIONS" {
			return
		}

		// if strings.Compare(apiname, "default") != 0 && strings.Compare(strings.ToLower(r.Method), "post") == 0 {

		// 	logrus.Info(hashcode, "===== start ", apiname, " =====")

		// 	w.Header().Add("Hashcode", hashcode)
		// 	w.Header().Add("Apiname", apiname)

		// 	h.ServeHTTP(w, r)

		// 	end := (time.Now().UnixNano() / int64(time.Millisecond)) - start
		// 	logrus.Info(hashcode, "===== finish ", apiname, " on ", end, " millis =====")

		// } else {
		// 	h.ServeHTTP(w, r)
		// }
	})
}

func CustomMatcher(key string) (string, bool) {
	return key, true
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				logrus.Error("panic go: ", err) // May be log this error? Send to newrelic?

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}

func CheckApiKey(apikey string) bool {
	key := "744804329777439e8c41085b003111174150e40c581fc1bd2b"
	return strings.Compare(key, apikey) == 0
}
