package vm

import (
	"flag"
	"os"
)

type Options struct {
	VersionFlag  bool
	JavaHomePath string

	MainClass string
	Args      []string
}

func ParseOptions() *Options {
	options := &Options{}
	flag.BoolVar(&options.VersionFlag, "version", false, "Displays version information and exit.")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		options.MainClass = args[0]
		options.Args = args[1:]
	}

	initOptions(options)
	return options
}

func initOptions(options *Options) {
	options.JavaHomePath = getJavaHome()
}

func getJavaHome() string {
	jh := ""
	if jh = os.Getenv("JAVA_HOME"); jh == "" {
		panic("$JAVA_HOME not set!")
	}
	return jh
}
