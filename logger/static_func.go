package logger

func Debug(logger Logger, desc string, args ...*LogArg) {
	if logger != nil {
		logger.Debug(desc, args...)
	}
}

func Info(logger Logger, desc string, args ...*LogArg) {
	if logger != nil {
		logger.Info(desc, args...)
	}
}

func Error(logger Logger, desc string, err interface{}) {
	if logger != nil {
		logger.Error(desc, ArgAny("err", err))
	}
}
func ErrorWithArg(logger Logger, desc string, err interface{}, args ...*LogArg) {
	if logger != nil {
		if args == nil {
			args = make([]*LogArg, 1)
			args = append(args, ArgAny("err", err))
		}
		logger.Error(desc, args...)
	}
}
func ErrorOnlyArg(logger Logger, desc string, args ...*LogArg) {
	if logger != nil {
		logger.Error(desc, args...)
	}
}

func Fatal(logger Logger, desc string, err interface{}) {
	if logger != nil {
		logger.Fatal(desc, ArgAny("err", err))
	}
}
func FatalWithArg(logger Logger, desc string, err interface{}, args ...*LogArg) {
	if logger != nil {
		if args == nil {
			args = make([]*LogArg, 1)
			args = append(args, ArgAny("err", err))
		}
		logger.Fatal(desc, args...)
	}
}
func FatalOnlyArg(logger Logger, desc string, args ...*LogArg) {
	if logger != nil {
		logger.Fatal(desc, args...)
	}
}

func Panic(logger Logger, desc string, err interface{}) {
	if logger != nil {
		logger.Panic(desc, ArgAny("err", err))
	}
}
func PanicWithArg(logger Logger, desc string, err interface{}, args ...*LogArg) {
	if logger != nil {
		if args == nil {
			args = make([]*LogArg, 1)
			args = append(args, ArgAny("err", err))
		}
		logger.Panic(desc, args...)
	}
}
func PanicOnlyArg(logger Logger, desc string, args ...*LogArg) {
	if logger != nil {
		logger.Panic(desc, args...)
	}
}

func Warn(logger Logger, desc string, err interface{}) {
	if logger != nil {
		logger.Warn(desc, ArgAny("err", err))
	}
}

func WarnWithArg(logger Logger, desc string, err interface{}, args ...*LogArg) {
	if logger != nil {
		if args == nil {
			args = make([]*LogArg, 1)
			args = append(args, ArgAny("err", err))
		}
		logger.Warn(desc, args...)
	}
}
func WarnOnlyArg(logger Logger, desc string, args ...*LogArg) {
	if logger != nil {
		logger.Warn(desc, args...)
	}
}
