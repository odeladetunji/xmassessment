package api 

import (
	EventService "xmservice.com/services"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	// Auth "xmservice.com/auth"
)

var eventService EventService.KafkaEventService;

type EventApi struct {

}

func (event *EventApi) Router(router *gin.Engine){
	var route *gin.RouterGroup = router.Group("/api/events");
	var eventApi EventApi;
    eventApi.CreateCompany(route);
}

func (event *EventApi) GetAllKafkaEvents(route *gin.RouterGroup){
	route.GET("/all", func(c *gin.Context) {

		data, err := eventService.GetAllKafkaEvents(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "Kafka events",
			"data": data,
			"status": http.StatusOK,
		});
		
	});
}