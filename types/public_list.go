// public_list
//
// 사전정보공표목록
//
// https://github.com/seoul-opengov/opengov#분야별-항목설명-별도-표시-없는-경우-not-null항목

package types

type PublicList struct {
	Nid            string `json:"nid"`             // 관리번호(PK)
	Category       string `json:"category"`        // 분야
	Title          string `json:"title"`           // 제목
	Writer         string `json:"writer"`          // 담당자
	DepartmentName string `json:"dept_nm"`         // 부서명
	RegisterDate   string `json:"regist_dt"`       // 등록일시(정보소통광장 등록일지)
	Taxonomy       string `json:"taxonomy"`        // 업무상세분류(1단계~4단계)
	TelephonNumber string `json:"telno,omitempty"` // 전화번호
	Copyright      string `json:"cpyrht"`          // 라이선스(CCL 적용, CC BY, CC BY-ND, CC BY-SA, CC BY-NC, CC BY-NC-SA, CC BY-NC-ND)
	Url            string `json:"url"`             // URL
}
