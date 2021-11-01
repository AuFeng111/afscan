go写的afscan，主要方便在webshell入口下的快速ssh、smb的弱口令探测。✔
------
####  刚起步，边学边写，代码偏基础，后续会重构。



Usage:</br>
./afscan.exe -model icmpalive -ip 192.168.201.1/24</br>
./afscan.exe -model icmpalive -ip 192.168.201.1/24 -t 500</br>
./afscan.exe -model portscan -ip 192.168.201.1/24 -port 22,445,1-10000</br>
./afscan.exe -model sshcrack -ip 192.168.201.1/24 -port 222(使用默认账号密码,指定端口爆破,port参数必须使用。)</br>
./afscan.exe -model sshcrack -ip 192.168.201.1/24 -port 22,222 -user root,admin -pass root,123456(指定用户名密码爆破，用逗号分割)</br>
./afscan.exe -model smbcrack -ip 192.168.201.1/24 -port 445 -user administrator,admin,guest,test(可指定密码,不指定默认使用默认密码)</br>
                       --by au7eng</br>
		       
#### usage:</br>		 
![Image text](https://raw.githubusercontent.com/AuFeng111/afscan/master/image.png)

</br>
#### 实战ssh爆破
</br>
![Image text](https://raw.githubusercontent.com/AuFeng111/afscan/master/image2.png)
