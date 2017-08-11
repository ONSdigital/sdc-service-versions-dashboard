package main // import "github.com/onsdigital/sdc-service-versions-dashboard"

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/cloudfoundry-community/go-cfenv"
	"golang.org/x/sync/syncmap"
)

const baseURL = "https://raw.githubusercontent.com/ONSdigital/sdc-service-versions"
const timeFormat = "Monday, 2 January 2006 15:04 MST"

type versionKey struct {
	Environment string
	Service     string
}

type templateData struct {
	Timestamp    string
	Environments map[string]string
	Services     []string
	Versions     *syncmap.Map
}

var environments = map[string]string{
	"CAT":            "cat",
	"CI":             "ci",
	"Demo":           "demo",
	"Development":    "dev",
	"Integration":    "int",
	"Pre-Production": "preprod",
	"Production":     "prod",
	"SIT":            "sit",
	"Test":           "test"}

var services = []string{
	"actionexportersvc",
	"actionsvc",
	"casesvc",
	"collectionexercisesvc",
	"collectioninstrumentsvc",
	"notifygatewaysvc",
	"partysvc",
	"samplesvc",
	"sdxgatewaysvc",
	"securemessagesvc",
	"surveysvc"}

// See https://stackoverflow.com/a/45612142
func (t templateData) Version(environment, service string) interface{} {
	version, ok := t.Versions.Load(versionKey{environment, service})
	if !ok {
		fmt.Printf("No value found for %v\n", versionKey{environment, service})
	}
	fmt.Printf("%v = %s\n", versionKey{environment, service}, version)
	return version
}

func main() {
	port := ":8080"
	appEnv, err := cfenv.Current()

	if err == nil {
		log.Println("Found Cloud Foundry environment")
		ps := appEnv.Port
		port = ":" + strconv.FormatInt(int64(ps), 10)
	} else {
		log.Println("No Cloud Foundry environment")
		if v := os.Getenv("PORT"); len(v) > 0 {
			port = v
		}
	}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}

func buildTemplateData() templateData {
	timestamp := time.Now().Format(timeFormat)
	versions := syncmap.Map{}

	for _, environment := range environments {
		for _, service := range services {
			go versionForEnvironment(environment, service, &versions)
			fmt.Print(".")
		}
	}

	return templateData{timestamp, environments, services, &versions}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, buildTemplateData())
}

func versionForEnvironment(environment, service string, versions *syncmap.Map) {
	version := "N/A"
	doc, err := goquery.NewDocument(fmt.Sprintf("%s/%s/services/%s.version", baseURL, environment, service))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" && !strings.Contains(s.Text(), "404: Not Found") {
			version = s.Text()
		}
	})

	versions.Store(versionKey{environment, service}, version)
}
