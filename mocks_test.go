package rabbitmq_core


import (
	"github.com/stretchr/testify/suite"
	"mocks_rabbitmq/mocks/amqptest"
	"mocks_rabbitmq/mocks/amqptest/server"
	"testing"
)

type RabbitmqTestSuite struct {
	suite.Suite
	fakeServer *server.AMQPServer
	// your define rabbitmq object
	//fakeRabbitMQ *RabbitMQ
}

func (suite *RabbitmqTestSuite) SetupSuite() {
	// mock server
	fakeServer := server.NewServer("amqp://localhost:35672/%2f")
	err := fakeServer.Start()
	if err != nil {
		return
	}
	mockConn, err := amqptest.Dial("amqp://localhost:35672/%2f")
	if err != nil {
		return
	}
	suite.fakeServer = fakeServer
	// init object
	//rabbitmObj := &RabbitMQ{
	//	Addr: "amqp://localhost:5672/%2f",
	//	PublishFactory: &PublishFactory{
	//		ReconnectTimes: 1,
	//		RepublishTimes: 1,
	//		ChannelContexts: make(map[string]*PbChannelContext),
	//	},
	//	ReceiverFactory: &ReceiverFactory{
	//	},
	//}
	//suite.fakeRabbitMQ = rabbitmObj
	MockRabbitMQFunc(mockConn)
}

//func (suite *RabbitmqTestSuite) TestRabbitmqPublish() {
//	// init object
//	params := &PublishParam {
//		Exchange: "test",
//		RoutingKey: "test",
//		QueueName: "test",
//		Durable: true,
//		Mode: Topic_Mode,
//		Reliable: true,
//	}
//
//	body := map[string]interface{}{
//		"test": "test",
//	}
//	byteBody, _ := json.Marshal(body)
//	err := suite.fakeRabbitMQ.Publish(params, byteBody)
//	suite.Equal(err, nil)
//}
//
//func (suite *RabbitmqTestSuite) TestRabbitmqConsumer() {
//	// publish info
//	go func() {
//		time.Sleep(5 * time.Second)
//		params := &PublishParam {
//			Exchange: "test",
//			RoutingKey: "test",
//			QueueName: "test",
//			Durable: true,
//			Mode: Topic_Mode,
//			Reliable: true,
//		}
//
//		body := map[string]interface{}{
//			"test": "test",
//		}
//		byteBody, _ := json.Marshal(body)
//		_ = suite.fakeRabbitMQ.Publish(params, byteBody)
//		//suite.Equal(err, nil)
//	}()
//
//	fn := func(msg []byte) bool {
//		fmt.Println(msg)
//		suite.Equal(string(msg), "{\"test\":\"test\"}")
//		return true
//	}
//	ctx, cancelFunc := context.WithCancel(context.Background())
//	go func() {
//		time.Sleep(10 * time.Second)
//		cancelFunc()
//	}()
//	err := suite.fakeRabbitMQ.Consumer(fn, ctx)
//	suite.Equal(err, nil)
//}


func TestRabbitmqSuite(t *testing.T) {
	suite.Run(t, new(RabbitmqTestSuite))
}

func (suite *RabbitmqTestSuite) TearDownSuite() {
	suite.fakeServer.Stop()
	UnMockRabbitMQFunc()
}