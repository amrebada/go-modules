package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port             int
	Engine           *gin.Engine
	TimeoutInSeconds int
	MaxHeaderBytes   int
	MainModule       *Module
	DB               *Database
}

func (s *Server) Start() error {
	h := &http.Server{
		Addr:           fmt.Sprintf(":%d", s.Port),
		Handler:        s.Engine,
		ReadTimeout:    time.Duration(s.TimeoutInSeconds) * time.Second,
		WriteTimeout:   time.Duration(s.TimeoutInSeconds) * time.Second,
		MaxHeaderBytes: s.MaxHeaderBytes,
	}
	return h.ListenAndServe()
}

func (s *Server) RegisterMainModule() {
	config := ConfigInstance()
	if config.IS_MIGRATE {
		err := s.MainModule.Migrate()
		if err != nil {
			panic(err)
		}
	}

	swagger := NewSwagger().
		SetInfo(SwaggerInfo{
			Title:       "Tradeling Framework",
			Description: "Tradeling Framework",
			Version:     "1.0.0",
			Contact: SwaggerContact{
				Name:  "Amr Abada",
				Email: "amr.ebada@tradeling.com",
				URL:   "tradeling.com",
			},
		}).
		AddServer(SwaggerServer{
			URL:         "http://b2b.localhost/api/module-templates",
			Description: "local_server",
		}).
		AddServer(SwaggerServer{
			URL:         "https://www.tradelingdev.com/api/module-templates",
			Description: "dev_server",
		}).
		AddServer(SwaggerServer{
			URL:         "https://www.tradelingstage.com/api/module-templates",
			Description: "stage_server",
		}).
		AddServer(SwaggerServer{
			URL:         "https://www.tradeling.com/api/module-templates",
			Description: "prod_server",
		}).SetShouldGenerateSwagger(config.IS_SWAGGER)

	s.MainModule.RegisterRoutes(s.Engine)

	swagger.GenerateSwagger()

}

var server *Server

func printAbout() {
	fmt.Println("=====================================================================")
	fmt.Println("=                 ", Colorize(magenta, "==== Obadas Framework ===="), "                      =")
	fmt.Println("= ", Colorize(yellow, "Designed and Developed by"), Colorize(green, "Amr Abada"), Colorize(yellow, " <amr.app.engine@gmail.com>"), " =")
	fmt.Println("=====================================================================")

}

func NewServer() *Server {
	if server == nil {
		printAbout()
		gin.SetMode(gin.ReleaseMode)
		engine := gin.New()
		engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s x-trace-id %s \"%s%s%s %s %s %s%d%s %s \"%s\" %s\" %s\n",
				param.TimeStamp.Format(time.RFC3339),
				param.Request.Header.Get("x-trace-id"),
				param.MethodColor(), param.Method, param.ResetColor(),
				param.Path,
				param.Request.Proto,
				param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
				param.Latency,
				param.Request.RemoteAddr,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
		engine.Use(gin.Recovery())
		server = &Server{
			Port:             ConfigInstance().Port,
			Engine:           engine,
			TimeoutInSeconds: 10,
			MaxHeaderBytes:   1 << 20, // shift binary 1 by 20  = 131kb
			DB:               NewDatabase(),
		}
		fmt.Printf("%s Server starting on Port %d\n", starting, server.Port)
	}
	return server
}
