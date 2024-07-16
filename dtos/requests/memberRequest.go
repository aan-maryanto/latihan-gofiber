package requests

type MemberRequest struct {
	Name       string `json:"name"`
	IdentityNo string `json:"identity_no"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
}
