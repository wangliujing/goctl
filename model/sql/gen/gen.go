package gen

import (
	"bytes"
	"fmt"
	modelutil "github.com/wangliujing/goctl/model/sql/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/wangliujing/goctl/config"
	"github.com/wangliujing/goctl/model/sql/model"
	"github.com/wangliujing/goctl/model/sql/parser"
	"github.com/wangliujing/goctl/model/sql/template"
	"github.com/wangliujing/goctl/util"
	"github.com/wangliujing/goctl/util/console"
	"github.com/wangliujing/goctl/util/format"
	"github.com/wangliujing/goctl/util/pathx"
	"github.com/wangliujing/goctl/util/stringx"
)

const pwd = "."

type (
	defaultGenerator struct {
		console.Console
		// source string
		dir           string
		pkg           string
		cfg           *config.Config
		isPostgreSql  bool
		ignoreColumns []string
	}

	// Option defines a function with argument defaultGenerator
	Option func(generator *defaultGenerator)

	code struct {
		importsCode string
		typesCode   string
		newCode     string
		tableName   string

		saveCode                       string
		batchSaveCode                  string
		deleteByIdCode                 string
		deleteByIdsCode                string
		findByIdCode                   string
		findByIdWithFieldCode          string
		findByIdUnscopedCode           string
		findByIdUnscopedWithFieldCode  string
		findByIdsCode                  string
		findByIdsWithFieldCode         string
		findByIdsUnscopedCode          string
		findByIdsUnscopedWithFieldCode string
		listCode                       string
		listUnscopedCode               string
		listWithFieldCode              string
		listUnscopedWithFieldCode      string
		countCode                      string
		countUnscopedCode              string
		pageCode                       string
		pageUnscopedCode               string
		pageWithFieldCode              string
		pageUnscopedWithFieldCode      string
	}

	codeTuple struct {
		modelCode        string
		mapperCode       string
		mapperCustomCode string
	}
)

// NewDefaultGenerator creates an instance for defaultGenerator
func NewDefaultGenerator(dir string, cfg *config.Config, opt ...Option) (*defaultGenerator, error) {
	if dir == "" {
		dir = pwd
	}
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	pkg := util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}

	generator := &defaultGenerator{dir: dir, cfg: cfg, pkg: pkg}
	var optionList []Option
	optionList = append(optionList, newDefaultOption())
	optionList = append(optionList, opt...)
	for _, fn := range optionList {
		fn(generator)
	}

	return generator, nil
}

// WithConsoleOption creates a console option.
func WithConsoleOption(c console.Console) Option {
	return func(generator *defaultGenerator) {
		generator.Console = c
	}
}

// WithIgnoreColumns ignores the columns while insert or update rows.
func WithIgnoreColumns(ignoreColumns []string) Option {
	return func(generator *defaultGenerator) {
		generator.ignoreColumns = ignoreColumns
	}
}

// WithPostgreSql marks  defaultGenerator.isPostgreSql true.
func WithPostgreSql() Option {
	return func(generator *defaultGenerator) {
		generator.isPostgreSql = true
	}
}

func newDefaultOption() Option {
	return func(generator *defaultGenerator) {
		generator.Console = console.NewColorConsole()
	}
}

func (g *defaultGenerator) StartFromDDL(filename string, withCache, strict bool, database string) error {
	modelList, err := g.genFromDDL(filename, withCache, strict, database)
	if err != nil {
		return err
	}

	return g.createFile(modelList)
}

func (g *defaultGenerator) StartFromInformationSchema(tables map[string]*model.Table, withCache, strict bool) error {
	m := make(map[string]*codeTuple)
	for _, each := range tables {
		table, err := parser.ConvertDataType(each, strict)
		if err != nil {
			return err
		}

		modelCode, err := g.genModel(*table)
		if err != nil {
			return err
		}

		mapperCode, err := g.genMapper(*table, withCache)
		if err != nil {
			return err
		}
		customCode, err := g.genMapperCustom(*table, withCache)
		if err != nil {
			return err
		}

		m[table.Name.Source()] = &codeTuple{
			modelCode:        modelCode,
			mapperCode:       mapperCode,
			mapperCustomCode: customCode,
		}
	}

	return g.createFile(m)
}

