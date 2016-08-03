// research_list
//
// 정책연구자료목록
//
// https://github.com/seoul-opengov/opengov#분야별-항목설명-별도-표시-없는-경우-not-null항목

package types

type ResearchList struct {
	Nid                    string `json:"nid"`            // 관리번호(PK)
	Title                  string `json:"title"`          // 제목
	RegisterDate           string `json:"regist_dt"`      // 등록일시(정보소통광장 등록일지)
	RelmCl                 string `json:"relm_cl"`        // 자료유형(정책연구자료, 논문, 간행물, 세미나)
	CreateYear             string `json:"creat_yr"`       // 생산년도
	Category               string `json:"category"`       // 분야(복지, 여성가족, 경제, 안전, 주택도시계획, 환경, 문화관광, 건강, 교통, 건설, 세금재정, 행정)
	Region                 string `json:"region"`         // 관련지역(서울시 전체, 서울 25개 자치구, 수도권, 전국, 해외)
	Isbn                   string `json:"isbn,omitempty"` // ISBN
	RelteArea              string `json:"relte_area"`     // 원본시스템명
	Writer                 string `json:"writer"`         // 담당자
	DocumentProductionDate string `json:"doc_prdctn_dt"`  // 자료생산일자
	Copyright              string `json:"cpyrht"`         // 라이선스('CCL 적용', 'CC BY', 'CC BY-ND', 'CC BY-SA', 'CC BY-NC', 'CC BY-NC-SA', 'CC BY-NC-ND')
	OthbsSe                string `json:"othbs_se"`       // 공개구분
	JobSe                  string `json:"job_se"`         // 작업구분(I, U, D)
	Url                    string `json:"url"`
}
