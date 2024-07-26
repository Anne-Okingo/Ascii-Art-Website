<h><b> ASCII-ART-WEB </h></b>

## Description 
* ASCII-art-web is a program written in Go, that consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version that takes a text from web as input and outputs a graphic representation of the text using ASCII ART characters. 
* This program supports numbers, letters, space, special characters, and newline characters. 

## Features ##
The primary functionality of Ascii-art-Web involves converting text string into ASCII art representation.
Features:
* Server Setup: The application runs a web server using http package.
* Handling Requests: GET /: Serves the template.html file located in the templates directory. POST /: Processes form submissions containing user input for generating ASCII art.
* Ascii-art Representation: Reads ASCII character representations from three files (standard.txt, thinkertoy.txt and shadow.txt). Processes user input to handle special characters like newline (\n).
## Error Handling:
This project handles errors in corellation to the below endpoints

    Not Found, if nothing is found, for example templates or banners.
    400 Bad Request, for incorrect requests.
    500 Internal Server Error, for unhandled errors.

* Checks for empty ASCII art bannerfiles (standard.txt, thinkertoy.txt, shadow.txt).
* Handles issues related to newline characters in the input string.

## Usage: how to run
* To use the programme, the user needs run the following commands:
```bash
git clone https://learn.zone01kisumu.ke/git/aokingo/ascii-art-web  
```
* Start the server on specified port:
```bash
go run . 
```
* Open web server and navigate to 
```bash
http://localhost:8080/
```
* input the text in the text area.
* select Banner file
* press "Generate Ascii-art"

## Authors
[aokingo](https://github.com/Anne-Okingo)

[shaokoth](https://github.com/shaokoth)

[vandisi](https://github.com/Vinolia-E)