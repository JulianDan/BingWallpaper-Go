package bingimage

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type BingImage struct {
    URL         string `json:"url"`
    StartDate   string `json:"startdate"`
    FullStartDate   string `json:"fullstartdate"`
    EndDate     string `json:"enddate"`
    URLBase     string `json:"urlbase"`
    Copyright  string `json:"copyright"`
}

func GetBingImage() (*BingImage, error) {
    url := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result struct {
        Images []BingImage `json:"images"`
    }
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, err
    }

    if len(result.Images) > 0 {
        return &result.Images[0], nil
    } else {
        return nil, fmt.Errorf("no image found")
    }
}
