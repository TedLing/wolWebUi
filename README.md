# wolWebUi
> wol网页界面 在线唤醒局域网设备 

# 使用说明
## 自行打包或使用编译的文件
 > window
 >> 双击运行 go_build_wolWebUi_.exe 即可 
 > 
 > linux
 >> ./wolWebUi 运行文件 
 > 
 > 配置文件位于 wolConf.json 可用于修改端口号

> linux 下作为服务启动
1. 创建文件
```shell
 nano /etc/systemd/system/wolwebui.service
```
2. 修改配置文件
```shell
[Unit]
Description=wolwebui
After=syslog.target
After=network.target
[Service]
RestartSec=2s
Type=notify
User=root
Group=root
WorkingDirectory=/app/wol_webui/
ExecStart=/app/wol_webui/wolWebUi
Restart=always
WatchdogSec=30s

[Install]
WantedBy=multi-user.target
```
3. 启动服务
```shell
# 启用
sudo systemctl enable wolwebui
# 启动
sudo systemctl start wolwebui
# 关闭
sudo systemctl stop wolwebui
```




# 使用框架
> gin + grom + layui  
> 引用项目 [github.com/ahmetozer/wakeonlan/share]