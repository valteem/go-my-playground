// https://stackoverflow.com/questions/76799052/golang-defer-is-not-always-executed

package reuse

import (
	"context"
	"fmt"
	"io"
)

func ReadFromConn(conn io.Reader, ctx context.Context, cancel context.CancelCauseFunc) {

	defer cancel(fmt.Errorf("closing reader"))

	tmp := make([]byte, 256)
	for {
		_, err := conn.Read(tmp)
		if err != nil {
			fmt.Println("error reading from connection")
			return
		}
		fmt.Println(tmp)
		select {
			case <- ctx.Done():
				return
			default:
				fmt.Println("still listening ...")
		}
	}

}