func (g *defaultGenerator) createFile(modelList map[string]*codeTuple) error {
	dirAbs, err := filepath.Abs(g.dir)
	if err != nil {
		return err
	}

	g.dir = dirAbs
	g.pkg = util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}
	err = pathx.MkdirIfNotExist(dirAbs + "/model")
	if err != nil {
		return err
	}

	for tableName, codes := range modelList {
		tn := stringx.From(tableName)

		modelFilename, err := format.FileNamingFormat(g.cfg.NamingFormat,
			fmt.Sprintf("%s_model", tn.Source()))
		if err != nil {
			return err
		}
		name := util.SafeString(modelFilename) + ".go"
		filename := filepath.Join(dirAbs+"/model", name)
		err = os.WriteFile(filename, []byte(codes.modelCode), os.ModePerm)
		if err != nil {
			return err
		}

		mapperFilename, err := format.FileNamingFormat(g.cfg.NamingFormat,
			fmt.Sprintf("%s_mapper", tn.Source()))
		if err != nil {
			return err
		}

		name = util.SafeString(mapperFilename) + ".go"
		filename = filepath.Join(dirAbs, name)
		err = os.WriteFile(filename, []byte(codes.mapperCode), os.ModePerm)
		if err != nil {
			return err
		}

		name = util.SafeString(mapperFilename) + "_custom.go"
		filename = filepath.Join(dirAbs, name)
		if pathx.FileExists(filename) {
			g.Warning("%s already exists, ignored.", name)
			continue
		}
		err = os.WriteFile(filename, []byte(codes.mapperCustomCode), os.ModePerm)
		if err != nil {
			return err
		}
	}
	g.Success("Done.")
	return nil
}

// ret1: key-table name,value-code
func (g *defaultGenerator) genFromDDL(filename string, withCache, strict bool, database string) (
	map[string]*codeTuple, error,
) {
	m := make(map[string]*codeTuple)
	tables, err := parser.Parse(filename, database, strict)
	if err != nil {
		return nil, err
	}

	for _, e := range tables {
		modelCode, err := g.genModel(*e)
		if err != nil {
			return nil, err
		}
		mapperCode, err := g.genMapper(*e, withCache)
		if err != nil {
			return nil, err
		}
		mapperCustomCode, err := g.genMapperCustom(*e, withCache)
		if err != nil {
			return nil, err
		}

		m[e.Name.Source()] = &codeTuple{
			modelCode:        modelCode,
			mapperCode:       mapperCode,
			mapperCustomCode: mapperCustomCode,
		}
	}

	return m, nil
}

// Table defines mysql table
type Table struct {
	parser.Table
	PrimaryCacheKey        Key
	UniqueCacheKey         []Key
	ContainsUniqueCacheKey bool
	ignoreColumns          []string
}

func (t Table) isIgnoreColumns(columnName string) bool {
	for _, v := range t.ignoreColumns {
		if v == columnName {
			return true
		}
	}
	return false
}

