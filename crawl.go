package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/meinside/opengov-go/types"

	"github.com/meinside/github-api-go" // $ go get -u github.com/meinside/github-api-go
)

// URL 다운로드
func getBytes(url string) (bytes []byte, err error) {
	var res *http.Response
	if res, err = http.Get(url); err == nil {
		defer res.Body.Close()
		bytes, err = ioutil.ReadAll(res.Body)
	}

	return bytes, err
}

// 전체 결재문서 수집
func crawlInfoList() []types.InfoList {
	result := []types.InfoList{}

	if contents, err := github.ListContents("seoul-opengov", "opengov", "info_list"); err == nil {
		for _, content := range contents {
			if strings.HasSuffix(content.Name, ".json") { // JSON only
				if bytes, err := getBytes(content.DownloadUrl); err == nil { // download,
					// 20160523 ~: array of JSON objects
					var infoLists []types.InfoList
					if err := json.Unmarshal(bytes, &infoLists); err == nil { // parse,
						result = append(result, infoLists...)
					} else {
						// ~ 20160522: array of string arrays
						var infoLists [][]string
						if err := json.Unmarshal(bytes, &infoLists); err == nil { // parse,
							for _, infoList := range infoLists[1:] { // XXX - omit first line (header)
								if len(infoList) < 11 { // XXX - malformed line
									continue
								}
								result = append(result, types.InfoList{
									PackageId:              infoList[0],
									DocumentProductionDate: infoList[1],
									TrackCardName:          infoList[2],
									Title:                  infoList[3],
									SourceDepartmentDocumentId: infoList[4],
									Writer:         infoList[5],
									OthndPeriod:    infoList[6],
									DepartmentName: infoList[7],
									OthbsSe:        infoList[8],
									Copyright:      infoList[9],
									Url:            infoList[10],
								})
							}
						} else {
							log.Printf("* Error while parsing info list: %s - %s\n", err, content.DownloadUrl)
						}
					}
				} else {
					log.Printf("* Error while retrieving info list: %s\n", err)
				}
			}
		}
	} else {
		log.Printf("* Error: %s\n", err)
	}

	return result
}

// 전체 사전정보공표목록 수집
func crawlPublicList() []types.PublicList {
	result := []types.PublicList{}

	if contents, err := github.ListContents("seoul-opengov", "opengov", "public_list"); err == nil {
		for _, content := range contents {
			if strings.HasSuffix(content.Name, ".json") { // JSON only
				if bytes, err := getBytes(content.DownloadUrl); err == nil { // download,
					// 20160523 ~: array of JSON objects
					var publicLists []types.PublicList
					if err := json.Unmarshal(bytes, &publicLists); err == nil { // parse,
						result = append(result, publicLists...)
					} else {
						// ~ 20160522: array of string arrays
						var publicLists [][]string
						if err := json.Unmarshal(bytes, &publicLists); err == nil { // parse,
							for _, publicList := range publicLists[1:] { // XXX - omit first line (header)
								if len(publicList) < 10 { // XXX - malformed line
									continue
								}
								result = append(result, types.PublicList{
									Nid:            publicList[0],
									Category:       publicList[1],
									Title:          publicList[2],
									Writer:         publicList[3],
									DepartmentName: publicList[4],
									RegisterDate:   publicList[5],
									Taxonomy:       publicList[6],
									TelephonNumber: publicList[7],
									Copyright:      publicList[8],
									Url:            publicList[9],
								})
							}
						} else {
							log.Printf("* Error while parsing public list: %s - %s\n", err, content.DownloadUrl)
						}
					}
				} else {
					log.Printf("* Error while retrieving public list: %s\n", err)
				}
			}
		}
	} else {
		log.Printf("* Error: %s\n", err)
	}

	return result
}

// 전체 정책연구자료목록 수집
func crawlResearchList() []types.ResearchList {
	result := []types.ResearchList{}

	if contents, err := github.ListContents("seoul-opengov", "opengov", "research_list"); err == nil {
		for _, content := range contents {
			if strings.HasSuffix(content.Name, ".json") { // JSON only
				if bytes, err := getBytes(content.DownloadUrl); err == nil { // download,
					// 20160523 ~: array of JSON objects
					var researchLists []types.ResearchList
					if err := json.Unmarshal(bytes, &researchLists); err == nil { // parse,
						result = append(result, researchLists...)
					} else {
						// ~ 20160522: array of string arrays
						var researchLists [][]string
						if err := json.Unmarshal(bytes, &researchLists); err == nil { // parse,
							for _, researchList := range researchLists[1:] { // XXX - omit first line (header)
								if len(researchList) < 15 { // XXX - malformed line
									continue
								}
								result = append(result, types.ResearchList{
									Nid:                    researchList[0],
									Title:                  researchList[1],
									RegisterDate:           researchList[2],
									RelmCl:                 researchList[3],
									CreateYear:             researchList[4],
									Category:               researchList[5],
									Region:                 researchList[6],
									Isbn:                   researchList[7],
									RelteArea:              researchList[8],
									Writer:                 researchList[9],
									DocumentProductionDate: researchList[10],
									Copyright:              researchList[11],
									OthbsSe:                researchList[12],
									JobSe:                  researchList[13],
									Url:                    researchList[14],
								})
							}
						} else {
							log.Printf("* Error while parsing research list: %s - %s\n", err, content.DownloadUrl)
						}
					}
				} else {
					log.Printf("* Error while retrieving research list: %s\n", err)
				}
			}
		}
	} else {
		log.Printf("* Error: %s\n", err)
	}

	return result
}
