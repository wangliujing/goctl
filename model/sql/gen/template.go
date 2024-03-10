package gen

import (
	"fmt"

	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util/pathx"
)

const (
	/*	category                              = "model"
		deleteTemplateFile                    = "delete.tpl"
		deleteMethodTemplateFile              = "interface-delete.tpl"
		fieldTemplateFile                     = "field.tpl"
		findOneTemplateFile                   = "find-one.tpl"
		findOneMethodTemplateFile             = "interface-find-one.tpl"
		findOneByFieldTemplateFile            = "find-one-by-field.tpl"
		findOneByFieldMethodTemplateFile      = "interface-find-one-by-field.tpl"
		findOneByFieldExtraMethodTemplateFile = "find-one-by-field-extra-method.tpl"
		importsTemplateFile                   = "import.tpl"
		importsWithNoCacheTemplateFile        = "import-no-cache.tpl"
		insertTemplateFile                    = "insert.tpl"
		insertTemplateMethodFile              = "interface-insert.tpl"
		modelGenTemplateFile                  = "model-gen.tpl"
		modelCustomTemplateFile               = "model.tpl"
		modelNewTemplateFile                  = "model-new.tpl"
		tableNameTemplateFile                 = "table-name.tpl"
		tagTemplateFile                       = "tag.tpl"
		typesTemplateFile                     = "types.tpl"
		updateTemplateFile                    = "update.tpl"
		updateMethodTemplateFile              = "interface-update.tpl"
		varTemplateFile                       = "var.tpl"
		errTemplateFile                       = "err.tpl"*/
	category                     = "model"
	importsTemplateFile          = "import.tpl"
	fieldTemplateFile            = "field.tpl"
	modelGenTemplateFile         = "model-gen.tpl"
	mapperGenTemplateFile        = "mapper-gen.tpl"
	mapperCustomTemplateFile     = "mapper-custom.tpl"
	mapperNewTemplateFile        = "mapper-new.tpl"
	tableNameTemplateFile        = "table-name.tpl"
	tagTemplateFile              = "tag.tpl"
	typesTemplateFile            = "types.tpl"
	saveTemplateFile             = "save.tpl"
	batchSaveTemplateFile        = "batch-save.tpl"
	saveMethodTemplateFile       = "interface-save.tpl"
	batchSaveMethodTemplateFile  = "interface-batch-save.tpl"
	deleteByIdTemplateFile       = "delete-by-id.tpl"
	deleteByIdMethodTemplateFile = "interface-delete-by-id.tpl"

	deleteByIdsTemplateFile       = "delete-by-ids.tpl"
	deleteByIdsMethodTemplateFile = "interface-delete-by-ids.tpl"

	findByIdTemplateFile       = "find-by-id.tpl"
	findByIdMethodTemplateFile = "interface-find-by-id.tpl"

	findByIdWithFieldTemplateFile       = "find-by-id-with-field.tpl"
	findByIdWithFieldMethodTemplateFile = "interface-find-by-id-with-field.tpl"

	findByIdUnscopedTemplateFile       = "find-by-id-unscoped.tpl"
	findByIdUnscopedMethodTemplateFile = "interface-find-by-id-unscoped.tpl"

	findByIdUnscopedWithFieldTemplateFile       = "find-by-id-unscoped-with-field.tpl"
	findByIdUnscopedWithFieldMethodTemplateFile = "interface-find-by-id-unscoped-with-field.tpl"

	findByIdsTemplateFile       = "find-by-ids.tpl"
	findByIdsMethodTemplateFile = "interface-find-by-ids.tpl"

	findByIdsWithFieldTemplateFile       = "find-by-ids-with-field.tpl"
	findByIdsWithFieldMethodTemplateFile = "interface-find-by-ids-with-field.tpl"

	findByIdsUnscopedTemplateFile       = "find-by-ids-unscoped.tpl"
	findByIdsUnscopedMethodTemplateFile = "interface-find-by-ids-unscoped.tpl"

	findByIdsUnscopedWithFieldTemplateFile       = "find-by-ids-unscoped-with-field.tpl"
	findByIdsUnscopedWithFieldMethodTemplateFile = "interface-find-by-ids-unscoped-with-field.tpl"

	listTemplateFile       = "list.tpl"
	listMethodTemplateFile = "interface-list.tpl"

	listUnscopedTemplateFile       = "list-unscoped.tpl"
	listUnscopedMethodTemplateFile = "interface-list-unscoped.tpl"

	listWithFieldTemplateFile       = "list-with-field.tpl"
	listWithFieldMethodTemplateFile = "interface-list-with-field.tpl"

	listUnscopedWithFieldTemplateFile       = "list-unscoped-with-field.tpl"
	listUnscopedWithFieldMethodTemplateFile = "interface-list-unscoped-with-field.tpl"

	countTemplateFile       = "count.tpl"
	countMethodTemplateFile = "interface-count.tpl"

	countUnscopedTemplateFile       = "count-unscoped.tpl"
	countUnscopedMethodTemplateFile = "interface-count-unscoped.tpl"

	pageTemplateFile       = "page.tpl"
	pageMethodTemplateFile = "interface-page.tpl"

	pageUnscopedTemplateFile       = "page-unscoped.tpl"
	pageUnscopedMethodTemplateFile = "interface-page-unscoped.tpl"

	pageWithFieldTemplateFile       = "page-with-field.tpl"
	pageWithFieldMethodTemplateFile = "interface-page-with-field.tpl"

	pageUnscopedWithFieldTemplateFile       = "page-unscoped-with-field.tpl"
	pageUnscopedWithFieldMethodTemplateFile = "interface-page-unscoped-with-field.tpl"
)

