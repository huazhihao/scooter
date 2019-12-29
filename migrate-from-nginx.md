# Migrate from nginx

## Basic mapping

`nginx.conf`

```
location /some/path/ {
    proxy_pass http://www.example.com/link/;
}
```

`scooter.yaml`

```yaml
- rules:
    - path: /some/path/
      url: "http://www.example.com/link/"
```

## Set headers

`nginx.conf`

```
location /some/path/ {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_pass http://localhost:8000;
}
```

`scooter.yaml`

```yaml
- rules:
    - path: /some/path/
      url: "http://localhost:8000"
      headers:
        - key: Host
          value: $local_host
        - key: X-Real-IP
          value: $remote_ip
```

## Bind address

`nginx.conf`

```
location /app1/ {
    proxy_bind 127.0.0.1;
    proxy_pass http://example.com/app1/;
}
```

`scooter.yaml`

```yaml
- name: api-gateway
  bind: "127.0.0.1"
  rules:
    - path: /
      url: "http://example.com/app1/"
```
