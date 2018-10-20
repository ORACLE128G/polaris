package worker

import "polaris/crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(
	request Request,
	result *ParseResult) error {
	engineRequest, err := DeserializeRequest(request)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineRequest)

	if err != nil {
		return err
	}
	*result = SerializedParseResult(engineResult)
	return nil
}
