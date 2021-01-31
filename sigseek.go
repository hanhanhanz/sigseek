package main
//import "regexp"
import "fmt"
import "os"
import "path/filepath"
import "bufio"
import "strings"
import "strconv"
import "flag"
import "os/exec"
import "sync"
//import "time"


func filetoslice(sfile string) []string {
	var seeds = []string{}
	var g *os.File
	var g2 *bufio.Scanner
	g,_ = os.Open(sfile) 
    g2 = bufio.NewScanner(g)


	for g2.Scan() {
		var line = g2.Text()
		seeds = append(seeds,line)
	}

	return seeds
		
}

func openandfind(path string, seed string) (int,error) {
	f, err := os.Open(path)
	if err != nil {
	    return 0, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
	    if strings.Contains(scanner.Text(), seed) {
	        return line, nil
	    }

	    line++
	}

	if err := scanner.Err(); err != nil {
	    return 0, err
	}
	return 0, err
}

func main() {
	apk := ""
	flag.StringVar(&(apk),"f","","specify apk to search")
	flag.Parse()
	if apk == "" {
        panic("no apk dude")
    }

	
	arg0 := "apktool"
	arg1 := "d"
	arg2 := apk
	arg3 := "-o"
	arg4 := fmt.Sprintf("ss-"+filepath.Base(apk))
	
	if _, err := os.Stat("ss-"+filepath.Base(apk)); os.IsNotExist(err) {
		cmd := exec.Command(arg0,arg1,arg2,arg3,arg4)
	    err2 := cmd.Run()
	    if err2 != nil {
	    	fmt.Println("lol")
	        panic(err)
	    }
	} else {
		fmt.Println("[!] using cached decompiler")
	}
	

	var seeds = []string{}
	seeds = filetoslice("wordlist-sslandroot.txt")
	
	var wg sync.WaitGroup
 	wg.Add(len(seeds))
 	body := make(chan string,1000)

    for _,seed := range seeds {
	    go func(seed string)  {

		    signa := seed
		    data := ""
			err := filepath.Walk("./"+"ss-"+filepath.Base(apk), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
					return err
				}
				
				if info.IsDir() {
					return nil
				}

				line,err := openandfind(path,signa)
				if err != nil {
					fmt.Println("error walking the path")
					fmt.Println(err)
					return err
				}

				if line != 0 {
					data = signa+" | "+path+" | "+ strconv.Itoa(line)
					fmt.Println(data)
					body <- data
					return nil
				}
				return nil
			})

			if err != nil {
				fmt.Println("error walking the path2 : ")
				fmt.Println(err)
				//wg.Done()
			
			} 
			wg.Done()
			//return
			
		}(seed)
	
	}
	wg.Wait()
}


