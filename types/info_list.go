// info_list
//
// 결재문서
//
// https://github.com/seoul-opengov/opengov#분야별-항목설명-별도-표시-없는-경우-not-null항목

package types

type InfoList struct {
	PackageId                  string `json:"package_id"`       // 문서관리번호(PK)
	DocumentProductionDate     string `json:"doc_prdctn_dt"`    // 자료생산일자
	TrackCardName              string `json:"trck_card_nm"`     // 단위과제카드명
	Title                      string `json:"title"`            // 제목
	SourceDepartmentDocumentId string `json:"src_dept_doc_id"`  // 문서번호 (예: 정보공개정책과-1234)
	Writer                     string `json:"writer"`           // 담당자
	OthndPeriod                string `json:"othnd_pd"`         // 문서보존기간(1년, 3년, 5년, 10년, 30년, 준영구, 영구)
	DepartmentName             string `json:"dept_nm"`          // 부서명
	OthbsSe                    string `json:"othbs_se"`         // 공개구분코드(공개, 부분공개, 비공개)
	Copyright                  string `json:"cpyrht,omitempty"` // 라이선스(CCL 적용, CC BY, CC BY-ND, CC BY-SA, CC BY-NC, CC BY-NC-SA, CC BY-NC-ND)
	Url                        string `json:"url,omitempty"`    // 원문공개URL
}
