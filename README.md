# Go Tinyurl

Go Tinyurl is a url shortener project. It's "Independent", "Testable" , and "Clean" based on articles: https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047

![Image](clean-arch.png "clean-arch")

## Technical stack
1. Go
2. Gin framework
3. Cassandra
4. Redis Caching
5. Docker

## Installation

Install Docker & Go (>1.13)

```bash
make cassandra
make redis
make startdb
make startredis
make migratedb
make dev
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)