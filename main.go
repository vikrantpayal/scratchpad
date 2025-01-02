package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function for the root route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := getPage()
	fmt.Fprintln(w, htmlContent)
}

func getFinTools() string {
	var finToolsHTML string = "Financial Tools"
	return finToolsHTML
}

func getPage() string {
	var htmlContent string

	htmlContent = "<!DOCTYPE html>"
	htmlContent += "<head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'>"
	htmlContent += "<title>Personal tools</title>"
	htmlContent += "<style></style>"
	htmlContent += "<body><main>Personal Tools<br>...<p>"
	htmlContent += "<div id='panel'>" + getFinTools() + "</div>"
	htmlContent += "</p>..."
	htmlContent += "</main></body>"
	htmlContent += "<footer><p>&copy; 2025 Vikrant</p></footer>"

	/*
			<style>
		    body {font-family: Arial, sans-serif;margin: 0;padding: 0;background-color:#f0f0f0;}
			header, footer {background-color: #333;color: #fff;text-align: center;padding: 1em;}
		    main {padding: 2em;}</style></head>
			<body><header>Welcome Vikrant!</header>
		    <main>
		        <p></p>
		    </main>
		    <footer><p>&copy; 2025 Vikrant's Toolkit</p></footer></body></html>"
	*/

	return htmlContent
}

// start the server
func main() {
	http.HandleFunc("/", homeHandler) // Register the handler for the root route
	port := ":8080"                   // Define the port to listen on

	log.Printf("Server starting on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) // Start the server
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
