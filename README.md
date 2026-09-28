# Groupie Tracker

## Description

The project main idea is to use the Groupie Trackers API and display information about different artists and bands in a user-friendly website.

The API contains information about the artists such as their names, members, images, creation dates, first albums, concert locations and concert dates.

The website also allows users to search for artists and view more detailed information about each artist.

The site also has the header as a way to go back the homepage if clicked on.

## Authors

- ahmsalem
- asami

## Usage: how to run

### Run directly with Go

1. From the root project directory, start the server:

```bash
go run .
```

2. Open http://localhost:8080 in a browser.

3. Browse the available artists.

4. Use the search bar to search for an artist.

5. Click on an artist to view their information, members, concert locations and dates.

## Implementation details: algorithm

1. Fetch the artist data from the Groupie Trackers API.
2. Fetch the locations, dates and relations data from their API endpoints.
3. Decode the JSON data into Go structs.
4. Store the API data so it can be used by the HTTP handlers.
5. Display all artists on the home page.
6. Allow the user to search for artists.
7. When an artist is selected, use their ID to find their information.
8. Combine the artist, location, date and relation data.
9. Send the data to the HTML templates.
10. Display the artist information in the browser.
11. Handle invalid pages and server errors using an error page.

## API

The project uses the following Groupie Trackers API:

```txt
https://groupietrackers.herokuapp.com/api
```

The API contains four main endpoints:

- `artists` - artist names, members, images, creation dates and first albums.
- `locations` - concert locations.
- `dates` - concert dates.
- `relation` - connects concert locations with their dates.

## Project Structure

```txt
groupie-tracker/
├── go.mod
├── main.go
├── README.md
├── handlers/
│   ├── artist.go
│   ├── errorpage.go
│   ├── home.go
│   └── search.go
├── structdata/
│   └── data.go
├── templates/
│   ├── artist.html
│   ├── errors.html
│   └── index.html
└── static/
    ├── images/
    │   └── groupie.png
    └── style/
        └── style.css
```