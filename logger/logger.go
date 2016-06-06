package logger

import (
	"os"

	"github.com/fatih/color"
	"github.com/kataras/iris/config"
	"github.com/mattn/go-colorable"
)

var (
	// Prefix is the prefix for the logger, default is [IRIS]
	Prefix = "[IRIS] "
	// bannersRan keeps track of the logger's print banner count
	bannersRan = 0
)

// Logger the logger
type Logger struct {
	config    config.Logger
	underline *color.Color
}

// attr takes a color integer and converts it to color.Attribute
func attr(sgr int) color.Attribute {
	return color.Attribute(sgr)
}

// New creates a new Logger from config.Logger configuration
func New(c config.Logger) *Logger {
	color.Output = colorable.NewColorable(c.Out)

	l := &Logger{c, color.New(attr(c.ColorBgDefault), attr(c.ColorFgDefault), color.Bold)}
	return l
}

// PrintBanner prints a text (banner) with BannerFgColor, BannerBgColor and a success message at the end
// It doesn't cares if the logger is disabled or not, it will print this
func (l *Logger) PrintBanner(banner string, sucessMessage string) {
	c := color.New(attr(l.config.ColorBgDefault), attr(l.config.ColorFgBanner), color.Bold)
	c.Println(banner)
	bannersRan++

	if sucessMessage != "" {
		c.Add(attr(l.config.ColorBgSuccess), attr(l.config.ColorFgSuccess), color.Bold)

		if bannersRan > 1 {
			c.Printf("Server[%#v]\n", bannersRan)

		}
		c.Println(sucessMessage)
	}

	c.DisableColor()
	c = nil
}

// ResetColors sets the colors to the default
// this func is called every time a success, info, warning, or danger message is printed
func (l *Logger) ResetColors() {
	l.underline.Add(attr(l.config.ColorBgDefault), attr(l.config.ColorFgBanner), color.Bold)
}

// SetEnable true enables, false disables the Logger
func (l *Logger) SetEnable(enable bool) {
	l.config.Disabled = !enable
}

// IsEnabled returns true if Logger is enabled, otherwise false
func (l *Logger) IsEnabled() bool {
	return !l.config.Disabled
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	if !l.config.Disabled {
		l.underline.Print(v...)
	}
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Printf(format, a...)
	}
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Println(a...)
	}
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Print(a...)

	}
	os.Exit(1)

}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Printf(format, a...)
	}

	os.Exit(1)

}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (l *Logger) Fatalln(a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Println(a...)
	}

	os.Exit(1)

}

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *Logger) Panic(a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Print(a...)
	}

	panic("")
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *Logger) Panicf(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Printf(format, a...)
	}
	panic("")
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *Logger) Panicln(a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Println(a...)
	}
	panic("")
}

// Sucessf calls l.Output to print to the logger with the Success colors.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Sucessf(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Add(attr(l.config.ColorBgSuccess), attr(l.config.ColorFgSuccess))
		l.underline.Printf(format, a...)
		l.ResetColors()
	}
}

// Infof calls l.Output to print to the logger with the Info colors.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Add(attr(l.config.ColorBgInfo), attr(l.config.ColorFgInfo))
		l.underline.Printf(format, a...)
		l.ResetColors()
	}
}

// Warningf calls l.Output to print to the logger with the Warning colors.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Add(attr(l.config.ColorBgWarning), attr(l.config.ColorFgWarning))
		l.underline.Printf(format, a...)
		l.ResetColors()
	}
}

// Dangerf calls l.Output to print to the logger with the Danger colors.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Dangerf(format string, a ...interface{}) {
	if !l.config.Disabled {
		l.underline.Add(attr(l.config.ColorBgDanger), attr(l.config.ColorFgDanger))
		l.underline.Printf(format, a...)
		l.ResetColors()
	}
}
