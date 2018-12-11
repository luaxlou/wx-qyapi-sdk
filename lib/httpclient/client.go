package httpclient

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

func GetJSON(url string, params map[string]string, resObj interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()

	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)



	if err != nil {
		return err
	}
	json.Unmarshal(body, resObj)




	return nil
}

func PostBody(url string, reqObj interface{}, resObj interface{}) error {
	client := &http.Client{}

	b, err := json.Marshal(reqObj)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}


	json.Unmarshal(body, resObj)

	log.Println(bytes.NewBuffer(body).String())


	return nil
}

func PostJSON(url string, params map[string]string, res interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	for k, v := range params {

		req.PostForm.Add(k, v)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, res)

	return nil
}
