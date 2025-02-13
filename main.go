package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Constantes para colores ANSI
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

type Endpoint struct {
	Name        string
	URL         string
	Environment string
}

func checkEndpoint(endpoint Endpoint) {
	resp, err := http.Get(endpoint.URL)
	envColor := colorBlue
	if endpoint.Environment == "PROD" {
		envColor = colorYellow
	}

	if err != nil {
		fmt.Printf("%s[%s]%s %s❌%s %s: Error al verificar %s: %v\n",
			envColor,
			endpoint.Environment,
			colorReset,
			colorRed,
			colorReset,
			endpoint.Name,
			endpoint.URL,
			err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("%s[%s]%s %s✅%s %s: OK (200)\n",
			envColor,
			endpoint.Environment,
			colorReset,
			colorGreen,
			colorReset,
			endpoint.Name)
	} else {
		fmt.Printf("%s[%s]%s %s❌%s %s: Error %d\n",
			envColor,
			endpoint.Environment,
			colorReset,
			colorRed,
			colorReset,
			endpoint.Name,
			resp.StatusCode)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando archivo .env")
		os.Exit(1)
	}

	endpoints := []Endpoint{
		// Endpoints de QA
		{"Login", os.Getenv("VITE_API_LOGIN"), "QA"},
		{"Logout", os.Getenv("VITE_API_LOGOUT"), "QA"},
		{"Forgot Password", os.Getenv("VITE_API_FORGOT_PASSWORD"), "QA"},
		{"Reset Password", os.Getenv("VITE_API_RESET_PASSWORD"), "QA"},
		{"Register", os.Getenv("VITE_API_REGISTER"), "QA"},
		{"Delete Profile", os.Getenv("VITE_API_DELETE_PROFILE"), "QA"},
		{"Update Profile", os.Getenv("VITE_API_UPDATE_PROFILE"), "QA"},
		{"Fetch Profile", os.Getenv("VITE_API_FETCH_PROFILE"), "QA"},
		{"Mark4Users", os.Getenv("VITE_API_MARK4USERS"), "QA"},
		{"Chatbot", os.Getenv("VITE_API_CHATBOT"), "QA"},
		{"Map Generator", os.Getenv("VITE_API_MAPGENERATOR"), "QA"},
		{"Get Category By ID", os.Getenv("VITE_API_GET_CATEGORY_BY_ID"), "QA"},
		{"Get Category", os.Getenv("VITE_API_GET_CATEGORY"), "QA"},
		{"Create Product", os.Getenv("VITE_API_CREATE_PRODUCT"), "QA"},
		{"Socket", os.Getenv("VITE_API_SOCKET"), "QA"},
		{"Fetch Chats", os.Getenv("VITE_API_FETCH_CHATS"), "QA"},
		{"Start Chat", os.Getenv("VITE_API_START_CHAT"), "QA"},
		{"Get User Items", os.Getenv("VITE_API_GET_USER_ITEMS"), "QA"},
		{"Update Product", os.Getenv("VITE_API_UPDATE_PRODUCT"), "QA"},
		{"Delete Product", os.Getenv("VITE_API_DELETE_PRODUCT"), "QA"},
		{"Get Products", os.Getenv("VITE_API_GET_PRODUCTS"), "QA"},
		{"Get Products By ID", os.Getenv("VITE_API_GET_PRODUCTS_BY_ID"), "QA"},
		// Endpoints de producción
		{"Login", os.Getenv("VITE_API_LOGIN_PROD"), "PROD"},
		{"Logout", os.Getenv("VITE_API_LOGOUT_PROD"), "PROD"},
		{"Forgot Password", os.Getenv("VITE_API_FORGOT_PASSWORD_PROD"), "PROD"},
		{"Reset Password", os.Getenv("VITE_API_RESET_PASSWORD_PROD"), "PROD"},
		{"Register", os.Getenv("VITE_API_REGISTER_PROD"), "PROD"},
		{"Delete Profile", os.Getenv("VITE_API_DELETE_PROFILE_PROD"), "PROD"},
		{"Update Profile", os.Getenv("VITE_API_UPDATE_PROFILE_PROD"), "PROD"},
		{"Fetch Profile", os.Getenv("VITE_API_FETCH_PROFILE_PROD"), "PROD"},
		{"Chatbot", os.Getenv("VITE_API_CHATBOT_PROD"), "PROD"},
		{"Map Generator", os.Getenv("VITE_API_MAPGENERATOR_PROD"), "PROD"},
		{"Mark4Users", os.Getenv("VITE_API_MARK4USERS_PROD"), "PROD"},
		{"Create Product", os.Getenv("VITE_API_CREATE_PROD"), "PROD"},
		{"Get Products", os.Getenv("VITE_API_GET_PRODUCTS_PROD"), "PROD"},
		{"Get Products By ID", os.Getenv("VITE_API_GET_PRODUCTS_ID_PROD"), "PROD"},
		{"Get Category", os.Getenv("VITE_API_GET_CATRGORY_PROD"), "PROD"},
		{"Get Category By ID", os.Getenv("VITE_API_GET_CATEGORY_ID_PROD"), "PROD"},
		{"Delete Product", os.Getenv("VITE_API_DELETE_PROD"), "PROD"},
		{"Get User Items", os.Getenv("VITE_API_PRODUCTS_BY_USER_PROD"), "PROD"},
		{"Update Product", os.Getenv("VITE_API_UPDATE_PROD"), "PROD"},
		{"Start Chat", os.Getenv("VITE_API_MESSAGES_PROD"), "PROD"},
	}

	fmt.Println("Iniciando monitoreo de servicios...")

	// Verificación inicial
	fmt.Println("\n--- Verificación inicial " + time.Now().Format("15:04:05") + " ---")
	for _, endpoint := range endpoints {
		checkEndpoint(endpoint)
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		fmt.Println("\n--- Verificación de estados " + time.Now().Format("15:04:05") + " ---")
		for _, endpoint := range endpoints {
			checkEndpoint(endpoint)
		}
	}
}
