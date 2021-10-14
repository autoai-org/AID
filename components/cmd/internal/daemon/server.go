package daemon

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/autoai-org/aid/internal/utilities"
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
)

// beforeResponse set global header to enable cors and set response header
func beforeResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("aid-version", "1.3.0 @ dev")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
		}
	}
}

func printIpAddress(port string) {
	// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	ifaces, err := net.Interfaces()
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		for _, addrs := range addrs {
			var ip string
			switch v := addrs.(type) {
			case *net.IPNet:
				ip = v.IP.String()
				if strings.Count(ip, ":") < 2 {
					fmt.Println("\t IP Address: " + ip + ":" + port)
				}
			case *net.IPAddr:
			}
		}
		if err != nil {
			utilities.Formatter.Error(err.Error())
		}
	}
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}

}

// RunServer starts the http(s) service
func RunServer(port string) {
	if port == "" {
		port = "17415"
		utilities.Formatter.Warn("Port not specified, using the default " + port)
	}
	welcomeFigure := figure.NewFigure("AID Project", "", true)
	utilities.Formatter.Info("Starting the server...")
	r := getRouter()
	welcomeFigure.Print()
	utilities.Formatter.Info("If you are on public network, open https://console.autoai.dev to connect to this server.")
	utilities.Formatter.Info("Listening on the following network: ")
	printIpAddress(port)
	err := r.Run("0.0.0.0:" + port)
	utilities.ReportError(err, "Cannot start server")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utilities.Formatter.Info("Server is shutting down gracefully...")
}
