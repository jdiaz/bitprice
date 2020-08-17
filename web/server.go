package web

import (
	"log"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
)

type CoindeskResponse struct {
	Bpi map[string]Currency `json: "bpi"`
	CharName string `json: "charName"`
	Disclaimer string `json: "disclaimer"`
	Time Time `json: "time"`
}

type Bpi struct {
	Currency Currency `json: "currency"`
}

type Currency struct {
	Code string `json: "code"`
	Description string `json: "description"`
	Rate string `json: "rate"`
	Symbol string `json: "symbol"`
	Rate_float float64 `json: "rate_float"`
}

type Time struct {
	Updated string `json: "updated"`
	UpdatedISO string `json: "updatedISO"`
	Updateduk string `json: "updateduk"`
}

const endpoint = "https://api.coindesk.com/v1/bpi/currentprice.json"

const htmlTemplate = `
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/mini.css/3.0.1/mini-default.min.css">
	</head>
	<body>
		<div class="container">
			<div class="row"></div>
			<div class="row">
				<div class="col-sm-12 col-lg-12">
					<div class="card fluid warning">
						<div class="section">
								<h1 class="doc">Bitcoin price: $%f</h1>
						</div>
						<div class="section">
							<p>
								Powered by <a href="https://www.coindesk.com/price/bitcoin">CoinDesk</a>
							</p>
						</div>
					</div>
				</div>
			</div>
			<div class="row"></div>
		</div>
	</body>
</html>
`

func FetchBTCPrice() float64 {
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(string(body))
	coindeskData := CoindeskResponse{}
	json.Unmarshal(body, &coindeskData)

	price := coindeskData.Bpi["USD"].Rate_float
	log.Printf("Bitcoin price: $%f\n", price)
	return price
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	price := FetchBTCPrice()
	io.WriteString(w, fmt.Sprintf(htmlTemplate, price))
}

func StartPriceServer(port int) {
	http.HandleFunc("/", landingPageHandler)
	log.Printf("Starting server on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d",port), nil)
}
