package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	repo UserRepository
}

type createUserInput struct {
	Name string `json:"name"`
	Whatsapp string `json:"whatsapp"`
	Age int `json:"age"`
	Dpp string `json:"dpp"`
	City string `json:"city"`
}

func NewHandlers(repo UserRepository) *Handlers {
	return &Handlers{repo: repo}
}

func (h *Handlers) CreateUser(c *gin.Context) {
	var input createUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
		return
	}
	var user User

	//considerando a possibilidade de erro, ele analisa se o JSON é valido
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
		return
	}
	if user.Name == "" || user.Whatsapp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and whatsapp are required"})
		return
	}

	t, err := time.Parse("01-02-2006", user.Dpp.Format("01-02-2006"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dpp format: " + err.Error()})
		return
	}
	user.Dpp = t
	
	u := User{
		Name: user.Name,
		Whatsapp: user.Whatsapp,
		Age: user.Age,
		Dpp: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC),
		City: user.City,
	}

	id, err := h.repo.Create(c.Request.Context(), u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db insert error: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "user": u})
}

func (h *Handlers) GetAllUsers(c *gin.Context) {
	users, err := h.repo.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db find error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handlers) GetUsersByDpp(c *gin.Context) {

	dppStr := c.Query("dpp")
	var dppPtr *time.Time

	if dppStr != "" {
		t, err := time.Parse("01-02-2006", dppStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dpp format: " + err.Error()})
			return
		}
		tt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		dppPtr = &tt
	}

	users, err := h.repo.FindByDpp(c.Request.Context(), *dppPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db find error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handlers) GetUsersByCity(c *gin.Context) {
	city := c.Query("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	users, err := h.repo.FindByCity(c.Request.Context(), city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db find error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handlers) GetUsersByAge(c *gin.Context) {
	ageStr := c.Query("age")
	age, err := strconv.Atoi(ageStr)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age format: " + err.Error()})
		return
	}

	users, err := h.repo.FindByAge(c.Request.Context(), age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db find error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
