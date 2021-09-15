package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type api1881 struct {
	url    string //https://services.api1881.no/
	key    string
	header string //HTTP_HEAD = { 'Ocp-Apim-Subscription-Key': None,}
}

var api *api1881

/*
Execute API Query
*/
func getHTTPMessageBody(resp *http.Response) string {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error parsing response body: %v", err)
		return "nil"
	}
	return string(b)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintf(w, "custom %d", status)
	}
}

/*
	def address_to_coordinates(self,address=None):
		if not address:
			return None
		params = { "address" : address }
		res = self.run_query(url_context="/geocoding/address", params=params)
		return res
*/
func addressToCoordinates(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query()["address"]
	if address == nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	//params := url.Values{}
	//params.Add("address", )
}

/*
	def coordinates_to_address(self, lat=None, lon=None):
		if not X or not Y:
			return None
		params = { "latitude" : lat, "longtitude": lon }
		res = self.run_query(url_context="/geocoding/address", params=params)
		return res
*/
func coordinatesToAddress(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusNotImplemented)
	return
}

/*
	def phonenumber_lookup(self, phone=None):
		if not phone:
			return None
		res = self.run_query(url_context="/lookup/phonenumber/{}".format(phone))
		return res
*/
func lookupPhoneNumber(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusNotImplemented)
	return
}

/*
	def orgnumber_lookup(self, orgnr=None):
		if not orgnr or not orgnr.isdigit():
			return None

		res = self.run_query(url_context="/lookup/organizationnumber/{}".format(orgnr))
		return res
*/
func lookupOrganization(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusNotImplemented)
	return
}

/*
	Execute API Query
	def run_query(self, url_context=None, params=None, data=None):
		if params:
			params = urllib.parse.urlencode( params, safe="" )
		if data:
			data = urllib.parse.urlencode( data )

		url = self.API_URL + url_context
		if params:
			url += "?{}".format(params)

		resp = get(url, headers=self.HTTP_HEAD)
		if resp.status_code == 200:
			print("{} - {}".format(resp.status_code, json.loads(resp.text)))
			return json.loads(resp.content)
		print("{} - {}".format(resp.status_code, resp.text))
		return None
*/
func executeAPIQuery(urlContext string, urlParam string) string {
	req, err := http.NewRequest("GET", api.url, nil)
	if err != nil {
		fmt.Errorf("Error creating GET request: %v", err)
		return ""
	}
	req.Header.Set(api.header, api.key)
	req.Header.Set("Content-Type", "text/json; charset=utf-8")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("Error creating GET request: %v", err))
		return ""
	}
	fmt.Println(resp.StatusCode)
	message := getHTTPMessageBody(resp)
	return message
}

/*
Defines REST API handlers
*/
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/maltego/addressToCoordinates", addressToCoordinates).Methods("GET")
	r.HandleFunc("/maltego/coordinatesToAddress", coordinatesToAddress).Methods("GET")
	return r
}

/*
 */
func main() {
	api := &api1881{}
	api.url = "https://services.api1881.no/"
	api.key = "6b21fb1aafd247d185ae6af255eb4049"
	api.header = "Ocp-Apim-Subscription-Key"

	r := newRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}
