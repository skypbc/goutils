package settings

type debugSettings struct {
	Active      bool
	Name        bool
	File        bool
	PathDepth   int
	Line        bool
	FuncName    bool
	Sep         string
	PathSep     string
	Format      bool
	PreNewLine  *bool
	PostNewLine *bool
}

type config struct {
	Active bool
	Debug  debugSettings
}

var defaultDebugSettings = debugSettings{
	Active:    false,
	Name:      true,
	File:      true,
	PathDepth: -1,
	Line:      true,
	FuncName:  true,
	Sep:       " =>",
	PathSep:   "/",
}

var Print config = config{Active: true, Debug: defaultDebugSettings}
