# Test with Go course notes

## Testing multiple cases / Shared setup and teardown

Extract any setup code e.g. database connection. Could also do the setup and the call the extracted test code from the main function:

```go
func TestUserStore (t *testing.T) {
	conn, err := db.Connect(url="localhost:3000", user="brodie", password="passwd")
	us := &UserStore {
		db: conn
}
	...
	t.Run("Create", testUserStore_create(us))
}

func testUserStore_create (us *UserStore) func(t *testing.T) {
	// Run some specific create tests
}
```

## Testing multiple cases / TestMain

```go
func TestMain(m * testing.M) {
	... setup code
	exitCode := m.Run()
	os.Exit(exitCode) // < Exits immediately, no running defer statements
}
```

The main function runs before all other tests.

Trade-offs

## Parallel Tests / Running tests in parallel

Useful for testing real world scenario where something will be accessed concurrently (website requests)

