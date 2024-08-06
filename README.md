# Tavern
Tavern is one of the easiest solutions for uploading files from the web or your other projects. Simply fill out your .env and run your executable :)

Supporting multiple storage options (local & SFTP) and multiple databases (sqlite & mysql/mariaDB) and more to follow.

Wanna see how to set it up? Read more below!

## How does it work?
Tavern exposes a HTTP server on port 8080, which is responsible for either retrieving files from the buckets or sending them to the buckets while executing some SQL queries.
```mermaid
graph LR
    A[Uploader] -->|Upload 1000x1000.png| D(Tavern)
    B[Downloader] -->|Download 1000x1000.png| D(Tavern)
    C[Uploader] -->|Upload 512x512.png| D(Tavern)
    D -->|Upload file| E[SFTP]
    D -->|Upload file| F[Local]
    D -->|Download file| E[SFTP]

    
    linkStyle 0 stroke:orange
    linkStyle 1 stroke:purple
    linkStyle 2 stroke:orange
    linkStyle 3 stroke:orange
    linkStyle 4 stroke:orange
    linkStyle 5 stroke:purple

```

## Setting up:
With these X easy steps you can have a file tavern of your own!
1. Download the executable from the releases (W.I.P, for now you have to build it yourself)
2. Grab the .env example from the git and save it in a folder you'll remember.
3. Fill out the .env example with your setup
4. Run `tavern encryption:generate` to get your encryption key
5. Run `tavern database:migrate` to setup your database structure
6. Run `tavern daemon:start` to start the webserver.

## Adding a new storage "bucket"
Using `tavern bucket:drivers` you can view the drivers that you are able to use. 
Then by simply using `tavern bucket:create` you can fill out the prompts to get your bucket online, this will also immediatly allow uploads to this bucket.

## Adding API keys
API keys are used for authenticating with your bucket, an API key can be easily generated with `tavern key:generate`, this will return your API key.

## The docs
the docs are available at https://tavern-gg.github.io (W.I.P)

## Supporting the creator
<a href="https://www.buymeacoffee.com/nietthijmen"><img src="https://img.buymeacoffee.com/button-api/?text=Support me&emoji=❤️&slug=nietthijmen&button_colour=FFDD00&font_colour=000000&font_family=Inter&outline_colour=000000&coffee_colour=ffffff" /></a>
