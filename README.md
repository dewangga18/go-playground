## Go Playground

Go learning examples with notes covering fundamentals, standard library packages, testing, concurrency, context, and relational databases.

**Module:** `goplayground`  
**Go version:** `1.24.1` in `go.mod`

---

### Project Structure

| Folder | Contents |
|--------|----------|
| [`basics/`](basics/) | Types, collections, control flow, functions, structs, interfaces, pointers, packages, and errors |
| [`libraries/`](libraries/) | Standard library examples: formatting, files, strings, time, encoding, I/O, and more |
| [`test/`](test/) | Unit tests, subtests, table-driven tests, benchmarks, assertions, and mocks |
| [`concurrency/`](concurrency/) | Goroutines, channels, timers, synchronization, and Go scheduler examples |
| [`context/`](context/) | Context values, cancellation, timeouts, and deadlines |
| [`database/`](database/) | MySQL and PostgreSQL connections, pooling, and inserts |
| [`docs/`](docs/) | English notes explaining the examples |

### Running Examples

Use Go compatible with `go.mod`, then run commands from the project root:

```bash
go mod download
go test ./test -v -count=1 -run '^TestSquare$'
```

Select tests with `-run`; `-count=1` bypasses cached results. Other starting points are `TestContext` in `./context` and `TestHelloWorld` in `./concurrency`.

> **Note:** Some tests intentionally fail or demonstrate deadlocks, including `TestZeroSquare`, `TestFatal`, and `TestTicker`. Run examples individually when learning.

For `basics/` and `libraries/`, uncomment the selected file's `main()` entry point, then run that file:

```bash
go run basics/0_hello_world.go
```

See [module and run commands](docs/0-init.md) for more details.

### Database Examples

The tests use local MySQL and PostgreSQL with database `test` and account `go` / `lang`. Follow [Database Setup](docs/0-database-setup.md) and [Go Database](docs/16-database.md) for connection details, the `customer` schema, and test commands.

> **Note:** Each successful insert test adds one row to `customer` and leaves it in the database.

---

### Learning Notes

Follow the numbered topics in order, or open the note matching the example you are studying. Read the concept, inspect its source example, then run it and compare the result with the explanation.

| Topic | Documentation |
|-------|---------------|
| 1. Fundamentals | [Types, variables, and conversions](docs/1-basic-fundamental.md) |
| 2. Collections | [Arrays, slices, and maps](docs/2-collections.md) |
| 3. Control flow | [Conditions and loops](docs/3-control-flow.md) |
| 4. Functions | [Parameters, returns, closures, and recursion](docs/4-functions.md) |
| 5. Handling | [Defer, panic, and recover](docs/5-handling.md) |
| 6. Structs | [Fields and methods](docs/6-struct.md) |
| 7. Interfaces | [Implementation and type assertions](docs/7-interface.md) |
| 8. Pointers | [Addresses, references, and receivers](docs/8-pointer.md) |
| 9. Packages | [Imports and exported symbols](docs/9-package.md) |
| 10. Errors | [Error values and custom errors](docs/10-error.md) |
| 11. Testing | [Unit tests, benchmarks, and Testify](docs/11-unit-test.md) |
| 12. Concurrency | [Goroutines, channels, and timers](docs/12-concurrency.md) |
| 13. Synchronization | [Race conditions, mutexes, and deadlocks](docs/13-concurrency-synchronization.md) |
| 14. Synchronization utilities | [Once, Pool, Cond, and atomic operations](docs/14-synchronization-utils.md) |
| 15. Context | [Values, cancellation, and deadlines](docs/15-context.md) |
| 16. Database | [Connections, pooling, and ExecContext](docs/16-database.md) |

### Setup and References

- [Modules, build, and run commands](docs/0-init.md)
- [Standard library reference](docs/0-standard-library.md)
- [Third-party package setup and usage](docs/0-third-party-packages.md)
- [MySQL and PostgreSQL setup](docs/0-database-setup.md)

Documentation writing guidelines are in [AGENTS.md](AGENTS.md).
