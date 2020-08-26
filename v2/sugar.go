package testversion

import "fmt"

func Trace(args ...interface{}) {
}

func Tracef(format string, args ...interface{}) {
}

func Traceln(args ...interface{}) {
}

func Debug(args ...interface{}) {
	jloggerInstance.debug(args...)
}

func Debugf(format string, args ...interface{}) {
	jloggerInstance.debugf(format, args...)
}

func Debugln(args ...interface{}) {
	jloggerInstance.debugln(args...)
}

func Info(args ...interface{}) {
	jloggerInstance.info(args...)
}

func Infof(format string, args ...interface{}) {
	jloggerInstance.infof(format, args...)
}

func Infoln(args ...interface{}) {
	jloggerInstance.infoln(args...)
}

func Warning(args ...interface{}) {
	jloggerInstance.warning(args...)
}

func Warningf(format string, args ...interface{}) {
	jloggerInstance.warningf(format, args...)
}

func Warningln(args ...interface{}) {
	jloggerInstance.warningln(args...)
}

func Error(args ...interface{}) {
	jloggerInstance.error(args...)
}

func Errorf(format string, args ...interface{}) {
	jloggerInstance.errorf(format, args...)
}

func Errorln(args ...interface{}) {
	jloggerInstance.errorln(args...)
}

func Fatal(args ...interface{}) {
	jloggerInstance.fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	jloggerInstance.fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	jloggerInstance.fatalln(args...)
}

//func (j *jlogger) trace(args ...interface{}) {
//	j.sugar.d
//}
//
//func (j *jlogger) tracef(format string, args ...interface{}) {
//}
//
//func (j *jlogger) traceln(args ...interface{}) {
//}

func (j *jlogger) debug(args ...interface{}) {
	j.sugar.Debug(args...)
}

func (j *jlogger) debugf(format string, args ...interface{}) {
	j.sugar.Debugf(format, args...)
}

func (j *jlogger) debugln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	j.sugar.Debug("%s", msg)
}

func (j *jlogger) info(args ...interface{}) {
	j.sugar.Info(args...)
}

func (j *jlogger) infof(format string, args ...interface{}) {
	j.sugar.Infof(format, args...)
}

func (j *jlogger) infoln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	j.sugar.Info("%s", msg)
}

func (j *jlogger) warning(args ...interface{}) {
	j.sugar.Warn(args...)
}

func (j *jlogger) warningf(format string, args ...interface{}) {
	j.sugar.Warnf(format, args...)
}

func (j *jlogger) warningln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	j.sugar.Warn("%s", msg)
}

func (j *jlogger) error(args ...interface{}) {
	j.sugar.Error(args...)
}

func (j *jlogger) errorf(format string, args ...interface{}) {
	j.sugar.Errorf(format, args...)
}

func (j *jlogger) errorln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	j.sugar.Error("%s", msg)
}

func (j *jlogger) fatal(args ...interface{}) {
	j.sugar.Fatal(args...)
}

func (j *jlogger) fatalf(format string, args ...interface{}) {
	j.sugar.Fatalf(format, args...)
}

func (j *jlogger) fatalln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	j.sugar.Fatal("%s", msg)
}
