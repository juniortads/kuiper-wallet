package company

type Company struct {
	Id		         string `ion:"Id"`
	Name             string `ion:"Name"`
	TrackingId       string `ion:"TrackingId"`
	ExternalId	     string `ion:"ExternalId"`
	MetadataId		 string `ion:"MetadataId"`
}
