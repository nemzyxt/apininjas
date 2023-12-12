package utils


const (
	BaseUrl string = "https://api.api-ninjas.com/v1/"
)

type Client struct {
	ApiKey string
}

func Init(apiKey string) Client {
	return Client {
			ApiKey: apiKey,
	}
}
