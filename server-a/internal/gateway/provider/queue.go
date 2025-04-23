package provider

import (
	"server/internal/queue/processor"
	"server/internal/queue/route"
	"server/internal/queue/tasks"

	"github.com/hibiken/asynq"
)

var (
	emailTask tasks.EmailTask
)

var (
	emailTaskProcessor *processor.EmailTaskProcessor
)

func ProvideQueueModule(client *asynq.Client, mux *asynq.ServeMux) {
	injectQueueModuleTask(client)
	injectQueueModuleProcessor()

	route.EmailTaskRoute(mux, emailTaskProcessor)
}

func injectQueueModuleTask(client *asynq.Client) {
	emailTask = tasks.NewEmailTask(client)
}

func injectQueueModuleProcessor() {
	emailTaskProcessor = processor.NewEmailTaskProcessor(base64Encryptor, smtpUtil)
}
