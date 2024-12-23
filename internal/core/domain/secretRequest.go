package domain

type SecretUpdateResponse struct {
	Message string `json:"message"`
}

type SecretGetResponse struct {
	Message string `json:"message"`
	Secret  Secret `json:"secret"`
}

type SecretGetParam struct {
	Action string `json:"action"`
}

type SecretUpdateParam struct {
	Action string `json:"action"`
	Secret Secret `json:"secret"`
}

type Secret struct {
	Password string `json:"password"`
}