## Poker Game API

A simple API for a poker game. 

## Documentation

### /api/deck/create

Creates a new deck with 52 cards sequentially ordered. Returns the details of the created deck like deck ID, shuffled, remaining card count.

### Parameters

| Name      | Required |  Type |
| :---        |    :----:   | :----: |
| shuffled      | No       | Query |
| cards   | No        | Query |

The optional parameter cards can be supplied to create a deck with specific cards only. It should contain the card codes in a comma separated format.

A cards code is combination of it's value and it's suites first character all in uppercase. Invalid card values that do not meet this specification will be ignored.

Examples:

| Card Value      | Suite |  Code |
| :---        |    :----:   | :----: |
| Ace      | SPADES       | AS |
| 5      | DIAMONDS       | 5D |
| KING      | HEARTS       | KH |

### /api/deck/open

Opens the specified deck and returns all the remaining cards.

### Parameters

| Name      | Required |  Type |
| :---        |    :----:   | :----: |
| deck_id      | Yes       | Query |

### /api/deck/draw

Draws cards from the specified deck and returns all the drawn cards.

### Parameters

| Name      | Required |  Type |
| :---        |    :----:   | :----: |
| deck_id      | Yes       | Query |
| count      | Yes       | Query |

## Build

Run the following command to start the application. 

The server starts listening on the port 6565 by default. This can be configured via changing the `PORT` option in the `.env` configuration file.

```console
$ go run app.go
```
