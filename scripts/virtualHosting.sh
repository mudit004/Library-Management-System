
sudo apt install apache2
sudo a2enmod proxy proxy_http
read -p "Enter your Server Name: " ServerName
read -p "Enter you Server Admin: " ServerAdmin
conf_Content=$(cat <<EOF
<VirtualHost *:80>
	ServerName "$ServerName"
	ServerAdmin "$ServerAdmin"
	ProxyPreserveHost On
	ProxyPass / http://127.0.0.1:8000/
	ProxyPassReverse / http://127.0.0.1:8000/
	TransferLog /var/log/apache2/mvc_access.log
	ErrorLog /var/log/apache2/mvc_error.log
</VirtualHost>
EOF
)

sudo echo $conf_Content >> ${ServerName}.conf 
sudo mv ${ServerName}.conf /etc
cd /etc
sudo a2ensite ${ServerName}.conf
sudo a2dissite 000-default.conf
sudo apache2ctl configtest
sudo systemctl restart apache2
sudo systemctl status apache2