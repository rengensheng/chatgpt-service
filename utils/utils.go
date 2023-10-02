package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
	tokenizer "github.com/samber/go-gpt-3-encoder"
)

var encoder *tokenizer.Encoder
var encoderError error

func init() {
	encoder, encoderError = tokenizer.NewEncoder()
	if encoderError != nil {
		fmt.Println("gpt encoder create error", encoderError.Error())
	}
}

func GetUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func CamelToSnake(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func HasSameKey(source []string, target []string) bool {
	for _, v := range source {
		if Includes(target, v) {
			return true
		}
	}
	return false
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func IsExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsExist(err) {
			return false
		}
		return false
	}
	return true
}

func GetFilePath(filepath string) string {
	if IsExists(filepath) {
		return filepath
	}
	err := os.Mkdir(filepath, 0755)
	if err != nil {
		return "./"
	}
	return filepath
}

func RemoveRepeat(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func Includes(v []string, target string) bool {
	for _, val := range v {
		if val == target {
			return true
		}
	}
	return false
}

func Sha256(input string) string {
	hash := sha256.Sum256([]byte(input))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}

func TokenEncoder(input string) ([]int, error) {
	if encoderError != nil {
		return nil, encoderError
	}
	result, err := encoder.Encode(input)
	if err != nil {
		fmt.Println("encode input text error", input, err.Error())
		return nil, err
	}
	return result, nil
}

func GetTokenCount(input string) int {
	if encoderError != nil {
		return len(input)
	}
	result, err := encoder.Encode(input)
	if err != nil {
		fmt.Println("encode input text error", input, err.Error())
		return len(input)
	}
	return len(result)
}
