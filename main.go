package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 把字节数读成人类友好的形式，比如 1536 -> "1.5 KiB"。
func humanSize(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div := float64(unit)
	exp := 0
	units := []string{"KiB", "MiB", "GiB", "TiB", "PiB"}
	for b >= int64(div*unit) && exp < len(units)-1 {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %s", float64(b)/div, units[exp])
}

func main() {
	args := os.Args[1:]
	human := true
	var paths []string
	for _, a := range args {
		switch a {
		case "-b", "--bytes":
			human = false
		case "-h", "--help":
			fmt.Println("go-fsize 查看文件大小")
			fmt.Println("用法: go-fsize [-b] <路径> ...")
			fmt.Println("  -b  直接打印字节数，不转可读形式")
			return
		default:
			paths = append(paths, a)
		}
	}
	if len(paths) == 0 {
		fmt.Println("给个文件路径")
		return
	}
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", p, err)
			continue
		}
		if human {
			fmt.Printf("%-12s %s\n", humanSize(info.Size()), filepath.Base(p))
		} else {
			fmt.Printf("%-12d %s\n", info.Size(), filepath.Base(p))
		}
	}
}
