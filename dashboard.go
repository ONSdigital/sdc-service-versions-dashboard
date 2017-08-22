package main // import "github.com/onsdigital/sdc-service-versions-dashboard"

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/cloudfoundry-community/go-cfenv"
)

const baseURL = "https://raw.githubusercontent.com/ONSdigital/sdc-service-versions"
const delimiter = ","
const timeFormat = "Monday, 2 January 2006 15:04 MST"

type version struct {
	Version string
	Commit  string
}

type versionKey struct {
	Environment string
	Service     string
}

type templateData struct {
	Timestamp    string
	Environments map[string]string
	Services     []string
	Versions     struct {
		sync.RWMutex
		m map[versionKey]version
	}
}

var versions = struct {
	sync.RWMutex
	m map[versionKey]version
}{m: make(map[versionKey]version)}

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
	"iacsvc",
	"notifygatewaysvc",
	"partysvc",
	"samplesvc",
	"sdxgatewaysvc",
	"securemessagesvc",
	"surveysvc"}

// See https://stackoverflow.com/a/45612142
func (t templateData) Version(environment, service string) version {
	t.Versions.RLock()
	defer t.Versions.RUnlock()
	return t.Versions.m[versionKey{environment, service}]
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

	buildTemplateData()
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}

func buildTemplateData() templateData {
	timestamp := time.Now().Format(timeFormat)

	for _, environment := range environments {
		for _, service := range services {
			go versionForEnvironment(environment, service)
			fmt.Print(".")
		}
	}

	return templateData{timestamp, environments, services, versions}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, buildTemplateData())
}

func getBodyContent(doc *goquery.Document) string {
	var bodyContent string
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" && !strings.Contains(s.Text(), "404: Not Found") {
			bodyContent = s.Text()
		}
	})

	return bodyContent
}

func versionForEnvironment(environment, service string) {
	version := version{Version: "N/A"}
	doc, err := goquery.NewDocument(fmt.Sprintf("%s/%s/services/%s.version", baseURL, environment, service))
	if err != nil {
		log.Fatal(err)
	}

	bodyContent := strings.Split(getBodyContent(doc), delimiter)
	if len(bodyContent) > 1 {
		version.Version = bodyContent[0]
		version.Commit = bodyContent[1]
	}

	versions.Lock()
	versions.m[versionKey{environment, service}] = version
	versions.Unlock()
}
