go写的afscan，主要方便在webshell入口下的快速ssh、smb的弱口令探测。

Usage:
./afscan.exe -model icmpalive -ip 192.168.201.1/24
./afscan.exe -model icmpalive -ip 192.168.201.1/24 -t 500
./afscan.exe -model portscan -ip 192.168.201.1/24 -port 22,445,1-10000
./afscan.exe -model sshcrack -ip 192.168.201.1/24 -port 222(使用默认账号密码,指定端口爆破,port参数必须使用。)
./afscan.exe -model sshcrack -ip 192.168.201.1/24 -port 22,222 -user root,admin -pass root,123456(指定用户名密码爆破，用逗号分割)
./afscan.exe -model smbcrack -ip 192.168.201.1/24 -port 445 -user administrator,admin,guest,test(可指定密码,不指定默认使用默认密码)
                       --by au7eng

Options:
  -ip string
        192.168.1-255.1-10
        192.168.1.1/24
        192.168.1.* (default "127.0.0.1")
  -model string
        icmpalive
        portscan
        sshcrack
        smbcrack
  -pass string
         (default "123456,admin,admin123,root,,pass123,pass@123,password,123123,654321,111111,123,1,admin@123,Admin@123,admin123!@#,P@ssw0rd!,P@ssw0rd,Passw0rd,qwe123,12345678,test,test123,123qwe!@#,123456789,123321,666666,a123456.,123456~a,123456!a,000000,1234567890,8888888,!QAZ2wsx,1qaz2wsx,abc123,abc123456,1qaz@WSX,a11111,a12345,Aa1234,Aa1234.,Aa12345,a123456,a123123,Aa123123,Aa123456,Aa12345.,sysadmin,system,1qaz!QAZ,2wsx@WSX,qwe123!@#,Aa123456!,A123456s!,sa123456,1q2w3e")
  -port string
         (default "22,80,445,3389,8000-9000")
  -t int
        thread (default 500)
  -user string
         (default "ubuntu,admin,test,user,root")