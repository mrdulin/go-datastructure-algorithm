package main

import (
	"context"
	"encoding/xml"
	"fmt"
	soap2 "github.com/hooklift/gowsdl/soap"
	"github.com/joho/godotenv"
	"github.com/mrdulin/googleads-go-lib/myservice"
	"github.com/tiaguinho/gosoap"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

const (
	WSDL string = "https://adwords.google.com/api/adwords/cm/v201809/CampaignService?wsdl"
)

var (
	ADWORDS_CLIENT_ID string
	ADWORDS_SECRET string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ADWORDS_CLIENT_ID = os.Getenv("ADWORDS_CLIENT_ID")
	ADWORDS_SECRET = os.Getenv("ADWORDS_SECRET")
	fmt.Println("ADWORDS_CLIENT_ID:",ADWORDS_CLIENT_ID)
	fmt.Println("ADWORDS_SECRET:", ADWORDS_SECRET)
}

type WithHeaderRoundTrip struct {
	r http.RoundTripper
}

func (rt WithHeaderRoundTrip) RoundTrip(r *http.Request) (*http.Response, error) {
	accessToken := os.Getenv("ADWORDS_ACCESS_TOKEN")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return rt.r.RoundTrip(r)
}

type GetCampaignResponse struct {
	XMLName xml.Name `xml:"getResponse"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Rval    struct{
		Text            string `xml:",chardata"`
		TotalNumEntries string `xml:"totalNumEntries"`
		PageType        string `xml:"Page.Type"`
		Entries         []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id"`
			Name   string `xml:"name"`
			Status string `xml:"status"`
			Budget string `xml:"budget"`
		} `xml:"entries"`
	} `xml:"rval"`
}

type GetCampaignResponseHeader struct {
	XMLName      xml.Name `xml:"ResponseHeader"`
	Text         string   `xml:",chardata"`
	Xmlns        string   `xml:"xmlns,attr"`
	RequestId    string   `xml:"requestId"`
	ServiceName  string   `xml:"serviceName"`
	MethodName   string   `xml:"methodName"`
	Operations   string   `xml:"operations"`
	ResponseTime string   `xml:"responseTime"`
}

type GetCampaignRequestHeader struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201809 RequestHeader"`
	ClientCustomerId string `xml:"clientCustomerId"`
	DeveloperToken string `xml:"developerToken"`
	UserAgent string `xml:"userAgent"`
	ValidateOnly bool `xml:"validateOnly,omitempty"`
	PartialFailure bool `xml:"partialFailure,omitempty"`
}


func main() {
	//httpClient := http.Client{
	//	Timeout: 5000 * time.Millisecond,
	//	Transport: WithHeaderRoundTrip{r: http.DefaultTransport},
	//}
	ctx := context.Background()
	conf := oauth2.Config{
		ClientID: ADWORDS_CLIENT_ID,
		ClientSecret: ADWORDS_SECRET,
		RedirectURL: "http://localhost:3000",
		Scopes: []string{"https://www.googleapis.com/auth/adwords"},
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	fmt.Println("Enter the authorization code: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := conf.Client(ctx, tok)

	//GosoapWay(httpClient)
	soap := soap2.NewClient(WSDL, soap2.WithHTTPClient(httpClient))
	soap.AddHeader(struct {
		Headers GetCampaignRequestHeader
	}{
		Headers: GetCampaignRequestHeader{
			ClientCustomerId: os.Getenv("ADWORDS_CLIENT_CUSTOMER_ID"),
			DeveloperToken: os.Getenv("ADWORDS_DEVELOPER_TOKEN"),
			UserAgent:  os.Getenv("ADWORDS_USER_AGENT"),
		},
	})

	campaignServiceInterface := myservice.NewCampaignServiceInterface(soap)
	res, err := campaignServiceInterface.Get(&myservice.Get{ServiceSelector: &myservice.Selector{
		Fields: []string{"Id", "Name", "Status"},
		Paging: &myservice.Paging{
			StartIndex: 0,
			NumberResults: 5,
		},
	}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %+v\n", res)
	fmt.Printf("rval: %+v\n", res.Rval)
}


func GosoapWay(httpClient *http.Client) {
	soap, err := gosoap.SoapClientWithConfig(WSDL, httpClient, &gosoap.Config{
		Dump: true,
	})
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	soap.HeaderName = "RequestHeader"
	soap.HeaderParams = gosoap.HeaderParams{
		"clientCustomerId": os.Getenv("ADWORDS_CLIENT_CUSTOMER_ID"),
		"developerToken": os.Getenv("ADWORDS_DEVELOPER_TOKEN"),
		"userAgent": os.Getenv("ADWORDS_USER_AGENT"),
	}
	res, err := soap.Call("get", gosoap.Params{
		"serviceSelector": []gosoap.Params{
			{"fields": "Id"},
			{"fields": "Name"},
			{"fields": "Status"},
			{
				"paging": gosoap.Params{
					"startIndex": "0",
					"numberResults": "5",
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}
	var rb GetCampaignResponse
	var rh GetCampaignResponseHeader
	err = xml.Unmarshal(res.Header, &rh)
	if err != nil {
		log.Fatalf("Unmarshal header error: %s", err)
	}
	err = xml.Unmarshal(res.Body, &rb)
	if err != nil {
		log.Fatalf("Unmarshal body error: %s", err)
	}
	fmt.Printf("rb: %+v\n", rb)
	fmt.Printf("rh: %+v\n", rh)
}


