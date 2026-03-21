package dto

type RootResponse struct {
	Status             int    `json:"status" example:"200"`
	Message            string `json:"message" example:"Welcome to the Whatsmiau API, a Evolution API alternative, it is working!"`
	Version            string `json:"version" example:"0.3.2"`
	ClientName         string `json:"clientName" example:"whatsmiau"`
	Documentation      string `json:"documentation" example:"https://doc.evolution-api.com"`
	WhatsappWebVersion string `json:"whatsappWebVersion" example:"2.3000.0"`
}
