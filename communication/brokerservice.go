package communication

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type BrokerService interface {
	Publish(data interface{}, topic string) error
	Subscribe(topic string, callback SubscriberCallback) error
	Connect() error
}

var QualityOfService byte = 0

type SubscriberCallback func(topic string, data []byte)

type mqttBrokerService struct {
	settings *MQTTSettings
	client   MQTT.Client
}

func NewMQTTBrokerService(settings *MQTTSettings) *mqttBrokerService {
	return &mqttBrokerService{
		settings: settings,
	}
}

func (service *mqttBrokerService) Connect() error {

	// MQTT.DEBUG = log.New(os.Stdout, "", 0)
	MQTT.ERROR = log.New(os.Stdout, "", 0)

	brokerURL := fmt.Sprintf("tcp://%v:%v", service.settings.Server, service.settings.Port)

	log.Printf("MQTT Broker Service is connecting to %v.\n", brokerURL)

	opts := MQTT.NewClientOptions().AddBroker(brokerURL)
	opts.SetClientID(service.settings.ClientID)
	opts.SetKeepAlive(20 * time.Second)
	opts.SetPingTimeout(10 * time.Second)

	service.client = MQTT.NewClient(opts)
	if token := service.client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(" > Connection failed")
		return token.Error()
	}

	log.Println(" > Connected")
	return nil
}

func (service *mqttBrokerService) Publish(data interface{}, topic string) error {
	value, err := json.Marshal(&data)

	if err != nil {
		return err
	}

	strValue := string(value)

	token := service.client.Publish(topic, QualityOfService, false, strValue)
	token.Wait()

	return token.Error()
}

func (service *mqttBrokerService) Subscribe(topic string, callback SubscriberCallback) error {
	token := service.client.Subscribe(topic, QualityOfService, func(client MQTT.Client, message MQTT.Message) {
		callback(message.Topic(), message.Payload())
	})

	token.Wait()
	return token.Error()
}

func (service *mqttBrokerService) Close() {
	service.client.Disconnect(250)
}
