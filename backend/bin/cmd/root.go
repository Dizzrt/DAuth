package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dizzrt/DAuth/backend/common/config"
	"github.com/Dizzrt/DAuth/backend/common/dlog"
	"github.com/Dizzrt/DAuth/backend/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// http://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=DAUTH
	greetingBanner = `
___________________________________________________________________________________________

██████╗  █████╗ ██╗   ██╗████████╗██╗  ██╗
██╔══██╗██╔══██╗██║   ██║╚══██╔══╝██║  ██║
██║  ██║███████║██║   ██║   ██║   ███████║
██║  ██║██╔══██║██║   ██║   ██║   ██╔══██║
██████╔╝██║  ██║╚██████╔╝   ██║   ██║  ██║
╚═════╝ ╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝

%s
___________________________________________________________________________________________
`

	// http://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=BYE
	byeBanner = `
	
██████╗ ██╗   ██╗███████╗
██╔══██╗╚██╗ ██╔╝██╔════╝
██████╔╝ ╚████╔╝ █████╗  
██╔══██╗  ╚██╔╝  ██╔══╝  
██████╔╝   ██║   ███████╗
╚═════╝    ╚═╝   ╚══════╝
                         
`
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dauth",
	Short: "Dauth is an identity provider and SSO platform",
	Run: func(cmd *cobra.Command, args []string) {
		start()
		_ = dlog.Sync()

		fmt.Printf("%s\n", byeBanner)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".bin" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".bin")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func start() {
	var s *server.Server

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		dlog.Infof("%s received", sig.String())
		if s != nil {
			_ = s.Shutdown(ctx)
		}
		cancel()
	}()

	s, err := server.NewServer(ctx)
	if err != nil {
		dlog.Panic("create server failed with error: %v", err)
	}

	if err := s.Run(ctx, config.ServerGetPort()); err != nil {
		if err != http.ErrServerClosed {
			dlog.Errorf("start server failed with error: %v", err)
			_ = s.Shutdown(ctx)
			cancel()
		}
	} else {
		msg := fmt.Sprintf("Version %s has started on port %d 🚀", "dev", 80)
		fmt.Printf(greetingBanner, msg)
	}

	// wait for ctrl-c
	<-ctx.Done()
}
