# Custom Go workspaces directives
The following directives should be put inside the `go.work` file when wanting to compile a specific chain code.

## Akash
```
// Custom Akash directives
use ./chains/akash
replace (
	github.com/cosmos/cosmos-sdk => github.com/ovrclk/cosmos-sdk v0.45.4-akash.1
)
```