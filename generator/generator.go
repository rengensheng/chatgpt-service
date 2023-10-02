package generator

import (
	"os"
	"strings"
	"text/template"

	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
	"github.com/hyahm/golog"
)

// 代码生成器

type TableInfo struct {
	TableName      string
	LowerTableName string // 首字母小写的表名称
	Fields         []Field
}

type Field struct {
	StructName   string // 结构体名称
	FieldType    string // 字段类型
	OriginalType string // 原始类型
	SourceName   string // 字段名称
	Required     bool   // 是否必须
	Conditions   string // 条件
	OrmInfo      string // x-orm同步时所需信息
	Comment      string // 表注释
}

var TypeMap = map[string]string{
	"varchar":  "string",
	"int":      "int64",
	"datetime": "string",
	"bigint":   "int64",
}

func GenerateByDatabase() (string, error) {
	engine := database.GetXOrmEngine()
	results, err := engine.SQL("show tables").QueryString()
	var routes []string
	if err != nil {
		golog.Info(err.Error())
		return "", err
	}
	engineTableInfo := database.GetTableInfoEngine()
	for _, tableInfoName := range results {
		tableName := tableInfoName["Tables_in_lowcode"]
		if tableName == "" {
			continue
		}
		firstLowerTableName := strings.ToLower(string(tableName[0])) + tableName[1:]
		snakeFileName := utils.CamelToSnake(tableName)
		modelFilePath := "./models/" + snakeFileName + "_entity.go"
		serviceFilePath := "./services/" + snakeFileName + "_service.go"
		routeFilePath := "./routes/" + snakeFileName + "_route.go"
		if utils.IsExists(modelFilePath) || utils.IsExists(serviceFilePath) || utils.IsExists(routeFilePath) {
			continue
		}
		fields, err := engineTableInfo.SQL("select * from columns where table_name = ?", tableName).QueryString()
		if err != nil {
			golog.Info(err.Error())
			continue
		}
		TableInfos := &TableInfo{
			TableName:      utils.UcFirst(tableName),
			LowerTableName: firstLowerTableName,
		}
		for _, field := range fields {
			fieldName := field["COLUMN_NAME"]
			fieldType, ok := TypeMap[field["DATA_TYPE"]]
			if !ok {
				fieldType = "string"
			}
			isRequired := false
			conditions := ""

			// 设置字段的相关信息
			ormInfo := field["COLUMN_TYPE"]

			// 设置主键
			isPk := field["COLUMN_KEY"]
			if isPk == "PRI" {
				ormInfo = ormInfo + " pk"
			}

			isNullable := field["IS_NULLABLE"]
			if isNullable == "NO" {
				ormInfo = ormInfo + " notnull"
			}

			// 设置json的非空条件
			if isNullable == "NO" && isPk != "PRI" {
				isRequired = true
				conditions = `required`
			}

			filedMaxLength := field["CHARACTER_MAXIMUM_LENGTH"]
			if filedMaxLength != "" {
				// 设置字段长度
				if conditions == "required" {
					conditions = conditions + ","
				}
				conditions = conditions + "max=" + filedMaxLength
			}

			if conditions != "" {
				conditions = " binding:\"" + conditions + "\""
			}

			fieldComment := field["COLUMN_COMMENT"]

			if fieldComment != "" {
				fieldComment = "  // " + fieldComment
			}

			fieldStruct := Field{
				StructName:   utils.UcFirst(utils.Case2Camel(fieldName)),
				SourceName:   fieldName,
				FieldType:    fieldType,
				Required:     isRequired,
				Conditions:   conditions,
				OrmInfo:      ormInfo,
				Comment:      fieldComment,
				OriginalType: field["COLUMN_TYPE"],
			}
			TableInfos.Fields = append(TableInfos.Fields, fieldStruct)
		}

		modelFileTmpl, err := os.ReadFile("./template/model.tmpl")
		if err != nil {
			golog.Info("model模板文件打开失败", err.Error())
		}
		serviceFileTmpl, err := os.ReadFile("./template/service.tmpl")
		if err != nil {
			golog.Info("service模板文件打开失败", err.Error())
		}
		routeFileTmpl, err := os.ReadFile("./template/route.tmpl")
		if err != nil {
			golog.Info("route模板文件打开失败", err.Error())
		}
		modelOutputFile, err := os.OpenFile(modelFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			golog.Info("无法成功打开文件", modelFilePath, err.Error())
		}
		serviceOutputFile, err := os.OpenFile(serviceFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			golog.Info("无法成功打开文件", serviceFilePath, err.Error())
		}
		routeOutputFile, err := os.OpenFile(routeFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			golog.Info("无法成功打开文件", routeFilePath, err.Error())
		}
		if err != nil {
			golog.Info(err.Error())
			return "", err
		}

		tpl, err := template.New("model").Parse(string(modelFileTmpl))
		if err != nil {
			golog.Info(err.Error())
		}
		err = tpl.Execute(modelOutputFile, TableInfos)
		if err != nil {
			golog.Info(err.Error())
		}

		tplService, err := template.New("service").Parse(string(serviceFileTmpl))
		if err != nil {
			golog.Info(err.Error())
		}
		err = tplService.Execute(serviceOutputFile, TableInfos)
		if err != nil {
			golog.Info(err.Error())
		}

		tplRoute, err := template.New("route").Parse(string(routeFileTmpl))
		if err != nil {
			golog.Info(err.Error())
		}
		err = tplRoute.Execute(routeOutputFile, TableInfos)

		if err != nil {
			golog.Info(err.Error())
		}
		routes = append(routes, TableInfos.TableName+"RouterRegistry(engine)")
		modelOutputFile.Close()
		serviceOutputFile.Close()
		routeOutputFile.Close()
	}
	return strings.Join(routes, ";"), nil
}
