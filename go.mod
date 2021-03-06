module github.com/OliverCardoza/traindown-cli

go 1.17

require (
	github.com/jedib0t/go-pretty/v6 v6.2.4
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v1.3.0
	github.com/traindown/traindown-go v1.1.1-0.20211226150333-1a4f0b99a169
)

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/rivo/uniseg v0.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/timtadh/data-structures v0.5.3 // indirect
	github.com/timtadh/lexmachine v0.2.2 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
)

// TODO: Remove once https://github.com/traindown/traindown-go/pull/3 merged.
replace github.com/traindown/traindown-go v1.1.1-0.20211226150333-1a4f0b99a169 => ../traindown-go
