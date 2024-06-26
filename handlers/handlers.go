package handlers

import (
	"net/http"

	"GoApi/auth"

	"GoApi/service"

	"GoApi/utils" // Import the utils package

	"github.com/gin-gonic/gin"
)

type Handler struct {
	addressBookService service.AddressBookService
}

func NewHandler(addressBookService service.AddressBookService) *Handler {
	return &Handler{
		addressBookService: addressBookService,
	}
}

func (h *Handler) HomePage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the HomePage!")
}

// GetAddressBookAll godoc
// @Summary Get all address books
// @Description get all address books
// @ID get-address-book-all
// @Produce  json
// @Success 200 {object} []models.AddressBook
// @Security BearerAuth
// @Router /getAddress [get]
func (h *Handler) GetAddressBookAll(c *gin.Context) {
	books, err := h.addressBookService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetAddressBookByUserName godoc
// @Summary Get by username address books
// @Description get by username address books
// @ID get-address-book-by-username
// @Produce  json
// @Success 200 {object} []models.AddressBook
// @Security BearerAuth
// @Router /GetAddressBookByUserName [get]
func (h *Handler) GetAddressBookByUserName(c *gin.Context) {
	// Extract username from JWT token
	username, err := utils.ExtractUsernameFromToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Use the username as needed, e.g., fetch data associated with the user
	//books, err := h.addressBookService.GetBooksByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": username, "books": nil})
}


// GenerateTokenHandler godoc
// @Summary Generate JWT token
// @Description generate a JWT token for a given username
// @ID generate-token
// @Produce  json
// @Param username query string true "Username"
// @Success 200 {object} map[string]string
// @Router /getToken [get]
func (h *Handler) GenerateTokenHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	token, err := auth.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

