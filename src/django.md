

# generate from exist table

python manage.py inspectdb <tbl_name>... > models.py


# mysql support issue

django.db.utils.NotSupportedError: MySQL 8.0.11 or later is required (found 5.7.20).

for old mysql version, use django4.1


pip3 install django==4.1



# mysql lib error on centos

rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022
yum -y install mysql-devel
pip install mysqlclient -i https://pypi.douban.com/simple/
export LD_LIBRARY_PATH=/usr/lib64/mysql/

