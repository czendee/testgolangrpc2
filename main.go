package main

import (
//	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"


       "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

        mul  "github.com/czendee/testgolangrpc2/genproto/go"
//        mul  "/genproto/go"

	"strconv"
	"fmt"	
	"strings"
	"time"
)

/**************************end handler.go ********************/

func Set(w http.ResponseWriter, hs ...Header) {
	for _, h := range hs {
		h.Populate(w.Header())
	}
}

type Header interface {
	Populate(http.Header)
}

type AccessControl struct {
	MaxAge           time.Duration
	Origin           string
	ExposedHeaders   []string
	AllowCredentials bool
	AllowedMethods   []string
	AllowedHeaders   []string
}

func (ac AccessControl) Populate(h http.Header) {
	if ac.Origin != "" {
		h.Set("Access-Control-Allow-Origin", ac.Origin)
	}
	if len(ac.AllowedMethods) > 0 {
		h.Set("Access-Control-Allow-Methods", strings.Join(ac.AllowedMethods, ", "))
	}
}

type ContentType string

func (ct ContentType) Populate(h http.Header) {
	h.Set("Content-Type", string(ct))
}
/**************************end handler.go ********************/
type server struct{}


// SayMultiplica implements multiplica.GreeterServer
func (s *server) SayMultiplica(ctx context.Context, in *mul.MultiplicaRequest) (*mul.MultiplicaReply, error) {
	
	
	elemento1 :=in.Numero
	var s1final float64 = 0
	
	if s1, err := strconv.ParseFloat(elemento1, 64); err == nil {
             fmt.Println(s1) // 3.1415927410125732
		s1final =s1;
       }
	
	elemento2 :=in.Veces
	var s2final float64 = 0
	
	if s2, err := strconv.ParseFloat(elemento2, 64); err == nil {
             fmt.Println(s2) // 3.1415927410125732
		s2final =s2;
       }
	
	resultado := s1final* s2final;
	 fmt.Println(resultado) 
	sresultado := fmt.Sprintf("%f", resultado)

//	return &mul.MultiplicaReply{Message: " " + in.Numero +" multiplica" + sresultado}, nil
        return &mul.MultiplicaReply{Message: " " + sresultado}, nil
}


func startGRPC(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	
	mul.RegisterGreeterServer(s, &server{})
	return s.Serve(lis)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Set(w, AccessControl{
			Origin:         "*",
			AllowedMethods: []string{"GET", "HEAD", "OPTIONS", "POST", "PUT", "DELETE", "PATCH"},
		})
		next.ServeHTTP(w, r)
	})
}

func startHTTP(httpPort, grpcPort string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()



	gwmuxmultiplica := runtime.NewServeMux()
	optsmul := []grpc.DialOption{grpc.WithInsecure()}
	if err := mul.RegisterGreeterHandlerFromEndpoint(ctx, gwmuxmultiplica, "127.0.0.1:"+grpcPort, optsmul); err != nil {
            
		return err
	}

	mux := http.NewServeMux()

	mux.Handle("/v2/", gwmuxmultiplica)

	http.ListenAndServe(":"+httpPort, cors(mux))
	return nil
}

func main() {
	errors := make(chan error)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50080"
	}

	if grpcPort == httpPort {
		panic("Can't listen on the same port")
	}

	go func() {
		errors <- startGRPC(grpcPort)
	}()

	go func() {
		errors <- startHTTP(httpPort, grpcPort)
	}()

	for err := range errors {
		log.Fatal(err)
		return
	}
}
