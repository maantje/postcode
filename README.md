# Dutch Postcode API

Welcome to the Dutch Postcode API written in go! This API provides a simple and efficient way to retrieve location information based on Dutch postcodes. Whether you're building a location-based application or need to enhance your data with geographical details, this API has got you covered.

## Features

- **Accurate Data**: Utilizes up-to-date Dutch postcode data from "Basisregistratie Adressen en Gebouwen (BAG)".
- **Easy Integration**: Simple RESTful API design for easy integration into your applications.
- **SQLite Database Export**: Easily extract addresses as an SQLite database from the Docker container.
- **Compact Image Size**: The Docker image is optimized for efficiency, with a small size of only 364.96 MB.


## Getting Started

1. Run the docker container

```bash
docker run -p 80:80 ghcr.io/maantje/postcode:latest
```

2. Wait for the database to be extracted

3. done!


## Usage

### Endpoint

```plaintext
GET /postcode/{postcode}/{house_number}
```

Replace {postcode} and {house_number} with the desired postcode and house number.

### Example

```curl
curl -X GET http://localhost/postcode/9743EG/258
```

### Response

```json
{
    "postcode": "9743EG",
    "house_number": 258,
    "municipality": "Groningen",
    "city": "Groningen",
    "street_name": "Goudlaan",
    "short_street_name": "",
    "long_street_name": "Goudlaan",
    "area": 124.0,
    "usage": "woonfunctie",
    "built_in": 1970,
    "latitude": 53.2312990867121,
    "longitude": 6.528704482471
}

```

### Extracting the database

The Docker image includes an SQLite database containing all Dutch addresses. Follow these steps to extract the database after starting the container.


1. Identify the container ID:
```bash
docker ps
```

2. Extract the SQLite database:

```bash
docker cp {CONTAINER ID}:/application/database/addresses.sqlite addresses.sqlite
```

You now have the addresses.sqlite file containing all Dutch addresses in your current directory.

## Thanks

Special thanks to @digitaldutch for their contribution to [https://github.com/digitaldutch/BAG_parser](https://github.com/digitaldutch/BAG_parser), which makes this repository possible.