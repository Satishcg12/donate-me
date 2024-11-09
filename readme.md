# Donate me

## Todo

- [x] Create a index page
- [x] Create a thank you page
- [ ] Create a error page
- [ ] Create about section
- [ ] Create recent donation section
- [ ] Create donation bar(with sse)

## Description

A simple donation website where users can donate money to the organization. The organization can create a campaign and users can donate money to the campaign. The organization can see the list of donors and the amount donated by the donors.

## Tech Stack

- go lang  [ Back-end ]
- sqlite [ Database ]
- templ [ Template Engine ]
- htmx [ Front-end ]
- tailwindcss [ Front-end ]
- Alpine.js [ Front-end ]
- chi [ Router ]
- goose [ Migration ]
- sqlc [ Query Builder ]
- esewa [ Payment Gateway ]

## Installation



## Quick Note

- To migrate the database, run the following command
```bash
# Migrate the database
goose -dir ./internal/db/migrations sqlite3 ./donation.db create add_donation_t
able  sql

# Run the migration
goose -dir ./internal/db/migrations sqlite3 ./donation.db up
```

- To run the server, run the following command
```bash
go run ./
```

- While developing, run the following command to watch the css file
```bash
tailwindcss -i ./assets/css/tailwind.css -o ./assets/css/style.css --watch
```

- To generate the sqlc file, run the following command
```bash
# Generate the sqlc file
sqlc generate
```




http://localhost:3000/api/v1/donation/success/3?data=eyJ0cmFuc2FjdGlvbl9jb2RlIjoiMDAwOTBBMCIsInN0YXR1cyI6IkNPTVBMRVRFIiwidG90YWxfYW1vdW50IjoiMzAwLjAiLCJ0cmFuc2FjdGlvbl91dWlkIjoiMTczMTE2NzkwNSIsInByb2R1Y3RfY29kZSI6IkVQQVlURVNUIiwic2lnbmVkX2ZpZWxkX25hbWVzIjoidHJhbnNhY3Rpb25fY29kZSxzdGF0dXMsdG90YWxfYW1vdW50LHRyYW5zYWN0aW9uX3V1aWQscHJvZHVjdF9jb2RlLHNpZ25lZF9maWVsZF9uYW1lcyIsInNpZ25hdHVyZSI6ImxmUnlsVE5jUGZPZGtJdkpLMnkxRXY1MHNGZU5ZNFllZjU3L1N5azJKbmM9In0=