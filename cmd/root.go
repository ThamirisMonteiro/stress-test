package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	neturl "net/url"
	"time"
)

var (
	url            string
	requests       int
	concurrency    int
	defaultTimeout = 5 * time.Second
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Sistema em CLI para realizar testes de carga em serviços web",
	Run: func(cmd *cobra.Command, args []string) {
		if !isValidURL(url) {
			fmt.Errorf("URL inválida")
			return
		}

		if requests <= 0 {
			fmt.Errorf("Número de requests inválido")
			return
		}

		if concurrency <= 0 {
			fmt.Errorf("Número de chamadas simultâneas inválido")
			return
		}

		executeLoadTest(url, requests, concurrency, defaultTimeout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "URL")
	rootCmd.PersistentFlags().IntVar(&requests, "requests", 100, "Número de Requests")
	rootCmd.PersistentFlags().IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")

	err := rootCmd.MarkPersistentFlagRequired("url")
	if err != nil {
		return
	}
}

func isValidURL(str string) bool {
	parsedURL, err := neturl.Parse(str)
	return err == nil && (parsedURL.Scheme == "http" || parsedURL.Scheme == "https")
}
