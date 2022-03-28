# Web-Scraping_golang

# Features of the project

1. User can create a new account.
2. User with a registered account can sign in to the system.
3. User can scrape multiple websites by passing multiple urls.
4. Scrapped content can be stored inside database.
5. All the stored contents can be viewed by anyone. For viewing the stored contents, searching(by url) and filtering(by date) can be used. 

# Technology Used

1. Golang
2. PostgreSQL
3. Docker

# Package Used

1. Gin ( go framework )
2. GORM ( ORM library )
3. Viper ( configuration package )
4. JWT ( for authentication )

# Documentation
  
- For documentation, rest file is used in the root directory of the project. Since I've used Visual Studio Code for my IDE, there's an extension called Rest Client which you need to install in your IDE. This extension helps you execute your rest file and test your system in your IDE.

# Steps to run

1. Install Docker in your system.
2. Clone this repo using the command "git clone https://github.com/bishan20/Web-Scraping_golang.git".
3. Open the project in Visual Studio Code.
4. For Windows, OPEN start.sh and SELECT " LF " in end of line sequence.
5. To run your project in docker, use command "docker-compose up" in your terminal.
6. To test all REST API, use the documentation given above.
7. After you are done with the testing, stop the server using the command "docker-compose down" in your terminal.



