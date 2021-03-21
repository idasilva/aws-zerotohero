package sns

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gopetbot/tidus/help"
	aws2 "github.com/idasilva/aws-zerotohero/lambda/aws"
)

type Topic struct {
	*sns.SNS
	*help.Logging
}

func (t *Topic) PublishMessage(message string, topic string) error {
	result, err := t.Publish(
		&sns.PublishInput{
			Message:  &message,
			TopicArn: &topic,
		})

	if err != nil {
		return err
	}

	t.Logger.Infof("send message to sns completed %s...",result)
	return nil
}

//NewNSN
func NewSNS() *Topic {
	config := aws2.NewRemote()
	sns := sns.New(
		session.Must(session.NewSession(config.Configuration)), config.Configuration,
	)
	return &Topic{
		sns,
		help.NewLog(),
	}
}
