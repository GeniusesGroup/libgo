/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"syscall"

	"../log"
)

func init() {
	log.Info("-----------------------------Achaemenid Server-----------------------------")
}

// Server represents needed data to serving some functionality such as networks, ...
// to both server and client apps!
type Server struct {
	State           int    // States locate in const of this file.
	RepoLocation    string // Just fill in supported OS
	Manifest        Manifest
	Cryptography    cryptography
	Networks        networks
	StreamProtocols streamProtocols
	Services        services
	Connections     connections
	Nodes           nodes
	Assets          assets
}

// Server State
const (
	ServerStateStop int = iota
	ServerStateRunning
	ServerStateStopping
	ServerStateStarting // plan to start
)

// Init method use to initialize server object with default data.
func (s *Server) Init() {
	defer s.PanicHandler()

	log.Info("Try to initialize server...")
	if s == nil {
		log.Fatal("Try to initialize nil server! Check codes!")
	}
	s.State = ServerStateStarting
	// Get UserGivenPermission from OS

	// Indicate repoLocation
	// TODO::: change to PersiaOS when it ready!
	var ex, err = os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	s.RepoLocation = filepath.Dir(ex)
	log.Info("App start in", s.RepoLocation)

	s.Assets.init(s)
	s.Services.init()
	s.Connections.init(s)
	s.Cryptography.init(s)
}

// Start will start the server and block caller until server shutdown.
func (s *Server) Start() {
	log.Info("Try to start server...")

	s.State = ServerStateRunning

	log.Info("Server start successfully, Now listen to any OS signals ...")

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	// Infinity loop block main goroutine to handle OS signals!
	for {
		var sig = make(chan os.Signal, 1)
		signal.Notify(sig)
		s.HandleOSSignals(sig)
	}
}

// PanicHandler recover from panics if exist to prevent server stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (s *Server) PanicHandler() {
	var r = recover()
	if r != nil {
		log.Warn("Panic Exception: ", r)
		log.Warn("Debug Stack: ", string(debug.Stack()))
	}
}

// HandleOSSignals use to handle OS signals! Caller will block until we get an OS signal
// https://en.wikipedia.org/wiki/Signal_(IPC)
func (s *Server) HandleOSSignals(sigChannel chan os.Signal) {
	var sig = <-sigChannel
	switch sig {
	case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL:
		log.Info("Caught signal to stop server")
		if s.State == ServerStateRunning {
			s.State = ServerStateStopping
			log.Info("Waiting for server to finish and release proccess, It will take up to 60 seconds")
			s.Shutdown()
			os.Exit(s.State)
		}
	case syscall.Signal(0x17): // syscall.SIGURG:
		log.Warn("Caught urgened signal: ", sig)
	case syscall.Signal(10): // syscall.SIGUSR1
		log.Info("Caught signal to reload server assets")
		s.Assets.LoadFromStorage()
	// case syscall.Signal(12): // syscall.SIGUSR2
	default:
		log.Warn("Caught un-managed signal: ", sig)
	}
}

// Shutdown use to graceful stop server!!
func (s *Server) Shutdown() {
	s.Cryptography.shutdown()
	s.Networks.shutdown()
	s.Assets.shutdown()

	// Wait to finish above logic, or kill app in --- second!
	// time.Sleep(10 * time.Second)

	// it must change to ServerStateStop(0) otherwise it means app can't close normally
	s.State = ServerStateStop

	log.Info("Server stopped")
	log.SaveToStorage("Achaemenid", s.RepoLocation)
}
