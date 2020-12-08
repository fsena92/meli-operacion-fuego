package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/fsena92/meli-operacion-fuego/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTopSecretWithValidSatellitesAndCoordinatesAndMessagesInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"kenobi",
				  "distance":485.41,
				  "message":[
					 "este",
					 "",
					 "",
					 "mensaje",
					 ""
				  ]
			   },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600.52,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"position":{"x":-100,"y":75},"message":"este es un mensaje secreto"}`, w.Body.String())
}

func TestTopSecretWithInvalidSatellitesCoordinatesAndMessagesInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"kenobi",
				  "distance":487,
				  "message":[
					 "este",
					 "",
					 "",
					 "mensaje",
					 ""
				  ]
			   },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"description":"Invalid distances with coordinates"}`, w.Body.String())
}

func TestTopSecretWithMissingSatelliteInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"Satellites missing or repeated"}`, w.Body.String())
}

func TestTopSecretWithRepeatedSatelliteInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
				{
					"name":"sato",
					"distance":487,
					"message":[
					   "este",
					   "",
					   "",
					   "mensaje",
					   ""
					]
				 },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"Satellites missing or repeated"}`, w.Body.String())
}

func TestTopSecretWithSatelliteNotConfiguredInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
				{
					"name":"yoda",
					"distance":487,
					"message":[
					   "este",
					   "",
					   "",
					   "mensaje",
					   ""
					]
				 },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"Satellites missing or repeated"}`, w.Body.String())
}

func TestTopSecretWithEmptyMessageInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"kenobi",
				  "distance":485.41,
				  "message":[
					   "este",
					   "",
					   "",
					   "mensaje",
					   ""
					]
			   },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[]
			   },
			   {
				  "name":"sato",
				  "distance":600.52,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"Invalid message or invalid message length in any satellite"}`, w.Body.String())
}

func TestTopSecretWithInvalidMessageLengthInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"kenobi",
				  "distance":485.41,
				  "message":[
					   "este",
					   "",
					   "",
					   "mensaje",
					   ""
					]
			   },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":["este",
				  "",
				  "",
				  "mensaje",
				  ""]
			   },
			   {
				  "name":"sato",
				  "distance":600.52,
				  "message":[
					 "este",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"Invalid message or invalid message length in any satellite"}`, w.Body.String())
}

func TestTopSecretWithMissingMessageInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"kenobi",
				  "distance":485.41
				  
			   },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600.52,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error with body request":"Key: 'Request.Satellites[0].Message' Error:Field validation for 'Message' failed on the 'required' tag"}`, w.Body.String())
}

func TestTopSecretWithMissingDistanceInRequest(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret", strings.NewReader(`{
			"satellites":[
			   {
				  "name":"kenobi",

				  "message":[
					   "este",
					   "",
					   "",
					   "mensaje",
					   ""
					]
			   },
			   {
				  "name":"skywalker",
				  "distance":265.75,
				  "message":[
					 "",
					 "es",
					 "",
					 "",
					 "secreto"
				  ]
			   },
			   {
				  "name":"sato",
				  "distance":600.52,
				  "message":[
					 "este",
					 "",
					 "un",
					 "",
					 ""
				  ]
			   }
			]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error with body request":"Key: 'Request.Satellites[0].Distance' Error:Field validation for 'Distance' failed on the 'required' tag"}`, w.Body.String())
}

//Top Secret Split
func TestTopSecretSplitPostWithSatelliteNotConfigured(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret_split/yoda", strings.NewReader(`{
					"distance":487,
					"message":[
					   "este",
					   "",
					   "",
					   "mensaje",
					   ""
					]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"yoda not registered in the satellites configured"}`, w.Body.String())
}

func TestTopSecretSplitPostWithSatelliteConfigured(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/topsecret_split/skywalker", strings.NewReader(`{
					"distance":265.75,
					"message":[
					"",
					"es",
					"",
					"",
					"secreto"
					]
		}`))
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTopSecretSplitGetWithoutSatellitesLoaded(t *testing.T) {
	router := gin.Default()
	api.Setup(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/topsecret_split", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"description":"Missing distances and messages from satellites"}`, w.Body.String())
}
