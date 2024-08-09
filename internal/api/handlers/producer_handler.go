package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/internal/model"
	"gitlab.ezrpro.in/godemo/internal/service"
	"gitlab.ezrpro.in/godemo/pkg"
)

type ProducerHandler struct {
	producerService service.IProducerService
}

func NewProducerHandler(producerService service.IProducerService) *ProducerHandler {
	return &ProducerHandler{
		producerService: producerService,
	}
}

func (p *ProducerHandler) SendTest(c *gin.Context) {

	topic := c.Param("topic")
	var project model.Project
	err := c.ShouldBindJSON(&project)

	msg, err := json.Marshal(project)
	partition, offset, err := p.producerService.Send(topic, string(msg))

	if err != nil {
		pkg.Failure(c, 500, err)
		return
	}
	pkg.Success(c, map[string]interface{}{"partition": partition, "offset": offset})
}
