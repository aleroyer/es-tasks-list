package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aleroyer/es-tasks-list/utils"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	esServer    string
	useSSL      bool
	detailed    bool
	startTime   bool
	cancellable bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "es-tasks-list",
	Short: "Show tasks list for ElasticSearch",
	Long: `Show ongoing tasks/queries like you can 
do on a MariaDB server.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := generateOptions()

		es, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{
				buildElasticSearchURL(),
			},
		})

		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}

		req := esapi.TasksListRequest{
			Actions:  args,
			Detailed: &detailed,
			Human:    true,
		}

		result, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error while getting tasks list: %s", err)
		}

		defer result.Body.Close()

		var data map[string]interface{}
		if err := json.NewDecoder(result.Body).Decode(&data); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}

		utils.PrintTasks(data, options)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().StringVarP(&esServer, "server", "s", "", "ElasticSearch URL (required)")
	RootCmd.Flags().BoolVar(&useSSL, "ssl", false, "Use HTTPS to connect")
	RootCmd.Flags().BoolVar(&detailed, "detailed", false, "Show task's description")
	RootCmd.Flags().BoolVar(&startTime, "start-time", false, "Show task's start time")
	RootCmd.Flags().BoolVar(&cancellable, "cancellable", false, "Show if a task is cancellable")
	RootCmd.MarkFlagRequired("server")
}

func buildElasticSearchURL() string {
	if useSSL {
		return fmt.Sprintf("https://%s/", esServer)
	}
	return fmt.Sprintf("http://%s/", esServer)
}

func generateOptions() map[string]bool {
	return map[string]bool{
		"description": detailed,
		"action":      detailed,
		"start_time":  startTime,
		"cancellable": cancellable,
	}
}
