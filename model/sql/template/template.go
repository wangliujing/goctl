package template

import (
	_ "embed"
)

//go:embed tpl/field.tpl
var Field string

//go:embed tpl/types.tpl
var Types string

//go:embed tpl/tag.tpl
var Tag string

//go:embed tpl/model-gen.tpl
var ModelGen string

//go:embed tpl/mapper-new.tpl
var MapperNew string

//go:embed tpl/mapper-custom.tpl
var MapperCustom string

//go:embed tpl/mapper-gen.tpl
var MapperGen string

//go:embed tpl/import.tpl
var Imports string

//go:embed tpl/save.tpl
var Save string

//go:embed tpl/batch-save.tpl
var BatchSave string

//go:embed tpl/interface-save.tpl
var SaveMethod string

//go:embed tpl/interface-batch-save.tpl
var BatchSaveMethod string

//go:embed tpl/delete-by-id.tpl
var DeleteById string

//go:embed tpl/interface-delete-by-id.tpl
var DeleteByIdMethod string

//go:embed tpl/delete-by-ids.tpl
var DeleteByIds string

//go:embed tpl/interface-delete-by-ids.tpl
var DeleteByIdsMethod string

//go:embed tpl/find-by-id.tpl
var FindById string

//go:embed tpl/interface-find-by-id.tpl
var FindByIdMethod string

//go:embed tpl/find-by-id-with-field.tpl
var FindByIdWithField string

//go:embed tpl/interface-find-by-id-with-field.tpl
var FindByIdWithFieldMethod string

//go:embed tpl/find-by-id-unscoped.tpl
var FindByIdUnscoped string

//go:embed tpl/interface-find-by-id-unscoped.tpl
var FindByIdUnscopedMethod string

//go:embed tpl/find-by-id-unscoped-with-field.tpl
var FindByIdUnscopedWithField string

//go:embed tpl/interface-find-by-id-unscoped-with-field.tpl
var FindByIdUnscopedWithFieldMethod string

//go:embed tpl/find-by-ids.tpl
var FindByIds string

//go:embed tpl/interface-find-by-ids.tpl
var FindByIdsMethod string

//go:embed tpl/find-by-ids-with-field.tpl
var FindByIdsWithField string

//go:embed tpl/interface-find-by-ids-with-field.tpl
var FindByIdsWithFieldMethod string

//go:embed tpl/find-by-ids-unscoped.tpl
var FindByIdsUnscoped string

//go:embed tpl/interface-find-by-ids-unscoped.tpl
var FindByIdsUnscopedMethod string

//go:embed tpl/find-by-ids-unscoped-with-field.tpl
var FindByIdsUnscopedWithField string

//go:embed tpl/interface-find-by-ids-unscoped-with-field.tpl
var FindByIdsUnscopedWithFieldMethod string

//go:embed tpl/list.tpl
var List string

//go:embed tpl/interface-list.tpl
var ListMethod string

//go:embed tpl/list-unscoped.tpl
var ListUnscoped string

//go:embed tpl/interface-list-unscoped.tpl
var ListUnscopedMethod string

//go:embed tpl/list-with-field.tpl
var ListWithField string

//go:embed tpl/interface-list-with-field.tpl
var ListWithFieldMethod string

//go:embed tpl/list-unscoped-with-field.tpl
var ListUnscopedWithField string

//go:embed tpl/interface-list-unscoped-with-field.tpl
var ListUnscopedWithFieldMethod string

//go:embed tpl/count.tpl
var Count string

//go:embed tpl/count-unscoped.tpl
var CountUnscoped string

//go:embed tpl/interface-count.tpl
var CountMethod string

//go:embed tpl/interface-count-unscoped.tpl
var CountUnscopedMethod string

//go:embed tpl/page.tpl
var Page string

//go:embed tpl/interface-page.tpl
var PageMethod string

//go:embed tpl/page-unscoped.tpl
var PageUnscoped string

//go:embed tpl/interface-page-unscoped.tpl
var PageUnscopedMethod string

//go:embed tpl/page-with-field.tpl
var PageWithField string

//go:embed tpl/interface-page-with-field.tpl
var PageWithFieldMethod string

//go:embed tpl/page-unscoped-with-field.tpl
var PageUnscopedWithField string

//go:embed tpl/interface-page-unscoped-with-field.tpl
var PageUnscopedWithFieldMethod string
