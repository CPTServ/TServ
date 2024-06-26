package push

import (
	"runtime"

	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/CPTServ/TServ/log"
	"github.com/CPTServ/TServ/storage/save"
)

func PushByte(conn *normal.Conn) error {
	// data, err := conn.Si.GetSec()
	// if err != nil {
	// 	log.Error(nil, "Byte data get error: %v", err)
	// 	return err
	// }
	// fmt.Println(data)

	defer runtime.GC()
	defer conn.Close()
	log.Info(nil, "Storage text start")
	err := save.SaveByte(conn)
	log.Info(nil, "Storage text done")
	if err != nil {
		return err
	}
	return nil
}
