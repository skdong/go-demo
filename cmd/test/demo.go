package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func RunServer(stopCh <-chan struct{}) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to my website")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	go func() {
		http.ListenAndServe(":80", nil)
	}()

	<-stopCh
	<-stopCh

	return nil
}

func ServerCommand(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "demo",
		Long: "just for demo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServer(stopCh)
		},
	}

	return cmd
}

func main() {
	cmd := ServerCommand(SetupSignalHandler())
	if err := cmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
