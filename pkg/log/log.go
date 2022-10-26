/*
Copyright © 2022 uscan team

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package log

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[uscan] INF ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLogger = log.New(os.Stdout, "[uscan] ERR ", log.Ldate|log.Ltime|log.Lmsgprefix)
	fatalLogger = log.New(os.Stdout, "[uscan] FTL ", log.Ldate|log.Ltime|log.Lmsgprefix)
)

func Info(msg ...any) {
	infoLogger.Println(msg...)
}

func Infof(format string, msg ...any) {
	infoLogger.Printf(format, msg...)
}
func Error(msg ...any) {
	errorLogger.Println(msg...)
}

func Errorf(format string, msg ...any) {
	errorLogger.Printf(format, msg...)
}

func Fatal(msg ...any) {
	fatalLogger.Fatalln(msg...)
}

func Fatalf(format string, msg ...any) {
	fatalLogger.Fatalf(format, msg...)
}
