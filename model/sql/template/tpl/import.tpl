import (
	{{if .containsDeletedAt}}"gorm.io/gorm"{{end}}
    "gorm.io/gorm/logger"
	{{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/wangliujing/foundation-framework/common/dto"
	"github.com/wangliujing/foundation-framework/orm"
    "gorm.io/gorm/clause"
)
