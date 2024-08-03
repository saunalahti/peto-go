# peto-go
Peto-media.fi scraper that modifies HTML data to JSON format. This is made only for educational purposes.

Peto-go fetches the peto-media.fi website every 30 seconds, and updates the in-memory cache if there is new content.

This can be used to develop for example a website or mobile app based on the data.

## Routes 
### /api/incidents
This endpoint shows the JSON data that is collected on the In-Memory Cache.

#### Severity types
If the incident data has not returned any integer, the `severity` field won't show up!

1 = High (Suuri)\
2 = Medium (Keskisuuri)\
3 = Small (Pieni)

#### Example single incident
```json
{
  "id": "string",
  "location": "string",
  "location_int": "string", # Field wont appear if `location` field is same.
  "incident": "string",
  "severity": 1, # Field wont appear if there is no severity reported.
  "date": "string"
}
```