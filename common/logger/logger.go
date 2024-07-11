package logger

import (
	"log"
	"os"
)

/**
* Crea una variable pública que configura un logger personalizado
* @author Cinthya Rondon
*/
var Log = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)