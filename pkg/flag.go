package pkg

import "flag"

type ParsedFlags struct {
	From   string
	To     string
	Since  string
	Until  string
	Author string
	Ignore string
}

func ParseFlags() ParsedFlags {
	var from string
	var to string
	var since string
	var until string
	var author string
	var ignore string

	flag.StringVar(&from, "from", "", "The hash of the commit to start from")
	flag.StringVar(&to, "to", "", "The hash of the commit to end at")
	flag.StringVar(&since, "since", "", "The date to start from")
	flag.StringVar(&until, "until", "", "The date to end at")
	flag.StringVar(&author, "author", "", "The author to filter by")
	flag.StringVar(&ignore, "ignore", "", "Regex of files to ignore")

	flag.Parse()

	if from == "" && since == "" {
		panic("You must provide either a from hash or a since date")
	}

	return ParsedFlags{
		From:   from,
		To:     to,
		Since:  since,
		Until:  until,
		Author: author,
		Ignore: ignore,
	}
}
