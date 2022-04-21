# traking-backend

The server part of the site with tracking the progress of its team, consisting of team leaders and their subordinates

![GO][go-version] ![Python][python-version]

---
## Installation

#### Requirements
* Golang 1.17  
* Python3.8+
* Linux, Windows or macOS

#### Installing
```
git clone https://github.com/Traking-work/traking-backend.git
cd traking-backend
```

#### Configure
To work, you must create a `.env` file in the main directory of the project and specify such variables as:
```
MONDO_DB_URL - link to mongodb database
SALT - a combination of characters to generate a password hash
SECRET_KEY - key for generating authentication tokens
FRONTEND_URL - the link from which the request will come from the frontend
```

Also, in the `configs/config.yml` file, specify your mongodb login and the name of the database

---
## Usage
The port on which the service will be launched is specified in the file `configs/config.yml`

To start, run
```
go build -o traking-backend cmd/app/main.go
./traking-backend
```

---
## Additionally
A `traking-backend.service` file was also created to run this bot on the server


[go-version]: https://img.shields.io/static/v1?label=GO&message=v1.17&color=blue
[python-version]: https://img.shields.io/static/v1?label=Python&message=v3.8&color=blue