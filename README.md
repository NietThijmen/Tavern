# File tavern.
A discord webhook alternative for storing your files. Fully customizable and easy to use.

## Warning:
this project is still in BETA & needs to be tested.
It has been tested on macOS on an M2 macbook. 

If you have any issues report the issue in an issue or try to create a PR as I've only been using Golang for 3 months


## Installation
1. Grab the latest release from the releases page & grab the executable for your system.
2. Copy the .env.example from Github & put it in the same folder named .env and edit it to your liking.
3. Run the executable and you're good to go!

## Building from source
1. Clone the repository
2. Make sure you have go installed
3. Run `build.sh` to build the project for all platforms. You can also run go build to build it for your current platform.

## Usage
1. Upload an image to the server by sending it to the main domain + SECRET_PATH (e.g. https://tavern.example.com/secret_path)
2. Read the response and use the link to your liking.

## Monitoring
You can enable monitoring for Tavern by editing .env ENABLE_PROMETHEUS to true and setting the PORT to your liking.
This allows you to read your metrics from /metrics on the port you set.

There are a few custom metrics:
- items_uploaded - The amount of items uploaded to the server.
- total_size - The total size of all the items uploaded to the server.
- saved_space - The amount of space saved by using Tavern (compression of images).

## Supporting the creator
<a href="https://www.buymeacoffee.com/nietthijmen"><img src="https://img.buymeacoffee.com/button-api/?text=Support me&emoji=❤️&slug=nietthijmen&button_colour=FFDD00&font_colour=000000&font_family=Inter&outline_colour=000000&coffee_colour=ffffff" /></a>
