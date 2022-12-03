# anylogbtc

## Constraint

- 

## API Reference

#### Get all items

```http
  POST /api/wallets/
```

Request body:

| field       | Type      | Description                  |
| :---------- | :-------- | :--------------------------- |
| `datetime`  | `string`  | **Required**. ISODATE string |
| `amount`    | `float64` | **Required**. btc amount     |
| `sender`    | `int`     | **Required**. id sender      |
| `recipient` | `int`     | **Required**. btc recipient  |

#### Get history

```http
  GET /api/wallets/{id}/history
```

| Parameter       | Type     | Description                       |
| :-------------- | :------- | :-------------------------------- |
| `id`            | `string` | **Required**. Id of item to fetch |
| `startDatetime` | `string` | **Required**. Start date time     |
| `endDatetime`   | `string` | **Required**. End date time       |

#### add(num1, num2)

Takes two numbers and returns the sum.


## Authors

- [@h4ckm03d](https://www.github.com/h4ckm03d)


