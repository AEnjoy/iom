package hashtool

import (
	"IOM/server/utils/debugTools"
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	debugTools.PrintLogsOnlyInDebugMode("hash str is ", str, ",and it MD5", md5str)
	return md5str
}
