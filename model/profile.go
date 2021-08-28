package model

type Profile struct {
	Name          string `json:"name,omitempty"`
	Gender        string `json:"gender,omitempty"`
	Age           int    `json:"age,omitempty"`
	Height        int    `json:"height,omitempty"`
	Weight        int    `json:"weight,omitempty"`
	Income        string `json:"income,omitempty"`
	Marriage      string `json:"marriage,omitempty"`
	Education     string `json:"education,omitempty"`
	Occupation    string `json:"occupation,omitempty"`
	Residence     string `json:"residence,omitempty"`
	Constellation string `json:"constellation,omitempty"`
	House         string `json:"house,omitempty"`
	Car           string `json:"car,omitempty"`
}
