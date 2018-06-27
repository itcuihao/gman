package commands

// commands 常用命令
var commands = map[string]string{
	"ls":         ls,
	"cd":         cd,
	"pwd":        pwd,
	"mkdir":      mkdir,
	"rm":         rm,
	"rmdir":      rmdir,
	"mv":         mv,
	"cp":         cp,
	"touch":      touch,
	"cat":        cat,
	"nl":         nl,
	"more":       more,
	"less":       less,
	"head":       head,
	"tail":       tail,
	"which":      which,
	"whereis":    whereis,
	"locate":     locate,
	"find":       find,
	"exec":       exec,
	"xargs":      xargs,
	"chmod":      chmod,
	"tar":        tar,
	"chgrp":      chgrp,
	"chown":      chown,
	"gzip":       gzip,
	"df":         df,
	"du":         du,
	"ln":         ln,
	"diff":       diff,
	"date":       date,
	"cal":        cal,
	"grep":       grep,
	"wc":         wc,
	"ps":         ps,
	"kill":       kill,
	"killall":    killall,
	"top":        top,
	"free":       free,
	"vmstat":     vmstat,
	"iostat":     iostat,
	"watch":      watch,
	"at":         at,
	"crontab":    crontab,
	"lsof":       lsof,
	"ifconfig":   ifconfig,
	"route":      route,
	"ping":       ping,
	"traceroute": traceroute,
	"netstat":    netstat,
	"ss":         ss,
	"telnet":     telnet,
	"rpc":        rpc,
	"scp":        scp,
	"wget":       wget,
}

// GetCommand get command
func GetCommand(c string) string {
	v, ok := commands[c]
	if !ok {
		return ""
	}

	return v
}
