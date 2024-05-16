# 使用 openssl 工具生成一个有效期为 365 天的自签名的证书和私钥
# 自签名证书不会被浏览器信任，出于安全考虑，在生产环境应该使用由受信任的证书颁发机构签发的证书
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -addext "subjectAltName=DNS:localhost"
