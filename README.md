# whereis
Service for displaying ones current location.

# Inspiration
This project is heavily inspired by https://github.com/lelandbatey/whereIAm checkout his other projects they're great!

# Required Environment
```
export LOCATION_API=http://localhost:8081/location
// PORT to run on
export PORT=8060
```

A location API is required, technically it can be anything that returns json in the below format. I'm using [current](https://github.com/zaquestion/current)
```
{
    "latitude": 47.38672163,
    "longitude": -122.23929259
}
```
