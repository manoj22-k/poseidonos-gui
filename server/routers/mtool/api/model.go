package api

type Request struct {
	Command string      `json:"command"`
	Rid     string      `json:"rid"`
	Param   interface{} `json:"param"`
}

type Response struct {
	Rid    string `json:"rid"`
	Result Result `json:"result"`
}

type Result struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
