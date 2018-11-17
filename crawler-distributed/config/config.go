package config

const (
	// all parsers
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ProfileParser  = "ProfileParser"
	NilParser     = "NilParser"

	// CrawlService
	CrawlServiceRpc = "CrawlService.Process"


	// worker0
	Worker0Host = ":9000"

	// worker1
	Worker1Host = ":9001"


	// itemsaver0
	ItemSaver0Host = ":9002"

	// itemsaver1
	ItemSaver1Host = ":9003"

	// ItemSaverService.Save

	ItemSaverServiceSave = "ItemSaverService.Save"

	ElasticsearchHosts = "http://192.168.1.9:9200"
)
