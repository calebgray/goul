package main

import (
	"flag"
	"github.com/calebgray/goul/goul/archive"
	"github.com/calebgray/goul/goul/arguments"
	"os"
)

// Global Flags
var pipe = flag.Bool("p", false, "extract files to pipe, no messages")
var list = flag.Bool("l", false, "list files (short format)")
var freshen = flag.Bool("f", false, "freshen existing files, create none")
var test = flag.Bool("t", false, "test compressed archive data")
var update = flag.Bool("u", false, "update files, create if necessary")
var comment = flag.Bool("z", false, "display archive comment only")
var timestamp = flag.Bool("T", false, "timestamp archive to latest")
var exclude = flag.Bool("x", false, "exclude files that follow (in xlist)")
var destination = flag.Bool("d", false, "extract files into exdir")

// Modifier Flags
var modFlags = flag.FlagSet{}
var never = modFlags.Bool("n", false, "never overwrite existing files")
var quiet = modFlags.Bool("q", false, "quiet mode")
var quieter = modFlags.Bool("qq", false, "quieter")
var overwrite = modFlags.Bool("o", false, "overwrite files WITHOUT prompting")
var autoconvert = modFlags.Bool("a", false, "auto-convert any text files")
var junk = modFlags.Bool("j", false, "junk paths (do not make directories)")
var astext = modFlags.Bool("aa", false, "treat ALL files as text")
var useescapes = modFlags.Bool("U", false, "use escapes for all non-ASCII Unicode")
var ignoreunicode = modFlags.Bool("UU", false, "ignore any Unicode fields")
var caseinsensitive = modFlags.Bool("C", false, "match filenames case-insensitively")
var lowercase = modFlags.Bool("L", false, "make (some) names lowercase")
var userpermissions = modFlags.Bool("X", false, "restore UID/GID info")
var vmsversions = modFlags.Bool("V", false, "retain VMS version numbers")
var filepermissions = modFlags.Bool("K", false, "keep setuid/setgid/tacky permissions")
var more = modFlags.Bool("M", false, "pipe through \"more\" pager")
var charsetMS = modFlags.Bool("O", false, "specify a character encoding for DOS, Windows and OS/2 archives")
var charsetNix = modFlags.Bool("I", false, "specify a character encoding for UNIX and other archives")

func main() {
	arguments.Run()

	if *pipe || *list || *freshen || *test || *update || *comment || *timestamp || *exclude || *destination || *never || *quiet || *quieter || *overwrite || *autoconvert || *junk || *astext || *useescapes || *ignoreunicode || *caseinsensitive || *lowercase || *userpermissions || *vmsversions || *filepermissions || *more || *charsetMS || *charsetNix {
		println("Not implemented, yet.")
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		archive.Unzip(flag.Args()[0:flag.NArg()-1], flag.Args()[flag.NArg()-1], arguments.Verbose)
	} else {
		archive.Unzip(flag.Args(), ".", arguments.Verbose)
	}
}
