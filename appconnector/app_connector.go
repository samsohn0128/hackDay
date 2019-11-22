package appconnector

type Connect struct {
	Actions []string
}

var instance ApplianceConnector

type ApplianceConnector interface {
	ReturnFalse(connectInfo Connect) error
	//Send(connectInfo Connect) error
}

func GetInstance() ApplianceConnector {
	if instance == nil{
		return &appConnectorStub{}
	}
	return instance
}
