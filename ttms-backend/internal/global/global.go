/**
 * @Author: lenovo
 * @Description:
 * @File:  global
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:32
 */

package global

import (
	ut "github.com/go-playground/universal-translator"
	"mognolia/internal/model/config"
	"mognolia/internal/pkg/logger"
)

var (
	RootDir  string
	Settings config.AllConfig
	Logger   *logger.Log   // 日志
	Trans    ut.Translator //翻译器
)
