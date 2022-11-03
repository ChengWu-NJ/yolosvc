systemctl stop yolosvc
cat > /etc/systemd/system/yolosvc.service << 'EEE'
[Unit]
Description=darknet yolo service
After=network-online.target syslog.target
#StartLimitIntervalSec=30
#StartLimitBurst=2

[Service]
Type=simple
WorkingDirectory=/root/go/src/github.com/ChengWu-NJ/yolosvc/bin/yolosvc
User=root
ExecStart=/root/go/src/github.com/ChengWu-NJ/yolosvc/bin/yolosvc/yolosvc
Restart=on-failure
RestartSec=5
LimitNOFILE=65536
SyslogIdentifier=yolosvc

[Install]
WantedBy=multi-user.target
EEE
systemctl daemon-reload
systemctl enable yolosvc
systemctl start yolosvc