func (g *defaultGenerator) genModel(in parser.Table) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}
	primaryKey, uniqueKey := genCacheKeys(in)
	var table Table
	table.Table = in
	table.PrimaryCacheKey = primaryKey
	table.UniqueCacheKey = uniqueKey
	table.ContainsUniqueCacheKey = len(uniqueKey) > 0
	table.ignoreColumns = g.ignoreColumns
	fieldsString, err := genFields(table, table.Fields)
	if err != nil {
		return "", err
	}
	text, err := pathx.LoadTemplate(category, modelGenTemplateFile, template.ModelGen)
	if err != nil {
		return "", err
	}

	t := util.With("model-gen").Parse(text).GoFmt(true)
	output, err := t.Execute(map[string]any{
		"time":                  in.ContainsTime(),
		"sql":                   in.ContainsSql(),
		"tableName":             in.Name.Source(),
		"upperStartCamelObject": in.Name.ToCamel(),
		"fields":                fieldsString,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func (g *defaultGenerator) genMapper(in parser.Table, withCache bool) (string, error) {

	primaryKey, uniqueKey := genCacheKeys(in)

	var table Table
	table.Table = in
	table.PrimaryCacheKey = primaryKey
	table.UniqueCacheKey = uniqueKey
	table.ContainsUniqueCacheKey = len(uniqueKey) > 0
	table.ignoreColumns = g.ignoreColumns

	importsCode, err := genImports(table)
	if err != nil {
		return "", err
	}

	saveCode, saveCodeMethod, err := genSave(table)
	if err != nil {
		return "", err
	}

	batchSaveCode, batchSaveCodeMethod, err := genBatchSave(table)
	if err != nil {
		return "", err
	}

	deleteByIdCode, deleteByIdCodeMethod, err := genDeleteById(table)
	if err != nil {
		return "", err
	}

	deleteByIdsCode, deleteByIdsCodeMethod, err := genDeleteByIds(table)
	if err != nil {
		return "", err
	}

	findByIdCode, findByIdCodeMethod, err := genFindById(table)
	if err != nil {
		return "", err
	}

	findByIdWithFieldCode, findByIdWithFieldCodeMethod, err := genFindByIdWithField(table)
	if err != nil {
		return "", err
	}

	findByIdUnscopedCode, findByIdUnscopedCodeMethod, err := genFindByIdUnscoped(table)
	if err != nil {
		return "", err
	}

	findByIdUnscopedWithFieldCode, findByIdUnscopedWithFieldCodeMethod, err := genFindByIdUnscopedWithField(table)
	if err != nil {
		return "", err
	}

	findByIdsCode, findByIdsCodeMethod, err := genFindByIds(table)
	if err != nil {
		return "", err
	}

	findByIdsWithFieldCode, findByIdsWithFieldCodeMethod, err := genFindByIdsWithField(table)
	if err != nil {
		return "", err
	}

	findByIdsUnscopedCode, findByIdsUnscopedCodeMethod, err := genFindByIdsUnscoped(table)
	if err != nil {
		return "", err
	}

	findByIdsUnscopedWithFieldCode, findByIdsUnscopedWithFieldCodeMethod, err := genFindByIdsUnscopedWithField(table)
	if err != nil {
		return "", err
	}

	listCode, listCodeMethod, err := genList(table)
	if err != nil {
		return "", err
	}

	listUnscopedCode, listUnscopedCodeMethod, err := genListUnscoped(table)
	if err != nil {
		return "", err
	}

	listWithFieldCode, listWithFieldCodeMethod, err := genListWithField(table)
	if err != nil {
		return "", err
	}

	listUnscopedWithFieldCode, listUnscopedWithFieldCodeMethod, err := genListUnscopedWithField(table)
	if err != nil {
		return "", err
	}

	countCode, countCodeMethod, err := genCount(table)
	if err != nil {
		return "", err
	}

	countUnscopedCode, countUnscopedCodeMethod, err := genCountUnscoped(table)
	if err != nil {
		return "", err
	}

	pageCode, pageCodeMethod, err := genPage(table)
	if err != nil {
		return "", err
	}

	pageUnscopedCode, pageUnscopedCodeMethod, err := genPageUnscoped(table)
	if err != nil {
		return "", err
	}

	pageWithFieldCode, pageWithFieldCodeMethod, err := genPageWithField(table)
	if err != nil {
		return "", err
	}

	pageUnscopedWithFieldCode, pageUnscopedWithFieldCodeMethod, err := genPageUnscopedWithField(table)
	if err != nil {
		return "", err
	}

	var list []string
	list = append(list, saveCodeMethod, batchSaveCodeMethod, deleteByIdCodeMethod, deleteByIdsCodeMethod,
		findByIdCodeMethod, findByIdWithFieldCodeMethod, findByIdUnscopedCodeMethod, findByIdUnscopedWithFieldCodeMethod,
		findByIdsCodeMethod, findByIdsWithFieldCodeMethod, findByIdsUnscopedCodeMethod, findByIdsUnscopedWithFieldCodeMethod,
		listCodeMethod, listUnscopedCodeMethod, listWithFieldCodeMethod, listUnscopedWithFieldCodeMethod, countCodeMethod,
		countUnscopedCodeMethod, pageCodeMethod, pageUnscopedCodeMethod, pageWithFieldCodeMethod, pageUnscopedWithFieldCodeMethod)

	typesCode, err := genTypes(table, strings.Join(modelutil.TrimStringSlice(list), pathx.NL), withCache)
	if err != nil {
		return "", err
	}

	newCode, err := genNew(table, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	tableName, err := genTableName(table)
	if err != nil {
		return "", err
	}

	code := &code{
		importsCode:                    importsCode,
		typesCode:                      typesCode,
		newCode:                        newCode,
		tableName:                      tableName,
		saveCode:                       saveCode,
		batchSaveCode:                  batchSaveCode,
		deleteByIdCode:                 deleteByIdCode,
		deleteByIdsCode:                deleteByIdsCode,
		findByIdCode:                   findByIdCode,
		findByIdWithFieldCode:          findByIdWithFieldCode,
		findByIdUnscopedCode:           findByIdUnscopedCode,
		findByIdUnscopedWithFieldCode:  findByIdUnscopedWithFieldCode,
		findByIdsCode:                  findByIdsCode,
		findByIdsWithFieldCode:         findByIdsWithFieldCode,
		findByIdsUnscopedCode:          findByIdsUnscopedCode,
		findByIdsUnscopedWithFieldCode: findByIdsUnscopedWithFieldCode,
		listCode:                       listCode,
		listUnscopedCode:               listUnscopedCode,
		listWithFieldCode:              listWithFieldCode,
		listUnscopedWithFieldCode:      listUnscopedWithFieldCode,
		countCode:                      countCode,
		countUnscopedCode:              countUnscopedCode,
		pageCode:                       pageCode,
		pageUnscopedCode:               pageUnscopedCode,
		pageWithFieldCode:              pageWithFieldCode,
		pageUnscopedWithFieldCode:      pageUnscopedWithFieldCode,
	}

	output, err := g.executeMapper(table, code)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (g *defaultGenerator) genMapperCustom(in parser.Table, withCache bool) (string, error) {
	text, err := pathx.LoadTemplate(category, mapperCustomTemplateFile, template.MapperCustom)
	if err != nil {
		return "", err
	}

	t := util.With("mapper-custom").
		Parse(text).
		GoFmt(true)
	output, err := t.Execute(map[string]any{
		"pkg":                   g.pkg,
		"withCache":             withCache,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (g *defaultGenerator) executeMapper(table Table, code *code) (*bytes.Buffer, error) {
	text, err := pathx.LoadTemplate(category, mapperGenTemplateFile, template.MapperGen)
	if err != nil {
		return nil, err
	}
	t := util.With("mapper").
		Parse(text).
		GoFmt(true)
	output, err := t.Execute(map[string]any{
		"pkg":                        g.pkg,
		"imports":                    code.importsCode,
		"types":                      code.typesCode,
		"new":                        code.newCode,
		"tableName":                  code.tableName,
		"data":                       table,
		"save":                       code.saveCode,
		"batchSave":                  code.batchSaveCode,
		"deleteById":                 code.deleteByIdCode,
		"deleteByIds":                code.deleteByIdsCode,
		"findById":                   code.findByIdCode,
		"findByIdWithField":          code.findByIdWithFieldCode,
		"findByIdUnscoped":           code.findByIdUnscopedCode,
		"findByIdUnscopedWithField":  code.findByIdUnscopedWithFieldCode,
		"findByIds":                  code.findByIdsCode,
		"findByIdsWithField":         code.findByIdsWithFieldCode,
		"findByIdsUnscoped":          code.findByIdsUnscopedCode,
		"findByIdsUnscopedWithField": code.findByIdsUnscopedWithFieldCode,
		"list":                       code.listCode,
		"listUnscoped":               code.listUnscopedCode,
		"listWithField":              code.listWithFieldCode,
		"listUnscopedWithField":      code.listUnscopedWithFieldCode,
		"count":                      code.countCode,
		"countUnscoped":              code.countUnscopedCode,
		"page":                       code.pageCode,
		"pageUnscoped":               code.pageUnscopedCode,
		"pageWithField":              code.pageWithFieldCode,
		"pageUnscopedWithField":      code.pageUnscopedWithFieldCode,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}

func wrapWithRawString(v string, postgreSql bool) string {
	if postgreSql {
		return v
	}

	if v == "`" {
		return v
	}

	if !strings.HasPrefix(v, "`") {
		v = "`" + v
	}

	if !strings.HasSuffix(v, "`") {
		v = v + "`"
	} else if len(v) == 1 {
		v = v + "`"
	}

	return v
}
