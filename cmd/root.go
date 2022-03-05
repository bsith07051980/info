package cmd

import (
	"os"
	"fmt"
	"net/url"
	"log"
	"strings"
	"encoding/json"

	"github.com/bsith07051980/info/website"
	"github.com/spf13/cobra"
	"github.com/likexian/whois"
	"github.com/likexian/whois-parser"
)


var rootCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info of something",
	Long: `Get info of a 
- website`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
	    u, err := url.Parse(args[0])
	    if err != nil {
		log.Fatal(err)
	    }
	    parts := strings.Split(u.Hostname(), ".")
	    domain := parts[0] + "." + parts[1]
	    res_raw, err := whois.Whois(domain)
	    if err != nil {
		log.Fatal(err)
	    }
	    res, err := whoisparser.Parse(res_raw)
	    // J stands for json
	    resJ, err := json.MarshalIndent(res.Domain.NameServers, "", " ")
	    if err != nil {
		log.Fatal(err)
	    }

	    fmt.Println("Name Servers: \n", string(resJ))
	    fmt.Println(website.Show())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.info.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


