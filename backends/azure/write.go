package azure

import (
	"context"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/jhump/protoreflect/desc"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/options"
	"github.com/batchcorp/plumber/writer"
)

// Write performs necessary setup and calls ServiceBus.Write() to write the actual message
func Write(opts *options.Options, md *desc.MessageDescriptor) error {
	ctx := context.Background()

	if err := writer.ValidateWriteOptions(opts, validateWriteOptions); err != nil {
		return errors.Wrap(err, "unable to validate write options")
	}

	writeValues, err := writer.GenerateWriteValues(md, opts)
	if err != nil {
		return errors.Wrap(err, "unable to generate write value")
	}

	client, err := NewClient(opts)
	if err != nil {
		return errors.Wrap(err, "unable to create client")
	}

	a := &ServiceBus{
		Options: opts,
		msgDesc: md,
		client:  client,
		log:     logrus.WithField("pkg", "azure/write.go"),
	}

	if opts.Azure.Queue != "" {
		queue, err := client.NewQueue(opts.Azure.Queue)
		if err != nil {
			return errors.Wrap(err, "unable to create new azure service bus queue client")
		}

		defer queue.Close(ctx)

		a.queue = queue
	} else {
		topic, err := client.NewTopic(opts.Azure.Topic)
		if err != nil {
			return errors.Wrap(err, "unable to create new azure service bus topic client")
		}

		defer topic.Close(ctx)

		a.topic = topic
	}

	for _, value := range writeValues {
		if err := a.Write(ctx, value); err != nil {
			a.log.Error(err)
		}
	}

	return nil
}

// Write writes a message to an ASB topic or queue, depending on which is specified
func (a *ServiceBus) Write(ctx context.Context, value []byte) error {
	if a.Options.Azure.Queue != "" {
		return a.writeToQueue(ctx, value)
	}

	return a.writeToTopic(ctx, value)
}

// writeToQueue writes the message to an ASB queue
func (a *ServiceBus) writeToQueue(ctx context.Context, value []byte) error {
	msg := servicebus.NewMessage(value)
	if err := a.queue.Send(ctx, msg); err != nil {
		return errors.Wrap(err, "message could not be published to queue")
	}

	a.log.Infof("Write message to queue '%s'", a.client.Name)

	return nil
}

// writeToTopic writes a message to an ASB topic
func (a *ServiceBus) writeToTopic(ctx context.Context, value []byte) error {
	msg := servicebus.NewMessage(value)
	if err := a.topic.Send(ctx, msg); err != nil {
		return errors.Wrap(err, "message could not be published to topic")
	}

	a.log.Infof("Write message to topic '%s'", a.client.Name)

	return nil
}

// validateWriteOptions ensures the correct CLI options are specified for the write action
func validateWriteOptions(opts *options.Options) error {
	if opts.Azure.Topic != "" && opts.Azure.Queue != "" {
		return errTopicOrQueue
	}
	return nil
}
