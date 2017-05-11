# GRAO

*Générateur de réponses aux appels d'offres*

## Nginx configuration

```nginx
server {
  listen 80;

  server_name localhost;
  auth_basic "Restricted";
  auth_basic_user_file /etc/nginx/.htpasswd;

  location /api/v1 {
    proxy_pass http://127.0.0.1:8090;
    proxy_redirect off;
  }

  location / {
    root   /var/www/html;
    index  index.html;
  }
}

```

## Systemd configuration

```systemd
[Unit]
AssertPathExists=/home/ubuntu
Description=GRAO core service
BindsTo=docd.service

[Service]
EnvironmentFile=/home/ubuntu/env
WorkingDirectory=/home/ubuntu
ExecStart=/bin/bash -c '/home/ubuntu/rao'
Restart=always
PrivateTmp=true
NoNewPrivileges=true
User=ubuntu

[Install]
WantedBy=multi-user.target
```

```SystemdAssertPathExists=/home/ubuntu
Description=Document conversion service

[Service]
WorkingDirectory=/home/ubuntu
ExecStart=/bin/bash -c '/home/ubuntu/docd'
Restart=always
PrivateTmp=true
NoNewPrivileges=true
User=ubuntu

[Install]
WantedBy=multi-user.target
```
