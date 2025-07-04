package auction

import "github.com/ervinmplayon/intercour-face-loggizle/logger"

/*
? This is declared at the package level, it has package scope. It can be used anywhere in this package.
? Now that I think about it, why don't we turn this file into "shared variables" that are usable across
? this package. Im a goddamn genius i know.
*/
var eight_ball_logger logger.Logger = &logger.StandardLogger{}

const (
	RoundRobinStrategy     = "roundrobin"
	FirstResponderStrategy = "firstresponder"
)
