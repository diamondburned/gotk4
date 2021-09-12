Move all the function stubs into its own file. The Cgo headers should still be
generated on top of each Go file.

This means we have to split the CgoHeaders map into a CgoStubs map as well.
