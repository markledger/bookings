#  Bookings Application

### Running the Application
From the root folder run: `go run ./cmd/web/ `

### Tidying modules

`go mod tidy`

go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.