package mysql

import (
	"strings"

	"github.com/odigos-io/odigos/common"
	"github.com/odigos-io/odigos/procdiscovery/pkg/process"
)

// This is an experimental feature, It is not a language
// but in order to avoid huge refactoring we are adding it as a language for now
type MySQLInspector struct{}

const MySQLProcessName = "mysqld"

func (j *MySQLInspector) Inspect(p *process.Details) (common.ProgramLanguageDetails, bool) {
	var programLanguageDetails common.ProgramLanguageDetails
	if strings.HasSuffix(p.ExeName, MySQLProcessName) || strings.HasSuffix(p.CmdLine, MySQLProcessName) {
		programLanguageDetails.Language = common.MySQLProgrammingLanguage
		return programLanguageDetails, true
	}

	return programLanguageDetails, false
}
