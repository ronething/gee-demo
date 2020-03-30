// author: ashing
// time: 2020/3/30 10:22 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import "github.com/ronething/gee"

func main() {
	g := gee.New()
	g.Static("/asserts", "/tmp/")
	g.Run("127.0.0.1:9998")

}