var templates = map[string]string{
	fieldTemplateFile:                            template.Field,
	importsTemplateFile:                          template.Imports,
	mapperGenTemplateFile:                        template.MapperGen,
	mapperCustomTemplateFile:                     template.MapperCustom,
	mapperNewTemplateFile:                        template.MapperNew,
	modelGenTemplateFile:                         template.ModelGen,
	tagTemplateFile:                              template.Tag,
	typesTemplateFile:                            template.Types,
	saveTemplateFile:                             template.Save,
	saveMethodTemplateFile:                       template.SaveMethod,
	deleteByIdTemplateFile:                       template.DeleteById,
	deleteByIdMethodTemplateFile:                 template.DeleteByIdMethod,
	deleteByIdsTemplateFile:                      template.DeleteByIds,
	deleteByIdsMethodTemplateFile:                template.DeleteByIdsMethod,
	findByIdTemplateFile:                         template.FindById,
	findByIdMethodTemplateFile:                   template.FindByIdMethod,
	findByIdWithFieldTemplateFile:                template.FindByIdWithField,
	findByIdWithFieldMethodTemplateFile:          template.FindByIdWithFieldMethod,
	findByIdUnscopedTemplateFile:                 template.FindByIdUnscoped,
	findByIdUnscopedMethodTemplateFile:           template.FindByIdUnscopedMethod,
	findByIdUnscopedWithFieldTemplateFile:        template.FindByIdUnscopedWithField,
	findByIdUnscopedWithFieldMethodTemplateFile:  template.FindByIdUnscopedWithFieldMethod,
	findByIdsTemplateFile:                        template.FindByIds,
	findByIdsMethodTemplateFile:                  template.FindByIdsMethod,
	findByIdsWithFieldTemplateFile:               template.FindByIdsWithField,
	findByIdsWithFieldMethodTemplateFile:         template.FindByIdsWithFieldMethod,
	findByIdsUnscopedTemplateFile:                template.FindByIdsUnscoped,
	findByIdsUnscopedMethodTemplateFile:          template.FindByIdsUnscopedMethod,
	findByIdsUnscopedWithFieldTemplateFile:       template.FindByIdsUnscopedWithField,
	findByIdsUnscopedWithFieldMethodTemplateFile: template.FindByIdsUnscopedWithFieldMethod,
	listTemplateFile:                             template.List,
	listMethodTemplateFile:                       template.ListMethod,
	listUnscopedTemplateFile:                     template.ListUnscoped,
	listUnscopedMethodTemplateFile:               template.ListUnscopedMethod,
	listWithFieldTemplateFile:                    template.ListWithField,
	listWithFieldMethodTemplateFile:              template.ListWithFieldMethod,
	listUnscopedWithFieldTemplateFile:            template.ListUnscopedWithField,
	listUnscopedWithFieldMethodTemplateFile:      template.ListUnscopedWithFieldMethod,
	pageTemplateFile:                             template.Page,

	pageMethodTemplateFile:                  template.PageMethod,
	pageUnscopedTemplateFile:                template.PageUnscoped,
	pageUnscopedMethodTemplateFile:          template.PageUnscopedMethod,
	pageWithFieldTemplateFile:               template.PageWithField,
	pageWithFieldMethodTemplateFile:         template.PageWithFieldMethod,
	pageUnscopedWithFieldTemplateFile:       template.PageUnscopedWithField,
	pageUnscopedWithFieldMethodTemplateFile: template.PageUnscopedWithFieldMethod,
}

// Category returns model const value
func Category() string {
	return category
}

// Clean deletes all template files
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates creates template files if not exists
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the deleted template files
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}

	return pathx.CreateTemplate(category, name, content)
}

// Update provides template clean and init
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
