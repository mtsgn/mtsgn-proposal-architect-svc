package glog

//var (
//	ZapLog *zap.SugaredLogger // 简易版日志文件
//	//logger *zap.Logger        // 这个日志强大一些, 目前还用不到
//
//	logLevel = zap.NewAtomicLevel()
//)
//
//func getLogWriter(logPath, logType string) zapcore.WriteSyncer {
//	if logType == "file" {
//		w := zapcore.AddSync(&lumberjack.Logger{
//			Filename:   logPath,
//			MaxSize:    512, // MB
//			LocalTime:  true,
//			Compress:   true,
//			MaxBackups: 8, // 最多保留 n 个备份
//		})
//		return w
//	} else {
//		return zapcore.Lock(os.Stdout) //标准输出
//	}
//}
//
//func getEncoder(logType string) zapcore.Encoder {
//	var encoder zapcore.Encoder
//	if logType == "file" {
//		c := zap.NewProductionEncoderConfig()
//		c.EncodeTime = zapcore.ISO8601TimeEncoder
//		encoder = zapcore.NewJSONEncoder(c)
//	} else {
//		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
//	}
//	return encoder
//}
//
//// InitLog 初始化日志文件
//func InitLog(logConfLevel, logPath, logType string) zapcore.WriteSyncer {
//	loglevel := zapcore.InfoLevel
//	switch logConfLevel {
//	case "INFO":
//		loglevel = zapcore.InfoLevel
//	case "ERROR":
//		loglevel = zapcore.ErrorLevel
//	}
//	setLevel(loglevel)
//	if "" == strings.TrimSpace(logPath) {
//		logPath = "./logs/boilerplate-api.log"
//	}
//	encoder := getEncoder(logType)           //获取编码方式
//	writer := getLogWriter(logPath, logType) //获取writer
//	core := zapcore.NewCore(encoder, writer, loglevel)
//	ZapLog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
//	return writer
//}
//
//func setLevel(level zapcore.Level) {
//	logLevel.SetLevel(level)
//}
//
//func Info(args ...interface{}) {
//	ZapLog.Named(funcName()).Info(args...)
//}
//
//func Infof(template string, args ...interface{}) {
//	ZapLog.Named(funcName()).Infof(template, args...)
//}
//
//func Warn(args ...interface{}) {
//	ZapLog.Named(funcName()).Warn(args...)
//}
//
//func Warnf(template string, args ...interface{}) {
//	ZapLog.Named(funcName()).Warnf(template, args...)
//}
//
//func Error(args ...interface{}) {
//	ZapLog.Named(funcName()).Error(args...)
//}
//
//func Errorf(template string, args ...interface{}) {
//	ZapLog.Named(funcName()).Errorf(template, args...)
//}
//
//func funcName() string {
//	pc, _, _, _ := runtime.Caller(2)
//	funcName := runtime.FuncForPC(pc).Name()
//	return fmt.Sprintf("%d_%s", time.Now().UnixNano(), path.Base(funcName))
//}
//
//func FileLoggerMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//
//		// 執行處理流程
//		c.Next()
//
//		// 請求結束後記錄 log
//		duration := time.Since(start)
//		log := fmt.Sprintf("[%s] %d %s %s  IP: %s  Time: %v",
//			time.Now().Format("2006-01-02 15:04:05"),
//			c.Writer.Status(),
//			c.Request.Method,
//			c.Request.URL.Path,
//			c.ClientIP(),
//			duration)
//		ZapLog.Info(log)
//	}
//}
